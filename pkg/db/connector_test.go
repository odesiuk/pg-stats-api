package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DNS(t *testing.T) {
	assert.Equal(t,
		"host=localhost port=5432 user=user password=pass dbname=db_name sslmode=disable",
		DSN("localhost", "5432", "user", "pass", "db_name"),
	)
}
