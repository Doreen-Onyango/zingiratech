package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/auth"
)

func main() {
	// Initialize Firebase Auth Service
	authService, err := auth.NewAuthService()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	wrapper, err := Config(authService)
	if err != nil {
		log.Fatalf("Configuration failed: %v", err)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrapper,
	}

	fmt.Println("Server running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server encountered an error: %v", err)
	}
}
