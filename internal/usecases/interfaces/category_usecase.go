package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ICategoryUseCase interface {
	GetCategory(id string) (*entities.Category, error)
	GetAllCategory() (*[]entities.Category, error)
}
