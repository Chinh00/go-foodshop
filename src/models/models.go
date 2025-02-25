package models

type (
	Food struct {
		FoodId int
		Name   string
		Price  float64
	}
	Category struct {
		CategoryId int
		Name       string
	}
)
