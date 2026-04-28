package models

// FacultyMember represents the 'faculty_members' table
type FacultyMember struct {
	FacultyMemberID int    `json:"faculty_member_id"`
	Name            string `json:"name"`
	FacultyRole     string `json:"faculty_role"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Extension       string `json:"extension"`
	LinkedinURL     string `json:"linkedin_url"`
	ImagePath       string `json:"image_path"`
}
