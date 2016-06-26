package util

import "github.com/project-domino/project-domino/models"

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

// GetCollections gets collections from the db whose ids are in a specified slice
func GetCollections(ids []uint) []models.Collection {
	var collections []models.Collection
	DB.Where("id in (?)", ids).Find(&collections)
	return collections
}
