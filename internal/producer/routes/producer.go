package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/kafka-golang/internal/producer/handlers"
)

type ProducerRoute struct {
	router          fiber.Router
	producerHandler handlers.ProducerHandler
}

func NewProducerRoute(
	router fiber.Router,
	producerHandler handlers.ProducerHandler,
) *ProducerRoute {
	return &ProducerRoute{
		router:          router,
		producerHandler: producerHandler,
	}
}

func (r *ProducerRoute) RegisterRoutes() {
	producerRoute := r.router.Group("/producers")

	producerRoute.Post("", r.producerHandler.SendMessage)
}
