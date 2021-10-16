package env

import (
	"fmt"
	"os"
)

// StringOrDefault parse env string var or return default.
func StringOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}

// MustGetString parse env string var or panic if not exists.
func MustGetString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("environment variable %s is required", key))
	}

	return val
}
