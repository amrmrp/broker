package kafka

import(
	"context"
	"log"
    "github.com/segmentio/kafka-go"
)


func StartKafkaConsumer(brokerAddress,topic string){
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic: topic,
		GroupID: "consumer-group-1",
	})

	defer reader.Close()

	for{
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Printf("Error reading message: %v",err)
			continue
		}
		log.Printf("Received message at %v %s",msg.Time,string(msg.Value))
	}
}