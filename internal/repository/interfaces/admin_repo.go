package interfaces

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
)

type IAdminRepository interface {
	FindByEmail(email string) (*entities.Admin, error)
}
