package usecase

import (
	"fmt"
	"github.com/docker/distribution/uuid"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go-foodshop/src/pkg/database"
	"go-foodshop/src/pkg/database/postgres"
	"go-foodshop/src/product/domain"
)

var _ Usecase = (*services)(nil)

var UseCaseSet = wire.NewSet(New)
var PgFoodRepositorySet = wire.NewSet(NewPgFoodRepository)

func NewPgFoodRepository(context database.DbContext) database.Repository[domain.Product] {
	pgRepo, err := postgres.NewPgRepository[domain.Product](context)
	if err != nil {
		fmt.Println(err.Error())
	}

	return pgRepo
}

type services struct {
	repository database.Repository[domain.Product]
}

func New(repository database.Repository[domain.Product]) Usecase {
	return &services{
		repository: repository,
	}
}

func (s *services) GetFoodsShop(context echo.Context) ([]*domain.Product, error) {
	data, err := s.repository.GetAllEntities(context)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *services) GetFoodById(context echo.Context, id int) (*domain.Product, error) {
	data, err := s.repository.GetEntityById(context, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (s *services) AddFood(context echo.Context, food *domain.Product) (*domain.Product, error) {
	food.Id = uuid.Generate()
	data, err := s.repository.AddEntity(context, food)
	if err != nil {
		return nil, err
	}
	return data, nil
}
