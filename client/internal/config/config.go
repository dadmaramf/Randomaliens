package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	HTTPClient      `yaml:"httpClient"`
	PostgresConnect `yaml:"postgresConnect"`
}

type HTTPClient struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type PostgresConnect struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
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
