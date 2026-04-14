package config

import "time"

type QueueConfig struct {
	Addr    string
	Timeout time.Duration
}
