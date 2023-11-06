package interfaces

import (
	"context"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IDishRepository interface {
	FindAll(c context.Context) ([]entities.Dish, error)
	FindByID(c context.Context, id string) (entities.Dish, error)
	Create(c context.Context, dish entities.Dish) error
	Update(c context.Context, id string, dish entities.Dish) (entities.Dish, error)
	Delete(c context.Context, id string) error
}
