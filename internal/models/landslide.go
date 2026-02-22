package models

import "time"

// Landslide represents the 'landslides' table
type Landslide struct {
	LandslideID   int             `json:"landslide_id"`
	LandslideDate time.Time       `json:"landslide_date"`
	Latitude      float64         `json:"latitude"`
	Longitude     float64         `json:"longitude"`
	ImagePath     string          `json:"image_path"`
}
