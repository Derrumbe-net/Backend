package models

// Project represents the 'projects' table
type Project struct {
	ProjectID     int    `json:"project_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	StartYear     uint16 `json:"start_year"`
	EndYear       uint16 `json:"end_year"`
	ProjectStatus string `json:"project_status"`
	ImagePath     string `json:"image_path"`
}
