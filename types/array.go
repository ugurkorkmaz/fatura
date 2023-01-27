package types

import "encoding/json"

type Array map[string][]string

// Get gets the first value associated with the given key.
// If there are no Array associated with the key, Get returns
// the empty string. To access multiple Array, use the map
// directly.
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

// Set sets the key to value. It replaces any existing
// Array.
func (a Array) Set(key, value string) {
	a[key] = []string{value}
}

// Add adds the value to key. It appends to any existing
// Array associated with key.
func (a Array) Add(key, value string) {
	a[key] = append(a[key], value)
}

// Del deletes the Array associated with key.
func (a Array) Del(key string) {
	delete(a, key)
}

// Has checks whether a given key is set.
func (a Array) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// Json returns the json string of the Array.
func (a Array) Json() string {
	data, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(data)
}
