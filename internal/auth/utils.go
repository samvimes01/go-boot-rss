package auth

import (
	"errors"
	"net/http"
	"strings"
)

func ParseApiKeyHeader(r *http.Request) (string, error) {
	authHeader := strings.Fields(r.Header.Get("Authorization"))
	if len(authHeader) != 2 || authHeader[0] != "ApiKey" {
		return "", errors.New("invalid auth header")
	}
	return authHeader[1], nil
}
