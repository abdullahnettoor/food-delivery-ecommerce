package interfaces

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type ICouponRepository interface {
	Create(coupon *entities.Coupon) error
	Update(id string, coupon *entities.Coupon) error
	UpdateStatus(id, status string) error
	Delete(id string) error
	Find(id string) (*entities.Coupon, error)
	FindAll() (*[]entities.Coupon, error)
	FindByCode(code string) (*entities.Coupon, error)

	FindAllForUser() (*[]entities.Coupon, error)
	FindAllAvailableForUser(userId string) (*[]entities.Coupon, error) 

	CreateRedeemed(userId, code string) error
	FindRedeemed(userId, code string) (*entities.RedeemedCoupon, error)
	FindRedeemedByUser(userId string) (*[]entities.RedeemedCoupon, error)
}
