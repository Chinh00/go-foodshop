package database

import "github.com/labstack/echo/v4"

type Repository[T any] interface {
	AddEntity(context echo.Context, entity *T) (*T, error)
	RemoveEntity(context echo.Context, id int) (*T, error)
	GetAllEntities(context echo.Context) ([]*T, error)
	GetEntityById(context echo.Context, id int) (*T, error)
}
