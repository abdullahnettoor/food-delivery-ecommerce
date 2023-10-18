package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Role  string
	Model interface{}
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("KEY"))

func CreateToken(c *fiber.Ctx, role string, expireAfter time.Duration, userModel interface{}) (string, error) {

	// Create the Custom Claims
	claims := &CustomClaims{
		role,
		userModel,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireAfter)), // Token expires in 24 hours
		},
	}

	// Generate token based on claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Retrieving token string
	tokenString, err := token.SignedString(secretKey)
	fmt.Printf("%v %v", tokenString, err)
	if err != nil {
		fmt.Println("Error occured while creating token:", err)
		return "", err
	}

	// Set user values to fiber context
	c.Locals(role+"Model", claims.Model)
	c.Locals("role", claims.Role)

	fmt.Println(role, "Model is : ", c.Locals(role+"Model"))
	return tokenString, nil
}

// Validate Token
func IsValidToken(tokenString string, c *fiber.Ctx) (bool, interface{}) {

	// Parse jwt token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("Error occured whilr fetching token")
		return false, nil
	}

	// Assign parsed data from token to calims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return false, nil
		}

		// Set user values to fiber Ctx
		c.Locals(claims.Role+"Model", claims.Model)
		c.Locals("role", claims.Role)

		return true, claims

	} else {
		fmt.Println("Error occured while parsing token:", err)
		return false, nil
	}
}
