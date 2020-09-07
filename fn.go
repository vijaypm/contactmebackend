package contactmebackend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/conf"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Config struct defines the configs we need
type config struct {
	FromName       string `conf:"FROM_NAME"`
	FromEmail      string `conf:"FROM_EMAIL"`
	ToName         string `conf:"TO_NAME"`
	ToEmail        string `conf:"TO_EMAIL"`
	SendgridAPIKey string `conf:"SENDGRID_API_KEY"`
}

// Email struct defines the data coming via REST API
type email struct {
	EmailAddress string `json:"email"`
	Name         string `json:"name"`
	Message      string `json:"message"`
}

// SendEmail function sents email with the data coming from REST API
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var c config
	conf.Load(&c)
	fmt.Println("From: ", c.FromEmail)
	fmt.Println("To: ", c.ToEmail)

	var e email
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}
	// Prepare the email content
	from := mail.NewEmail(c.FromName, c.FromEmail)
	subject := "Contacted via SendEmail function"
	to := mail.NewEmail(c.ToName, c.ToEmail)
	plainTextContent := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", e.Name, e.EmailAddress, e.Message)
	htmlContent := fmt.Sprintf("Name: %s<br>Email: %s<br>Message: %s", e.Name, e.EmailAddress, e.Message)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// Generate a SendGrid Send Client
	client := sendgrid.NewSendClient(c.SendgridAPIKey)
	// Send the message
	response, err := client.Send(message)
	if response.StatusCode != http.StatusAccepted {
		if response.StatusCode != http.StatusOK {
			// Request failed
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something bad happened!"))
		}
	} else {
		// Request successful
		// Setting header to allow cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Success"))
	}
}
