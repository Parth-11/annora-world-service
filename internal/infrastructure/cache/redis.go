package cache

import (
	"github.com/Parth-11/annora-world-service/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg config.CacheConfig) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:       cfg.Addr,
			Password:   cfg.Passwd,
			DB:         cfg.DB,
			MaxRetries: cfg.MaxRetries,
		})
}
