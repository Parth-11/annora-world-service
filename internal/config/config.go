package config

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

func Load() *Config {
	return &Config{}
}
