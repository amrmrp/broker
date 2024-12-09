package broker

type Broker interface {
	Produce(topic string, message []byte) error
	Consume(topic string, handler func(message []byte) error) error
	Close() error
}
