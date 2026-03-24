package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Parth-11/annora-world-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(ctx context.Context, config config.PostgresConfig) *pgxpool.Pool {

	pg_config, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		log.Fatal("Error Unable to connect to the database: ", err)
	}

	pg_config.MaxConns = int32(config.MaxOpenConnections)
	pg_config.MinConns = 5
	pg_config.MaxConnLifetime = time.Hour

	DB, err := pgxpool.NewWithConfig(ctx, pg_config)

	if err != nil {
		log.Fatal("Error unable to connect to database: ", err)
	}
	fmt.Println("Success: Connected to database")

	return DB
}
