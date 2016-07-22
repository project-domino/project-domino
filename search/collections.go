package search

import (
	"strings"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Collections returns all collections that match a given query
func Collections(q string, items int, page int) ([]models.Collection, error) {
	var collections []models.Collection

	searchQuery, err := ParseQuery(q)
	if err != nil {
		return collections, err
	}
	// qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	if q != "" {
		if err := db.DB.
			Preload("Tags").
			Where(queryFormat, qText).
			Where("published = ?", true).
			Find(&collections).
			Limit(items).
			Offset(page * items).
			Error; err != nil {
			return collections, err
		}
	}
	return collections, nil
}
