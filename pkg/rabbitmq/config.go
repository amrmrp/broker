package rabbitmq

import "async-entity-fetcher/pkg/config"

type Config struct {
	RabbitMQ struct {
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

func (configs *Config) GetRabbitConfig() {
	var configStruct = configs
	var configHandler config.Configs
	configHandler.SetConfigs("./../../configs/connection-config.yaml", &configStruct)
}
