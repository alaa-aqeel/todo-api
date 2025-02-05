package helpers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestResponseHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)

	resp.Header("Content-Type", "application/json")

	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", recorder.Header().Get("Content-Type"))
	}
}

func TestResponseJson(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)

	data := map[string]string{"message": "Hello, World!"}
	resp.Json(http.StatusOK, data)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code to be 200, got %d", recorder.Code)
	}

	var result map[string]string
	if err := json.NewDecoder(recorder.Body).Decode(&result); err != nil {
		t.Fatalf("Expected valid JSON, got error %v", err)
	}
	if result["message"] != "Hello, World!" {
		t.Errorf("Expected message 'Hello, World!', got '%s'", result["message"])
	}
}

func TestResponseSuccess(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)

	resp.Success("Operation successful", map[string]string{"key": "value"})

	var result map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&result)

	if result["message"] != "Operation successful" {
		t.Errorf("Expected success message, got '%s'", result["message"])
	}
	if result["status"] != "success" {
		t.Errorf("Expected status to be 'success', got '%s'", result["status"])
	}
}

func TestResponseError(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)

	resp.Error("Something went wrong", "Invalid input")

	var result map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&result)

	if result["message"] != "Something went wrong" {
		t.Errorf("Expected error message, got '%s'", result["message"])
	}
	if result["status"] != "error" {
		t.Errorf("Expected status to be 'error', got '%s'", result["status"])
	}
}

func TestResponseValidationErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)

	errors := map[string]string{"field": "This field is required"}
	resp.ValidationErrors(errors)

	var result map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&result)

	if result["message"] != "Invalid data" {
		t.Errorf("Expected message 'Invalid data', got '%s'", result["message"])
	}
	if result["status"] != "error" {
		t.Errorf("Expected status to be 'error', got '%s'", result["status"])
	}
}

func TestResponseNotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	resp := helpers.Response(recorder)
	resp.NotFound()

	var result map[string]interface{}
	json.NewDecoder(recorder.Body).Decode(&result)

	if result["message"] != "Not found" {
		t.Errorf("Expected message 'Not found', got '%s'", result["message"])
	}
	if result["status"] != "error" {
		t.Errorf("Expected status to be 'error', got '%s'", result["status"])
	}
}
