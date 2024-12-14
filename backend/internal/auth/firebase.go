package auth

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var firebaseAuth *auth.Client

func InitFirebase() error {
	ctx := context.Background()
	
	// Get the absolute path to the project root
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../..")
	
	// Construct the path to the service account key
	credentialsFile := filepath.Join(projectRoot, "config", "firebase", "serviceAccountKey.json")
	
	opt := option.WithCredentialsFile(credentialsFile)
	
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %v", err)
	}
	
	auth, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("error getting Auth client: %v", err)
	}
	
	firebaseAuth = auth
	log.Printf("Firebase Auth initialized successfully using credentials from: %s", credentialsFile)
	return nil
}

// VerifyIDToken verifies the Firebase ID token
func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := firebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
} 