package main

import (
	// stdlib
	"flag"
	"log"
	"net/http"

	"golang.org/x/tools/godoc/vfs/httpfs"

	// Internal dependencies

	"github.com/project-domino/project-domino-server/common"
	"github.com/project-domino/project-domino-server/middleware"

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
	debug     = flag.Bool("debug", false, "Enables debugging.")
	serveOn   = flag.String("serveOn", ":80", "The address to serve on.")
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
	db.LogMode(*debug)
	SetupDatabase(db)

	// Set up routes.
	r := mux.NewRouter()

	// Assets
	assetFS, err := NewAssetFilesystem(*assetPath)
	if err != nil {
		log.Fatal(err)
	}
	r.Methods("GET").Handler(http.FileServer(httpfs.New(assetFS)))
	common.LoadTemplates(assetFS)

	// Authentication Routes
	// r.Methods("POST").Path("/login").HandlerFunc(api.LoginHandler)
	// r.Methods("POST").Path("/register").HandlerFunc(api.RegisterHandler)
	// r.Methods("POST").Path("/logout").HandlerFunc(api.LogoutHandler)

	// View Routes
	// r.Methods("GET").Path("/").HandlerFunc(view.HomeHandler)
	// r.Methods("GET").Path("/subjects").HandlerFunc(view.SubjectListHandler)
	// r.Methods("GET").Path("/subject/{subjectName}").HandlerFunc(view.SubjectHandler)
	// r.Methods("GET").Path("/subject/{subjectName}/{topicName}").HandlerFunc(view.TopicHandler)
	// r.Methods("GET").Path("/subject/{subjectName}/{topicName}/{noteURLID}/{noteName}").HandlerFunc(view.NoteHandler) // Note name for url readability
	// r.Methods("GET").Path("/subject/{subjectName}/{topicName}/{noteURLID}").HandlerFunc(view.NoteHandler)
	// r.Methods("GET").Path("/user/{userName}").HandlerFunc(view.UserHandler)
	// r.Methods("GET").Path("/university/{shortName}").HandlerFunc(view.UniversityHandler)
	// r.Methods("GET").Path("/search").HandlerFunc(view.SearchHandler)

	// Start serving.
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseFunc(middleware.DatabaseMiddleware(db))
	n.UseFunc(middleware.LoginMiddleware)
	n.UseHandler(r)
	n.Run(*serveOn)
}
