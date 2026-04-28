package email

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/Derrumbe-net/Backend/internal/models"
	"gopkg.in/gomail.v2"
)

// We embed the email templates into the binary
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
	fullPath := "email_templates/" + templateName + ".html"

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

// SendVerificationEmail triggers the 'email_verification' template
func (s *EmailService) SendVerificationEmail(to string, token string) error {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	data := struct {
		Email            string
		VerificationLink string
	}{
		Email:            to,
		VerificationLink: fmt.Sprintf("%s/api/admins/verify?token=%s", frontendURL, token),
	}

	return s.send(to, "Verify your DerrumbeNet email address", "email_verification", data)
}

// SendNewAdminAlert triggers the 'new_admin' template for the Super Admin
func (s *EmailService) SendNewAdminAlert(adminEmail string) error {
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")
	if superAdminEmail == "" {
		return fmt.Errorf("SUPERADMIN_EMAIL not configured")
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	data := struct {
		Email   string
		CMSLink string
	}{
		Email:   adminEmail,
		CMSLink: frontendURL + "/cms",
	}

	return s.send(superAdminEmail, "New Admin Signup Request", "new_admin", data)
}

// SendAdminWelcome triggers the 'admin_welcome' template once authorized
func (s *EmailService) SendAdminWelcome(to string) error {
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	data := struct {
		Email   string
		CMSLink string
	}{
		Email:   to,
		CMSLink: frontendURL + "/cms",
	}

	return s.send(to, "Your Admin Access Has Been Approved", "admin_welcome", data)
}

// SendReportAlert triggers the 'report_submitted' template for the Super Admin
func (s *EmailService) SendReportAlert(r *models.Report) error {
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")
	if superAdminEmail == "" {
		return fmt.Errorf("SUPERADMIN_EMAIL not configured")
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	// Safely unwrap ALL optional database fields
	safeLat := 0.0
	if r.Latitude != nil {
		safeLat = *r.Latitude
	}

	safeLng := 0.0
	if r.Longitude != nil {
		safeLng = *r.Longitude
	}

	safeDesc := "Not provided"
	if r.Description != nil {
		safeDesc = *r.Description
	}

	safeName := "Anonymous"
	if r.ReporterName != nil {
		safeName = *r.ReporterName
	}

	safePhone := "Not provided"
	if r.ReporterPhone != nil {
		safePhone = *r.ReporterPhone
	}

	safeEmail := "Not provided"
	if r.ReporterEmail != nil {
		safeEmail = *r.ReporterEmail
	}

	safeAddress := "Not provided"
	if r.PhysicalAddress != nil {
		safeAddress = *r.PhysicalAddress
	}

	data := struct {
		ID              int
		ReportedAt      string
		City            string
		Latitude        float64
		Longitude       float64
		PhysicalAddress string
		Description     string
		ReporterName    string
		ReporterPhone   string
		ReporterEmail   string
		CMSLink         string
	}{
		ID:              r.ReportID,
		ReportedAt:      r.ReportedAt.Format("Jan 02, 2006 15:04"),
		City:            r.City,
		Latitude:        safeLat,
		Longitude:       safeLng,
		PhysicalAddress: safeAddress,
		Description:     safeDesc,
		ReporterName:    safeName,
		ReporterPhone:   safePhone,
		ReporterEmail:   safeEmail,
		CMSLink:         frontendURL + "/cms",
	}

	subject := fmt.Sprintf("New Report Submitted (#%d)", r.ReportID)
	return s.send(superAdminEmail, subject, "report_submitted", data)
}
