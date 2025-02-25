package database

import "github.com/labstack/echo/v4"

type Repository[T any] interface {
	//AddEntity(entity *T) (*T, error)
	//RemoveEntity(entity *T) (*T, error)
	GetAllEntities(context echo.Context) ([]*T, error)
	//GetEntityById(id string) (*T, error)
}
