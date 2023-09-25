package http

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrInvalidAuthHeader error = errors.New("invalid auth header")
)

func BearerAuth(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrInvalidAuthHeader
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}
