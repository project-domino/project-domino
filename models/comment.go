package models

import "github.com/jinzhu/gorm"

// A Comment is a response to a note
type Comment struct {
	gorm.Model

	UserID uint
	User   User

	NoteID uint
	Note   Note

	// No recursive structs in go
	CommentID uint

	Body string

	// Favorite Info
	FavoriteUsers []User `gorm:"many2many:favoritecollection_user;"`

	// Ranking Info
	Upvotes       uint
	Downvotes     uint
	UpvoteUsers   []User `gorm:"many2many:upvotecomment_user;"`
	DownvoteUsers []User `gorm:"many2many:downvotecomment_user;"`
}
