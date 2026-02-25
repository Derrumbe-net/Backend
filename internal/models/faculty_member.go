package models

// FacultyMember represents the 'faculty_members' table
type FacultyMember struct {
	FacultyMemberID int    `json:"faculty_member_id"`
	Name            string `json:"name"`
	Role            string `json:"role"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Extension       string `json:"extension"`
	LinkedinURL     string `json:"linkedin_url"`
	ProfileImagePath string `json:"profile_image_path"`
}
