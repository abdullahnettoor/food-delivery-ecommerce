package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type IOfferRepository interface {
	Create(offer *entities.CategoryOffer) error
	Update(id string, offer *entities.CategoryOffer) error
	UpdateStatus(id, status string) error
	FindAll() (*[]entities.CategoryOffer, error)
	FindByID(id string) (*entities.CategoryOffer, error)
	FindAllForSeller(sellerId string) (*[]entities.CategoryOffer, error)
	FindBySellerAndCategory(sellerId, categoryId string) (*entities.CategoryOffer, error)
}
