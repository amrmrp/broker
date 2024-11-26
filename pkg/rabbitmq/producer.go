package rabbitmq

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/wagslane/go-rabbitmq"
)

type Config struct {
	RabbitMQ struct {
		URL      string `yaml:"url"`
		Exchange struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		} `yaml:"exchange"`
		Queue struct {
			Name       string `yaml:"name"`
			RoutingKey string `yaml:"routing_key"`
		} `yaml:"queue"`
	} `yaml:"rabbitmq"`
}

type MyMessage struct {
	ID      string              `json:"id"`
	Command string              `json:"command"`
	Data    map[string][]string `json:"data"`
	Time    time.Time           `json:"time"`
}

func CreateRabbitProducer(message map[string][]string, routeKey string, topic string) {
	
	var config Config

	conn, err := rabbitmq.NewConn(
		config.RabbitMQ.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	messages := MyMessage{
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
		log.Fatalf("failed to serialize message: %v", err)
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(config.RabbitMQ.Exchange.Name),
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
		rabbitmq.WithPublishOptionsExchange(topic),
	)
	if err != nil {
		log.Println(err)
	}
}
