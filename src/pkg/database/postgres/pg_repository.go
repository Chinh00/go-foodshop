package postgres

import (
	"github.com/labstack/echo/v4"
	"go-foodshop/src/pkg/database"
	"gorm.io/gorm"
)

type Options func(db *gorm.DB)

type pgRepository[T any] struct {
	db *gorm.DB
}

func (pg *pgRepository[T]) AddEntity(context echo.Context, entity *T) (*T, error) {
	err := pg.db.AutoMigrate(entity)
	if err != nil {
		return nil, err
	}
	pg.db.Create(entity)
	return entity, nil
}

func (pg *pgRepository[T]) RemoveEntity(context echo.Context, id int) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (pg *pgRepository[T]) GetAllEntities(context echo.Context) ([]*T, error) {
	data := make([]*T, 0)
	pg.db.Find(&data)
	return data, nil
}

func (pg *pgRepository[T]) GetEntityById(context echo.Context, id int) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func NewPgRepository[T any](context database.DbContext) (database.Repository[T], error) {
	dbContext, err := context.GetDatabase()
	if err != nil {
		return nil, err
	}
	return &pgRepository[T]{
		db: dbContext,
	}, nil
}
