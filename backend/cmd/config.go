package main

import (
	"fmt"
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/middlewares"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/routes"
	"github.com/Doreen-Onyango/zingiratech/backend/internal/utils"
)

// Config initializes the application configuration and returns a configured http.Handler.
func Config() (http.Handler, error) {
	if err := utils.LoadTemplates(); err != nil {
		return nil, fmt.Errorf("error loading templates: %w", err)
	}

	mux := http.NewServeMux()
	routes.InitRoutes(mux)

	wrappedMux := middlewares.RouteChecker(mux)

	fmt.Println("HTTP server configured successfully")
	return wrappedMux, nil
}
