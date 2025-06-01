package services

import "github.com/wisaitas/kafka-golang/internal/producer/api/requests"

type ProducerService interface {
	SendMessage(req requests.SendMessageRequest) error
}

type producerService struct {
}

func NewProducerService() ProducerService {
	return &producerService{}
}

func (s *producerService) SendMessage(req requests.SendMessageRequest) error {
	return nil
}
