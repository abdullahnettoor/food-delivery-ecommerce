package middlewares

import (
	"fmt"
	"strings"

	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// Authourize admin
func AuthenticateAdmin(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing Admin")

	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("KEY")

	// Check if it is admin
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}

	admin := claims.(*jwttoken.CustomClaims).Model
	role := claims.(*jwttoken.CustomClaims).Role
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).
			JSON(res.UnauthorizedAccess)
	}

	c.Locals("AdminModel", admin)

	fmt.Println("MW: Admin Authorised")
	return c.Next()
}

// Authenticate seller
func AuthenticateSeller(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing Seller")

	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("KEY")

	// Check if it is seller
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}
	seller := claims.(*jwttoken.CustomClaims).Model
	role := claims.(*jwttoken.CustomClaims).Role
	if role != "seller" {
		return c.Status(fiber.StatusForbidden).
			JSON(res.UnauthorizedAccess)
	}

	fmt.Println("Seller is", seller)

	c.Locals("SellerModel", seller)

	fmt.Println("MW: Seller Authorised")
	return c.Next()
}

// Authenticate user
func AuthenticateUser(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing User")

	tokenString := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	var secretKey = viper.GetString("KEY")

	// Check if it is user
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}
	user := claims.(*jwttoken.CustomClaims).Model
	role := claims.(*jwttoken.CustomClaims).Role
	if role != "user" {
		return c.Status(fiber.StatusForbidden).
			JSON(res.UnauthorizedAccess)
	}

	c.Locals("UserModel", user)

	fmt.Println(c.Locals("UserModel"))
	fmt.Println("MW: User Authorised")
	return c.Next()
}
