package main

import (
	"context"
	"log"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
)

func main() {
	config := config.Load()

	_ = database.New(config)
	if err := database.Migrate(context.Background(), config); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	log.Println("App started.")
}
