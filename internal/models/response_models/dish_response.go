package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type DishListRes struct {
	Status   string          `json:"status"`
	Message  string          `json:"message"`
	DishList []entities.Dish `json:"dishList"`
}

type SingleDishRes struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Dish    entities.Dish `json:"dish"`
}
