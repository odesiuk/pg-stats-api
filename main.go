package main

import (
	"log"

	"github.com/odesiuk/pg-stats-api/cmd"
	"github.com/odesiuk/pg-stats-api/pkg/db"
)

func main() {
	// get config from ENV.
	cfg := cmd.ParseConfigFromENV()

	// get DB connection.
	dbConn, err := db.NewConnectionFromENV("PG")
	if err != nil {
		log.Fatal("DB connection ERROR:", err)
	}

	app := cmd.Setup(dbConn, cfg)
	log.Println("Exit", app.Listen(":"+cfg.Port))
}
