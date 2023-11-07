package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IUserRepository interface {
	FindAll() (*[]entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(id string) (*entities.User, error)
	Create(user *entities.User) (*entities.User, error)
	Update(id string, user *entities.User) (*entities.User, error)
	Verify(phone string) error
	Block(id string) error
	Unblock(id string) error
	DeleteByPhone(phone string) error
}
