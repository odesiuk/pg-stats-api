package main

import (
	"log"

	"github.com/odesiuk/pg-stats-api/cmd"
	"github.com/odesiuk/pg-stats-api/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// get config from ENV.
	cfg := cmd.ParseConfigFromENV()

	dbConf := postgres.New(postgres.Config{
		DSN:                  utils.GetPgDSNFromENV("PG"),
		PreferSimpleProtocol: true,
	})

	// get DB connection.
	dbConn, err := gorm.Open(dbConf, &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection ERROR:", err)
	}

	app := cmd.Setup(dbConn, cfg)
	log.Println("Exit", app.Listen(":"+cfg.Port))
}
