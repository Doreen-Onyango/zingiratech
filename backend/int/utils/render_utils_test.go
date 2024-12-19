package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Renders HTML template with provided status code and error message
func TestRenderServerErrorTemplateWithValidInput(t *testing.T) {
	w := httptest.NewRecorder()
	statusCode := http.StatusBadRequest
	errMsg := "Invalid request"

	RenderServerErrorTemplate(w, statusCode, errMsg)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, resp.StatusCode)
	}

	if !strings.Contains(string(body), errMsg) {
		t.Errorf("Response body does not contain error message: %s", errMsg)
	}

	if !strings.Contains(string(body), fmt.Sprintf("Error %d", statusCode)) {
		t.Errorf("Response body does not contain status code: %d", statusCode)
	}
}

// Handle empty error message
func TestRenderServerErrorTemplateWithEmptyError(t *testing.T) {
	w := httptest.NewRecorder()
	statusCode := http.StatusInternalServerError
	errMsg := ""

	RenderServerErrorTemplate(w, statusCode, errMsg)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		t.Errorf("Expected status code %d, got %d", statusCode, resp.StatusCode)
	}

	if !strings.Contains(string(body), fmt.Sprintf("Error %d", statusCode)) {
		t.Errorf("Response body does not contain status code: %d", statusCode)
	}

	if !strings.Contains(string(body), "<p></p>") {
		t.Errorf("Response body should contain empty paragraph tag")
	}
}
