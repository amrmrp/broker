package config

type KafkaConfig struct {
	Read struct {
		BROKERS  []string `yaml:"brokers"`
		TOPIC    string   `yaml:"topic"`
		GROUP_ID string   `yaml:"group_id"`
		PROTOCOL string   `yaml:"protocol"`
	} `yaml:"kafka"`
}

type RabbitMQConfig struct {
	Read struct {
		URL      string `yaml:"url"`
		Exchange struct {
			Name string `yaml:"name"`
			Type string `yaml:"type"`
		} `yaml:"exchange"`
		Queue struct {
			Name        string   `yaml:"name"`
			RoutingKeys []string `yaml:"routing_keys"`
		} `yaml:"queue"`
	} `yaml:"rabbitmq"`
}

type Configs struct {
	Kafka    *KafkaConfig
	RabbitMQ *RabbitMQConfig
}
