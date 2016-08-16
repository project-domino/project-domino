package notifications

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// Comment creates a notification if someone has replied to a user's comment
func Comment(db *gorm.DB, actor models.User, subject models.User, comment models.Comment) error {

	title := fmt.Sprintf("%s replied to your %s.", actor.UserName, comment.Type)
	link := fmt.Sprintf("/note/%d", comment.NoteID)

	if err := db.Create(&models.Notification{
		SubjectID: subject.ID,
		ActorID:   actor.ID,
		Type:      CommentNotificationType,
		Title:     title,
		Link:      link,
	}).Error; err != nil {
		return err
	}

	return nil
}
