package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	wrapper, err := Config()
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
