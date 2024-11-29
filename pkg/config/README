### README: Config Package Usage

#### Overview
This package handles reading a YAML configuration file, unmarshaling it into a dynamic structure, and providing easy access to the data.

---

### **Usage Example**

1. **Define Your Configuration Structure**  
   Define a struct that matches the structure of your YAML configuration file. Example:

   ```go
   type Config struct {
       RabbitMQ struct {
           URL      string `yaml:"url"`
           Exchange struct {
               Name string `yaml:"name"`
               Type string `yaml:"type"`
           } `yaml:"exchange"`
           Queue struct {
               Name       string `yaml:"name"`
               RoutingKey string `yaml:"routing_key"`
           } `yaml:"queue"`
       } `yaml:"rabbitmq"`
   }
   ```

2. **Initialize Config and Read YAML**  
   Create a variable of type `Config`, and pass it to the `SetConfigs` function:

   ```go
   var configStruct = &Config{}
   var configHandler config.Configs
   configHandler.SetConfigs("./path/to/config.yaml", configStruct)
   ```

3. **Access Config Data**  
   After loading the YAML, access the fields as follows:

   ```go
   fmt.Println(configStruct.RabbitMQ.Exchange.Name) // prints the Exchange name
   ```

4. **Sample YAML File:**

   ```yaml
   rabbitmq:
     url: "amqp://admin:admin@localhost"
     exchange:
       name: "amq.topic"
       type: "topic"
     queue:
       name: "my_queue"
       routing_key: "invoice.cmd.created"
   ```

---

### **Package Methods**

- `SetConfigs(location string, structure interface{})`:  
  Reads the YAML file from the provided location and unmarshals it into the provided structure.

- `GetConfigs()`:  
  Accesses the loaded configuration and prints or returns the data.

Make sure to handle errors when accessing or modifying your configuration for robustness.