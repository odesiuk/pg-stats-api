package main

import (
	"log"

	"github.com/odesiuk/pg-stats-api/cmd"
)

func main() {
	// get config from ENV.
	cfg := cmd.ParseConfigFromENV()

	if err := cmd.Start(cfg); err != nil {
		log.Fatal(err)
	}

	log.Println("Exit without error")
}
