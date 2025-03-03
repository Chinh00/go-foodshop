package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go-foodshop/src/infra/database/postgres"
	configs "go-foodshop/src/pkg/config"
	"os"
)

type (
	Config struct {
		configs.App             `yaml:"app"`
		configs.HTTP            `yaml:"http"`
		configs.Log             `yaml:"logger"`
		postgres.PostgresConfig `yaml:"postgres"`
	}
)

func NewConfig() (*Config, error) {
	config := &Config{}

	dir, err := os.Getwd()
	err = cleanenv.ReadConfig(dir+"/config.yml", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
