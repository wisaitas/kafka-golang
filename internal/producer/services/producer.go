package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wisaitas/kafka-golang/internal/producer"
	"github.com/wisaitas/kafka-golang/internal/producer/api/requests"
	"github.com/wisaitas/kafka-golang/internal/producer/api/responses"
	"github.com/wisaitas/kafka-golang/pkg"
)

type ProducerService interface {
	SendMessage(ctx context.Context, req requests.SendMessageRequest) (responses.SendMessageResponse, error)
}

type producerService struct {
	producer pkg.ProducerUtil
}

func NewProducerService(
	producer pkg.ProducerUtil,
) ProducerService {
	return &producerService{
		producer: producer,
	}
}

func (s *producerService) SendMessage(ctx context.Context, req requests.SendMessageRequest) (responses.SendMessageResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return responses.SendMessageResponse{}, err
	}

	message := pkg.Message{
		Brokers: []string{
			fmt.Sprintf("%s:%d", producer.Config.Kafka.Broker.Host, producer.Config.Kafka.Broker.Port),
		},
		Topic: producer.Config.Kafka.Topic.Test,
		Value: jsonData,
	}

	if err := s.producer.SendMessage(ctx, message); err != nil {
		return responses.SendMessageResponse{}, err
	}
	return responses.SendMessageResponse{
		Message: req.Message,
	}, nil
}
