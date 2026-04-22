package auth

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	AdminIDKey    contextKey = "admin_id"
	AdminEmailKey contextKey = "admin_email"
)

// JWTMiddleware validates the JWT token in the Authorization header
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "CHANGE_THIS_SECRET_KEY" // Fallback matching PHP code for now
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		adminID, ok := claims["admin_id"].(float64)
		if !ok {
			http.Error(w, "Admin ID not found in token", http.StatusUnauthorized)
			return
		}

		email, _ := claims["email"].(string)

		// Add admin_id and email to request context
		ctx := context.WithValue(r.Context(), AdminIDKey, int(adminID))
		if email != "" {
			ctx = context.WithValue(ctx, AdminEmailKey, email)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
