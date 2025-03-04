package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-foodshop/cmd/product/config"
	"go-foodshop/src/pkg/database"
	"go-foodshop/src/product/domain"
	"go-foodshop/src/product/usecase"
	"net/http"
)

type App struct {
	config     *config.Config
	dbContext  database.DbContext
	repository database.Repository[domain.Product]
	usecase    usecase.Usecase
	echo       *echo.Echo
	controller *Controller
}

func (a *App) StartApplication() error {
	a.controller.MapFoodApiV1(a.echo)
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	err := a.echo.Start(fmt.Sprintf("%s:%d", a.config.HTTP.Host, a.config.HTTP.Port))
	if err != nil {
		return err
	}
	return nil
}

func NewApp(repository database.Repository[domain.Product], dbContext database.DbContext, usecase usecase.Usecase, config *config.Config, echo *echo.Echo, controller *Controller) *App {
	return &App{
		repository: repository,
		usecase:    usecase,
		config:     config,
		echo:       echo,
		controller: controller,
		dbContext:  dbContext,
	}
}
