package auth

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// AuthService handles Firebase authentication
type AuthService struct {
	client *auth.Client
}

// NewAuthService initializes the AuthService
func NewAuthService() (*AuthService, error) {
	ctx := context.Background()

	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")

	credentialsFile := os.Getenv("FIREBASE_CREDENTIALS")
	if credentialsFile == "" {
		credentialsFile = filepath.Join(projectRoot, "config", "firebase", "serviceAccountKey.json")
	}

	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	log.Printf("Firebase Auth initialized successfully using credentials from: %s", credentialsFile)
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
