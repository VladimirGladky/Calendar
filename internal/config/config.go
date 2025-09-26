package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"port" env-default:"4040"`
	Host string `yaml:"host" env-default:"localhost"`
}

func NewConfig() (*Config, error) {

	var cfg Config
	err := cleanenv.ReadConfig("./config/config.yaml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
