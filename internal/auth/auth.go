package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get Api key extract an Api key from
// the headers of an http request
// Example:
// Authorization: API Key  {insert api key here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization header info found")

	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed Authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed First part of auth header")

	}

	return vals[1], nil
}
