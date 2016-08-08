package api

import (
	"github.com/gin-gonic/gin"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// VoteNote handles a users request to vote on a note
func VoteNote(c *gin.Context) {
	// Get variables from request context
	user := c.MustGet("user").(models.User)
	note := c.MustGet("note").(models.Note)
	direction := c.PostForm("dir")

	// Change vote
	tx := db.DB.Begin()
	var err error
	err = tx.Model(&note).Association("UpvoteUsers").Delete(user).Error
	err = tx.Model(&note).Association("DownvoteUsers").Delete(user).Error
	if direction == "1" {
		err = tx.Model(&note).Association("UpvoteUsers").Append(user).Error
	} else if direction == "-1" {
		err = tx.Model(&note).Association("DownvoteUsers").Append(user).Error
	}
	if err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}
	tx.Commit()

	// Reset note ranking
	upvoteUsers := db.DB.Model(&note).Association("UpvoteUsers").Count()
	downvoteUsers := db.DB.Model(&note).Association("DownvoteUsers").Count()
	if db.DB.Error != nil {
		c.AbortWithError(500, err)
		return
	}
	ranking := upvoteUsers - downvoteUsers
	if err = db.DB.Exec(
		"UPDATE notes SET ranking = ? WHERE id = ?",
		ranking,
		note.ID).
		Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
}

// VoteCollection handles a users request to vote on a Collection
func VoteCollection(c *gin.Context) {
	// Get variables from request context
	user := c.MustGet("user").(models.User)
	collection := c.MustGet("collection").(models.Collection)
	direction := c.PostForm("dir")

	// Change vote
	tx := db.DB.Begin()
	var err error
	err = tx.Model(&collection).Association("UpvoteUsers").Delete(user).Error
	err = tx.Model(&collection).Association("DownvoteUsers").Delete(user).Error
	if direction == "1" {
		err = tx.Model(&collection).Association("UpvoteUsers").Append(user).Error
	} else if direction == "-1" {
		err = tx.Model(&collection).Association("DownvoteUsers").Append(user).Error
	}
	if err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}
	tx.Commit()

	// Reset collection ranking
	upvoteUsers := db.DB.Model(&collection).Association("UpvoteUsers").Count()
	downvoteUsers := db.DB.Model(&collection).Association("DownvoteUsers").Count()
	if db.DB.Error != nil {
		c.AbortWithError(500, err)
		return
	}
	ranking := upvoteUsers - downvoteUsers
	if err = db.DB.Exec(
		"UPDATE collections SET ranking = ? WHERE id = ?",
		ranking,
		collection.ID).
		Error; err != nil {
		c.AbortWithError(500, err)
		return
	}
}