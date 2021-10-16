package db

import (
	"fmt"
	"log"

	"github.com/odesiuk/pg-stats-api/pkg/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

// NewConnectionFromENV get connection from ENV vars by prefix.
func NewConnectionFromENV(prefix string) (*gorm.DB, error) {
	host := env.StringOrDefault(prefix+dbHostSuffix, localhost)
	port := env.StringOrDefault(prefix+dbPortSuffix, pgDefaultPort)
	user := env.MustGetString(prefix + dbUserSuffix)
	pass := env.MustGetString(prefix + dbPassSuffix)
	dbName := env.MustGetString(prefix + dbNameSuffix)

	log.Printf("[ DB ] - DSN: %s \n", DSN(host, port, user, "***", dbName))

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN(host, port, user, pass, dbName),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}

// DSN connection string.
func DSN(host, port, user, password, dbName string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
