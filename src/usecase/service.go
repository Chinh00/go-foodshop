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

func (s services) GetFoodById(context echo.Context, id string) (*models.Food, error) {
	//TODO implement me
	panic("implement me")
}
