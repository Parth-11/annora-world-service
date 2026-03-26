package config

import "time"

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
	}
}
