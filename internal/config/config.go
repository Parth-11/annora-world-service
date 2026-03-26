package config

import (
	"time"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Cache    CacheConfig
	Queue    QueueConfig
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("PORT", "4000"),
			ReadTimeout:  getDuration("READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getDuration("WRITE_TIMEOUT", 5*time.Second),
		},
		Postgres: PostgresConfig{
			URL:                mustGetEnv("DATABASE_URL"),
			MaxOpenConnections: getInt("DB_MAX_OPEN_CONS", 10),
		},
		Cache: CacheConfig{
			Addr:       getEnv("REDIS_URL", "localhost:6379"),
			Passwd:     getEnv("REDIS_PASSWORD", ""),
			DB:         getInt("REDIS_DB", 0),
			MaxRetries: getInt("REDIS_MAX_RETRIES", 3),
		},
		Queue: QueueConfig{
			Addr:    getEnv("KAFKA_ADDR", "localhost:9092"),
			Timeout: getDuration("KAFKA_TIMEOUT", 5*time.Second),
		},
	}
}
