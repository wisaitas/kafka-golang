package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/kafka-golang/internal/producer/api/requests"
	"github.com/wisaitas/kafka-golang/internal/producer/api/responses"
	"github.com/wisaitas/kafka-golang/internal/producer/services"
	"github.com/wisaitas/share-pkg/response"
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
	var req requests.SendMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resp, err := h.producerService.SendMessage(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(response.Response[responses.SendMessageResponse]{
		Code:    fiber.StatusOK,
		Message: "send message success",
		Data:    resp,
	})
}
