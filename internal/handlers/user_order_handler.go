package handlers

import (
	"fmt"

	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PlaceOrder(c *fiber.Ctx) error {
	body := struct {
		PayementMethod string    `json:"paymentMethod"`
		AddressID      uuid.UUID `json:"adddressId"`
	}{}

	c.BodyParser(&body)
	user := c.Locals("UserModel").(map[string]any)
	userId, _ := uuid.Parse(user["userId"].(string))

	cartItems := []struct {
		CartID       uuid.UUID `json:"cartId" gorm:"type:uuid;notNull"`
		DishID       uuid.UUID `json:"dishId" gorm:"type:uuid;"`
		Quantity     uint      `json:"dishQuantity" gorm:"type:uint"`
		Stock        uint      `json:"dishStock" gorm:"type:uint"`
		RestaurantID uuid.UUID `json:"restaurantId" gorm:"type:uuid;foreignKey:restaurant.id;notNull"`
		Name         string    `json:"dishName" gorm:"notNull"`
		Description  string    `json:"dishDescription"`
		Price        float64   `json:"dishPrice" gorm:"notNull"`
		Category     uint      `json:"category" gorm:"type:bigInt;foreignKey:category.id;notNull"`
		IsVeg        bool      `json:"isVeg" gorm:"default:false"`
		Availability bool      `json:"isAvailable" gorm:"type:boolean;default:true"`
	}{}

	result := initializers.DB.Raw(`
		SELECT
		c.cart_id,
		c.dish_id,
		c.quantity,
		d.quantity as stock,
		d.restaurant_id,
		d.name,
		d.description,
		d.price,
		d.is_veg,
		d.category,
		d.availability
		FROM cart_items c
		INNER JOIN dishes d
		ON d.id = c.dish_id
		WHERE cart_id = ?`,
		userId).Scan(&cartItems)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed!",
			"message": "Couldn't Place the order, Cart is Empty",
		})
	}

	if body.PayementMethod != "COD" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed!",
			"message": "Payment method must be COD",
		})
	}

	orderId, _ := uuid.NewUUID()

	order := models.Order{
		ID:             orderId,
		RestaurantID:   cartItems[0].RestaurantID,
		AddressID:      body.AddressID,
		PaymentMethod:  body.PayementMethod,
		UserID:         userId,
		ItemCount:      uint(len(cartItems)),
		Status:         "Order Placed",
		PayementStatus: "Success",
	}
	var totalPrice float64
	orderItems := []models.OrderItem{}
	for _, dish := range cartItems {
		item := models.OrderItem{
			OrderID:  orderId,
			DishID:   dish.DishID,
			Quantity: dish.Quantity,
			Price:    dish.Price,
		}
		totalPrice += dish.Price * float64(dish.Quantity)
		orderItems = append(orderItems, item)

		fmt.Println("Stock", dish.Stock, "Quantity", dish.Quantity)

		err := initializers.DB.Exec(`
		UPDATE dishes
		SET quantity = ?
		WHERE id = ?
		`, dish.Stock-dish.Quantity, dish.DishID).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "failed!",
				"message": "DB Error",
				"error":   err,
			})
		}
	}

	// Discount is not applied yet, It'll be added soon
	order.TotalPrice = totalPrice

	err := initializers.DB.Create(&order).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   err,
		})
	}
	err = initializers.DB.Create(&orderItems).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   err,
		})
	}

	err = initializers.DB.Exec(`
	DELETE FROM carts
	WHERE id = ?`,
		userId).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":     "success",
		"user":       user,
		"message":    "Order Placed Successfully",
		"order":      order,
		"orderItems": orderItems,
	})
}
