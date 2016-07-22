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
	// qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	if err := db.DB.
		Preload("Tags").
		Where(queryFormat, qText).
		Where("published = ?", true).
		Find(&notes).
		Limit(items).
		Offset(page * items).
		Error; err != nil {
		return notes, err
	}

	return notes, nil
}
