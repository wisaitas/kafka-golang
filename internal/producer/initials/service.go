package initials

import "github.com/wisaitas/kafka-golang/internal/producer/services"

type service struct {
	producerService services.ProducerService
}

func newService(
	util *util,
) *service {
	return &service{
		producerService: services.NewProducerService(util.producer),
	}
}
