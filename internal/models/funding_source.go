package models

import "time"

// FundingSource represents the 'funding_sources' table
type FundingSource struct {
	FundingID    int       `json:"funding_id"`
	Name         string    `json:"name"`
	WebsiteURL   string    `json:"website_url"`
	ImagePath    string    `json:"image_path"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at"`
}
