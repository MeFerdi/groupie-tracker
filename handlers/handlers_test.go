package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to execute a request and return the response
func executeRequest(req *http.Request, handler http.HandlerFunc) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func TestArtistsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artists", nil)
	rr := executeRequest(req, ArtistsHandler)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := "<expected_content>" // Set expected content based on actual data
	if string(body) != expected {
		t.Errorf("expected body %v; got %v", expected, string(body))
	}
}

func TestArtistHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist/1", nil)
	rr := executeRequest(req, ArtistHandler)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := "<expected_content>" // Set expected content based on actual data
	if string(body) != expected {
		t.Errorf("expected body %v; got %v", expected, string(body))
	}
}

func TestLocationHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/locations/1", nil)
	rr := executeRequest(req, LocationHandler)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := "<expected_content>" // Set expected content based on actual data
	if string(body) != expected {
		t.Errorf("expected body %v; got %v", expected, string(body))
	}
}

func TestDateHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/dates/1", nil)
	rr := executeRequest(req, DateHandler)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := "<expected_content>" // Set expected content based on actual data
	if string(body) != expected {
		t.Errorf("expected body %v; got %v", expected, string(body))
	}
}

func TestRelationHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/relation/1", nil)
	rr := executeRequest(req, RelationHandler)

	resp := rr.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}

	expected := "<expected_content>" // Set expected content based on actual data
	if string(body) != expected {
		t.Errorf("expected body %v; got %v", expected, string(body))
	}
}
