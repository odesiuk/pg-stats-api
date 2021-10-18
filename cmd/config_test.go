package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigFromENV(t *testing.T) {
	tests := map[string]struct {
		env map[string]string
		exp Config
	}{
		"default": {
			env: map[string]string{
				"PORT":               "",
				"MIN_QUERY_DURATION": "",
			},
			exp: Config{
				AppName:          "PG-STATS-API",
				Port:             "8080",
				MinQueryDuration: 2000,
			},
		},
		"new": {
			env: map[string]string{
				"PORT":               "2021",
				"MIN_QUERY_DURATION": "777",
			},
			exp: Config{
				AppName:          "PG-STATS-API",
				Port:             "2021",
				MinQueryDuration: 777,
			},
		},
	}

	// non-concurrency test because of os.Setenv.
	for _, tt := range tests {
		for key, value := range tt.env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}
		assert.Equal(t, tt.exp, ParseConfigFromENV())
	}
}
