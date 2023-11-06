package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type IUserUseCase interface {
	Login(user entities.User) error
}
