package db

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(c *config.DbConfig) (*gorm.DB, error) {

	// dbUri = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.User, c.Password, c.Name, c.Port)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := syncDatabase(db); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to DB")
	return db, nil
}

func syncDatabase(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.Admin{},
		&entities.Seller{},
	// &models.Admin{},
	// &models.User{},
	// &models.Address{},
	// &models.Restaurant{},
	// &models.Dish{},
	// &models.Category{},
	// &models.Cart{},
	// &models.CartItem{},
	// &models.Order{},
	// &models.OrderItem{},
	)
}
