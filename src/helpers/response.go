package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	responseWriter http.ResponseWriter
	statusCode     int
}

func Response(w http.ResponseWriter) *response {
	return &response{
		responseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (resp *response) Header(key, value string) *response {
	resp.responseWriter.Header().Set(key, value)

	return resp
}

func (resp *response) Status(code int) *response {
	resp.statusCode = code

	return resp
}

func (resp *response) Json(statusCode int, data any) {
	resp.Header("Content-Type", "Application/json")
	resp.responseWriter.WriteHeader(statusCode)
	json.NewEncoder(resp.responseWriter).Encode(data)
}

func (resp *response) Success(msg string, data any) {
	resp.Json(resp.statusCode, Map{
		"message": msg,
		"data":    data,
		"status":  "success",
	})
}

// Error writes a JSON response with a 400 status code and the given message and errors in the body.
func (resp *response) Error(msg string, errors any) {
	resp.Json(resp.statusCode, Map{
		"message": msg,
		"errors":  errors,
		"status":  "error",
	})
}

func (resp *response) ValidationErrors(errors any) {
	resp.Json(http.StatusBadRequest, Map{
		"message": "Invalid data",
		"errors":  errors,
		"status":  "error",
	})
}

func (resp *response) NotFound() {
	resp.Json(http.StatusNotFound, Map{
		"message": "Not found",
		"status":  "error",
	})
}

func (resp *response) NotFoundModel(model string) {
	resp.Json(http.StatusNotFound, Map{
		"message": "Not found " + model,
		"status":  "error",
	})
}

func (r *response) NoContent() {
	r.Json(http.StatusNoContent, Map{
		"message": "No content",
		"status":  "success",
	})
}
