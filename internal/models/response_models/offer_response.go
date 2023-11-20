package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type OfferListRes struct {
	Status    string                   `json:"status"`
	Message   string                   `json:"message"`
	OfferList []entities.CategoryOffer `json:"offerList"`
}
