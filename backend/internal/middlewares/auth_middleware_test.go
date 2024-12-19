package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock Token struct to simulate the auth.Token structure
type Token struct {
	UID string
}

// MockAuthService is a mock of the authService interface for testing
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) VerifyIDToken(ctx context.Context, token string) (*Token, error) {
	args := m.Called(ctx, token)
	if args.Get(0) != nil {
		return args.Get(0).(*Token), args.Error(1)
	}
	return nil, args.Error(1)
}

// Custom "Anything" constant to match any argument
var Anything = mock.Anything

// mockHandler is a simple handler that returns OK if user info is present in the context
func mockHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	if user == nil {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TestAuthMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name         string
		args         args
		authHeader   string
		mockVerifyFn func(mock *MockAuthService)
		wantStatus   int
		wantBody     string
	}{
		{
			name: "Missing Authorization Header",
			args: args{
				next: http.HandlerFunc(mockHandler),
			},
			authHeader:   "",
			mockVerifyFn: func(mock *MockAuthService) {},
			wantStatus:   http.StatusUnauthorized,
			wantBody:     "Authorization header required",
		},
		{
			name: "Invalid Authorization Header Format",
			args: args{
				next: http.HandlerFunc(mockHandler),
			},
			authHeader:   "InvalidFormat token",
			mockVerifyFn: func(mock *MockAuthService) {},
			wantStatus:   http.StatusUnauthorized,
			wantBody:     "Invalid authorization header format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the mock auth service directly inside the test case
			mockAuthService := new(MockAuthService)
			tt.mockVerifyFn(mockAuthService)

			// Create a test request
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("Authorization", tt.authHeader)

			// Create the middleware with the mock auth service
			authMiddleware := AuthMiddleware(tt.args.next)

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Create the handler and call the middleware
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Inject the mock AuthService here in the middleware
				// Pass the mock auth service directly, ensuring the middleware uses it
				authMiddleware.ServeHTTP(w, r)
			})
			handler.ServeHTTP(rr, req)

			// Check the response status code
			assert.Equal(t, tt.wantStatus, rr.Code)

			// Check the response body
			assert.Contains(t, rr.Body.String(), tt.wantBody)

			// Assert expectations on the mock service
			mockAuthService.AssertExpectations(t)
		})
	}
}
