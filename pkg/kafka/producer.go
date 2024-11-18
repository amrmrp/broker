package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Data struct{}

type MyMessage struct {
	ID      string              `json:"id"`
	Command string              `json:"command"`
	Data    map[string][]string `json:"data"`
	Time    time.Time           `json:"time"`
}

func CreateKafkaProducer(message map[string][]string, routeKey string, topic string, partition int) {
	/*
		-------------------------------------------------------------------------
		| to produce messages and initial message structure
		-------------------------------------------------------------------------
	*/
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

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

	/*
		-------------------------------------------------------------------------
		| new connection and produce job
		-------------------------------------------------------------------------
	*/
	conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
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

	fmt.Println("Starting HelloWorld Application...")
	log.Println("success produce 😊")
}
