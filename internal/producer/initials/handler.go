package initials

import "github.com/wisaitas/kafka-golang/internal/producer/handlers"

type handler struct {
	producerHandler handlers.ProducerHandler
}

func newHandler(
	service *service,
) *handler {
	return &handler{
		producerHandler: handlers.NewProducerHandler(service.producerService),
	}
}
