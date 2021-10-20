package env

import (
	"fmt"
	"os"
	"strconv"
)

// Uint64OrDefault parse ENV var by key with default value (uint64).
// panic in case of ParseUint error.
func Uint64OrDefault(key string, def uint64) uint64 {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	n, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("parsing env varriable `%s` error: %s", key, err))
	}

	return n
}
