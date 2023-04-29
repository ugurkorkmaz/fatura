package fatura

import "encoding/json"

type Array map[string][]string

// Get returns the first value associated with the given key.
func (a Array) Get(key string) string {
	if a == nil {
		return ""
	}
	vs := a[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to the specified value. It replaces any existing values.
func (a Array) Set(key, value string) {
	a[key] = []string{value}
}

// Add adds the specified value to the slice of values associated with the key.
// If the key does not already exist in the map, it creates a new key-value pair.
func (a Array) Add(key, value string) {
	a[key] = append(a[key], value)
}

// Del deletes the key-value pair associated with the specified key.
func (a Array) Del(key string) {
	delete(a, key)
}

// Has checks whether the map contains the specified key.
func (a Array) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// Json returns a JSON string representation of the map.
func (a Array) Json() string {
	data, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(data)
}
