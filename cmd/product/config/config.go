package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	configs "go-foodshop/src/pkg/config"
	"go-foodshop/src/pkg/database"
	"os"
)

type (
	Config struct {
		configs.App             `yaml:"app"`
		configs.HTTP            `yaml:"http"`
		configs.Log             `yaml:"logger"`
		database.PostgresConfig `yaml:"postgres"`
	}
)

func NewConfig() (*Config, error) {
	config := &Config{}

	dir, err := os.Getwd()
	err = cleanenv.ReadConfig(dir+"/config.yml", config)
	if err != nil {
		return nil, err
	}
	println(config.PostgresConfig.Username)
	return config, nil
}
