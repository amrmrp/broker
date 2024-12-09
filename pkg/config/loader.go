package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) (*Configs, error) {
	file, err := os.ReadFile(path)
	if(err != nil){
		return nil,err
	}

	var config Configs
	err = yaml.Unmarshal(file,&config)
	if err != nil{
		return nil, err
	}

	return &config, nil
}