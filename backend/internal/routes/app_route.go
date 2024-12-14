package routes

import (
	"log"
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/handlers"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/utils"
)

// InitRoutes initializes all application routes and serves static files.
func InitRoutes(mux *http.ServeMux) error {
	dir, err := utils.GetProjectRootPath("frontend", "static")
	if err != nil {
		return err
	}

	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	registerRoutes(mux)

	log.Println("Routes initialized successfully")
	return nil
}

// registerRoutes sets up route handlers for the application.
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Contact page is under construction.", http.StatusNotImplemented)
	})
}
