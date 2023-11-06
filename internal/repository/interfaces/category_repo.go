package interfaces

import (
	"context"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type ICategoryRepository interface {
	FindByID(c context.Context, id string) (entities.Category, error)
	FindAll(c context.Context) ([]entities.Category, error)
	Create(c context.Context, category entities.Category) (entities.Category, error)
}
