package producer

var (
	Config = &ProducerConfig{}
)

type ProducerConfig struct {
	Server Server
	Kafka  Kafka
}

type Server struct {
	Env  string `config:"dev"`
	Port int    `config:"8081"`
}

type Kafka struct {
	Broker Broker
	Topic  Topic
}

type Broker struct {
	Host string `config:"localhost"`
	Port int    `config:"9092"`
}

type Topic struct {
	Test string `config:"test"`
}
