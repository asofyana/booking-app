package utils

import (
	"fmt"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DefaultPassword string `yaml:"defaultPassword"`
	Language        string `yaml:"language"`
}

var config Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		yamlFile, _ := os.ReadFile("config.yaml")
		err := yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			fmt.Println("Error loading config")
		}

	})
	return config
}
