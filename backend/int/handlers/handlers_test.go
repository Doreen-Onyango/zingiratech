package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		template string
	}{
		{name: "NotFoundHandler", expected: http.StatusNotFound, template: "404.page.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
			resp := httptest.NewRecorder()
			NotFoundHandler(resp, req)
			if resp.Code != tt.expected {
				t.Errorf("expected status %v, got %v", tt.expected, resp.Code)
			}
		})
	}
}

func TestUnauthorizedHandler(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		template string
	}{
		{name: "UnauthorizedHandler", expected: http.StatusNotFound, template: "401.page.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/unauthorized", nil)
			resp := httptest.NewRecorder()
			UnauthorizedHandler(resp, req)
			if resp.Code != tt.expected {
				t.Errorf("expected status %v, got %v", tt.expected, resp.Code)
			}
		})
	}
}

func TestForbiddenHandler(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		template string
	}{
		{name: "ForbiddenHandler", expected: http.StatusForbidden, template: "403.page.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/forbidden", nil)
			resp := httptest.NewRecorder()
			ForbiddenHandler(resp, req)
			if resp.Code != tt.expected {
				t.Errorf("expected status %v, got %v", tt.expected, resp.Code)
			}
		})
	}
}

func TestInternalServerHandler(t *testing.T) {
	tests := []struct {
		name     string
		expected int
		template string
	}{
		{name: "InternalServerHandler", expected: http.StatusInternalServerError, template: "500.page.html"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/error", nil)
			resp := httptest.NewRecorder()
			InternalServerHandler(resp, req)
			if resp.Code != tt.expected {
				t.Errorf("expected status %v, got %v", tt.expected, resp.Code)
			}
		})
	}
}
