package main

import (
	// Standard Library
	"flag"
	"log"

	// Extended Standard Library
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"

	// Internal Dependencies
	"github.com/project-domino/project-domino/common"
	"github.com/project-domino/project-domino/handlers"
	"github.com/project-domino/project-domino/handlers/api"
	"github.com/project-domino/project-domino/middleware"
	"github.com/project-domino/project-domino/models"

	// Third-Party Dependencies
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vharitonsky/iniflags"

	// Database Drivers
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

	// Create router.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.DatabaseMiddleware(db))
	r.Use(middleware.LoginMiddleware)

	// Load assets and templates.
	// TODO: There's a better way...
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
	if err := common.LoadTemplates(assetFS); err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(common.Views)
	r.StaticFS("/assets/", httpfs.New(assetFS))

	// Authentication Routes
	r.GET("/login", handlers.Simple("login.html"))
	r.GET("/register", handlers.Simple("register.html"))
	r.POST("/login", api.LoginHandler)
	r.POST("/register", api.RegisterHandler)
	r.POST("/logout", api.LogoutHandler)

	// Home Route
	r.GET("/", handlers.Simple("home.html"))

	// Collection Routes - names for URL readibility
	// r.GET("/collection/:collectionID", views.CollectionHandler)
	// r.GET("/collection/:collectionID/:collectionName", views.CollectionHandler)
	// r.GET("/collection/:collectionID/note/:noteID", views.CollectionNoteHandler)
	// r.GET("/collection/:collectionID/note/:noteID/:noteName", views.CollectionNoteHandler)

	// Note Routes
	r.GET("/note/:note-id", handlers.TODO)
	r.GET("/note/:note-id/:note-name", handlers.TODO)

	// User Routes
	r.GET("/user/:username", handlers.TODO)

	// University Route
	r.GET("/uni/:uni-short-name", handlers.TODO)

	// Search Route
	r.GET("/search", handlers.TODO)

	// Writer-panel routes
	r.GET("/writer-panel",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin),
		handlers.Simple("writer-panel.html"))
	r.GET("/writer-panel/note",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin),
		handlers.Simple("new-note.html"))
	r.GET("/writer-panel/tag",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin),
		handlers.Simple("new-tag.html"))

	// New routes
	r.POST("/new/note",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin),
		api.NewNote)
	// r.PUT("/new/note, api.EditNoteHandler)
	// r.GET("/new/collection, views.NewCollectionHandler)
	// r.POST("/new/collection, api.NewCollectionHandler)
	// r.PUT("/new/collection, api.EditCollectionHandler)

	// Debug Routes
	debug := r.Group("/debug")
	debug.GET("/editor", handlers.Simple("editor.html"))
	debug.GET("/new/note", handlers.Simple("new-note.html"))

	// Start serving.
	if err := r.Run(); err != nil {
		panic(err)
	}
}
