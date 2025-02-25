package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	configs "go-foodshop/src/pkg/config"
	"os"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
	}
)

func NewConfig() (*Config, error) {
	config := &Config{}

	dir, err := os.Getwd()
	err = cleanenv.ReadConfig(dir+"/cmd/api/config.yml", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
