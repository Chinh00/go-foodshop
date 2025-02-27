//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	config2 "go-foodshop/cmd/api/config"
	"go-foodshop/src/infra/database/postgres"
	"go-foodshop/src/usecase"
)

func InitApp(config *config2.Config) (*App, error) {
	panic(wire.Build(
		NewApp,
		echo.New,
		CatPostgresConfig,
		usecase.UseCaseSet,
		usecase.PgFoodRepositorySet,
		NewController,
	))
}
func CatPostgresConfig(config *config2.Config) *postgres.PostgresConfig {
	return &postgres.PostgresConfig{
		Host:     config.PostgresConfig.Host,
		Port:     config.PostgresConfig.Port,
		Username: config.PostgresConfig.Username,
		Password: config.PostgresConfig.Password,
		Database: config.PostgresConfig.Database,
	}
}
