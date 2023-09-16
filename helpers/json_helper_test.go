package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJson(t *testing.T) {
	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	// Define your test data (e.g., a struct to encode to JSON)
	testData := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	// Call your WriteJson function
	err := WriteJson(w, http.StatusOK, testData)

	// Check for errors
	if err != nil {
		t.Errorf("WriteJson returned an error: %v", err)
	}

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check the Content-Type header
	expectedContentType := "application/json"
	if contentType := w.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected Content-Type header to be %s, but got %s", expectedContentType, contentType)
	}

	// Decode the response body and check its content
	var response struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check the decoded response
	if response != testData {
		t.Errorf("Expected response %v, but got %v", testData, response)
	}
}
