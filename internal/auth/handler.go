package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Derrumbe-net/Backend/internal/models"
)

type AdminHandler struct {
	Service *AuthService
}

func NewAdminHandler(service *AuthService) *AdminHandler {
	return &AdminHandler{Service: service}
}

// --- DTOs ---

type SignUpAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAuthorizationRequest struct {
	IsAuthorized bool `json:"is_authorized"`
}

type AdminResponse struct {
	AdminID         int    `json:"admin_id"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	IsAuthorized    bool   `json:"is_authorized"`
	IsEmailVerified bool   `json:"is_email_verified"`
}

func toAdminResponse(a *models.Admin) AdminResponse {
	return AdminResponse{
		AdminID:         a.AdminID,
		Email:           a.Email,
		Name:            a.Name,
		IsAuthorized:    a.IsAuthorized,
		IsEmailVerified: a.IsEmailVerified,
	}
}

func toAdminResponses(admins []models.Admin) []AdminResponse {
	var responses []AdminResponse
	for _, a := range admins {
		responses = append(responses, toAdminResponse(&a))
	}
	return responses
}

// --- Handlers ---

func (h *AdminHandler) SignUpAdmin(w http.ResponseWriter, r *http.Request) {
	var req SignUpAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	id, token, err := h.Service.SignUp(req.Email, req.Password, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"admin_id":           id,
		"verification_token": token,
		"message":            "Admin created. Please verify your email.",
	})
}

func (h *AdminHandler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req LoginAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AdminHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	if err := h.Service.VerifyEmail(token); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email verified successfully"})
}

func (h *AdminHandler) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := h.Service.GetAllAdmins()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(toAdminResponses(admins))
}

func (h *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	admin, err := h.Service.GetAdmin(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if admin == nil {
		http.Error(w, "Admin not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(toAdminResponse(admin))
}

func (h *AdminHandler) UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateAuthorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateAuthorization(id, req.IsAuthorized); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	// Check if requester is Super Admin
	requesterEmail, ok := r.Context().Value(AdminEmailKey).(string)
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")

	if !ok || superAdminEmail == "" || strings.ToLower(requesterEmail) != strings.ToLower(superAdminEmail) {
		http.Error(w, "Unauthorized: Only Super Admin can delete admins", http.StatusForbidden)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.DeleteAdmin(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
