package helpers_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/alaa-aqeel/todo/src/helpers"
)

func TestRequestQueryArags(t *testing.T) {
	query := url.Values{}
	query.Set("name", "John")
	query.Set("age", "30")
	request, _ := http.NewRequest("GET", "/test?"+query.Encode(), nil)

	queryMap := helpers.RequestQueryArags(request)

	if queryMap.GetString("name") != "John" {
		t.Errorf("Expected name to be 'John', got '%v'", queryMap["name"])
	}
	if queryMap.GetString("age") != "30" {
		t.Errorf("Expected age to be '30', got '%v'", queryMap["age"])
	}
}

func TestMapGet(t *testing.T) {
	m := helpers.Map{"name": "John"}

	value := m.Get("name")
	if value != "John" {
		t.Errorf("Expected 'John', got '%v'", value)
	}

	nilValue := m.Get("nonexistent")
	if nilValue != nil {
		t.Errorf("Expected nil, got '%v'", nilValue)
	}
}

// TestMapGetString checks the Map GetString method
func TestMapGetString(t *testing.T) {
	m := helpers.Map{"name": "  John "}

	value := m.GetString("name")
	if value != "John" {
		t.Errorf("Expected 'John', got '%v'", value)
	}

	emptyValue := m.GetString("nonexistent")
	if emptyValue != "" {
		t.Errorf("Expected empty string, got '%v'", emptyValue)
	}
}

// TestMapHas checks the Map Has method
func TestMapHas(t *testing.T) {
	m := helpers.Map{"name": "John"}

	if !m.Has("name") {
		t.Error("Expected true for existing key 'name'")
	}

	if m.Has("nonexistent") {
		t.Error("Expected false for nonexistent key")
	}
}
