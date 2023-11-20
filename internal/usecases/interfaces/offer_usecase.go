package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
)

type IOfferUseCase interface {
	CreateOffer(sellerId string, offer *req.CreateOfferReq) error
	UpdateOffer(id, sellerId string, offer *req.UpdateOfferReq) error
	UpdateOfferStatus(id, status string) error
	GetAllOffer() (*[]entities.CategoryOffer, error)
	GetOffersBySeller(sellerId string) (*[]entities.CategoryOffer, error)
}
