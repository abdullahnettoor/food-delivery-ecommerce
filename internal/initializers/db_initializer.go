package initializers

import (
	"os"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DB_URI")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(`Error connecting to DB`)
	}

}

func SyncDatabase() {
	DB.AutoMigrate(&models.Admin{}, &models.User{}, &models.Restaurant{}, &models.Dish{})
}
