package config

import (
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

type Configs struct{ Config interface{} }

func (config *Configs) SetConfigs(location string, structure interface{}) {

	// Read YAML file
	file, err := os.ReadFile(location)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Unmarshal into the Config struct
	err = yaml.Unmarshal(file, structure)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML content: %v", err)
	}

	config.Config = structure
}

func (config *Configs) GetConfigs() {
	// Access and print interface
	fmt.Printf("configs: %s\n", config)
}