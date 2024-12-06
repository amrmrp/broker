package broker

type Broker interface {
	Produce(topic string, message []byte) error
	Consume(topic string, handler func(message []byte) error) error
	Close() error
}

func NewBroker(brokerType string) Broker {
	switch brokerType {
	case "kafka":
		return &kafka.KafkaBroker{}
	case "rabbitmq":
		return &rabbitmq.RabbitMQBroker{}
	default:
		return nil
	}
}
