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
	qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	preloadedDB := db.DB.
		Preload("Tags").
		Where("collections.published = ?", true)

	// If there are tag selectors, add them to the query
	if len(qSelectors[TagSelector]) > 0 {
		preloadedDB = preloadedDB.
			Joins("JOIN collection_tag ON collection_tag.collection_id = collections.id").
			Joins("JOIN tags ON collection_tag.tag_id = tags.id").
			Where("tags.name IN (?)", qSelectors[TagSelector])
	}

	if qText != "" {
		preloadedDB = preloadedDB.Where("collections."+queryFormat, qText)
	}

	// If there are author selectors, add them to the query
	if len(qSelectors[AuthorSelector]) > 0 {
		var userIDs []uint
		if err := db.DB.
			Model(&models.User{}).
			Where("user_name IN (?)", qSelectors[AuthorSelector]).
			Pluck("ID", &userIDs).
			Error; err != nil {
			return collections, err
		}
		preloadedDB = preloadedDB.Where("collections.author_id IN (?)", userIDs)
	}

	if err := preloadedDB.
		Limit(items).
		Offset((page - 1) * items).
		Find(&collections).
		Error; err != nil {
		return collections, err
	}

	return collections, nil
}
