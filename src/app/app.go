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
}

func NewApp(repository database.Repository[models.Food], usecase usecase2.Usecase, config *configs.Config) *App {
	return &App{
		repository: repository,
		usecase:    usecase,
		config:     config,
	}
}

func NewApplication(configs *configs.Config) error {
	echo := echo.New()

	i := 1
	repository := database.NewInMemoryDB[models.Food]([]*models.Food{
		{
			FoodId: i,
			Name:   "Food 1",
			Price:  1.0,
		},
		{
			FoodId: i + 1,
			Name:   "Food 2",
			Price:  1.0,
		},
		{
			FoodId: i + 2,
			Name:   "Food 3",
			Price:  1.0,
		},
	})

	usecase := usecase2.New(repository)
	NewController(usecase).MapFoodApiV1(echo)
	app := NewApp(repository, usecase, configs)

	echo.Start(fmt.Sprintf("%s:%d", app.config.Host, app.config.Port))
	return nil
}
