package main

import (
	"log"
	"net/smtp"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Open database connection.
	opened := false
	var err error
	var db *gorm.DB
	for !opened {
		log.Printf("Connecting to DB at %s...", Config.Database.URL)
		db, err = gorm.Open(
			Config.Database.Type,
			Config.Database.URL,
		)
		opened = err == nil
		time.Sleep(time.Second)
	}
	defer db.Close()
	db.LogMode(Config.Database.Debug)
	if err != nil {
		log.Fatal(err)
	}

	// Get smtp settings
	auth := smtp.PlainAuth(
		Config.SMTP.Identity,
		Config.SMTP.Username,
		Config.SMTP.Password,
		Config.SMTP.Host,
	)

	// Start polling db
	for _ = range time.Tick(time.Second) {
		// Get next email from the db
		email, err := GetEmail(db)

		// If there is an email and no error, send the email
		if err != nil {
			log.Fatal(err)
		} else {
			if email.ID != 0 {
				err := SendEmail(email, auth, Config.SMTP.Address)
				if err != nil {
					log.Println(err)
				}
				if err := MarkSent(db, email); err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}
