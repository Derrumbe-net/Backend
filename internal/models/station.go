package models

import "time"

// Station represents the 'stations' table
type Station struct {
	StationID               int       `json:"station_id"`
	Name                    string    `json:"name"`
	LandUnit                *string   `json:"land_unit"`
	GeologicalUnit          *string   `json:"geological_unit"`
	Susceptibility          *string   `json:"susceptibility"`
	Depth                   *string   `json:"depth"`
	LandslideForecast       *float64  `json:"landslide_forecast"`
	ImagePath               *string   `json:"image_url"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	Elevation               *int      `json:"elevation"`
	Slope                   *float64  `json:"slope"`
	IsAvailable             bool      `json:"is_available"`
	Collaborator            *string   `json:"collaborator"`
	StationInstallationDate *time.Time `json:"station_installation_date"`
	WC1Max                  *float64  `json:"wc1_max"`
	WC2Max                  *float64  `json:"wc2_max"`
	WC3Max                  *float64  `json:"wc3_max"`
	WC4Max                  *float64  `json:"wc4_max"`
}
