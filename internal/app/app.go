package app

import (
	//  "fmt"
	//   "async-entity-fetcher/pkg/example"
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Start initializes the application
func Start() {

	// to produce messages
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
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
