package database

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-foodshop/src/models"
	"net/http"
)

type in_memorydb struct {
	items []*models.Food
}

func NewInMemoryDB(items []*models.Food) *in_memorydb {
	return &in_memorydb{
		items: items,
	}
}

func (receiver *in_memorydb) GetAllEntities(context echo.Context) ([]*models.Food, error) {
	var result []*models.Food

	for _, value := range receiver.items {
		result = append(result, value)
	}
	return result, nil
}
func (receiver *in_memorydb) GetEntityById(context echo.Context, id int) (*models.Food, error) {
	var result *models.Food = nil
	for _, value := range receiver.items {
		if value.FoodId == id {
			result = value
		}
	}
	if result == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}
	return result, nil
}

func (receiver *in_memorydb) AddEntity(context echo.Context, entity *models.Food) (*models.Food, error) {
	items_size := len(receiver.items)
	entity.FoodId = items_size + 1
	err := append(receiver.items, entity)
	if err != nil {
		fmt.Println(err)
	}
	return entity, nil
}

func (receiver *in_memorydb) RemoveEntity(context echo.Context, id int) (*models.Food, error) {
	panic("implement me")
}
