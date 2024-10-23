package controller

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/Manikandan-Parasuraman/smtp-mail-sender/models"
)

// Send email using Google SMTP
func SendMail(input models.MailInput) error {
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USER")
    smtpPass := os.Getenv("SMTP_PASS")

    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

    // Load and parse HTML template
    t, err := template.ParseFiles("templates/email.html")
    if err != nil {
        return fmt.Errorf("error parsing template: %v", err)
    }

    // Prepare email content
    var body bytes.Buffer
    err = t.Execute(&body, input)
    if err != nil {
        return fmt.Errorf("error executing template: %v", err)
    }

    // Email headers and body
    from := smtpUser
    to := []string{input.Email}
    subject := "Subject: Emergency Alert - Service Outage\n"
    mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    
	// Convert the body to string
    bodyStr := body.String()

    // Concatenate headers and body
    msg := subject + mime + bodyStr    
	
	
	message := []byte(subject + msg)

    addr := smtpHost + ":" + smtpPort
    err = smtp.SendMail(addr, auth, from, to, message)
    if err != nil {
        return fmt.Errorf("error sending mail: %v", err)
    }
    return nil
}
