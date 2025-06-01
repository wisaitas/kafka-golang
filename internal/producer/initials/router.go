package initials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/wisaitas/kafka-golang/internal/producer"
	"github.com/wisaitas/kafka-golang/internal/producer/routes"
)

type route struct {
	producerRoutes *routes.ProducerRoute
}

func newRoute(
	router *fiber.App,
	handler *handler,
) {
	if producer.Config.Server.Env != producer.EnvDev {
		router.Get("/swagger/*", swagger.New(
			swagger.Config{},
		))
	}

	apiRoute := router.Group("/api/v1")

	route := &route{
		producerRoutes: routes.NewProducerRoute(
			apiRoute,
			handler.producerHandler,
		),
	}

	route.registerRoutes()
}

func (r *route) registerRoutes() {
	r.producerRoutes.RegisterRoutes()
}
