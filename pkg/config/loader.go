package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Configs, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Configs
	configMap := map[string]interface{}{
		"Kafka":    &config.Kafka,
		"RabbitMQ": &config.RabbitMQ,
	}

	for name, target := range configMap {
		err = yaml.Unmarshal(file, target)
		if err != nil {
			log.Fatalf("Error unmarshaling %s config: %v", name, err)
		}
	}

	return &config, nil
}
