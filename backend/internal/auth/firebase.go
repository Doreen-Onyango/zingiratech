package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// FirebaseCredentials represents the structure of Firebase service account credentials
type FirebaseCredentials struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// AuthService handles Firebase authentication
type AuthService struct {
	client *auth.Client
}

// NewAuthService initializes the AuthService
func NewAuthService() (*AuthService, error) {
	ctx := context.Background()

	// Try to get credentials from environment variable first
	credentials := os.Getenv("FIREBASE_CREDENTIALS_JSON")
	var opt option.ClientOption

	if credentials != "" {
		// Use credentials from environment variable
		credBytes := []byte(credentials)
		opt = option.WithCredentialsJSON(credBytes)
	} else {
		// Fall back to constructing credentials from individual environment variables
		creds := FirebaseCredentials{
			Type:                    "service_account",
			ProjectID:               os.Getenv("FIREBASE_PROJECT_ID"),
			PrivateKeyID:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
			PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
			ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
			ClientID:                os.Getenv("FIREBASE_CLIENT_ID"),
			AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
			TokenURI:                "https://oauth2.googleapis.com/token",
			AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
			ClientX509CertURL:       os.Getenv("FIREBASE_CLIENT_CERT_URL"),
		}

		// Convert credentials struct to JSON
		credJSON, err := json.Marshal(creds)
		if err != nil {
			return nil, fmt.Errorf("error marshaling credentials: %v", err)
		}

		opt = option.WithCredentialsJSON(credJSON)
	}

	// Initialize Firebase app with credentials
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	// Create the Auth client
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	log.Printf("Firebase Auth initialized successfully")
	return &AuthService{client: authClient}, nil
}

// VerifyIDToken verifies the Firebase ID token
func (as *AuthService) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := as.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}

// ExtractClaims extracts custom claims from a verified token
func (as *AuthService) ExtractClaims(token *auth.Token) map[string]interface{} {
	return token.Claims
}
