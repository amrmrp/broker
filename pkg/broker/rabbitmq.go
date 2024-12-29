package broker

import (
	"encoding/json"
	"log"
	"time"
    "github.com/amrmrp/broker/pkg/errors"
	"github.com/amrmrp/broker/pkg/config"
	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

type RabbitMQ struct {
	config *config.RabbitMQConfig
}

type MessageRabbitMQ struct {
	ID      string              `json:"id"`
	Command string              `json:"command"`
	Data    map[string][]string `json:"data"`
	Time    time.Time           `json:"time"`
}

func NewRabbitMQ(config *config.RabbitMQConfig) *RabbitMQ {

	return &RabbitMQ{config: config}

}

func (rabbitmqInterface *RabbitMQ) Produce(message map[string][]string, routeKey string) {
	/*
		-------------------------------------------------------------------------
		| Initial config and new connection
		-------------------------------------------------------------------------
	*/
	conn, err := rabbitmq.NewConn(
		rabbitmqInterface.config.Read.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {

		errors.Mssage("rabbitMQ connection failed.").Error()
	}
	defer conn.Close()

	messages := MessageRabbitMQ{
		ID:      uuid.New().String(),
		Command: routeKey,
		Data:    message,
		Time:    time.Now(),
	}

	/*
		-------------------------------------------------------------------------
		| Serialize the object to JSON
		-------------------------------------------------------------------------
	*/
	messagesSerialize, err := json.Marshal(messages)
	if err != nil {
		errors.Mssage("rabbit marshal message failed.").Error()
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(rabbitmqInterface.config.Read.Exchange.Name),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	err = publisher.Publish(
		[]byte(messagesSerialize),
		[]string{routeKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(rabbitmqInterface.config.Read.Exchange.Name),
	)
	if err != nil {
		errors.Mssage("rabbit publish message failed.").Error()
	}
}

func (rabbitmqInterface *RabbitMQ) Consume(queueName string, routeKey string) {
	/*
		-------------------------------------------------------------------------
		| Initial config and new connection
		-------------------------------------------------------------------------
	*/
	conn, err := rabbitmq.NewConn(
		rabbitmqInterface.config.Read.URL,
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
		rabbitmq.WithConsumerOptionsExchangeName(rabbitmqInterface.config.Read.Exchange.Name),
		rabbitmq.WithConsumerOptionsExchangeKind(rabbitmqInterface.config.Read.Exchange.Type),
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
