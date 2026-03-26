package config

import (
	"time"

	"github.com/segmentio/kafka-go"
)

type QueueConfig struct {
	Addr      string
	Timeout   time.Duration
	Transport kafka.RoundTripper
}
