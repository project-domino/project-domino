package main

import (
	// Standard Library
	"fmt"
	"log"

	// Internal Dependencies
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/email"
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

	// Setup email worker
	sendgridAPIKey := viper.GetString("sendgrid.api.key")
	if sendgridAPIKey != "" {
		if err := email.Init(sendgridAPIKey); err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Sendgrid API key not found. Email not setup.")
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

	// Reset Password Routes
	m.Group("/reset-password").
		GET("/", handlers.Simple("reset-password.html")).
		POST("/", api.SendPasswordResetCode).
		PUT("/", api.ResetPassword)

	// Email verification routes
	m.Group("/email", middleware.RequireAuth()).
		GET("/verify", handlers.Simple("email-verify.html")).
		POST("/verify", api.SendEmailVerification).
		GET("/verify/:verificationCode", redirect.EmailVerify).
		GET("/conf", handlers.Simple("email-verify-conf.html"))

	// View Routes
	m.GET("/",
		middleware.LoadRankItems(),
		middleware.LoadFeaturedItems(),
		handlers.Simple("home.html"))

	m.Group("/account",
		middleware.RequireAuth()).
		GET("/", redirect.Account).
		PUT("/", api.EditUser).
		PUT("/password", api.ChangePassword).
		GET("/profile",
			handlers.Simple("account-profile.html")).
		GET("/security",
			handlers.Simple("account-security.html")).
		GET("/notifications",
			handlers.Simple("account-notifications.html"))

	m.GET("/search/:searchType",
		middleware.LoadRankItems(),
		middleware.LoadSearchItems(),
		handlers.Simple("search.html"))

	m.Group("/u/:username",
		middleware.LoadUser("Notes", "Collections")).
		GET("/", redirect.User).
		GET("/notes",
			middleware.LoadRequestUser("UpvoteNotes", "DownvoteNotes"),
			middleware.LoadNotes(middleware.LoadNotesAuthor, "Tags"),
			handlers.Simple("user-notes.html")).
		GET("/collections",
			middleware.LoadRequestUser("UpvoteCollections", "DownvoteCollections"),
			middleware.LoadCollections(middleware.LoadCollectionsAuthor, "Tags"),
			handlers.Simple("user-collections.html"))

	m.Group("/note",
		middleware.LoadRankItems(),
		middleware.LoadNote("Author", "Tags"),
		middleware.VerifyNotePublic()).
		GET("/:noteID", handlers.Simple("individual-note.html")).
		GET("/:noteID/:note-name", handlers.Simple("individual-note.html"))

	m.Group("/collection",
		middleware.LoadCollection("Author", "Tags"),
		middleware.VerifyCollectionPublic(),
		middleware.LoadRankItems()).
		GET("/:collectionID",
			handlers.Simple("collection.html")).
		GET("/:collectionID/note/:noteID",
			middleware.LoadNote("Author", "Tags"),
			handlers.Simple("collection-note.html")).
		GET("/:collectionID/note/:noteID/:noteName",
			middleware.LoadNote("Author", "Tags"),
			handlers.Simple("collection-note.html"))

	m.Group("/writer-panel",
		middleware.RequireAuth(),
		middleware.RequireUserType(models.Writer, models.Admin),
		middleware.LoadRequestUser("Notes", "Collections")).
		GET("/", redirect.WriterPanel).
		GET("/note",
			handlers.Simple("new-note.html")).
		GET("/note/:noteID/edit",
			middleware.LoadNote("Author", "Tags"),
			middleware.VerifyNoteOwner(),
			handlers.Simple("edit-note.html")).
		GET("/collection",
			handlers.Simple("new-collection.html")).
		GET("/collection/:collectionID/edit",
			middleware.LoadCollection("Author", "Tags"),
			middleware.VerifyCollectionOwner(),
			handlers.Simple("edit-collection.html")).
		GET("/tag",
			handlers.Simple("new-tag.html"))

	// API
	m.Group("/api/v1").
		GET("/notifications",
			middleware.RequireAuth(),
			middleware.LoadNotifications(),
			handlers.JSON("notifications")).
		PUT("/notification/:notificationID/read",
			middleware.RequireAuth(),
			api.MarkNotificationRead).
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
		PUT("/note/:noteID/vote",
			middleware.RequireAuth(),
			middleware.LoadNote(),
			api.VoteNote).
		GET("/note/:noteID/comments/:commentType",
			middleware.LoadComments("User"),
			handlers.JSON("comments")).
		POST("/note/:noteID/comments/:commentType",
			middleware.RequireAuth(),
			api.NewComment).
		POST("/collection",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.NewCollection).
		PUT("/collection/:collectionID",
			middleware.RequireAuth(),
			middleware.RequireUserType(models.Writer, models.Admin),
			api.EditCollection).
		PUT("/collection/:collectionID/vote",
			middleware.RequireAuth(),
			middleware.LoadCollection(),
			api.VoteCollection).
		PUT("/comment/:commentID/vote",
			middleware.RequireAuth(),
			middleware.LoadComment(),
			api.VoteComment).
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
