package main

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino-server/models"
)

// SetupDatabase initializes the database with empty tables of all the needed
// types.
func SetupDatabase(db *gorm.DB) {
	db.CreateTable(&models.User{})
	db.CreateTable(&models.AuthToken{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Note{})
	db.CreateTable(&models.Outline{})
	db.CreateTable(&models.Subject{})
	db.CreateTable(&models.University{})
}
