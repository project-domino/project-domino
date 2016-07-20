package search

import (
	"strings"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
)

// Users returns all users that match a given query
func Users(q string, items int, page int) ([]models.User, error) {
	var users []models.User

	searchQuery, err := ParseQuery(q)
	if err != nil {
		return users, err
	}
	// qSelectors := searchQuery.Selectors
	qText := strings.Join(searchQuery.Text, " & ")

	if err := db.DB.Where(queryFormat, qText).
		Find(&users).
		Limit(items).
		Offset(page * items).
		Error; err != nil {
		return users, err
	}
	return users, nil
}
