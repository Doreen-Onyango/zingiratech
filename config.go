package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/int/auth"
	"github.com/Doreen-Onyango/zingiratech/backend/int/middlewares"
	"github.com/Doreen-Onyango/zingiratech/backend/int/routes"
	"github.com/Doreen-Onyango/zingiratech/backend/int/utils"
)

// Config initializes the application configuration and returns a configured http.Handler.
func Config(authService *auth.AuthService) (http.Handler, error) {
	// Load templates
	if err := utils.LoadTemplates(); err != nil {
		return nil, fmt.Errorf("error loading templates: %w", err)
	}
	log.Println("Templates loaded successfully.")

	// Initialize routes
	mux := http.NewServeMux()
	if err := routes.InitRoutes(mux); err != nil {
		return nil, fmt.Errorf("error initializing routes: %w", err)
	}
	log.Println("Routes initialized successfully.")

	// Wrap the routes with middleware
	wrappedMux := middlewares.RouteChecker(mux)
	log.Println("Middleware applied successfully.")

	log.Println("HTTP server configured successfully.")
	return wrappedMux, nil
}
