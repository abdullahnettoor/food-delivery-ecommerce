package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type SellerLoginRes struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type SellerDishListRes struct {
	Status   string          `json:"status"`
	Message  string          `json:"message"`
	DishList []entities.Dish `json:"dishList"`
}
