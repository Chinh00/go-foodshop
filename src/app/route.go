package app

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-foodshop/docs"
	usecase2 "go-foodshop/src/usecase"
)

// @title FoodShop API
// @version 1.0
// @description API tài liệu cho FoodShop
// @host localhost:8080
// @BasePath /api/v1

type controller struct {
	usecase usecase2.Usecase
}

func NewController(usecase usecase2.Usecase) *controller {
	return &controller{
		usecase: usecase,
	}
}

// @title FoodShop API
// @version 1.0
// @description API tài liệu cho FoodShop
// @host localhost:8080
// @BasePath /api/v1

func (controller *controller) MapFoodApiV1(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/api/v1")
	api.GET("/foods", controller.HandleGetFoods)
}

// HandleGetFoods godoc
// @Summary Lấy danh sách món ăn
// @Description Trả về danh sách các món ăn
// @Tags foods
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/foods [get]
func (controller *controller) HandleGetFoods(context echo.Context) error {
	data, err := controller.usecase.GetFoodsShop(context)
	if err != nil {
		return err
	}
	return context.JSON(200, data)
}
