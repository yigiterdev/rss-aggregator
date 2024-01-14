package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
// Example:
// Authorization: ApiKey {API_KEY}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("Authorization header is missing")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("Authorization header is invalid")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Authorization header is invalid")
	}

	return vals[1], nil
}
