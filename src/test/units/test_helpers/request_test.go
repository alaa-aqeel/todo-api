package helpers_test

import (
	"bytes"
	"io"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestRequestBody(t *testing.T) {
	jsonData := `{"name":"John Doe","age":30}`
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")

	r := helpers.Request(req)
	body := r.Body()

	if body["name"] != "John Doe" {
		t.Errorf("Expected name to be 'John Doe', got '%v'", body["name"])
	}
	if int(body["age"].(float64)) != 30 { // JSON numbers are decoded as float64
		t.Errorf("Expected age to be 30, got '%v'", body["age"])
	}
}

func TestRequestQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=John&age=30", nil)
	r := helpers.Request(req)

	query := r.Query("name")
	if query.Get("name") != "John" {
		t.Errorf("Expected name to be 'John', got '%s'", query.Get("name"))
	}
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer token123")

	r := helpers.Request(req)
	header := r.Header("Authorization")

	if header != "Bearer token123" {
		t.Errorf("Expected Authorization header to be 'Bearer token123', got '%s'", header)
	}
}

func TestRequestInput(t *testing.T) {
	form := url.Values{}
	form.Set("key", "value")

	req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r := helpers.Request(req)
	input := r.Input("key")

	if input != "value" {
		t.Errorf("Expected input to be 'value', got '%s'", input)
	}
}

func TestParseBody(t *testing.T) {
	type Payload struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	jsonData := `{"name":"Alice","age":25}`
	body := io.NopCloser(bytes.NewBufferString(jsonData))

	result, err := helpers.ParseBody[Payload](body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Name != "Alice" {
		t.Errorf("Expected Name to be 'Alice', got '%s'", result.Name)
	}
	if result.Age != 25 {
		t.Errorf("Expected Age to be 25, got '%d'", result.Age)
	}
}
