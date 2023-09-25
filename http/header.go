package http

import (
	"net/http"
	"strconv"
)

func GetHeaderInt(h http.Header, key string, defaultValue int) int {
	v := h.Get(key)
	if v == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return i
}
