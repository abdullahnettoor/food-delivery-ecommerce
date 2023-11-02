package interfaces

import (
	"context"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type ISellerRepository interface {
	FindAll(c context.Context) ([]entities.Seller, error)
	FindByID(c context.Context, id string) (entities.Seller, error)
	Create(c context.Context, seller entities.Seller) (entities.Seller, error)
	UpdateByID(c context.Context, id string, seller entities.Seller) (entities.Seller, error)
	Delete(c context.Context, seller entities.Seller) error
	FindByQuery(c context.Context, query string) ([]entities.Seller, error)
}
