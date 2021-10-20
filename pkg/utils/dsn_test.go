package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DNS(t *testing.T) {
	assert.Equal(t,
		"host=localhost port=5432 user=user password=pass dbname=db_name sslmode=disable",
		PgDSN("localhost", "5432", "user", "pass", "db_name"),
	)
}

func TestGetPgDSNFromENV(t *testing.T) {
	tests := map[string]struct {
		prefix string
		env    map[string]string
		exp    string
	}{
		"default": {
			prefix: "MY_PG",
			env: map[string]string{
				"MY_PG_USER":     "my_user",
				"MY_PG_PASSWORD": "my_pass",
				"MY_PG_DATABASE": "my_db",
			},
			exp: "host=localhost port=5432 user=my_user password=my_pass dbname=my_db sslmode=disable",
		},
		"new": {
			prefix: "PG",
			env: map[string]string{
				"PG_HOST":     "777.777.777",
				"PG_PORT":     "777",
				"PG_USER":     "my_user",
				"PG_PASSWORD": "my_pass",
				"PG_DATABASE": "my_db",
			},
			exp: "host=777.777.777 port=777 user=my_user password=my_pass dbname=my_db sslmode=disable",
		},
	}

	// non-concurrency test because of os.Setenv.
	for _, tt := range tests {
		for key, value := range tt.env {
			if err := os.Setenv(key, value); err != nil {
				t.Fatal(err)
			}
		}

		assert.Equal(t, tt.exp, GetPgDSNFromENV(tt.prefix))
	}
}
