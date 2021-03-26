// Package helpers implements commonly used functions (response API)//
package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponseToken is
type APIResponseToken struct {
	AccessToken  string      `json:"access_token"`
	ExpiresIn    int         `json:"expires_in"`
	TokenType    string      `json:"token_type"`
	Scope        interface{} `json:"scope"`
	RefreshToken string      `json:"refresh_token"`
}

// APIResponseResource
type APIResponseResource struct {
	AccessToken  string      `json:"access_token"`
	ClientID     string      `json:"client_id"`
	UserID       string      `json:"user_id"`
	FullName     string      `json:"full_name"`
	Npm          string      `json:"npm"`
	Expires      interface{} `json:"expires"`
	RefreshToken string      `json:"refresh_token"`
}

// APIResponseError is
type APIResponseError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Response is
func Response(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ResponseBadRequest is
func ResponseBadRequest(w http.ResponseWriter, httpStatus int, err error) {
	apiResponse := new(APIResponseError)
	apiResponse.Error = "invalid request"
	apiResponse.ErrorDescription = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(apiResponse)
}
