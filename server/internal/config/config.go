package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	HTTPClient `yaml:"httpClient"`
}

type HTTPClient struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Protocol string `yaml:"protocol"`
}

func MustReadConfig() *Config {

	file, err := os.ReadFile("config/local_config.yaml")

	if err != nil {
		log.Fatal("not read yaml file ", err)
	}

	var config Config

	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("not convert yaml", err)
	}

	return &config
}
