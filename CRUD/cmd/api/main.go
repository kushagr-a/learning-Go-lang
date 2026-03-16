package main

import (
	"fmt"
	"log"
	"note-api/internal/config"
	"note-api/internal/db"
	"note-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	client, database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	defer func() {
		if err := db.DisconnectDB(client); err != nil {
			log.Fatalf("db disconnect error: %v", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}

}
