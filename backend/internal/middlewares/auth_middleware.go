package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authService, err := auth.NewAuthService()
		if err != nil {
			log.Fatalf("Failed to initialize Firebase: %v", err)
		}
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extract the token
		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		if idToken == authHeader {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Verify the token
		token, err := authService.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add the verified user info to the request context
		ctx := context.WithValue(r.Context(), "user", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
