package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB    ConfigDB    `yaml:"db"`
	App   ConfigApp   `yaml:"app"`
	Kafka KafkaConfig `yaml:"kafka"`
}

type ConfigApp struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ConfigDB struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

func Load(cfgPath string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	return &cfg, nil
}
