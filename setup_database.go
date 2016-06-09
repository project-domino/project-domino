package main

import (
	"github.com/jinzhu/gorm"
	"github.com/project-domino/project-domino/models"
)

// SetupDatabase initializes the database with empty tables of all the needed
// types.
func SetupDatabase(db *gorm.DB) {
	db.CreateTable(&models.User{})
	db.CreateTable(&models.AuthToken{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Note{})
	db.CreateTable(&models.Collection{})
	db.CreateTable(&models.Textbook{})
	db.CreateTable(&models.University{})
	db.CreateTable(&models.Tag{})
	db.CreateTable(&models.TagDepends{})
}
