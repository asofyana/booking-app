package utils

import (
	"fmt"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DefaultPassword string `yaml:"defaultPassword"`
	DefaultLanguage string `yaml:"defaultLanguage"`
	Port            string `yaml:"port"`
	DbFile          string `yaml:"dbFile"`
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
