package initials

import "github.com/wisaitas/kafka-golang/pkg"

type util struct {
	producer pkg.ProducerUtil
}

func newUtil() *util {
	return &util{
		producer: pkg.NewProducerUtil(),
	}
}
