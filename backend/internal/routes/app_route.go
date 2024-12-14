package routes

import (
	"log"
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/handlers"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/middlewares"
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
	// Public routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/signup", handlers.SignupHandler)

	// Protected routes
	protectedRoutes := http.NewServeMux()
	protectedRoutes.HandleFunc("/dashboard", handlers.DashboardHandler)
	protectedRoutes.HandleFunc("/schedule-pickup", handlers.SchedulePickupHandler)
	
	// Apply auth middleware to protected routes
	mux.Handle("/dashboard/", middlewares.AuthMiddleware(protectedRoutes))
	mux.Handle("/api/", middlewares.AuthMiddleware(protectedRoutes))
}
