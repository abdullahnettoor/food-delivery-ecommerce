package di

import (
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/config"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/api"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/db"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/infrastructure/handlers"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/repository"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases"
	cld "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/cloudinary"
)

func InitializeAPI(c *config.DbConfig, imgUploaderCfg *config.ImgUploaderCfg) (*api.ServerHttp, error) {

	gormDB, err := db.ConnectPostgres(c)
	if err != nil {
		return nil, err
	}

	err = cld.ConnectCloudinary(imgUploaderCfg)
	if err != nil {
		return nil, err
	}

	adminRepo := repository.NewAdminRepository(gormDB)
	categoryRepo := repository.NewCategoryRepository(gormDB)
	dishRepo := repository.NewDishRepository(gormDB)
	sellerRepo := repository.NewSellerRepository(gormDB)
	userRepo := repository.NewUserRepository(gormDB)
	cartRepo := repository.NewCartRepository(gormDB)
	orderRepo := repository.NewOrderRepository(gormDB)

	adminUcase := usecases.NewAdminUsecase(adminRepo, userRepo, sellerRepo, categoryRepo)
	sellerUcase := usecases.NewSellerUsecase(sellerRepo, dishRepo)
	userUcase := usecases.NewUserUsecase(userRepo, dishRepo, sellerRepo)
	dishUcase := usecases.NewDishUsecase(dishRepo, categoryRepo)
	cartUcase := usecases.NewCartUsecase(cartRepo, dishRepo)
	orderUcase := usecases.NewOrderUsecase(cartRepo, orderRepo, dishRepo)

	sellerHandler := handlers.NewSellerHandler(sellerUcase, dishUcase)
	userHandler := handlers.NewUserHandler(userUcase, dishUcase)
	adminHandler := handlers.NewAdminHandler(adminUcase)
	cartHandler := handlers.NewCartHandler(cartUcase)
	orderHandler := handlers.NewOrderHandler(orderUcase)

	serverHttp := api.NewServerHttp(adminHandler, sellerHandler, userHandler, cartHandler, orderHandler)

	return serverHttp, nil
}
