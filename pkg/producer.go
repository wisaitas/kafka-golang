package pkg

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer interface {
	SendMessage(ctx context.Context, message Message) error
}

type producer struct {
	producer *kafka.Writer
}

func NewProducer(brokers []string, topic string) Producer {
	return &producer{
		producer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
		}),
	}
}

func (p *producer) SendMessage(ctx context.Context, message Message) error {
	return p.producer.WriteMessages(ctx, mappingMessage(message))
}
