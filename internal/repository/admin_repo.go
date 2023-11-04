package repository

import (
	"errors"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.IAdminRepository {
	return &AdminRepository{DB: DB}
}

func (a *AdminRepository) FindByEmail(email string) (*entities.Admin, error) {
	var adminDetails entities.Admin
	query := a.DB.Raw(`
		SELECT * FROM admins WHERE email = ?`,
		email).Scan(&adminDetails)

	if query.RowsAffected == 0 {
		return nil, errors.New("no admin registered with this email")
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return &adminDetails, nil
}
