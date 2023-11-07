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
func AuthorizeAdmin(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing Admin")

	tokenString := StripToken(c.Get("Authorization"))

	var secretKey = viper.GetString("KEY")

	// Check if it is admin
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}

	admin := claims.(*jwttoken.CustomClaims).Model

	c.Locals("AdminModel", admin)

	fmt.Println("MW: Admin Authorised")
	return c.Next()
}

// Authorize restaurant
func AuthorizeRestaurant(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing Restaurant")

	tokenString := StripToken(c.Get("Authorization"))

	var secretKey = viper.GetString("KEY")

	// Check if it is restaurant
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}
	restaurant := claims.(*jwttoken.CustomClaims).Model

	c.Locals("RestaurantModel", restaurant)

	fmt.Println("MW: Restaurant Authorised")
	return c.Next()
}

// Authorize user
func AuthorizeUser(c *fiber.Ctx) error {
	fmt.Println("MW: Authorizing User")

	tokenString := StripToken(c.Get("Authorization"))

	var secretKey = viper.GetString("KEY")

	// Check if it is user
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).
			JSON(res.UnauthorizedAccess)
	}
	user := claims.(*jwttoken.CustomClaims).Model

	c.Locals("UserModel", user)

	fmt.Println(c.Locals("UserModel"))
	fmt.Println("MW: User Authorised")
	return c.Next()
}

func StripToken(tokenHeader string) string {
	if tokenHeader == "" {
		return ""
	}
	token := strings.Split(tokenHeader, " ")
	if len(token) != 2 {
		return ""
	}
	fmt.Println("Token is", token)
	return token[1]
}
