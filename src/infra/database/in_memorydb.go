package database

import "github.com/labstack/echo/v4"

type in_memorydb[T any] struct {
	items []*T
}

func NewInMemoryDB[T any](items []*T) *in_memorydb[T] {
	return &in_memorydb[T]{
		items: items,
	}
}

func (receiver *in_memorydb[T]) GetAllEntities(context echo.Context) ([]*T, error) {
	var result []*T

	for _, value := range receiver.items {
		result = append(result, value)
	}
	return result, nil
}
