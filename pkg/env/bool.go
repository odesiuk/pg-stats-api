package env

import (
	"fmt"
	"os"
	"strconv"
)

// BoolOrDefault parse env boolean var or return default.
func BoolOrDefault(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		fmt.Printf("evn '%s'(bool) parcing error ", key)
	}

	return b
}
