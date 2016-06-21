package main

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// SetupDatabase initializes the database with empty tables of all the needed
// types.
func SetupDatabase(db *gorm.DB) error {
	setupTable(db, &models.User{})
	setupTable(db, &models.AuthToken{})
	setupTable(db, &models.Comment{})
	setupTable(db, &models.Note{})
	setupTable(db, &models.Collection{})
	setupTable(db, &models.Textbook{})
	setupTable(db, &models.University{})
	setupTable(db, &models.Tag{})
	return db.Error
}

// Creates a table for a specified struct if one doesn't exist
func setupTable(db *gorm.DB, val interface{}) {
	if !db.HasTable(val) {
		db.CreateTable(val)
	}
}
