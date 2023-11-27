package repository

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository/interfaces"
	"gorm.io/gorm"
)

type couponRepo struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.ICouponRepository {
	return &couponRepo{db}
}

func (repo *couponRepo) Create(coupon *entities.Coupon) error {
	return repo.DB.Create(coupon).Error
}

func (repo *couponRepo) Update(id string, coupon *entities.Coupon) error {
	return repo.DB.Save(&coupon).Error
}

func (repo *couponRepo) UpdateStatus(id, status string) error {
	res := repo.DB.Exec(`
	UPDATE coupons
	SET status = ?
	WHERE id =?`,
		status, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return e.ErrNotFound
	}
	return nil
}

func (repo *couponRepo) Delete(id string) error {
	return repo.DB.Delete(&entities.Coupon{}, id).Error
}

func (repo *couponRepo) Find(id string) (*entities.Coupon, error) {
	var coupon entities.Coupon
	res := repo.DB.First(&coupon, id)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}
	return &coupon, nil
}

func (repo *couponRepo) FindAll() (*[]entities.Coupon, error) {
	var couponList []entities.Coupon
	res := repo.DB.Raw(`
	SELECT * 
	FROM coupons
	WHERE status <> 'DELETED'`).Scan(&couponList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
	}
	return &couponList, nil
}

func (repo *couponRepo) FindAllForUser() (*[]entities.Coupon, error) {
	var couponList []entities.Coupon
	res := repo.DB.Raw(`
	SELECT * 
	FROM coupons
	WHERE status = 'ACTIVE'`).Scan(&couponList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
	}
	return &couponList, nil
}

func (repo *couponRepo) FindAllAvailableForUser(userId string) (*[]entities.Coupon, error) {
	var couponList []entities.Coupon
	res := repo.DB.Raw(`
	SELECT * 
	FROM coupons
	WHERE status = 'ACTIVE'
	AND code NOT IN (
		SELECT coupon_code 
		FROM redeemed_coupons
		WHERE user_id = ?)`, userId).Scan(&couponList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
	}
	return &couponList, nil
}

func (repo *couponRepo) FindByCode(code string) (*entities.Coupon, error) {
	var coupon entities.Coupon
	res := repo.DB.Raw(`
	SELECT *
	FROM coupons
	WHERE code = ?
	AND status <> 'DELETED'`,
		code).Scan(&coupon)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}
	return &coupon, nil
}

func (repo *couponRepo) CreateRedeemed(userId, code string) error {
	return repo.DB.Exec(`
	INSERT INTO redeemed_coupons
	(user_id, coupon_code)
	VALUES 
	(?,?)`,
		userId, code).Error
}

func (repo *couponRepo) FindRedeemed(userId, code string) (*entities.RedeemedCoupon, error) {
	var redeemedCoupon entities.RedeemedCoupon
	res := repo.DB.Raw(`
	SELECT *
	FROM redeemed_coupons
	WHERE user_id = ?
	AND coupon_code = ?`,
		userId, code).Scan(&redeemedCoupon)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}
	return &redeemedCoupon, nil
}

func (repo *couponRepo) FindRedeemedByUser(userId string) (*[]entities.RedeemedCoupon, error) {
	var redeemedCoupons []entities.RedeemedCoupon
	res := repo.DB.Raw(`
	SELECT * 
	FROM redeemed_coupons
	WHERE user_id = ?`,
		userId).Scan(&redeemedCoupons)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, e.ErrIsEmpty
	}
	return &redeemedCoupons, nil
}
