package repository

import (
	"time"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) interfaces.IOfferRepository {
	return &offerRepository{db}
}

func (repo *offerRepository) Create(offer *entities.CategoryOffer) error {
	return repo.DB.Create(offer).Error
}

func (repo *offerRepository) Update(id string, offer *entities.CategoryOffer) error {
	res := repo.DB.Save(&offer)
	if res.RowsAffected == 0 {
		return e.ErrNotFound
	}

	return res.Error
}

func (repo *offerRepository) UpdateStatus(id, status string) error {
	res := repo.DB.Exec(`
	UPDATE category_offers
	SET status = ?
	WHERE id = ?`,
		status, id)
	if res.RowsAffected == 0 {
		return e.ErrNotFound
	}

	return res.Error
}

func (repo *offerRepository) FindAll() (*[]entities.CategoryOffer, error) {
	var offerList []entities.CategoryOffer
	res := repo.DB.Raw(`
	SELECT * 
	FROM category_offers
	WHERE start_date < ?
	AND end_date > ?`,
		time.Now(), time.Now()).Scan(&offerList)
	if res.Error != nil {
		return nil, res.Error
	}

	return &offerList, nil
}

func (repo *offerRepository) FindAllForSeller(sellerId string) (*[]entities.CategoryOffer, error) {
	var offerList []entities.CategoryOffer
	res := repo.DB.Raw(`
	SELECT * 
	FROM category_offers
	WHERE seller_id = ?`,
		sellerId).Scan(&offerList)

	if res.Error != nil {
		return nil, res.Error
	}

	return &offerList, nil
}

func (repo *offerRepository) FindByID(id string) (*entities.CategoryOffer, error) {
	var offer entities.CategoryOffer
	res := repo.DB.Raw(`
	SELECT * 
	FROM category_offers
	WHERE id = ?`,
		id).Scan(&offer)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	return &offer, nil
}
