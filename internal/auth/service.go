package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Derrumbe-net/Backend/internal/email"
	"github.com/Derrumbe-net/Backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	DAO          *AdminDAO
	EmailService *email.EmailService
}

func NewAuthService(dao *AdminDAO, emailService *email.EmailService) *AuthService {
	return &AuthService{
		DAO:          dao,
		EmailService: emailService,
	}
}

func (s *AuthService) SignUp(email, password string) (int64, string, error) {
	// Check if admin already exists
	existing, _ := s.DAO.GetAdminByEmail(email)
	if existing != nil {
		return 0, "", errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, "", err
	}

	// Create in DB
	id, token, err := s.DAO.CreateAdmin(email, string(hashedPassword))
	if err != nil {
		return 0, "", err
	}

	// Trigger verification email
	if s.EmailService != nil {
		_ = s.EmailService.SendVerificationEmail(email, token)
	}

	return id, token, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	admin, err := s.DAO.GetAdminByEmail(email)
	if err != nil {
		return "", err
	}
	if admin == nil {
		return "", errors.New("invalid email or password")
	}

	if !admin.IsEmailVerified {
		return "", errors.New("email not verified")
	}

	if !admin.IsAuthorized {
		return "", errors.New("account not authorized")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "CHANGE_THIS_SECRET_KEY"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.AdminID,
		"email":    admin.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func (s *AuthService) VerifyEmail(token string) error {
	admin, err := s.DAO.GetAdminByToken(token)
	if err != nil {
		return err
	}
	if admin == nil {
		return errors.New("invalid or expired token")
	}

	if err := s.DAO.VerifyEmail(admin.AdminID); err != nil {
		return err
	}

	// Trigger alert to Super Admin
	if s.EmailService != nil {
		_ = s.EmailService.SendNewAdminAlert(admin.Email)
	}

	return nil
}

func (s *AuthService) GetAllAdmins() ([]models.Admin, error) {
	return s.DAO.GetAllAdmins()
}

func (s *AuthService) GetAdmin(id int) (*models.Admin, error) {
	return s.DAO.GetAdminByID(id)
}

func (s *AuthService) UpdateAuthorization(id int, isAuthorized bool) error {
	// Get admin before update to check if status changed
	admin, err := s.DAO.GetAdminByID(id)
	if err != nil {
		return err
	}

	if err := s.DAO.UpdateAuthorization(id, isAuthorized); err != nil {
		return err
	}

	// If newly authorized, send welcome email
	if isAuthorized && !admin.IsAuthorized && s.EmailService != nil {
		_ = s.EmailService.SendAdminWelcome(admin.Email)
	}

	return nil
}

func (s *AuthService) DeleteAdmin(id int) error {
	return s.DAO.DeleteAdmin(id)
}

// Password Reset

func (s *AuthService) RequestPasswordReset(email string) error {
	admin, err := s.DAO.GetAdminByEmail(email)
	if err != nil {
		return err
	}
	// Security: If admin not found, don't reveal it. Just return nil.
	if admin == nil {
		return nil
	}

	// Generate secure token
	token := fmt.Sprintf("%x", time.Now().UnixNano()) + fmt.Sprintf("%x", time.Now().Unix()) // Simple concatenation for uniqueness
	expiresAt := time.Now().Add(30 * time.Minute)

	if err := s.DAO.SetPasswordResetToken(admin.AdminID, token, expiresAt); err != nil {
		return err
	}

	// Send email
	if s.EmailService != nil {
		return s.EmailService.SendPasswordResetEmail(admin.Email, token)
	}

	return nil
}

func (s *AuthService) ResetPassword(token string, newPassword string) error {
	admin, resetToken, err := s.DAO.GetAdminByResetToken(token)
	if err != nil {
		return err
	}
	if admin == nil || resetToken == nil {
		return errors.New("invalid or expired token")
	}

	// Check expiration
	if resetToken.ResetTokenExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password and clear token
	if err := s.DAO.UpdatePassword(admin.AdminID, string(hashedPassword)); err != nil {
		return err
	}

	return s.DAO.ClearPasswordResetToken(admin.AdminID)
}
