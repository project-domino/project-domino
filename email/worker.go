package email

import (
	"log"

	"github.com/project-domino/project-domino/db"
	"github.com/project-domino/project-domino/models"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func worker(emails <-chan models.Email) {
	for e := range emails {
		// Create email
		from := mail.NewEmail("no-reply", "no-reply@notebox.org")
		subject := e.Subject
		to := mail.NewEmail(e.User.UserName, e.User.Email)
		content := mail.NewContent("text/plain", e.Body)
		m := mail.NewV3MailInit(from, subject, to, content)

		// Create sendgrid request
		request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
		request.Method = "POST"
		request.Body = mail.GetRequestBody(m)
		response, err := sendgrid.API(request)
		if err != nil {
			log.Printf("Email Error -\nresponse - %v\nerror - %v", response, err)
			markDropped(e)
		} else {
			markSent(e)
		}
	}
}

func markSent(e models.Email) {
	e.Sent = true
	e.Dropped = false

	db.DB.Save(&e)
}

func markDropped(e models.Email) {
	e.Sent = false
	e.Dropped = true

	db.DB.Save(&e)
}
