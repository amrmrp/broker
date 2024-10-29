package kafka

import (
    "github.com/segmentio/kafka-go"
	"context"
	"log"

)

func CreateKafkaProducer(brokerAddress,topic string,message []byte)error{

	writer := kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err  := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key: []byte("Key"),
			Value: message,
		},
)

	if err != nil {
		log.Printf("Failed to write message to Kafka: %v",err)
	}

	return err
}