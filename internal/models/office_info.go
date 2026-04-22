package models

import "time"

// OfficeInfo represents the 'office_info' table
type OfficeInfo struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	PhoneExt       string    `json:"phone_ext"`
	OfficeLocation string    `json:"office_location"`
	FacebookURL    string    `json:"facebook_url"`
	UpdatedAt      time.Time `json:"updated_at"`
}
