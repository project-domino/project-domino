package search

import (
	"strings"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Notes returns all notes that match a given query
func Notes(q string, items int, page int) ([]models.Note, error) {
	var notes []models.Note

	searchQuery, err := ParseQuery(q)
	if err != nil {
		return notes, err
	}
	qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	preloadedDB := db.DB.
		Preload("Tags").
		Where("notes.published = ?", true)

	// If there are tag selectors, add them to the query
	if len(qSelectors[TagSelector]) > 0 {
		preloadedDB = preloadedDB.
			Joins("JOIN note_tag ON note_tag.note_id = notes.id").
			Joins("JOIN tags ON note_tag.tag_id = tags.id").
			Where("tags.name IN (?)", qSelectors[TagSelector])
	}

	if qText != "" {
		preloadedDB = preloadedDB.Where("notes."+queryFormat, qText)
	}

	// If there are author selectors, add them to the query
	if len(qSelectors[AuthorSelector]) > 0 {
		var userIDs []uint
		if err := db.DB.
			Model(&models.User{}).
			Where("user_name IN (?)", qSelectors[AuthorSelector]).
			Pluck("ID", &userIDs).
			Error; err != nil {
			return notes, err
		}
		preloadedDB = preloadedDB.Where("notes.author_id IN (?)", userIDs)
	}

	if err := preloadedDB.
		Limit(items).
		Offset((page - 1) * items).
		Find(&notes).
		Error; err != nil {
		return notes, err
	}

	return notes, nil
}
