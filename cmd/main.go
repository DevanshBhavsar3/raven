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
	config := config.Load()

	db := database.New(config)
	if err := database.Migrate(context.Background(), config); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	repos := internal.NewRepositories(db)
	svcs := internal.NewServices(repos)
	handlers := internal.NewHandlers(svcs)
	router := internal.NewRouter(handlers)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
