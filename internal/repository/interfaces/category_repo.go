package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type ICategoryRepository interface {
	FindByID(id string) (*entities.Category, error)
	FindAll() (*[]entities.Category, error)
	Create(category *entities.Category) error
	Update(id, name string) error
}
