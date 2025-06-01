package producer

var (
	Config = &ProducerConfig{}
)

type ProducerConfig struct {
	Server Server
}

type Server struct {
	Env     string   `config:"env"`
	Port    int      `config:"port"`
	Brokers []Broker `config:"brokers"`
	Topic   string   `config:"topic"`
}

type Broker struct {
	Host string `config:"host"`
	Port int    `config:"port"`
}
