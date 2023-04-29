package fatura_test

import (
	"testing"

	"github.com/ugurkorkmaz/fatura"
)

func TestArray(t *testing.T) {
	arr := make(fatura.Array)

	// Test Set() and Get()
	arr.Set("name", "John")
	if arr.Get("name") != "John" {
		t.Errorf("Expected value: %s, got: %s", "John", arr.Get("name"))
	}

	// Test Add()
	arr.Add("phone", "123456789")
	if len(arr["phone"]) != 1 || arr["phone"][0] != "123456789" {
		t.Errorf("Expected value: %v, got: %v", []string{"123456789"}, arr["phone"])
	}
	arr.Add("phone", "987654321")
	if len(arr["phone"]) != 2 || arr["phone"][1] != "987654321" {
		t.Errorf("Expected value: %v, got: %v", []string{"123456789", "987654321"}, arr["phone"])
	}

	// Test Del()
	arr.Del("name")
	if arr.Has("name") {
		t.Errorf("Expected key not to exist: %s", "name")
	}

	// Test Json()
	expectedJSON := `{"phone":["123456789","987654321"]}`
	if arr.Json() != expectedJSON {
		t.Errorf("Expected JSON string: %s, got: %s", expectedJSON, arr.Json())
	}
}
