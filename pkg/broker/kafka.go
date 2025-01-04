package broker

import (
	"github.com/amrmrp/broker/pkg/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/amrmrp/broker/pkg/errors"
)

type Kafka struct {
	config  *config.KafkaConfig
}

type KafkaMessage struct {
	ID      string              `json:"id"`
	Command string              `json:"command"`
	Data    map[string][]string `json:"data"`
	Time    time.Time           `json:"time"`
}

func NewKafka(config  *config.KafkaConfig) *Kafka {

	return &Kafka{config :config}
}

func (kafkaInterface *Kafka) Consume(topic string, partition int) {

	// to consume messages
	conn, err := kafka.DialLeader(context.Background(), kafkaInterface.config.Read.PROTOCOL, kafkaInterface.config.Read.BROKERS[0], topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(1, 1e9) // fetch 10KB min, 1MB max

	b := make([]byte, 200e3) // 200KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}

func (kafkaInterface *Kafka) Produce(message map[string][]string, routeKey string, topic string, partition int) {
	/*
		-------------------------------------------------------------------------
		| to produce messages and initial message structure
		-------------------------------------------------------------------------
	*/
	conn, err := kafka.DialLeader(context.Background(), kafkaInterface.config.Read.PROTOCOL, kafkaInterface.config.Read.BROKERS[0], topic, partition)
	if err != nil {
		errors.Mssage("failed to dial leader").Error()
	}

	messages := KafkaMessage{
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
		errors.Mssage("failed to serialize message").Error()
	}

	/*
		-------------------------------------------------------------------------
		| new connection and produce job
		-------------------------------------------------------------------------
	*/
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
			Key:   []byte(routeKey),
			Value: messagesSerialize,
		},
	)

	if err != nil {
		errors.Mssage("failed to produce messages").Error()
	}

	if err := conn.Close(); err != nil {
		errors.Mssage("failed to produce connection").Error()
	}

	log.Println("success produce 😊")
}
