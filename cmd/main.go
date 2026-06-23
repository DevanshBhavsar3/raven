package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DevanshBhavsar3/raven/internal"
	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
)

func main() {
	cfg := config.Load()

	db := database.New(cfg)
	if err := database.Migrate(context.Background(), cfg); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	repos := internal.NewRepositories(db)
	svcs := internal.NewServices(cfg, repos)
	handlers := internal.NewHandlers(svcs)
	router := internal.NewRouter(handlers)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Server.Port), router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
