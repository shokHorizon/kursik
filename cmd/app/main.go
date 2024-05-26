package app

import (
	"context"
	"github.com/shokHorizon/kursik/config"
	"github.com/shokHorizon/kursik/database"
	"os"
)

func main() {
	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// Create ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// make graceful shutdown chan
	shutdown := make(chan os.Signal, 1)
	// launch goroutine for graceful shutdown
	go func() {
		<-shutdown
		cancel()
	}()

	// Postgres connection
	db, err := database.ConnectPostgres(*cfg.Repository.Postgres)
	if err != nil {
		panic(err)
	}

	// Create server
	coursesRepo := NewCoursesRepository()

	server :=

}
