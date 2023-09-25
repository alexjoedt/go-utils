package http

import (
	"net/url"
	"strconv"
	"strings"
)

func ReadString(qs url.Values, key string, defaulValue string) string {
	s := qs.Get(key)
	if s == "" {
		return defaulValue
	}

	return s
}

// ReadCSV reads a comma seperated string from the given key and returns a
// slice of the values
func ReadCSV(qs url.Values, key string, defaulValue []string) []string {
	csv := qs.Get(key)
	if csv == "" {
		return defaulValue
	}

	return strings.Split(csv, ",")
}

// ReadBoolParam reads the key from the url query and returns true if
// the param is equal 'true' otherwise it returns false.
func ReadBoolParam(qs url.Values, key string) bool {
	return qs.Get(key) == "true"
}

func ReadInt(qs url.Values, key string, defaultValue int) int {
	s := qs.Get(key)
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return i
}
