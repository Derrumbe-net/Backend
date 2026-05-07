package models

type Municipality struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Stage          string `json:"stage"`
	StartYear      *int   `json:"start_year"`
	RenovationYear *int   `json:"renovation_year"`
}
