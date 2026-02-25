package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/Derrumbe-net/Backend/internal/email" 
)

func main() {
	// Docker won't find the file.
	// It's okay, as the file is injected directly into the container.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Parse Environment Variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	testReceiver := "jose.rivera471@upr.edu"

	// Environment variables are loaded as strings, port is expected as int
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	log.Println("Initializing email service...")
	mailer := email.NewEmailService(smtpHost, smtpPort, smtpUser, smtpPass)

	log.Printf("Sending test emails to: %s\n", testReceiver)

	// Test: Admin Welcome Email
	adminData := email.AdminWelcomeData{
		Email:   testReceiver,
		CMSLink: "https://www.youtube.com/",
	}
	
	err = mailer.SendAdminWelcome(testReceiver, adminData)
	if err != nil {
		log.Printf("Failed to send Admin Welcome: %v\n", err)
	} else {
		log.Println("Admin Welcome email sent successfully!")
	}

	// Test: Email Verification Email
	emailVerificationData := email.EmailVerificationData{
		Email:   testReceiver,
		VerificationLink: "https://www.youtube.com/",
	}

	err = mailer.SendEmailVerification(testReceiver, emailVerificationData)
	if err != nil {
		log.Printf("Failed to send Email Verification: %v\n", err)
	} else {
		log.Println("Email Verification email sent successfully!")
	}

	// Test: New Admin Request Notification Email
	newAdminRequestNotificationData := email.NewAdminRequestNotificationData{
		Email:   testReceiver,
		CMSLink: "https://www.youtube.com/",
	}

	err = mailer.SendNewAdminRequestNotification(testReceiver, newAdminRequestNotificationData)
	if err != nil {
		log.Printf("Failed to send New Admin Request Notification: %v\n", err)
	} else {
		log.Println("New Admin Request Notificaiton email sent successfully!")
	}

	// Test: Report Submitted Email
	reportData := email.ReportSubmittedData{
		ID:              105,
		ReportedAt:      time.Now().Format("Jan 02, 2006 03:04 PM"),
		City:            "Mayagüez",
		Latitude:        18.2013,
		Longitude:       -67.1452,
		PhysicalAddress: "Carr 108, Km 2.3",
		Description:     "Large rockfall blocking the right lane.",
		ReporterName:    "Juan del Pueblo",
		ReporterPhone:   "787-555-5555",
		ReporterEmail:   "juan.pueblo@example.com",
		CMSLink:         "https://www.youtube.com/",
	}

	err = mailer.SendReportNotification(testReceiver, reportData)
	if err != nil {
		log.Printf("Failed to send Report Notification: %v\n", err)
	} else {
		log.Println("Report Notification email sent successfully!")
	}

	log.Println("Done!")
}
