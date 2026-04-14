package queue

import (
	"net"

	"github.com/Parth-11/annora-world-service/internal/config"
	"github.com/rabbitmq/amqp091-go"
)

func NewClient(cfg config.QueueConfig) (*amqp091.Connection, error) {
	return amqp091.DialConfig(
		cfg.Addr,
		amqp091.Config{
			Dial: func(network, addr string) (net.Conn, error) {
				dailer := net.Dialer{Timeout: cfg.Timeout}
				return dailer.Dial(network, addr)
			},
		},
	)
}
