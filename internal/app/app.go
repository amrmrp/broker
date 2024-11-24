package app

import "async-entity-fetcher/pkg/rabbitmq"

//  "fmt"
//   "async-entity-fetcher/pkg/example"

// Start initializes the application
func Start() {

	/*
	data := make(map[string][]string)
	data["entity1"] = []string{"item1", "item2", "item3"}
	data["entity2"] = []string{"item1", "item2", "item3"}
	data["entity3"] = []string{"item1", "item2", "item3"}
	rabbitmq.CreateRabbitProducer(data,"invoice.cmd.created","master")
	*/

    
//	rabbitmq.NewRabbitConsumer("service1", "invoice.cmd.created", "master")

	
	/*
		data := make(map[string][]string)
		data["entity1"] = []string{"item1", "item2", "item3"}
		data["entity2"] = []string{"item1", "item2", "item3"}
		data["entity3"] = []string{"item1", "item2", "item3"}

		kafka.CreateKafkaProducer(data, "invoice.cmd.created", "master-exchange", 0)
	*/

	//kafka.StartKafkaConsumer("master-exchange", 0)
}
