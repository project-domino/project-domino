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

	// Database Driver
	_ "github.com/lib/pq"
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

	m.Group("/account",
		middleware.RequireAuth()).
		GET("/", redirect.Account).
		GET("/profile",
			middleware.AddPageName("profile"),
			handlers.Simple("account-profile.html")).
		GET("/security",
			middleware.AddPageName("security"),
			handlers.Simple("account-security.html")).
		GET("/notifications",
			middleware.AddPageName("notifications"),
			handlers.Simple("account-notifications.html"))

	m.GET("/search/:searchType",
		middleware.LoadSearchItems(),
		middleware.LoadSearchVars(),
		handlers.Simple("search.html"))

	m.Group("/u/:username",
		middleware.LoadUser("Notes", "Collections", "Notes.Tags", "Collections.Tags")).
		GET("/", redirect.User).
		GET("/notes", handlers.Simple("user-notes.html")).
		GET("/collections", handlers.Simple("user-collections.html"))

	m.Group("/note",
		middleware.LoadNote("Author", "Tags")).
		GET("/:noteID", handlers.Simple("individual-note.html")).
		GET("/:noteID/:note-name", handlers.Simple("individual-note.html"))

	m.Group("/collection").
		GET("/:collectionID",
			middleware.LoadCollection("Author", "Tags"),
			handlers.Simple("collection.html")).
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
		GET("/note",
			middleware.AddPageName("new-note"),
			handlers.Simple("new-note.html")).
		GET("/note/:noteID/edit",
			middleware.LoadNote("Author", "Tags"),
			middleware.VerifyNoteOwner(),
			handlers.Simple("edit-note.html")).
		GET("/collection",
			middleware.AddPageName("new-collection"),
			handlers.Simple("new-collection.html")).
		GET("/collection/:collectionID/edit",
			middleware.LoadCollection("Author", "Tags"),
			middleware.VerifyCollectionOwner(),
			handlers.Simple("edit-collection.html")).
		GET("/tag",
			middleware.AddPageName("new-tag"),
			handlers.Simple("new-tag.html"))

	// API
	m.Group("/api/v1").
		GET("/search/:searchType",
			middleware.LoadSearchItems(),
			api.Search).
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
