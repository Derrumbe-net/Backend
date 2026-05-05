package models 

import "time"

// AdminToken represents the 'admin_tokens' table
type AdminToken struct {
	TokenID           int        `json:"token_id"`
	AdminID           int        `json:"admin_id"`
	VerificationToken string     `json:"verification_token"`
	ResetToken        *string    `json:"reset_token"`
	ResetTokenExpiresAt *time.Time `json:"reset_token_expires_at"`
	TokenExpiresAt    time.Time  `json:"token_expires_at"`
}
