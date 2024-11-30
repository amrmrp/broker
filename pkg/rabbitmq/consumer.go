package rabbitmq

import (
	"log"
	"github.com/wagslane/go-rabbitmq"
)

func NewRabbitConsumer(queueName string, routeKey string) {
	/*
		-------------------------------------------------------------------------
		| Initial config and new connection
		-------------------------------------------------------------------------
	*/
	var config Config
	config.GetRabbitConfig()

	conn, err := rabbitmq.NewConn(
		config.RabbitMQ.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		queueName,
		rabbitmq.WithConsumerOptionsRoutingKey(routeKey),
		rabbitmq.WithConsumerOptionsExchangeName(config.RabbitMQ.Exchange.Name),
		rabbitmq.WithConsumerOptionsExchangeKind(config.RabbitMQ.Exchange.Type),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		log.Printf("consumed: %v", string(d.Body))
		// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
}
