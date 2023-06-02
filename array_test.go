package fatura_test

import (
	"testing"

	"github.com/ugurkorkmaz/fatura"
)

func TestArray(t *testing.T) {
	arr := fatura.Array[int, string]{
		1: "apple",
		2: "banana",
		3: "orange",
	}

	// Get method test
	if arr.Get(1) != "apple" {
		t.Error("arr.Get(1) != \"apple\"")
	}

	// Set method test
	arr.Set(1, "pineapple")
	if arr.Get(1) != "pineapple" {
		t.Error("arr.Get(1) != \"pineapple\"")
	}

	// Del method test
	arr.Del(1)
	if arr.Get(1) != "" {
		t.Error("arr.Get(1) != \"\"")
	}

	// Has method test
	if arr.Has(1) {
		t.Error("arr.Has(1) == true")
	}

	// Json method test
	if arr.Json() != "{\"2\":\"banana\",\"3\":\"orange\"}" {
		t.Error("arr.Json() != \"{\"2\":\"banana\",\"3\":\"orange\"}\"")
	}

	// Encode method test
	if arr.Encode() != "2=banana&3=orange" {
		t.Error("arr.Encode() != \"2=banana&3=orange\"")
	}

}
