### Final Documentation: **Inter-Service Queue Management Package**

This package, hosted at **`github.com/amrmrp/broker`**, is designed to manage message queues for **inter-microservice communication** using **Kafka** and **RabbitMQ**.

---

### Project Structure
```
.
├── README.md                # Project Documentation
├── go.mod                   # Go module file
├── go.sum                   # Dependency checksums
├── configs
│   ├── example.yaml         # Example configuration
│   └── schema.yaml          # Schema for validation
├── pkg
│   ├── broker
│   │   ├── broker.go        # Interface for queue brokers
│   │   ├── factory.go       # Factory for broker initialization
│   │   ├── kafka.go         # Kafka implementation
│   │   └── rabbitmq.go      # RabbitMQ implementation
│   ├── config
│   │   ├── config.go        # Configuration definitions
│   │   ├── loader.go        # YAML configuration loader
│   │   └── README.md        # Config usage docs
│   └── errors
│       └── errors.go        # Custom error handling
└── main.go                  # Usage entry point
```

---

### Configuration Example (`configs/example.yaml`)
```yaml
kafka:
  brokers: ["localhost:9092"]
  topic: "service-topic"
  group_id: "service-group"
  protocol: "tcp"

rabbitmq:
  url: "amqp://guest:guest@localhost:5672/"
  exchange:
    name: "service-exchange"
    type: "direct"
  queue:
    name: "service-queue"
    routing_key: "service-key"
```

---

### Code Usage Example

#### 1. **Load Configuration and Initialize Brokers**
```go
package main

import (
	"log"
	"github.com/amrmrp/broker/pkg/config"
	"github.com/amrmrp/broker/pkg/broker"
)

func main() {
	// Load configurations
	configs, err := config.LoadConfig("./configs/example.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize brokers via factory
	brokerManager, err := broker.NewBrokerManager(configs)
	if err != nil {
		log.Fatalf("Failed to initialize brokers: %v", err)
	}

	// Produce sample data
	data := map[string][]string{
		"user_events": {"login", "logout"},
	}

	brokerManager.KafkaBroker.Produce(data, "user-event", "service-topic", 0)
	brokerManager.RabbitMQBroker.Produce(data, "service-key")
}
```

---

### Key Features
1. **Unified Interface**: Kafka and RabbitMQ implementations adhere to a single `Broker` interface.
2. **Dynamic Configuration**: Easily switch or configure messaging systems through YAML files.
3. **Scalable**: Designed for microservice architectures requiring reliable inter-service messaging.
4. **Extensible**: Additional brokers (e.g., AWS SQS, Google Pub/Sub) can be integrated.

---

### Factory Design (`broker/factory.go`)
The factory dynamically initializes Kafka or RabbitMQ brokers based on the provided configuration.

```go
func NewBrokerManager(cfg *config.Configs) (*BrokerManager, error) {
	kafkaBroker := NewKafka(&cfg.Kafka)
	rabbitMQBroker := NewRabbitMQ(&cfg.RabbitMQ)

	return &BrokerManager{
		KafkaBroker:    kafkaBroker,
		RabbitMQBroker: rabbitMQBroker,
	}, nil
}
```

---

### Benefits for Microservices
- Decouples services via reliable message queues.
- Supports both **publish/subscribe** (RabbitMQ) and **partitioned topics** (Kafka).
- Ensures modularity, reusability, and maintainability.

This package adheres to **Go project standards** and best practices and is a robust solution for inter-service message queue management.