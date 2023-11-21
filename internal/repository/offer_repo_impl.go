package repository

import (
	"strings"
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

	query := "UPDATE category_offers SET status = '" + status

	if strings.ToUpper(status) == "CLOSED" {
		query += "', end_date = '" + time.Now().Format("2006-01-02 15:04:05.999999-07:00")
	}

	query += "' WHERE id = " + id

	res := repo.DB.Exec(query)
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
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
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

func (repo *offerRepository) FindBySellerAndCategory(sellerId, categoryId string) (*entities.CategoryOffer, error) {
	var offer entities.CategoryOffer
	res := repo.DB.Raw(`
	SELECT * 
	FROM category_offers
	WHERE seller_id = ?
	AND category_id = ?`,
		sellerId, categoryId).Scan(&offer)

	if res.Error != nil {
		return nil, res.Error
	}

	return &offer, nil
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
