package main

import (
	"log"
	"net/smtp"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/project-domino/project-domino/db"
)

func main() {
	db.Open(Config.Database.Type, Config.Database.URL, Config.Database.Debug)
	defer db.DB.Close()

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
		email, err := GetEmail(db.DB)

		// If there is an email and no error, send the email
		if err != nil {
			log.Fatal(err)
		} else {
			if email.ID != 0 {
				err := SendEmail(email, auth, Config.SMTP.Address)
				if err != nil {
					log.Println(err)
				}
				if err := MarkSent(db.DB, email); err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}
