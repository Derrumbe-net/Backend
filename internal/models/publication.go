package models

import "time"

// Publication represents the 'publications' table
type Publication struct {
	PublicationID  int       `json:"publication_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	PublicationURL string    `json:"publication_url"`
	ImagePath      string    `json:"image_path"`
	PublishedDate  time.Time `json:"published_date"`
}
