package utils

import (
	"fmt"
	"log"

	"github.com/odesiuk/pg-stats-api/pkg/env"
)

const (
	dbHostSuffix = "_HOST"
	dbPortSuffix = "_PORT"
	dbUserSuffix = "_USER"
	dbPassSuffix = "_PASSWORD"
	dbNameSuffix = "_DATABASE"

	pgDefaultPort = "5432"
	localhost     = "localhost"
)

// GetPgDSNFromENV get connection string from ENV vars by prefix.
func GetPgDSNFromENV(prefix string) string {
	host := env.StringOrDefault(prefix+dbHostSuffix, localhost)
	port := env.StringOrDefault(prefix+dbPortSuffix, pgDefaultPort)
	user := env.MustGetString(prefix + dbUserSuffix)
	pass := env.MustGetString(prefix + dbPassSuffix)
	dbName := env.MustGetString(prefix + dbNameSuffix)

	log.Printf("[ %s_DB ] - DSN: %s \n", prefix, PgDSN(host, port, user, "***", dbName))

	return PgDSN(host, port, user, pass, dbName)
}

// PgDSN connection string.
func PgDSN(host, port, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
