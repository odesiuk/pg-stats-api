package cmd

import "github.com/odesiuk/pg-stats-api/pkg/env"

const appName = "PG-STATS-API"

// Config - env variables.
type Config struct {
	AppName          string
	Port             string
	MinQueryDuration uint64
}

// ParseConfigFromENV creates config, getting values from env.
func ParseConfigFromENV() Config {
	return Config{
		AppName:          appName,
		Port:             env.StringOrDefault("PORT", "8080"),
		MinQueryDuration: env.Uint64OrDefault("MIN_QUERY_DURATION", 2000),
	}
}
