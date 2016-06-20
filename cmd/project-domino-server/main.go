package main

import (
	// Standard Library
	"fmt"
	"log"
	"os"

	// Extended Standard Library
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"

	// Internal Dependencies
	"github.com/project-domino/project-domino/common"
	"github.com/project-domino/project-domino/handlers"
	"github.com/project-domino/project-domino/handlers/api"
	"github.com/project-domino/project-domino/middleware"
	"github.com/project-domino/project-domino/models"
	"github.com/spf13/viper"

	// Third-Party Dependencies
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Database Drivers
	_ "github.com/denisenkom/go-mssqldb" // MS SQL
	_ "github.com/go-sql-driver/mysql"   // MySQL, MariaDB
	_ "github.com/lib/pq"                // PostgreSQL
	_ "github.com/mattn/go-sqlite3"      // SQLite 3.x.y
)

func main() {
	// Open database connection.
	db, err := gorm.Open(
		viper.GetString("database.type"),
		viper.GetString("database.url"),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(viper.GetBool("database.debug"))
	SetupDatabase(db)

	// Enable/disable gin's debug mode.
	if viper.GetBool("http.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.DatabaseMiddleware(db))
	r.Use(middleware.LoginMiddleware)

	// Load assets and templates.
	// TODO: There's a better way...
	var assetFS vfs.FileSystem
	if viper.GetBool("assets.dev") {
		assetFS = vfs.OS("assets/dist")
	} else {
		var err error
		assetFS, err = NewZipFileSystem(viper.GetString("assets.path"))
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

	// View Routes
	r.GET("/", handlers.Simple("home.html"))
	r.GET("/u/:username", handlers.TODO)
	r.GET("/uni/:uni-short-name", handlers.TODO)
	r.GET("/search", handlers.TODO)

	r.GET("/note/:note-id", handlers.TODO)
	r.GET("/note/:note-id/:note-name", handlers.TODO)

	r.GET("/collection/:collectionID", handlers.TODO)
	r.GET("/collection/:collectionID/note/:noteID", handlers.TODO)
	r.GET("/collection/:collectionID/note/:noteID/:noteName", handlers.TODO)

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

	// API
	r.GET("/search/tag", api.SearchTags)

	r.POST("/note",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin),
		api.NewNote)
	r.PUT("/note", handlers.TODO)

	r.POST("/collection", handlers.TODO)
	r.PUT("/collection", handlers.TODO)

	// Debug Routes
	debug := r.Group("/debug")
	debug.GET("/editor", handlers.Simple("editor.html"))
	debug.GET("/env", func(c *gin.Context) { c.JSON(200, os.Environ()) })
	debug.GET("/config", func(c *gin.Context) { c.JSON(200, viper.AllSettings()) })
	debug.GET("/new/note", handlers.Simple("new-note.html"))

	// Start serving.
	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("http.port"))); err != nil {
		panic(err)
	}
}
