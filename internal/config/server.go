package config

import "time"

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Time
	WriteTimeout time.Time
}
