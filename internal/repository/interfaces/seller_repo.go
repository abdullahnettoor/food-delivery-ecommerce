package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type ISellerRepository interface {
	FindAll() (*[]entities.Seller, error)
	FindByID(id string) (*entities.Seller, error)
	FindByEmail(email string) (*entities.Seller, error)
	Create(seller *entities.Seller) error
	Verify(id string) error
	Block(id string) error
	Unblock(id string) error
	// TODO: FindByQuery(query string) ([]*entities.Seller, error)
	// TODO: UpdateByID(id string, seller *entities.Seller) (*entities.Seller, error)
	// TODO: Delete(seller *entities.Seller) error
}
