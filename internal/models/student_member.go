package models

// StudentMember represents the 'student_members' table
type StudentMember struct {
	StudentMemberID int    `json:"student_member_id"`
	Name            string `json:"name"`
}
