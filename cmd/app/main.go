package main

import (
	"log"

	"github.com/flaiers/flallet/internal/app"
	"github.com/flaiers/flallet/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.New(cfg).Listen(cfg.Addr))
}
