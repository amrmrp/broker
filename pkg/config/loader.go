package config

import (
	"os"
	"github.com/amrmrp/broker/pkg/errors"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Configs, error) {
	file, err := os.ReadFile(path)
	if err != nil {

		errors.Mssage("I couldn't read the config file.").Error()
	}

	var config Configs
	configMap := map[string]interface{}{
		"Kafka":    &config.Kafka,
		"RabbitMQ": &config.RabbitMQ,
	}

	for name, target := range configMap {
		err = yaml.Unmarshal(file, target)
		if err != nil {
			errors.Mssage("Error unmarshaling config:" + name).Error()	
		}
	}

	return &config, nil
}
