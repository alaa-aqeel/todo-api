package helpers

import (
	"net/http"
	"strings"
)

type Map map[string]interface{}

func RequestQueryArags(r *http.Request) Map {
	queryMap := make(Map)
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			queryMap[key] = values[0] // Get the first value for each query parameter
		}
	}
	return queryMap
}

func (m Map) Get(key string) interface{} {
	value, exists := m[key]
	if exists && value != "" {
		return value
	}
	return nil
}

func (m Map) GetString(key string) string {
	if value, ok := m[key].(string); ok {
		return strings.TrimSpace(value)
	}
	return ""
}

func (m Map) Has(key string) bool {
	value, exists := m[key]

	return exists && value != ""
}
