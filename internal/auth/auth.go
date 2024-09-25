package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", errors.New("no authentication info found")
	}

	mappedAuth := strings.Split(auth, " ")
	if len(mappedAuth) != 2 {
		return "", errors.New("wrong auth header")
	}
	if mappedAuth[0] != "ApiKey" {
		return "", errors.New("wrong Auth Identifier")
	}

	return mappedAuth[1], nil
}
