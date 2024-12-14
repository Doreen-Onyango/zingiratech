package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/handlers"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/utils"
)

var validRoutes = map[string]bool{
	"/":      true,
	"/about": true,
}

// Supported static file extensions
var validExtensions = []string{".css", ".js", ".jpg", ".png", ".gif", ".svg"}

// InitRoutes initializes all application routes.
func InitRoutes(mux *http.ServeMux) {
	dir, err := utils.GetProjectRootPath("frontend", "static")
	if err != nil {
		log.Fatalf("Error finding static directory: %v", err)
	}
	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
}

// RouteChecker acts as middleware to validate and handle requests.
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

		if _, ok := validRoutes[r.URL.Path]; !ok {
			handlers.NotFoundHandler(w, r)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v (Request: %s)", err, r.URL.Path)
				handlers.InternalServerHandler(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// isValidExtension checks if a file path has a valid extension.
func isValidExtension(path string) bool {
	for _, ext := range validExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
