package fatura

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Key interface {
	string | int
}

type Value interface {
	[]string | string | int | bool | any
}

// Generic array type
type Array[K Key, V Value] map[K]V

func (a Array[K, V]) Get(key K) V {
	return a[key]
}

func (a Array[K, V]) Set(key K, value V) {
	a[key] = value
}

func (a Array[K, V]) Del(key K) {
	delete(a, key)
}

func (a Array[K, V]) Has(key K) bool {
	_, ok := a[key]
	return ok
}

func (a Array[K, V]) Json() string {
	data, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(data)
}

func (a Array[K, V]) Encode() string {
	output := []string{}
	for key, value := range a {
		if encodeValue(value) != "" && encodeKey(key) != "" {
			output = append(output, encodeKey(key)+"="+encodeValue(value))
		}
	}

	return strings.Join(output, "&")
}

func encodeValue(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case []string:
		return strings.Join(v, ",")
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

func encodeKey(key any) string {
	switch k := key.(type) {
	case string:
		return k
	case int:
		return strconv.Itoa(k)
	default:
		return ""
	}
}
