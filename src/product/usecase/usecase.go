package usecase

import (
	"github.com/labstack/echo/v4"
	"go-foodshop/src/product/domain"
)

type Usecase interface {
	GetFoodsShop(context echo.Context) ([]*domain.Product, error)
	GetFoodById(context echo.Context, id int) (*domain.Product, error)
	AddFood(context echo.Context, food *domain.Product) (*domain.Product, error)
}
