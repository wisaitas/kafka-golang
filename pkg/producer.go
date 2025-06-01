package pkg

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type ProducerUtil interface {
	SendMessage(ctx context.Context, message Message) error
}

type producerUtil struct {
}

func NewProducerUtil() ProducerUtil {
	return &producerUtil{}
}

func (p *producerUtil) SendMessage(ctx context.Context, message Message) error {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(message.Brokers...),
		Topic:    message.Topic,
		Balancer: &kafka.LeastBytes{}, // ใช้ LeastBytes balancer สำหรับการกระจาย message
	}
	defer writer.Close()

	kafkaMessage := mappingMessage(message)

	err := writer.WriteMessages(ctx, kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %w", err)
	}

	return nil
}
