package handlers

import (
	"fmt"
	"time"

	req "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/request_models"
	res "github.com/abdullahnettoor/food-delivery-eCommerce/internal/models/response_models"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/usecases/interfaces"
	jwttoken "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/jwt_token"
	requestvalidation "github.com/abdullahnettoor/food-delivery-eCommerce/pkg/request_validation"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type UserHandler struct {
	usecase interfaces.IUserUseCase
}

func NewUserHandler(uCase interfaces.IUserUseCase) *UserHandler {
	return &UserHandler{uCase}
}

func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	var signUpReq req.UserSignUpReq

	if err := c.BodyParser(&signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	user, err := h.usecase.SignUp(&signUpReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	secret := viper.GetString("KEY")
	fmt.Println("Key is", secret)
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserLoginRes{
			Status: "success",
			Token:  token,
		})
}

func (h *UserHandler) SendOtp(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)

	if err := h.usecase.SendOtp(user["phone"].(string)); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status":  "success",
			"message": "otp sent successfully",
		})
}

func (h *UserHandler) VerifyOtp(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)

	var req req.UserVerifyOtpReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	err := h.usecase.VerifyOtp(user["phone"].(string), &req)
	fmt.Println("Error is", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"status":  "success",
			"message": "otp verified successfully",
		})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginReq req.UserLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}
	if err := requestvalidation.ValidateRequest(loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err,
			})
	}

	user, err := h.usecase.Login(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	secret := viper.GetString("KEY")
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserLoginRes{
			Status: "success",
			Token:  token,
		})
}

func (h *UserHandler) GetDishesPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l", "5")

	dishList, err := h.usecase.GetDishesPage(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dishes",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.DishListRes{
			Status:   "success",
			Message:  "successfully fetched dish",
			DishList: *dishList,
		})

}

func (h *UserHandler) GetDish(c *fiber.Ctx) error {
	dishId := c.Params("id")

	dish, err := h.usecase.GetDish(dishId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dish",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SingleDishRes{
			Status:  "success",
			Message: "successfully fetched dish",
			Dish:    *dish,
		})
}

func (h *UserHandler) SearchDish(c *fiber.Ctx) error {
	searchQuery := c.Query("q")

	dishList, err := h.usecase.SearchDish(searchQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch dish",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.DishListRes{
			Status:   "success",
			Message:  "successfully fetched dish",
			DishList: *dishList,
		})
}
