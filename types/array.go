package types

// Package types provides a custom type for handling maps with string arrays.
// This type allows you to easily manipulate arrays associated with specific keys.
import "encoding/json"

// Array is a custom type for handling maps with string arrays.
type Array map[string][]string

// Get retrieves the first value associated with the given key from the map.
// If there is no array associated with the key, Get returns an empty string.
// If you need to access multiple values, you should use the map directly.
func (a Array) Get(key string) string {
	// If the map is empty, return an empty string.
	if a == nil {
		return ""
	}
	// Get the values associated with the key.
	vs := a[key]
	// If the values slice is empty, return an empty string.
	if len(vs) == 0 {
		return ""
	}
	// Return the first value associated with the key.
	return vs[0]
}

// Set sets the key to the specified value. It replaces any existing values.
func (a Array) Set(key, value string) {
	// Set the key to a slice containing only the value.
	a[key] = []string{value}
}

// Add adds the specified value to the slice of values associated with the key.
// If the key does not already exist in the map, it creates a new key-value pair.
func (a Array) Add(key, value string) {
	// Append the value to the slice of values associated with the key.
	a[key] = append(a[key], value)
}

// Del deletes the key-value pair associated with the specified key.
func (a Array) Del(key string) {
	delete(a, key)
}

// Has checks whether the map contains the specified key.
func (a Array) Has(key string) bool {
	// Check if the key exists in the map.
	_, ok := a[key]
	return ok
}

// Json returns a JSON string representation of the map.
func (a Array) Json() string {
	// Marshal the map to a JSON string.
	data, err := json.Marshal(a)
	// If there is an error while marshalling, return an empty string.
	if err != nil {
		return ""
	}
	// Return the JSON string representation of the map.
	return string(data)
}
