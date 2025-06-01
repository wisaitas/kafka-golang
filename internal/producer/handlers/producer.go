package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/kafka-golang/internal/producer/services"
)

type ProducerHandler interface {
	SendMessage(c *fiber.Ctx) error
}

type producerHandler struct {
	producerService services.ProducerService
}

func NewProducerHandler(
	producerService services.ProducerService,
) ProducerHandler {
	return &producerHandler{
		producerService: producerService,
	}
}

func (h *producerHandler) SendMessage(c *fiber.Ctx) error {
	return nil
}
