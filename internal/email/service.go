package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

// We embed the email templates into the binary
// The line below is not a comment, and is necessary for embeds
//go:embed email_templates/*.html
var templateFS embed.FS

type EmailService struct {
	dialer      *gomail.Dialer
	senderEmail string
}

func NewEmailService(host string, port int, user string, pass string) *EmailService {
	return &EmailService{
		dialer:      gomail.NewDialer(host, port, user, pass),
		senderEmail: user,
	}
}

func (s *EmailService) send(to string, subject string, templateName string, data interface{}) error {
	// The template path is relative to our embed root
	fullPath := "email_templates/" + templateName

	templateFile, err := template.ParseFS(templateFS, fullPath)
	if err != nil {
		return fmt.Errorf("failed to find embedded template %s: %w", templateName, err)
	}

	var body bytes.Buffer
	if err = templateFile.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to create message body using template: %w", err)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", s.senderEmail)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body.String())

	if err := s.dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email to %s: %w", to, err)
	}

	return nil
}
