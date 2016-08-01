package search

import (
	"strings"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Tags returns all tags that match a given query
func Tags(q string, items int, page int) ([]models.Tag, error) {
	var tags []models.Tag

	searchQuery, err := ParseQuery(q)
	if err != nil {
		return tags, err
	}
	// qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	if err := db.DB.Where(queryFormat, qText).
		Limit(items).
		Offset((page - 1) * items).
		Find(&tags).
		Error; err != nil {
		return tags, err
	}

	return tags, nil
}
