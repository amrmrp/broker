### README: Kafka Producer and Consumer Example

#### Overview
This example showcases dynamic Kafka message production and consumption using Go. Messages are generated dynamically from a map and consumed using a Kafka consumer.

---

### **Producer**
#### Steps:
1. **Prepare Map**: Map keys represent entities, values are slices of strings.
2. **Send Messages**:
   ```go
   kafka.CreateKafkaProducer(data, "invoice.cmd.created", "master-exchange", 0)
   ```
   - **Parameters**:
     - `data`: Map with dynamic keys and values.
     - `"invoice.cmd.created"`: Kafka topic.
     - `"master-exchange"`: Exchange name.
     - `0`: Partition.

#### Example:
```go
data := map[string][]string{
    "entity1": {"item1", "item2"},
    "entity2": {"item3", "item4"},
}

kafka.CreateKafkaProducer(data, "invoice.cmd.created", "master-exchange", 0)
```

---

### **Consumer**
#### Steps:
1. **Start Kafka Consumer**:
   ```go
   kafka.StartKafkaConsumer("master-exchange", 0)
   ```
   - **Parameters**:
     - `"master-exchange"`: Name of the exchange.
     - `0`: Partition.

#### Example:
```go
kafka.StartKafkaConsumer("master-exchange", 0)
```

---

### **Usage**
- Produce messages using `CreateKafkaProducer`.
- Consume messages dynamically using `StartKafkaConsumer`. 

This architecture ensures seamless, scalable message handling in Kafka-based systems.