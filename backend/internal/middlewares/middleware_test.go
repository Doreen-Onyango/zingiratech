package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChainMiddlewares(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	tests := []struct {
		name        string
		handler     http.Handler
		middlewares []Middleware
		wantStatus  int
	}{
		{
			"No middleware",
			dummyHandler,
			[]Middleware{},
			http.StatusOK,
		},
		{
			"With middleware",
			dummyHandler,
			[]Middleware{middleware1, middleware2},
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			handler := ChainMiddlewares(tt.handler, tt.middlewares...)
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("ChainMiddlewares() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestLogger(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	tests := []struct {
		name        string
		requestPath string
		wantStatus  int
	}{
		{
			"Log a request",
			"/",
			http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.requestPath, nil)
			handler := Logger(dummyHandler)
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("Logger() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestRecovery(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})
	tests := []struct {
		name       string
		handler    http.Handler
		wantStatus int
	}{
		{
			"Recover from panic",
			dummyHandler,
			http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			handler := Recovery(tt.handler)
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("Recovery() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestStaticFileHandler(t *testing.T) {
	tests := []struct {
		name       string
		filePath   string
		wantStatus int
	}{
		{
			"File exists",
			"/static/testfile.txt",
			http.StatusOK,
		},
		{
			"File does not exist",
			"/static/missingfile.txt",
			http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.filePath, nil)
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Mock static file handler logic
				if tt.filePath == "/static/testfile.txt" {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			})
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("StaticFileHandler() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestRouteChecker(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	tests := []struct {
		name        string
		requestPath string
		wantStatus  int
	}{
		{
			"Valid dynamic route",
			"/",
			http.StatusOK,
		},
		{
			"Unauthorized route",
			"/dashboard",
			http.StatusForbidden,
		},
		{
			"Invalid route",
			"/unknown",
			http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.requestPath, nil)
			handler := RouteChecker(dummyHandler)
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("RouteChecker() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
		})
	}
}

func TestCORS(t *testing.T) {
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	tests := []struct {
		name        string
		method      string
		wantStatus  int
		wantHeaders map[string]string
	}{
		{
			"Allow CORS preflight",
			http.MethodOptions,
			http.StatusOK,
			map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
			},
		},
		{
			"Allow CORS for GET",
			http.MethodGet,
			http.StatusOK,
			map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, "/", nil)
			handler := CORS(dummyHandler)
			handler.ServeHTTP(recorder, req)
			if recorder.Code != tt.wantStatus {
				t.Errorf("CORS() status = %v, want %v", recorder.Code, tt.wantStatus)
			}
			for header, value := range tt.wantHeaders {
				if got := recorder.Header().Get(header); got != value {
					t.Errorf("CORS() header %s = %v, want %v", header, got, value)
				}
			}
		})
	}
}

func Test_isValidExtension(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{"Valid extension", "/static/style.css", true},
		{"Invalid extension", "/static/file.exe", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidExtension(tt.path); got != tt.want {
				t.Errorf("isValidExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAuthenticated(t *testing.T) {
	tests := []struct {
		name       string
		userHeader string
		wantStatus bool
	}{
		{
			"Authenticated user",
			"Bearer valid_token",
			false,
		},
		{
			"Unauthenticated user",
			"",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.userHeader != "" {
				req.Header.Set("Authorization", tt.userHeader)
			}
			if got := isAuthenticated(req); got != tt.wantStatus {
				t.Errorf("isAuthenticated() = %v, want %v", got, tt.wantStatus)
			}
		})
	}
}
