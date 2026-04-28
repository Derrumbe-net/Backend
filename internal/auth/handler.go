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
}

type LoginAdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateAuthorizationRequest struct {
	IsAuthorized bool `json:"isAuthorized"`
}

type AdminResponse struct {
	AdminID         int    `json:"admin_id"`
	Email           string `json:"email"`
	IsAuthorized    bool   `json:"isAuthorized"`
	IsEmailVerified bool   `json:"is_email_verified"`
}

func toAdminResponse(a *models.Admin) AdminResponse {
	return AdminResponse{
		AdminID:         a.AdminID,
		Email:           a.Email,
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email and password are required"})
		return
	}

	id, token, err := h.Service.SignUp(req.Email, req.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (h *AdminHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token is required"})
		return
	}

	if err := h.Service.VerifyEmail(token); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email verified successfully"})
}

func (h *AdminHandler) GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := h.Service.GetAllAdmins()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toAdminResponses(admins))
}

func (h *AdminHandler) GetAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	admin, err := h.Service.GetAdmin(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if admin == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Admin not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toAdminResponse(admin))
}

func (h *AdminHandler) UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	var req UpdateAuthorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if err := h.Service.UpdateAuthorization(id, req.IsAuthorized); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	// Check if requester is Super Admin
	requesterEmail, ok := r.Context().Value(AdminEmailKey).(string)
	superAdminEmail := os.Getenv("SUPERADMIN_EMAIL")

	if !ok || superAdminEmail == "" || strings.ToLower(requesterEmail) != strings.ToLower(superAdminEmail) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized: Only Super Admin can delete admins"})
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	if err := h.Service.DeleteAdmin(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
