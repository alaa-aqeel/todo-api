package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type request struct {
	request *http.Request
}

func Request(r *http.Request) *request {
	return &request{
		request: r,
	}
}

func (r *request) Body() map[string]interface{} {
	var body map[string]interface{}
	if err := json.NewDecoder(r.request.Body).Decode(&body); err != nil {

		return nil
	}
	return body
}

func (r *request) Query(key string) url.Values {
	return r.request.URL.Query()
}

func (r *request) Header(key string) string {
	return r.request.Header.Get(key)
}

func (r *request) Input(key string) string {
	return r.request.PostFormValue(key)
}

func (r *request) Get(key string) string {
	return r.request.FormValue(key)
}

func ParseBody[T any](body io.ReadCloser) (*T, error) {
	var t T
	err := json.NewDecoder(body).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
