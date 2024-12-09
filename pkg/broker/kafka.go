package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	Data interface{}
}

type Config struct {
	Kafka struct {
		BROKERS  []string `yaml:"brokers"`
		TOPIC    string   `yaml:"topic"`
		GROUP_ID string   `yaml:"group_id"`
		PROTOCOL string   `yaml:"protocol"`
	} `yaml:"kafka"`
}

type Message struct {
	ID      string              `json:"id"`
	Command string              `json:"command"`
	Data    map[string][]string `json:"data"`
	Time    time.Time           `json:"time"`
}

func NewKafka(brokers []string, topic string) *Kafka {

	var config *Config
	return &Kafka{config}
}

func (kafkaInterface *Kafka) Consume(topic string, partition int) {

	// to consume messages
	conn, err := kafka.DialLeader(context.Background(), kafkaInterface.PROTOCOL, kafkaInterface.BROKERS[0], topic, partition)
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
	conn, err := kafka.DialLeader(context.Background(), kafkaInterface.PROTOCOL, kafkaInterface.BROKERS[0], topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	messages := Message{
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
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	log.Println("success produce 😊")
}
