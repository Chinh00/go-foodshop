package app

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-foodshop/cmd/product/docs"
	"go-foodshop/src/product/domain"
	"go-foodshop/src/product/usecase"
	"log/slog"
	"strconv"
)

type Controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

// @title FoodShop API
// @version 1.0
// @description API tài liệu cho FoodShop
// @host localhost:8081
// @BasePath /api/v1

func (controller *Controller) MapFoodApiV1(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/api/v1")
	api.GET("/foods", controller.HandleGetFoods)
	api.GET("/foods/:id", controller.HandleGetFoodById)
	api.POST("/foods", controller.HandleCreateFood)

}

// HandleGetFoods godoc
// @Summary Lấy danh sách món ăn
// @Description Trả về danh sách các món ăn
// @Tags foods
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /foods [get]
func (controller *Controller) HandleGetFoods(context echo.Context) error {
	slog.Info("Request start", "http_method", "HandleGetFoods", "url", context.Request().RequestURI)
	data, err := controller.usecase.GetFoodsShop(context)
	if err != nil {
		return err
	}
	return context.JSON(200, data)
}

// HandleGetFoodById godoc
// @Summary Lấy thông tin món ăn theo id
// @Description Trả về thông tin món ăn theo id
// @Tags foods
// @Accept json
// @Param id path int true "ID của món ăn"
// @Produce json
// @Success 200 {object} int
// @Router /foods/{id} [get]
func (controller *Controller) HandleGetFoodById(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return err
	}
	data, err := controller.usecase.GetFoodById(context, id)
	if err != nil {
		return err
	}
	return context.JSON(200, data)
}

// HandleCreateFood godoc
// @Summary Tạo mới món ăn
// @Description Tạo món ăn mới
// @Tags foods
// @Accept json
// @Param food body domain.Product true "Tạo món ăn thành công"
// @Produce json
// @Success 200 {object} int
// @Router /foods [post]
func (controller *Controller) HandleCreateFood(context echo.Context) error {
	food := &domain.Product{}
	err := context.Bind(food)
	if err != nil {
		return err
	}
	food, err = controller.usecase.AddFood(context, food)
	if err != nil {
		return err
	}
	return context.JSON(200, food)
}
