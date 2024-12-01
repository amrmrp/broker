package kafka

import "async-entity-fetcher/pkg/config"

type Config struct {
	Kfka struct {
		BROKERS   string `yaml:"brokers"`
		TOPIC     string `yaml:"topic"`
		GROUP_ID  string `yaml:"group_id"`
	} `yaml:"kafka"`
}

func(configs *Config) GetRabbitConfig(){
	var configStruct = configs
	var configHandler config.Configs
	configHandler.SetConfigs("./../../configs/config.yaml", &configStruct)
}