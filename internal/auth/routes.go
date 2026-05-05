package auth

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *AdminHandler) {
	protected := func(h http.HandlerFunc) http.Handler {
		return JWTMiddleware(h)
	}

	// Admin Public Routes
	mux.HandleFunc("POST /admins/signup", handler.SignUpAdmin)
	mux.HandleFunc("POST /admins/login", handler.LoginAdmin)
	mux.HandleFunc("GET /admins/verify", handler.VerifyEmail)
	mux.HandleFunc("POST /admins/password-reset/request", handler.RequestPasswordReset)
	mux.HandleFunc("POST /admins/password-reset/confirm", handler.ConfirmPasswordReset)

	// Admin Protected Routes (Require JWT)
	mux.Handle("GET /admins", protected(handler.GetAllAdmins))
	mux.Handle("GET /admins/{id}", protected(handler.GetAdmin))
	mux.Handle("PUT /admins/{id}/isAuthorized", protected(handler.UpdateAuthorization))
	mux.Handle("DELETE /admins/{id}", protected(handler.DeleteAdmin))
}
