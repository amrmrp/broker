package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume(topic string, partition int) {

	// to consume messages
	conn, err := kafka.DialLeader(context.Background(), config.Kfka.PROTOCOL, config.Kfka.BROKERS[0], topic, partition)
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
