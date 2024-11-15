package app

import (
	//  "fmt"
	//   "async-entity-fetcher/pkg/example"
	"context"
	"log"
	"time"
    "encoding/json"
	"github.com/segmentio/kafka-go"
)

type MyMessage struct {
    ID      string `json:"id"`
    Command string `json:"command"`
    Data    string `json:"data"`
}
// Start initializes the application
func Start() {

	// to produce messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

    msg := MyMessage{
        ID:      "12345",
        Command: "create",
        Data:    "Hello, Kafka!",
    }

    // Serialize the object to JSON
    value, err := json.Marshal(msg)
    if err != nil {
        log.Fatalf("failed to serialize message: %v", err)
    }

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{
            Key: []byte("service.cmd.create"),
            Value: value,
        },
	//	kafka.Message{Value: []byte("test")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}


    log.Println("success produce 😊")

	/*
	   fmt.Println("Application is running...")
	   result := example.SayHello("World")
	   fmt.Println(result)
	*/
}
