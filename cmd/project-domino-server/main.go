package main

import (
	// Standard Library
	"fmt"

	// Internal Dependencies
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/handlers"
	"github.com/project-domino/project-domino/handlers/api"
	"github.com/project-domino/project-domino/handlers/redirect"
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
	var err error
	db.DB, err = gorm.Open(
		viper.GetString("database.type"),
		viper.GetString("database.url"),
	)
	Must(err)
	defer db.DB.Close()
	db.DB.LogMode(viper.GetBool("database.debug"))
	Must(SetupDatabase(db.DB))

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
	r.Use(middleware.ErrorHandler())
	Must(SetupAssets(r))

	// Routes that require user object
	m := r.Group("/")
	m.Use(middleware.Login())

	// Authentication Routes
	m.GET("/login", handlers.Simple("login.html"))
	m.GET("/register", handlers.Simple("register.html"))
	m.POST("/login", api.Login)
	m.POST("/register", api.Register)
	m.POST("/logout", api.Logout)

	// View Routes
	m.GET("/", handlers.Simple("home.html"))
	m.GET("/uni/:uni-short-name", handlers.TODO)
	m.GET("/search", handlers.TODO)

	m.Group("/u/:username",
		middleware.LoadUser("Notes", "Collections")).
		GET("/", handlers.Simple("user-notes.html")).
		GET("/notes", handlers.Simple("user-notes.html")).
		GET("/collections", handlers.Simple("user-collections.html"))

	m.Group("/note",
		middleware.LoadNote("Author", "Tags")).
		GET("/:noteID", handlers.Simple("individual-note.html")).
		GET("/:noteID/:note-name", handlers.Simple("individual-note.html"))

	m.Group("/collection").
		GET("/:collectionID", handlers.TODO).
		GET("/:collectionID/note/:noteID",
			middleware.LoadNote("Author", "Tags"),
			middleware.LoadCollection("Author", "Tags"),
			handlers.Simple("collection-note.html")).
		GET("/:collectionID/note/:noteID/:noteName",
			middleware.LoadNote("Author", "Tags"),
			middleware.LoadCollection("Author", "Tags"),
			handlers.Simple("collection-note.html"))

	m.Group("/writer-panel",
		middleware.RequireAuth(),
		middleware.RequireUserType(models.Writer, models.Admin),
		middleware.LoadRequestUser("Notes", "Collections")).
		GET("/", redirect.WriterPanel).
		GET("/note", handlers.Simple("new-note.html")).
		GET("/note/:noteID/edit",
			middleware.LoadNote("Author", "Tags"),
			middleware.VerifyNoteOwner(),
			handlers.Simple("edit-note.html")).
		GET("/collection", handlers.Simple("new-collection.html")).
		GET("/collection/:collectionID/edit",
			middleware.LoadCollection("Author", "Tags"),
			middleware.VerifyCollectionOwner(),
			handlers.Simple("edit-collection.html")).
		GET("/tag", handlers.Simple("new-tag.html"))

	// API
	m.Group("/api/v1").
		GET("/search/tag", api.SearchTags).
		GET("/search/note", api.SearchNotes).
		POST("/note",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewNote).
		PUT("/note/:noteID",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.EditNote).
		POST("/collection",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewCollection).
		PUT("/collection/:collectionID",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.EditCollection).
		POST("/tag",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewTag)

	// Debug Routes
	if viper.GetBool("http.debug") {
		m.Group("/debug").
			GET("/editor", handlers.Simple("editor.html")).
			GET("/error", func(c *gin.Context) {
				errors.Debug.Apply(c)
			}).
			GET("/config", func(c *gin.Context) {
				handlers.RenderData(c, "debug.html", "data", viper.AllSettings())
			}).
			GET("/new/note", handlers.Simple("new-note.html"))
	}

	// Start serving.
	Must(r.Run(fmt.Sprintf(":%d", viper.GetInt("http.port"))))
}
