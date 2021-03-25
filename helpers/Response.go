// Package helpers implements commonly used functions (response API)//
package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponse is
type APIResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// APIResponse is
type APIResponseNotFound struct {
	Id         string `json:"id"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

// Response is
func Response(w http.ResponseWriter, httpStatus int, data interface{}) {
	apiResponse := new(APIResponse)
	apiResponse.Status = httpStatus
	apiResponse.Data = data

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(apiResponse)
}

func ResponseBadRequest(w http.ResponseWriter, httpStatus int, err error) {
	apiResponse := new(APIResponseNotFound)
	apiResponse.Error = err.Error()
	apiResponse.StatusCode = httpStatus

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(apiResponse)
}

// ResponseNotFound is
func ResponseNotFound(w http.ResponseWriter) {
	apiResponse := new(APIResponseNotFound)
	apiResponse.Error = "Not Found"
	apiResponse.StatusCode = 404

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(apiResponse)
}
