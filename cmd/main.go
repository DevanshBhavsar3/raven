package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
	"github.com/DevanshBhavsar3/raven/internal/handlers"
	"github.com/DevanshBhavsar3/raven/internal/repositories"
	"github.com/DevanshBhavsar3/raven/internal/router"
	"github.com/DevanshBhavsar3/raven/internal/services"
)

func main() {
	config := config.Load()

	db := database.New(config)
	if err := database.Migrate(context.Background(), config); err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	repos := repositories.NewRepositories(db)
	svcs := services.NewServices(repos)
	handlers := handlers.NewHandlers(svcs)
	router := router.NewRouter(handlers)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), router); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
