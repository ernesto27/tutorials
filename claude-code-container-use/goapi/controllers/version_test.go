package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	// Create a temporary version.json file for testing
	versionContent := `{"version":"1.0.0"}`
	err := os.WriteFile("version.json", []byte(versionContent), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("version.json")

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var actualResponse VersionResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &actualResponse); err != nil {
		t.Fatal(err)
	}

	expectedVersion := "1.0.0"
	if actualResponse.Version != expectedVersion {
		t.Errorf("handler returned unexpected version: got %v want %v",
			actualResponse.Version, expectedVersion)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, "application/json")
	}
}

func TestVersionHandlerFileNotFound(t *testing.T) {
	// Ensure version.json doesn't exist
	os.Remove("version.json")

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestVersionHandlerInvalidJSON(t *testing.T) {
	// Create a version.json file with invalid JSON
	invalidJSON := `{"version": invalid}`
	err := os.WriteFile("version.json", []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("version.json")

	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}