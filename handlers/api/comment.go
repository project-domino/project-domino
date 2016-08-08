package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// NewComment creates a comment with specified values
func NewComment(c *gin.Context) {
	// Get variables from context
	user := c.MustGet("user").(models.User)
	noteID := c.Param("noteID")
	commentType := c.Param("commentType")
	parentID := c.PostForm("parentID")
	body := c.PostForm("body")

	var comment models.Comment

	// Get note
	var note models.Note
	if err := db.DB.
		Where("ID = ?", noteID).
		Find(&note).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			errors.NoteNotFound.Apply(c)
		} else {
			errors.InternalError.Apply(c)
		}
		return
	}

	// Get comment
	var parentComment models.Comment
	if parentID != "" {
		if err := db.DB.
			Where("ID = ?", parentID).
			Where("note_id = ?", note.ID).
			Find(&parentComment).
			Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				errors.CommentNotFound.Apply(c)
				return
			}
		}
		if parentComment.ParentID != 0 {
			errors.CommentNesting.Apply(c)
			return
		}
	}

	// Verify valid comment type
	if (commentType != models.QuestionComment) && (commentType != models.SuggestionComment) {
		errors.NotFound.Apply(c)
		return
	}

	comment.UserID = user.ID
	comment.NoteID = note.ID
	comment.Type = commentType
	comment.ParentID = parentComment.ID
	comment.Body = body

	if err := db.DB.Create(&comment).Error; err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, comment)
}
