package helpers

import "strings"

type ErrorValidation map[string][]string

func NewErrorValidation() ErrorValidation {
	return make(ErrorValidation)
}

func (e ErrorValidation) Add(key string, err string) {
	e[strings.ToLower(key)] = append(e[key], err)
}

func (e ErrorValidation) Get(key string) []string {
	return e[key]
}

func (e ErrorValidation) Has(key string) bool {
	_, ok := e[strings.ToLower(key)]
	return ok
}

func (e ErrorValidation) Keys() []string {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}
	return keys
}

func (e ErrorValidation) Values() []string {
	values := make([]string, 0, len(e))
	for _, v := range e {
		values = append(values, v...)
	}
	return values
}

func (e ErrorValidation) HasErrors() bool {
	return len(e.Values()) > 0
}
