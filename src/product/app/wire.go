//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	config2 "go-foodshop/cmd/product/config"
	"go-foodshop/src/pkg/database"
	"go-foodshop/src/product/usecase"
)

func InitApp(config *config2.Config) (*App, error) {
	panic(wire.Build(
		NewApp,
		echo.New,
		database.DbContextSet,
		usecase.UseCaseSet,
		usecase.PgFoodRepositorySet,
		CatPostgresConfig,
		NewController,
	))
}
func CatPostgresConfig(config *config2.Config) *database.PostgresConfig {
	return &config.PostgresConfig
}
