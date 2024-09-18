package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func TestRenderError(t *testing.T) {
	Init()
	tests := []struct {
		name           string
		status         int
		message        string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test 404 Not Found",
			status:         http.StatusNotFound,
			message:        "Page Not Found",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Error 404",
		},
		{
			name:           "Test 500 Internal Server Error",
			status:         http.StatusInternalServerError,
			message:        "Internal Server Error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Error 500",
		},
		{
			name:           "Test 403 Forbidden",
			status:         http.StatusForbidden,
			message:        "Access Denied",
			expectedStatus: http.StatusForbidden,
			expectedBody:   "Error 403",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			renderError(w, tt.status, tt.message)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, w.Code)
			}

			if !strings.Contains(w.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q; got %q", tt.expectedBody, w.Body.String())
			}

			if !strings.Contains(w.Body.String(), tt.message) {
				t.Errorf("expected body to contain %q; got %q", tt.message, w.Body.String())
			}
		})
	}
}

func TestInit(t *testing.T) {
	// Temporarily replace the global errorTemplate
	originalTemplate := errorTemplate
	defer func() { errorTemplate = originalTemplate }()

	// Reset errorTemplate to nil
	errorTemplate = nil

	// Call init() manually
	Init()

	// Check if errorTemplate is not nil after init
	if errorTemplate == nil {
		t.Error("errorTemplate is nil after init")
	}

	testCases := []struct {
		name         string
		code         int
		message      string
		expectedBody string
	}{
		{"Not Found Error", 404, "Test Error", "Error 404"},
		{"Internal Server Error", 500, "Server Error", "Error 500"},
		{"Forbidden Error", 403, "Access Denied", "Error 403"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := errorTemplate.Execute(w, struct {
				Code    int
				Message string
			}{
				Code:    tc.code,
				Message: tc.message,
			})
			if err != nil {
				t.Errorf("Error executing template: %v", err)
			}

			if !strings.Contains(w.Body.String(), tc.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tc.expectedBody, w.Body.String())
			}

			if !strings.Contains(w.Body.String(), tc.message) {
				t.Errorf("Expected body to contain %q, got %q", tc.message, w.Body.String())
			}
		})
	}

	// Test with invalid template path
	errorTemplate = nil
	Init()
	if errorTemplate == nil {
		t.Error("Fallback template was not created when given an invalid path")
	}
}
