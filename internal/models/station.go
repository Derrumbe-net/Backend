package models

import "time"

// Station represents the 'stations' table
type Station struct {
	StationID               int       `json:"station_id"`
	Name                    string    `json:"name"`
	Depth                   string    `json:"depth"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	IsAvailable             bool      `json:"is_available"`
	StationInstallationDate time.Time `json:"station_installation_date"`

	// Optional fields converted to pointers to handle database NULLs
	LandUnit          *string  `json:"land_unit"`
	GeologicalUnit    *string  `json:"geological_unit"`
	Susceptibility    *string  `json:"susceptibility"`
	LandslideForecast *float64 `json:"landslide_forecast"`
	ImagePath         *string  `json:"image_path"`
	Elevation         *int     `json:"elevation"`
	Slope             *float64 `json:"slope"`
	Collaborator      *string  `json:"collaborator"`
}
