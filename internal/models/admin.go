package models

// Admin represents the 'admins' table
type Admin struct {
	AdminID         int       `json:"admin_id"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Name            string    `json:"name"`
	IsAuthorized    bool      `json:"is_authorized"`
	IsEmailVerified bool      `json:"is_email_verified"`
}
