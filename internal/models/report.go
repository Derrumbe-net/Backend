package models

import "time"

// Report represents the 'reports' table
type Report struct {
	ReportID    int       `json:"report_id"`
	LandslideID int       `json:"landslide_id"` // Required
	ReportedAt  time.Time `json:"reported_at"`  // Required
	City        string    `json:"city"`         // Required

	// Pointers to handle optional/NULL states
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	PhysicalAddress *string  `json:"physical_address"`
	ReporterName    *string  `json:"reporter_name"`
	ReporterPhone   *string  `json:"reporter_phone"`
	ReporterEmail   *string  `json:"reporter_email"`
	Description     *string  `json:"description"`
	ImagePath       *string  `json:"image_path"`
	IsValidated     *bool    `json:"is_validated"`
}
