package repository

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type favouriteRepository struct {
	DB *gorm.DB
}

func NewFavouriteRepository(db *gorm.DB) interfaces.IFavoriteRepository {
	return &favouriteRepository{db}
}

func (repo *favouriteRepository) Create(userId, dishId string, item *entities.Favourite) error {
	var favItem entities.Favourite
	res := repo.DB.Raw(`
	SELECT * 
	FROM favourites
	WHERE dish_id = ?
	AND user_id = ?
	`,
		dishId, userId).Scan(&favItem)
	if res.RowsAffected != 0 {
		return e.ErrConflict
	}
	if res.Error != nil {
		return res.Error
	}

	return repo.DB.Create(&item).Error
}

func (repo *favouriteRepository) FindByUserId(userId string) (*[]entities.Favourite, error) {
	var favList []entities.Favourite
	res := repo.DB.Raw(`
	SELECT * 
	FROM favourites
	WHERE user_id = ?
	ORDER BY id DESC
	`,
		userId).Scan(&favList)
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &favList, nil
}

func (repo *favouriteRepository) Delete(userId, dishId string) error {
	res := repo.DB.Exec(`
	DELETE FROM favourites
	WHERE user_id = ?
	AND dish_id = ?`,
		userId, dishId)
	if res.RowsAffected == 0 {
		return e.ErrNotFound
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}
