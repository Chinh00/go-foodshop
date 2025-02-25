package usecase

import (
	"github.com/labstack/echo/v4"
	"go-foodshop/src/infra/database"
	"go-foodshop/src/models"
)

var _ Usecase = (*services)(nil)

type services struct {
	repository database.Repository[models.Food]
}

func New(repository database.Repository[models.Food]) Usecase {
	return &services{
		repository: repository,
	}
}

func (s *services) GetFoodsShop(context echo.Context) ([]*models.Food, error) {
	data, err := s.repository.GetAllEntities(context)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *services) GetFoodById(context echo.Context, id int) (*models.Food, error) {
	data, err := s.repository.GetEntityById(context, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *services) AddFood(context echo.Context, food *models.Food) (*models.Food, error) {
	data, err := s.repository.AddEntity(context, food)
	if err != nil {
		return nil, err
	}
	return data, nil
}
