package helpers_test

import (
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestNewErrorValidation(t *testing.T) {
	errors := helpers.NewErrorValidation()
	if errors == nil {
		t.Error("Expected a new ErrorValidation map, got nil")
	}
}

func TestAddError(t *testing.T) {
	errors := helpers.NewErrorValidation()
	errors.Add("Email", "Email is required")

	if len(errors["email"]) != 1 {
		t.Errorf("Expected 1 error for 'email', got %d", len(errors["email"]))
	}
	if errors["email"][0] != "Email is required" {
		t.Errorf("Expected error message 'Email is required', got '%s'", errors["email"][0])
	}
}

func TestGetError(t *testing.T) {
	errors := helpers.NewErrorValidation()
	errors.Add("Name", "Name is required")

	result := errors.Get("name")

	if len(result) != 1 {
		t.Errorf("Expected 1 error for 'name', got %d", len(result))
	}
	if result[0] != "Name is required" {
		t.Errorf("Expected error message 'Name is required', got '%s'", result[0])
	}
}

func TestHasError(t *testing.T) {
	errors := helpers.NewErrorValidation()
	errors.Add("Password", "Password is required")

	if !errors.Has("password") {
		t.Error("Expected Has to return true for 'password'")
	}
	if errors.Has("unknown") {
		t.Error("Expected Has to return false for 'unknown'")
	}
}

func TestKeys(t *testing.T) {
	errors := helpers.NewErrorValidation()
	errors.Add("Email", "Invalid email")
	errors.Add("Password", "Password is required")

	keys := errors.Keys()

	if len(keys) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(keys))
	}
}

func TestValues(t *testing.T) {
	errors := helpers.NewErrorValidation()
	errors.Add("Email", "Invalid email")
	errors.Add("Password", "Password is required")

	values := errors.Values()

	if len(values) != 2 {
		t.Errorf("Expected 2 error messages, got %d", len(values))
	}
}

func TestHasErrors(t *testing.T) {
	errors := helpers.NewErrorValidation()

	if errors.HasErrors() {
		t.Error("Expected HasErrors to return false for an empty map")
	}

	errors.Add("Email", "Invalid email")

	if !errors.HasErrors() {
		t.Error("Expected HasErrors to return true when errors are present")
	}
}
