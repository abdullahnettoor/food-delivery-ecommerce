package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Name  string
	Email string
	Role  string
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("KEY"))

// func CreateToken(name, email, role string, expiryDuration time.Duration) (ts string, err error) {
// 	claims := &CustomClaims{
// 		name,
// 		email,
// 		role,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryDuration)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

// 	ts, err = token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		fmt.Println(err, ": Inside JWT")
// 		return "", err
// 	}
// 	return ts, nil
// }

func CreateToken(c *fiber.Ctx, name, email, role string, expireAfter time.Duration) (string, error) {

	// Create the Custom Claims
	claims := &CustomClaims{
		name,
		email,
		role,
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
	c.Locals("name", claims.Name)
	c.Locals("email", claims.Email)
	c.Locals("role", claims.Role)
	return tokenString, nil
}

// Validate Token
func IsValidToken(tokenString string, c *fiber.Ctx) bool {

	// Parse jwt token with custom claims
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		fmt.Println("Error occured whilr fetching token")
		return false
	}

	// Assign parsed data from token to calims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		// Check if token is expired
		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return false
		}

		// Set user values to fiber Ctx
		c.Locals("userEmail", claims.Email)
		c.Locals("username", claims.Name)
		c.Locals("role", claims.Role)

		return true

	} else {
		fmt.Println("Error occured while parsing token:", err)
		return false
	}
}
