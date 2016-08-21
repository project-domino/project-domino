package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// LoadRankItems loads the notes, collections, and comments the request user
// has upvoted or downvoted. It only loads the ids of these items for efficiency.
func LoadRankItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)

		if user.ID != 0 {
			var err error

			preloadedDB := db.DB.Model(&user).Select("id")

			err = preloadedDB.Association("UpvoteNotes").Find(&user.UpvoteNotes).Error
			err = preloadedDB.Association("DownvoteNotes").Find(&user.DownvoteNotes).Error
			err = preloadedDB.Association("UpvoteCollections").Find(&user.UpvoteCollections).Error
			err = preloadedDB.Association("DownvoteCollections").Find(&user.DownvoteCollections).Error
			err = preloadedDB.Association("UpvoteComments").Find(&user.UpvoteComments).Error
			err = preloadedDB.Association("DownvoteComments").Find(&user.DownvoteComments).Error

			if err != nil {
				errors.DB.Apply(c)
				return
			}

			c.Set("user", user)
		}

		c.Next()
	}
}
