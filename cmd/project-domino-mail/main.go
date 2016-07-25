package main

import (
	// Standard Library

	// Internal Dependencies
	"log"
	"net/smtp"
	"time"

	// Third-Party Dependencies

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// Database Driver
	_ "github.com/lib/pq"
)

func main() {
	// Open database connection.
	db, err := gorm.Open(
		viper.GetString("database.type"),
		viper.GetString("database.url"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(viper.GetBool("database.debug"))
	if err != nil {
		log.Fatal(err)
	}

	// Get smtp settings
	auth := smtp.PlainAuth(
		viper.GetString("smtp.identity"),
		viper.GetString("smtp.username"),
		viper.GetString("smtp.password"),
		viper.GetString("host"),
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
				err := SendEmail(email, auth, viper.GetString("smtp.address"))
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
