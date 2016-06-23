package main

import (
	// Standard Library
	"fmt"
	"os"

	// Internal Dependencies
	"github.com/project-domino/project-domino/handlers"
	"github.com/project-domino/project-domino/handlers/api"
	"github.com/project-domino/project-domino/handlers/view"
	"github.com/project-domino/project-domino/middleware"
	"github.com/project-domino/project-domino/models"

	// Third-Party Dependencies
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

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
	Must(err)
	defer db.Close()
	db.LogMode(viper.GetBool("database.debug"))
	Must(SetupDatabase(db))

	// Enable/disable gin's debug mode.
	if viper.GetBool("http.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create and set up router.
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.DatabaseMiddleware(db))
	r.Use(middleware.LoginMiddleware)
	Must(SetupAssets(r))

	// Authentication Routes
	r.GET("/login", handlers.Simple("login.html"))
	r.GET("/register", handlers.Simple("register.html"))
	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.POST("/logout", api.Logout)

	// View Routes
	r.GET("/", handlers.Simple("home.html"))
	r.GET("/u/:username", handlers.TODO)
	r.GET("/uni/:uni-short-name", handlers.TODO)
	r.GET("/search", handlers.TODO)

	r.Group("/note").
		GET("/:note-id", handlers.TODO).
		GET("/:note-id/:note-name", handlers.TODO)

	r.Group("/collection").
		GET("/:collectionID", handlers.TODO).
		GET("/:collectionID/note/:noteID", handlers.TODO).
		GET("/:collectionID/note/:noteID/:noteName", handlers.TODO)

	r.Group("/writer-panel",
		middleware.RequireAuth,
		middleware.RequireUserType(models.Writer, models.Admin)).
		GET("/", view.WriterPanelSimple("writer-panel.html")).
		GET("/note", view.WriterPanelSimple("new-note.html")).
		GET("/note/:noteID/edit", view.EditNote).
		GET("/tag", view.WriterPanelSimple("new-tag.html"))

	// API
	r.Group("/api/v1").
		GET("/search/tag", api.SearchTags).
		POST("/note",
			middleware.RequireAuth,
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewNote).
		PUT("/note/:noteID",
			middleware.RequireAuth,
			middleware.RequireUserType(models.Writer, models.Admin),
			api.EditNote).
		POST("/tag",
			middleware.RequireAuth,
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewTag).
		POST("/collection", handlers.TODO).
		PUT("/collection", handlers.TODO)

	// Debug Routes
	debug := r.Group("/debug")
	debug.GET("/editor", handlers.Simple("editor.html"))
	debug.GET("/env", func(c *gin.Context) { c.JSON(200, os.Environ()) })
	debug.GET("/config", func(c *gin.Context) { c.JSON(200, viper.AllSettings()) })
	debug.GET("/new/note", handlers.Simple("new-note.html"))

	// Start serving.
	Must(r.Run(fmt.Sprintf(":%d", viper.GetInt("http.port"))))
}
