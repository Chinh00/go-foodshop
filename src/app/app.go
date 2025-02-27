package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	configs "go-foodshop/cmd/api/config"
	"go-foodshop/src/infra/database"
	"go-foodshop/src/models"
	usecase2 "go-foodshop/src/usecase"
)

type App struct {
	config     *configs.Config
	repository database.Repository[models.Food]
	usecase    usecase2.Usecase
	echo       *echo.Echo
	controller *Controller
}

func (a *App) StartApplication() error {
	a.controller.MapFoodApiV1(a.echo)
	a.echo.Start(fmt.Sprintf("%s:%d", a.config.HTTP.Host, a.config.HTTP.Port))
	return nil
}

func NewApp(repository database.Repository[models.Food], usecase usecase2.Usecase, config *configs.Config, echo *echo.Echo, controller *Controller) *App {
	return &App{
		repository: repository,
		usecase:    usecase,
		config:     config,
		echo:       echo,
		controller: controller,
	}
}
