package main

import (
	"log"

	"github.com/MirzaKarabulut/fill-labs/cmd"
	"github.com/MirzaKarabulut/fill-labs/internal/db"
)

func main() {
	if err := db.Connection(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	cmd.CreateServer(8080)
}
