### README: RabbitMQ Producer and Consumer Example

#### Overview
This demonstrates the dynamic creation of RabbitMQ messages and consumption using a producer-consumer model with `amq.topic` exchange.

---

### **Producer**
#### Steps:
1. **Prepare Data**:
   - Create a map where keys represent entities, and values are slices of strings.

2. **Send Messages**:
   ```go
   rabbitmq.CreateRabbitProducer(data, "invoice.cmd.created", "amq.topic")
   ```
   - **Parameters**:
     - `data`: Map with dynamic keys and values.
     - `"invoice.cmd.created"`: Routing key for the messages.
     - `"amq.topic"`: Exchange name.

#### Example:
```go
data := make(map[string][]string)
data["entity1"] = []string{"item1", "item2", "item3"}
data["entity2"] = []string{"item1", "item2", "item3"}
data["entity3"] = []string{"item1", "item2", "item3"}

rabbitmq.CreateRabbitProducer(data, "invoice.cmd.created", "amq.topic")
```

---

### **Consumer**
#### Steps:
1. **Set Up Consumer**:
   ```go
   rabbitmq.NewRabbitConsumer("service1", "invoice.cmd.created", "amq.topic")
   ```

2. **Parameters**:
   - `service1`: Consumer identifier.
   - `invoice.cmd.created`: Routing key to filter messages.
   - `amq.topic`: Exchange to consume from.

#### Example:
```go
err := rabbitmq.NewRabbitConsumer("service1", "invoice.cmd.created", "amq.topic")
if err != nil {
    log.Fatalf("Failed to start consumer: %v", err)
}
```

This setup allows your application to consume and process messages dynamically routed via the `amq.topic` exchange.