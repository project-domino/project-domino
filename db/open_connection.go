package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Open opens a database connection
func Open(dbType string, dbURL string, debug bool) {
	opened := false
	for !opened {
		var err error
		log.Printf("Connecting to DB at %s...", dbURL)
		DB, err = gorm.Open(
			dbType,
			dbURL,
		)
		opened = err == nil
		log.Println(err)
		time.Sleep(time.Second)
	}
	DB.LogMode(true)
}
