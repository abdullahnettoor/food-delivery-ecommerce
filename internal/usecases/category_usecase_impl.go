package usecases

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	i "github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
)

type categoryUcase struct {
	CategoryRepo interfaces.ICategoryRepository
}

func NewCategoryUsecase(CategoryRepo interfaces.ICategoryRepository) i.ICategoryUseCase {
	return &categoryUcase{CategoryRepo}
}

func (uc *categoryUcase) GetCategory(id string) (*entities.Category, error) {
	return uc.CategoryRepo.FindByID(id)
}

func (uc *categoryUcase) GetAllCategory() (*[]entities.Category, error) {
	return uc.CategoryRepo.FindAll()
}
