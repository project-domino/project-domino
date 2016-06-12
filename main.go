package main

import (
	// stdlib
	"flag"
	"log"
	"net/http"
	"os"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"

	// Internal dependencies

	"github.com/project-domino/project-domino/common"
	"github.com/project-domino/project-domino/handlers/api"
	"github.com/project-domino/project-domino/handlers/debug"
	"github.com/project-domino/project-domino/handlers/view"
	"github.com/project-domino/project-domino/middleware"

	// 3rd-party dependencies
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/vharitonsky/iniflags"

	// Database drivers
	_ "github.com/denisenkom/go-mssqldb" // MS SQL
	_ "github.com/go-sql-driver/mysql"   // MySQL, MariaDB
	_ "github.com/lib/pq"                // PostgreSQL
	_ "github.com/mattn/go-sqlite3"      // SQLite 3.x.y
)

var (
	assetPath = flag.String("assetPath", "assets.zip", "The zip file to load assets from.")
	dbAddr    = flag.String("dbAddr", "domino.db", "The database's address or path.")
	dbType    = flag.String("dbType", "sqlite3", "The database's type.")
	dbDebug   = flag.Bool("dbDebug", false, "Enables debugging on the database.")
	dev       = flag.Bool("dev", false, "Load assets from a directory instead of a .zip file.")
	serveOn   = flag.String("serveOn", "default", "The address to serve on.")
)

func init() {
	iniflags.Parse()
}

func main() {
	// Open database connection.
	db, err := gorm.Open(*dbType, *dbAddr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(*dbDebug)
	SetupDatabase(db)

	// Set up routes.
	r := mux.NewRouter()

	// Assets
	var assetFS vfs.FileSystem
	if *dev {
		assetFS = vfs.OS("assets/dist")
	} else {
		var err error
		assetFS, err = NewZipFileSystem(*assetPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	r.Methods("GET").Path("/assets/{file}").Handler(http.FileServer(httpfs.New(assetFS)))
	if err := common.LoadTemplates(assetFS); err != nil {
		log.Fatal(err)
	}

	// Authentication Routes
	r.Methods("GET").Path("/login").HandlerFunc(view.LoginHandler)
	r.Methods("GET").Path("/register").HandlerFunc(view.RegisterHandler)
	r.Methods("POST").Path("/login").HandlerFunc(api.LoginHandler)
	r.Methods("POST").Path("/register").HandlerFunc(api.RegisterHandler)
	r.Methods("POST").Path("/logout").HandlerFunc(api.LogoutHandler)

	// View Routes
	r.Methods("GET").Path("/").HandlerFunc(view.HomeHandler)
	r.Methods("GET").Path("/subjects").HandlerFunc(view.SubjectListHandler)
	r.Methods("GET").Path("/subject/{subjectName}").HandlerFunc(view.SubjectHandler)
	r.Methods("GET").Path("/subject/{subjectName}/{topicName}").HandlerFunc(view.TopicHandler)
	r.Methods("GET").Path("/subject/{subjectName}/{topicName}/{noteID}").HandlerFunc(view.NoteHandler)
	// noteName for url readability
	r.Methods("GET").Path("/subject/{subjectName}/{topicName}/{noteID}/{noteName}").HandlerFunc(view.NoteHandler)
	r.Methods("GET").Path("/user/{userName}").HandlerFunc(view.UserHandler)
	r.Methods("GET").Path("/university/{shortName}").HandlerFunc(view.UniversityHandler)
	r.Methods("GET").Path("/search").HandlerFunc(view.SearchHandler)

	// Debug Routes
	r.Methods("GET").Path("/debug/editor").HandlerFunc(debug.EditorHandler)

	// Set up
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseFunc(middleware.DatabaseMiddleware(db))
	n.UseFunc(middleware.LoginMiddleware)
	n.UseHandler(r)

	// Start serving.
	if *serveOn == "default" {
		if port := os.Getenv("PORT"); port != "" {
			*serveOn = ":" + port
		} else {
			*serveOn = ":80"
		}
	}
	n.Run(*serveOn)
}
