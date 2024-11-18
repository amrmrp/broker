### README: Kafka Producer Example with Dynamic Map Data

#### Overview
This example demonstrates creating a Kafka producer in Go that dynamically sends messages using a map where keys represent entities, and values are slices of strings. The messages are published to a Kafka topic with specified routing.

#### Steps
1. **Initialize Map**: Keys (`entity1`, `entity2`, etc.) hold slices of values (e.g., `item1`, `item2`).
2. **Kafka Producer Call**:
   ```go
   kafka.CreateKafkaProducer(data, "invoice.cmd.created", "master-exchange", 0)
   ```
   - `data`: Map containing entities and items.
   - `invoice.cmd.created`: Kafka topic.
   - `master-exchange`: Kafka exchange.
   - `0`: Partition number.

#### Example
```go
data := make(map[string][]string)
data["entity1"] = []string{"item1", "item2", "item3"}
data["entity2"] = []string{"item1", "item2", "item3"}
data["entity3"] = []string{"item1", "item2", "item3"}

kafka.CreateKafkaProducer(data, "invoice.cmd.created", "master-exchange", 0)
```

This setup ensures dynamic publishing of multiple messages to Kafka topics.