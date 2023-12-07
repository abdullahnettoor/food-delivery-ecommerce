package repository

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(DB *gorm.DB) interfaces.ISellerRepository {
	return &SellerRepository{DB: DB}
}

func (repo *SellerRepository) FindAll() (*[]entities.Seller, error) {
	var sellerList []entities.Seller

	if err := repo.DB.Raw(`
	SELECT * 
	FROM sellers
	WHERE status <> 'Deleted'`).
		Scan(&sellerList).Error; err != nil {
		return nil, err
	}

	return &sellerList, nil
}

func (repo *SellerRepository) FindByID(id string) (*entities.Seller, error) {
	var seller entities.Seller

	res := repo.DB.Raw(`
	SELECT *
	FROM sellers
	WHERE id = ?`, id).
		Scan(&seller)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &seller, nil
}

func (repo *SellerRepository) FindByEmail(email string) (*entities.Seller, error) {
	var seller entities.Seller

	res := repo.DB.Raw(`
	SELECT *
	FROM sellers
	WHERE email = ?`, email).
		Scan(&seller)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &seller, nil
}

func (repo *SellerRepository) Create(seller *entities.Seller) error {
	query := repo.DB.Raw(`
	SELECT *
	FROM sellers 
	WHERE email = ?`, seller.Email)
	if query.Error != nil {
		return query.Error
	}
	if query.RowsAffected > 0 {
		return e.ErrConflict
	}
	if err := repo.DB.Create(&seller).Error; err != nil {
		return err
	}
	return nil
}

func (repo *SellerRepository) Verify(id string) error {
	err := repo.DB.Exec(`
	UPDATE sellers
	SET status = 'Verified'
	WHERE id = ?`, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *SellerRepository) Block(id string) error {
	err := repo.DB.Exec(`
	UPDATE sellers
	SET status = 'Blocked'
	WHERE id = ?`, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *SellerRepository) Unblock(id string) error {
	err := repo.DB.Exec(`
	UPDATE sellers
	SET status = 'Verified'
	WHERE id = ?`, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *SellerRepository) SearchByStatus(search, status string) (*[]entities.Seller, error) {
	var sellersList []entities.Seller

	query := fmt.Sprintf("SELECT * FROM sellers WHERE ((name ILIKE '%%%s%%') OR (description ILIKE '%%%s%%')) AND status ILIKE '%%%s%%'", search, search, status)

	res := repo.DB.Raw(query).Scan(&sellersList)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &sellersList, nil
}

func (repo *SellerRepository) FindPageWise(page, limit uint) (*[]entities.Seller, error) {
	var sellersList []entities.Seller

	offset := (page - 1) * uint(limit)

	query := `
	SELECT *
	FROM sellers
	WHERE status ILIKE 'verified'
	OFFSET ? LIMIT ?`

	res := repo.DB.Raw(query, offset, limit).Scan(&sellersList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &sellersList, nil
}

func (repo *SellerRepository) FindVerifiedByID(id string) (*entities.Seller, error) {
	var seller entities.Seller

	res := repo.DB.Raw(`
	SELECT *
	FROM sellers
	WHERE id = ?
	AND status ILIKE 'verified'`, id).
		Scan(&seller)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &seller, nil
}
