package usecase

import (
	"github.com/labstack/echo/v4"
	"go-foodshop/src/models"
)

type Usecase interface {
	GetFoodsShop(context echo.Context) ([]*models.Food, error)
	GetFoodById(context echo.Context, id int) (*models.Food, error)
	AddFood(context echo.Context, food *models.Food) (*models.Food, error)
}
