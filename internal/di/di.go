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
	favRepo := repository.NewFavouriteRepository(gormDB)
	orderRepo := repository.NewOrderRepository(gormDB)
	offerRepo := repository.NewOfferRepository(gormDB)
	couponRepo := repository.NewCouponRepository(gormDB)

	adminUcase := usecases.NewAdminUsecase(adminRepo, userRepo, sellerRepo, categoryRepo)
	sellerUcase := usecases.NewSellerUsecase(sellerRepo)
	userUcase := usecases.NewUserUsecase(userRepo, sellerRepo)
	categoryUcase := usecases.NewCategoryUsecase(categoryRepo)
	offerUcase := usecases.NewOfferUsecase(offerRepo)
	dishUcase := usecases.NewDishUsecase(dishRepo, offerUcase)
	favUcase := usecases.NewFavouriteUsecase(favRepo, dishUcase)
	cartUcase := usecases.NewCartUsecase(cartRepo, dishUcase)
	orderUcase := usecases.NewOrderUsecase(cartRepo, orderRepo, dishUcase, couponRepo)

	sellerHandler := handlers.NewSellerHandler(sellerUcase, dishUcase)
	userHandler := handlers.NewUserHandler(userUcase)
	adminHandler := handlers.NewAdminHandler(adminUcase)
	categoryHandler := handlers.NewCategoryHandler(categoryUcase)
	dishHandler := handlers.NewDishHandler(dishUcase)
	favHandler := handlers.NewFavHandler(favUcase)
	cartHandler := handlers.NewCartHandler(cartUcase)
	orderHandler := handlers.NewOrderHandler(orderUcase)
	offerHandler := handlers.NewOfferHandler(offerUcase)

	serverHttp := api.NewServerHttp(
		adminHandler,
		sellerHandler,
		userHandler,
		categoryHandler,
		dishHandler,
		cartHandler,
		orderHandler,
		offerHandler,
		favHandler,
	)

	return serverHttp, nil
}
