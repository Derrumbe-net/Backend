package main

import (
	"strconv"

	"github.com/Derrumbe-net/Backend/internal/auth"
	"github.com/Derrumbe-net/Backend/internal/content"
	"github.com/Derrumbe-net/Backend/internal/db"
	"github.com/Derrumbe-net/Backend/internal/email"
	"github.com/Derrumbe-net/Backend/internal/landslide"
	"github.com/Derrumbe-net/Backend/internal/radar"
	"github.com/Derrumbe-net/Backend/internal/report"
	"github.com/Derrumbe-net/Backend/internal/station"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Initialize Database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbPort == "" {
		dbPort = "3306"
	}

	database, err := db.New(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize Email Service
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpPort, _ := strconv.Atoi(smtpPortStr)

	var emailService *email.EmailService
	if smtpHost != "" && smtpUser != "" {
		emailService = email.NewEmailService(smtpHost, smtpPort, smtpUser, smtpPass)
		log.Println("Email Service initialized")
	} else {
		log.Println("Email Service NOT initialized (missing SMTP config)")
	}

	// Initialize DAOs, Services and Handlers
	adminDAO := auth.NewAdminDAO(database)
	authService := auth.NewAuthService(adminDAO, emailService)
	adminHandler := auth.NewAdminHandler(authService)

	landslideDAO := landslide.NewLandslideDAO(database)
	landslideService := landslide.NewLandslideService(landslideDAO)
	landslideHandler := landslide.NewLandslideHandler(landslideService)

	contentDAO := content.NewContentDAO(database)
	contentService := content.NewContentService(contentDAO)
	contentHandler := content.NewContentHandler(contentService)

	stationDAO := station.NewStationDAO(database)
	stationService := station.NewStationService(stationDAO)
	stationHandler := station.NewStationHandler(stationService)

	reportDAO := report.NewReportDAO(database)
	reportService := report.NewReportService(reportDAO, emailService)
	reportHandler := report.NewReportHandler(reportService)

	// Ensure upload directories exist
	cacheDir := "./radar_cache"
	os.MkdirAll(cacheDir, os.ModePerm)
	os.MkdirAll("./uploads/landslides", os.ModePerm)
	os.MkdirAll("./uploads/projects", os.ModePerm)
	os.MkdirAll("./uploads/publications", os.ModePerm)
	os.MkdirAll("./uploads/stations", os.ModePerm)
	os.MkdirAll("./uploads/reports", os.ModePerm)
	os.MkdirAll("./uploads/funding", os.ModePerm)
	os.MkdirAll("./uploads/faculty", os.ModePerm)

	go radar.StartWorker(cacheDir)

	mux := http.NewServeMux()

	// Radar Routes (Internal)
	mux.HandleFunc("GET /radar/frames", radar.GetFramesHandler(cacheDir))
	mux.Handle("GET /radar/images/", http.StripPrefix("/radar/images/", http.FileServer(http.Dir(cacheDir))))

	// Register Feature Routes
	auth.RegisterRoutes(mux, adminHandler)
	landslide.RegisterRoutes(mux, landslideHandler)
	content.RegisterRoutes(mux, contentHandler)
	station.RegisterRoutes(mux, stationHandler)
	report.RegisterRoutes(mux, reportHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Backend API starting on %s...", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
