package helpers_test

import (
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
	"golang.org/x/crypto/bcrypt"
)

func TestMakeHash(t *testing.T) {
	value := "secret-key"

	hashedValue, err := helpers.MakeHash(value)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(hashedValue) == 0 {
		t.Fatal("Expected a non-empty hashed value")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value)); err != nil {
		t.Fatal("Hashed value did not match the original value")
	}
}

func TestVerifyHash(t *testing.T) {
	value := "secret-key"
	wrongValue := "wrong-secret-key"

	hashedValue, err := helpers.MakeHash(value)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test correct password
	if !helpers.VerifyHash(value, hashedValue) {
		t.Error("Expected true for matching password and hash")
	}

	// Test incorrect password
	if helpers.VerifyHash(wrongValue, hashedValue) {
		t.Error("Expected false for non-matching password and hash")
	}
}
