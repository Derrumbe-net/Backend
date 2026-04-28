package main

import (
    "log"
    "net/http"
    "os"
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
    "github.com/rs/cors"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment variables")
    }

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

    mux.HandleFunc("GET /radar/frames", radar.GetFramesHandler(cacheDir))
    mux.Handle("GET /radar/images/", http.StripPrefix("/radar/images/", http.FileServer(http.Dir(cacheDir))))

    auth.RegisterRoutes(mux, adminHandler)
    landslide.RegisterRoutes(mux, landslideHandler)
    content.RegisterRoutes(mux, contentHandler)
    station.RegisterRoutes(mux, stationHandler)
    report.RegisterRoutes(mux, reportHandler)

    // Setup CORS configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"}, // For development; specify your frontend URL in production
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    })

    handler := c.Handler(mux)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    addr := "0.0.0.0:" + port
    log.Printf("Backend API starting on %s...", addr)

    // 4. Use the 'handler' instead of 'mux'
    if err := http.ListenAndServe(addr, handler); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
