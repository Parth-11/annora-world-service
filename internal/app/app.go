package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Parth-11/annora-world-service/internal/config"
	"github.com/Parth-11/annora-world-service/internal/handlers/https"
	worldhandler "github.com/Parth-11/annora-world-service/internal/handlers/https/world"
	"github.com/Parth-11/annora-world-service/internal/infrastructure/db"
	worldrepo "github.com/Parth-11/annora-world-service/internal/repository/world"
	worldservice "github.com/Parth-11/annora-world-service/internal/service/world"
)

type App struct {
	Server *http.Server
	Config *config.Config
}

func New(cfg *config.Config) (*App, error) {
	//setup infra
	ctx := context.Background()
	db := db.NewPostgresDB(ctx, cfg.Postgres)

	world_repo := worldrepo.NewWorldRepo(db)

	world_service := worldservice.NewService(world_repo)

	world_handler := worldhandler.NewHandler(world_service)

	router := https.NewRouter(world_handler)
	return &App{
		Config: cfg,
		Server: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler:      router,
			ReadTimeout:  cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
	}, nil
}

func (a *App) Start() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("World Service listening on %s\n", a.Server.Addr)

		if err := a.Server.ListenAndServe(); err != nil && err == http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()

	<-stop
	log.Println("Shutdown signal recieved")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed %v\n", err)
		return err
	}

	log.Println("World service stopped gracefully")
	return nil
}
