package models

import "github.com/jinzhu/gorm"

// These constants are the valid values for Comment.Type
const (
	QuestionComment   string = "question"
	SuggestionComment        = "suggestion"
)

// A Comment is a response to a note
type Comment struct {
	gorm.Model

	Type string

	UserID uint
	User   User

	NoteID uint
	Note   Note

	ParentID uint      `json:"-" xml:"-"`
	Children []Comment `gorm:"-"`

	Body string

	// Ranking Info
	Ranking       int
	UpvoteUsers   []User `gorm:"many2many:upvotecomment_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotecomment_user;"`
}

// InList checks if a given comment is in a list
func (c Comment) InList(l []Comment) bool {
	for _, e := range l {
		if c.ID == e.ID {
			return true
		}
	}
	return false
}
