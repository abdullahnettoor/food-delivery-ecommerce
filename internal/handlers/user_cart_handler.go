package handlers

import (
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/initializers"
	"github.com/abdullahnettoor/food-delivery-ecommerce/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddToCart(c *fiber.Ctx) error {
	params := c.Params("id")
	dish := models.Dish{}
	dbCart := models.Cart{}
	dbCartItems := []models.CartItem{}
	user := c.Locals("UserModel").(map[string]any)

	dishId, err := uuid.Parse(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "UUID Error",
			"error":   err,
		})
	}

	result := initializers.DB.Raw(`
		SELECT 
		*
		FROM dishes
		WHERE id = ? AND
		deleted_at IS NULL AND
		quantity > 0 AND
		availability = true`,
		dishId).Scan(&dish)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Selected dish is not available",
		})
	}

	userId, err := uuid.Parse(user["userId"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "UUID Error",
			"error":   err,
		})
	}

	result = initializers.DB.Raw(`SELECT * FROM carts WHERE id = ?`, userId).Scan(&dbCart)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}
	if result.RowsAffected != 0 && dbCart.RestaurantID != dish.RestaurantID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed!",
			"message": "You can't add dishes from different restaurant at a time",
		})
	}
	initializers.DB.Save(models.Cart{ID: userId, RestaurantID: dish.RestaurantID}).Scan(&dbCart)

	result = initializers.DB.Raw(`
	SELECT * FROM cart_items WHERE cart_id = ?`,
		dbCart.ID).Scan(&dbCartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}
	for _, d := range dbCartItems {
		if d.DishID == dishId {
			d.Quantity = d.Quantity + 1
			initializers.DB.Save(&d)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":   "success",
				"message":  "Dish quantity incremented in cart",
				"cartItem": d,
				"user":     c.Locals("UserModel"),
			})
		}
	}

	cartItem := models.CartItem{
		CartID:   dbCart.ID,
		DishID:   dish.ID,
		Quantity: 1,
	}

	result = initializers.DB.Create(&cartItem)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"message":  "Dish added to cart",
		"cartItem": cartItem,
		"user":     c.Locals("UserModel"),
	})
}

func ViewCart(c *fiber.Ctx) error {
	cartDishes := []struct {
		models.Dish
		Quantity uint `json:"dishQuantity" gorm:"notNull"`
	}{}
	user := c.Locals("UserModel").(map[string]any)
	userId, err := uuid.Parse(user["userId"].(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "UUID Error",
			"error":   err,
		})
	}

	result := initializers.DB.Raw(`
		SELECT 
		d.*,
		c.quantity
		FROM cart_items c 
		INNER JOIN dishes d 
		ON d.id = c.dish_Id 
		WHERE c.cart_id = ?`,
		userId).Scan(&cartDishes)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   err,
		})
	}
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"cart":   "Cart is Empty"})
	}

	var totalPrice float64
	for _, dish := range cartDishes {
		totalPrice += (float64(dish.Price) * float64(dish.Quantity))
	}

	cart := map[string]any{
		"id":           userId,
		"restaurantId": cartDishes[0].RestaurantID,
		"dishes":       cartDishes,
		"totalPrice":   totalPrice,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"cart":   cart,
		"user":   c.Locals("UserModel"),
	})
}

func DeleteCartItem(c *fiber.Ctx) error {
	dishId := c.Params("id")
	user := c.Locals("UserModel").(map[string]any)

	result := initializers.DB.Exec(`DELETE FROM cart_items WHERE cart_id = ? AND dish_id = ?`, user["userId"], dishId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "DB Error",
			"error":   result.Error,
		})
	}
	result = initializers.DB.Exec(`
	DELETE FROM carts
	WHERE NOT EXISTS (
		SELECT 1 FROM cart_items
		WHERE cart_items.cart_id = ?
		)
	AND carts.id = ?`,
		user["userId"], user["userId"])
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "DB Error",
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Dish deleted from cart successfully",
		"user":    user,
	})
}

func DecrementCartItem(c *fiber.Ctx) error {
	dishId := c.Params("id")
	dbCartItem := models.CartItem{}
	user := c.Locals("UserModel").(map[string]any)

	result := initializers.DB.Raw(`
		SELECT * 
		FROM cart_items 
		WHERE cart_id = ? 
		AND dish_id = ?`,
		user["userId"], dishId).Scan(&dbCartItem)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed!",
			"message": "DB Error",
			"error":   result.Error,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "failed!",
			"message": "Item already deleted",
		})
	}

	if dbCartItem.Quantity == 1 {
		return DeleteCartItem(c)
	}

	dbCartItem.Quantity = dbCartItem.Quantity - 1
	initializers.DB.Save(&dbCartItem)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "success",
		"message":  "Dish quantity reduced from cart successfully",
		"cartItem": dbCartItem,
		"user":     user,
	})
}
