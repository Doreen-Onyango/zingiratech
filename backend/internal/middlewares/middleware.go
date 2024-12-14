package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/auth"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/handlers"
)

// Route configuration for dynamic and static routes
var routes = map[string]struct {
	RequiresAuth bool
}{
	"/":          {RequiresAuth: false},
	"/about":     {RequiresAuth: false},
	"/signup":    {RequiresAuth: false},
	"/login":     {RequiresAuth: false},
	"/dashboard": {RequiresAuth: true},
}

// Supported static file extensions
var validExtensions = []string{".css", ".js", ".jpg", ".png", ".gif", ".svg"}

// Middleware type for chaining
type Middleware func(http.Handler) http.Handler

// Middleware chaining function
func ChainMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
}

// Logger middleware for request logging
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware to handle panics gracefully
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v (Request: %s %s)", err, r.Method, r.URL.Path)
				handlers.InternalServerHandler(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// StaticFileHandler serves static files
func StaticFileHandler(staticDir string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))
	return http.StripPrefix("/static/", fs)
}

// RouteChecker middleware validates dynamic and static routes
func RouteChecker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static/") {
			if !isValidExtension(r.URL.Path) {
				handlers.ForbiddenHandler(w, r)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		if route, exists := routes[r.URL.Path]; exists {
			if route.RequiresAuth && !isAuthenticated(r) {
				handlers.ForbiddenHandler(w, r)
				return
			}
		} else {
			handlers.NotFoundHandler(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// CORS middleware to handle cross-origin requests
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Helper function to check valid file extensions
func isValidExtension(path string) bool {
	for _, ext := range validExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}

// Helper function to check if a user is authenticated
func isAuthenticated(r *http.Request) bool {
	authService, err := auth.NewAuthService()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return false
	}

	idToken := strings.TrimPrefix(authHeader, "Bearer ")
	_, err = authService.VerifyIDToken(r.Context(), idToken)
	return err == nil
}
