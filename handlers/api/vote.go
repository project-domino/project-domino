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

	// Reset note ranking
	upvoteUsers := db.DB.Model(&note).Association("UpvoteUsers").Count()
	downvoteUsers := db.DB.Model(&note).Association("DownvoteUsers").Count()
	if db.DB.Error != nil {
		c.AbortWithError(500, err)
		return
	}
	ranking := upvoteUsers - downvoteUsers
	if err = tx.Exec(
		"UPDATE notes SET ranking = ? WHERE id = ?",
		ranking,
		note.ID).
		Error; err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	tx.Commit()
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

	// Reset collection ranking
	upvoteUsers := db.DB.Model(&collection).Association("UpvoteUsers").Count()
	downvoteUsers := db.DB.Model(&collection).Association("DownvoteUsers").Count()
	if db.DB.Error != nil {
		c.AbortWithError(500, err)
		return
	}
	ranking := upvoteUsers - downvoteUsers
	if err = tx.Exec(
		"UPDATE collections SET ranking = ? WHERE id = ?",
		ranking,
		collection.ID).
		Error; err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	tx.Commit()
}

// VoteComment handles a users request to vote on a comment
func VoteComment(c *gin.Context) {
	// Get variables from request context
	user := c.MustGet("user").(models.User)
	comment := c.MustGet("comment").(models.Comment)
	direction := c.PostForm("dir")

	// Change vote
	tx := db.DB.Begin()
	var err error
	err = tx.Model(&comment).Association("UpvoteUsers").Delete(user).Error
	err = tx.Model(&comment).Association("DownvoteUsers").Delete(user).Error
	if direction == "1" {
		err = tx.Model(&comment).Association("UpvoteUsers").Append(user).Error
	} else if direction == "-1" {
		err = tx.Model(&comment).Association("DownvoteUsers").Append(user).Error
	}
	if err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	// Reset comment ranking
	upvoteUsers := db.DB.Model(&comment).Association("UpvoteUsers").Count()
	downvoteUsers := db.DB.Model(&comment).Association("DownvoteUsers").Count()
	if db.DB.Error != nil {
		c.AbortWithError(500, err)
		return
	}
	ranking := upvoteUsers - downvoteUsers
	if err = tx.Exec(
		"UPDATE comments SET ranking = ? WHERE id = ?",
		ranking,
		comment.ID).
		Error; err != nil {
		tx.Rollback()
		c.AbortWithError(500, err)
		return
	}

	tx.Commit()
}
