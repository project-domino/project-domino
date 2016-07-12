package db

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/errors"
	"github.com/project-domino/project-domino/models"
)

// GetTags gets tags from the db whose ids are in a specified slice
func GetTags(ids []uint) ([]models.Tag, error) {
	var tags []models.Tag
	if err := DB.Where("id in (?)", ids).
		Find(&tags).Error; err != nil && err != gorm.ErrRecordNotFound {
		return tags, err
	}
	return tags, nil
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

// LoadCollectionNotes loads the notes into a collection
func LoadCollectionNotes(c *models.Collection) error {
	// Find collection note relationships
	var collectioNotes []models.CollectionNote
	if err := DB.Preload("Note").
		Where("collection_id = ?", c.ID).
		Order("order").
		Find(&collectioNotes).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// Load notes into collection
	for _, collectionNote := range collectioNotes {
		c.Notes = append(c.Notes, collectionNote.Note)
	}

	return nil
}
