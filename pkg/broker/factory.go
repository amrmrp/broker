package broker

import "github.com/amrmrp/broker/pkg/config"

type BrokerManager struct {
	KafkaBroker *Kafka
	RabbitMQBroker *RabbitMQ
}

func NewBrokerManager(config *config.Configs) (*BrokerManager, error) {

	kafkaBroker := NewKafka(config.Kafka)
	rabbitMQBroker := NewRabbitMQ(config.RabbitMQ)

	return &BrokerManager{
		KafkaBroker : kafkaBroker,
		RabbitMQBroker : rabbitMQBroker,
	}, nil
}
