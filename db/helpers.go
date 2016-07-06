package db

import (
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// GetTags gets tags from the db whose ids are in a specified slice
func GetTags(ids []uint) []models.Tag {
	var tags []models.Tag
	DB.Where("id in (?)", ids).Find(&tags)
	return tags
}

// GetNotes gets notes from the db whose ids are in a specified slice
func GetNotes(ids []uint) []models.Note {
	var notes []models.Note
	DB.Where("id in (?)", ids).Find(&notes)
	return notes
}

// VerifyNotes checks if notes with given ids exist
func VerifyNotes(ids []uint) error {
	for _, id := range ids {
		var note models.Note
		DB.Where("id = ?", id).Find(&note)
		if note.ID == 0 {
			return errors.NoteNotFound
		}
	}
	return nil
}
