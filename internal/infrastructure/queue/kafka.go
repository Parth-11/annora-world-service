package queue

import (
	"github.com/Parth-11/annora-world-service/internal/config"
	"github.com/segmentio/kafka-go"
)

func NewConn(cfg config.QueueConfig) *kafka.Client {
	return &kafka.Client{
		Addr:      kafka.TCP(cfg.Addr),
		Timeout:   cfg.Timeout,
		Transport: cfg.Transport,
	}
}
