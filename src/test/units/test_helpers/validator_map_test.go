package helpers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestNewValidatorMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	validator := helpers.NewValidatorMap(req)

	if validator == nil {
		t.Error("Expected NewValidatorMap to return a non-nil instance")
	}
}

func TestValidateMapParseBody(t *testing.T) {
	body := map[string]interface{}{"name": "John"}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	validator := helpers.NewValidatorMap(req)

	if !validator.Valid() {
		t.Error("Expected parseBody to parse the valid JSON body without errors")
	}
}

func TestValidateMapWithRules(t *testing.T) {
	body := map[string]interface{}{"email": "invalid_email"}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	validator := helpers.NewValidatorMap(req).
		Rules(map[string]interface{}{
			"email": "required,email",
		})

	if validator.Valid() {
		t.Error("Expected validation to fail for an invalid email")
	}

	errors := validator.Errors()
	if !errors.Has("email") {
		t.Error("Expected validation error for 'email'")
	}
}

func TestValidateMapValidated(t *testing.T) {
	body := map[string]interface{}{"name": "John", "age": 30}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	validator := helpers.NewValidatorMap(req).
		Rules(map[string]interface{}{
			"name": "required",
			"age":  "required,number",
		})

	if !validator.Valid() {
		t.Error("Expected validation to pass")
	}

	validatedData := validator.Validated()
	if validatedData["name"] != "John" {
		t.Errorf("Expected name to be 'John', got '%v'", validatedData["name"])
	}
	if validatedData["age"].(float64) != 30 {
		t.Errorf("Expected age to be 30, got '%v'", validatedData["age"])
	}
}

func TestValidateMapErrorHandling(t *testing.T) {
	body := map[string]interface{}{"name": ""}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	validator := helpers.NewValidatorMap(req).
		Rules(map[string]interface{}{
			"name": "required",
		})

	if validator.Valid() {
		t.Error("Expected validation to fail when 'name' is empty")
	}

	if !validator.Errors().Has("name") {
		t.Error("Expected error for missing 'name'")
	}
}
