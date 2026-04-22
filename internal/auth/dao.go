package auth

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Derrumbe-net/Backend/internal/models"
)

type AdminDAO struct {
	DB *sql.DB
}

func NewAdminDAO(db *sql.DB) *AdminDAO {
	return &AdminDAO{DB: db}
}

// CreateAdmin creates a new admin and a verification token
func (dao *AdminDAO) CreateAdmin(email, hashedPassword, name string) (int64, string, error) {
	tx, err := dao.DB.Begin()
	if err != nil {
		return 0, "", err
	}
	defer tx.Rollback()

	// Insert into admins table
	res, err := tx.Exec("INSERT INTO admins (email, password, name, is_email_verified, is_authorized) VALUES (?, ?, ?, 0, 0)",
		email, hashedPassword, name)
	if err != nil {
		return 0, "", err
	}

	adminID, err := res.LastInsertId()
	if err != nil {
		return 0, "", err
	}

	// Generate verification token
	token := fmt.Sprintf("%x", time.Now().UnixNano()) // Simple token for now, could be more robust
	expiresAt := time.Now().Add(24 * time.Hour)

	_, err = tx.Exec("INSERT INTO admin_tokens (admin_id, verification_token, token_expires_at) VALUES (?, ?, ?)",
		adminID, token, expiresAt)
	if err != nil {
		return 0, "", err
	}

	if err := tx.Commit(); err != nil {
		return 0, "", err
	}

	return adminID, token, nil
}

func (dao *AdminDAO) GetAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	query := "SELECT admin_id, email, password, name, is_authorized, is_email_verified FROM admins WHERE email = ?"
	err := dao.DB.QueryRow(query, email).Scan(&admin.AdminID, &admin.Email, &admin.Password, &admin.Name, &admin.IsAuthorized, &admin.IsEmailVerified)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (dao *AdminDAO) GetAdminByID(id int) (*models.Admin, error) {
	var admin models.Admin
	query := "SELECT admin_id, email, password, name, is_authorized, is_email_verified FROM admins WHERE admin_id = ?"
	err := dao.DB.QueryRow(query, id).Scan(&admin.AdminID, &admin.Email, &admin.Password, &admin.Name, &admin.IsAuthorized, &admin.IsEmailVerified)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (dao *AdminDAO) GetAdminByToken(token string) (*models.Admin, error) {
	var admin models.Admin
	query := `
		SELECT a.admin_id, a.email, a.name, a.is_authorized, a.is_email_verified 
		FROM admins a 
		JOIN admin_tokens t ON a.admin_id = t.admin_id 
		WHERE t.verification_token = ? AND t.token_expires_at > NOW()`
	err := dao.DB.QueryRow(query, token).Scan(&admin.AdminID, &admin.Email, &admin.Name, &admin.IsAuthorized, &admin.IsEmailVerified)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (dao *AdminDAO) VerifyEmail(adminID int) error {
	tx, err := dao.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE admins SET is_email_verified = 1 WHERE admin_id = ?", adminID)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM admin_tokens WHERE admin_id = ?", adminID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (dao *AdminDAO) GetAllAdmins() ([]models.Admin, error) {
	rows, err := dao.DB.Query("SELECT admin_id, email, name, is_authorized, is_email_verified FROM admins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []models.Admin
	for rows.Next() {
		var a models.Admin
		if err := rows.Scan(&a.AdminID, &a.Email, &a.Name, &a.IsAuthorized, &a.IsEmailVerified); err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}
	return admins, nil
}

func (dao *AdminDAO) UpdateAuthorization(id int, isAuthorized bool) error {
	_, err := dao.DB.Exec("UPDATE admins SET is_authorized = ? WHERE admin_id = ?", isAuthorized, id)
	return err
}

func (dao *AdminDAO) DeleteAdmin(id int) error {
	_, err := dao.DB.Exec("DELETE FROM admins WHERE admin_id = ?", id)
	return err
}
