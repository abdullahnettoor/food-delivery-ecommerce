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

// @Summary		Sign up as a user
// @Description	Register a new user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.UserSignUpReq	true	"User sign-up request"
// @Success		200	{object}	res.UserLoginRes	"Successfully signed up"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/signup [post]
func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	var signUpReq req.UserSignUpReq

	if err := c.BodyParser(&signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(signUpReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	user, err := h.usecase.SignUp(&signUpReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to signup user",
				Error:   err.Error(),
			})
	}

	secret := viper.GetString("KEY")
	fmt.Println("Key is", secret)
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to create token",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserLoginRes{
			Status:  "success",
			Message: "verify otp to see home",
			Token:   token,
		})
}

// @Summary		Send OTP
// @Description	Send OTP to the user's registered phone number
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.CommonRes	"Successfully sent OTP"
// @Failure		400	{object}	res.CommonRes	"Bad Request"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/sendOtp [post]
func (h *UserHandler) SendOtp(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)

	if err := h.usecase.SendOtp(user["phone"].(string)); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to send otp",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully sent otp",
		})
}

// @Summary		Verify OTP
// @Description	Verify OTP for the user's registered phone number
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.UserVerifyOtpReq	true	"User OTP verification request"
// @Success		200	{object}	res.CommonRes			"Successfully verified OTP"
// @Failure		400	{object}	res.CommonRes			"Bad Request"
// @Failure		401	{object}	res.CommonRes			"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/verifyOtp [post]
func (h *UserHandler) VerifyOtp(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)

	var req req.UserVerifyOtpReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	err := h.usecase.VerifyOtp(user["phone"].(string), &req)
	fmt.Println("Error is", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to verify otp",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully verified ",
		})
}

// @Summary		User Login
// @Description	Authenticate and log in as a user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.UserLoginReq	true	"User login request"
// @Success		200	{object}	res.UserLoginRes	"Successfully logged in"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginReq req.UserLoginReq

	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to parse body",
			})
	}
	if err := requestvalidation.ValidateRequest(loginReq); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   fmt.Sprint(err),
				Message: "failed. invalid fields",
			})
	}

	user, err := h.usecase.Login(&loginReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to login",
				Error:   err.Error(),
			})
	}

	secret := viper.GetString("KEY")
	token, _, err := jwttoken.CreateToken(secret, time.Hour*24, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to create token",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserLoginRes{
			Status:  "success",
			Message: "successfully logged in",
			Token:   token,
		})
}

// @Summary		Get paginated list of dishes
// @Description	Retrieve a paginated list of dishes for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			p	query		string			false	"Page number (default: 1)"
// @Param			l	query		string			false	"Number of items per page"
// @Success		200	{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/dishes [get]
func (h *UserHandler) GetDishesPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l")

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

// @Summary		Get a dish
// @Description	Retrieve a specific dish by ID for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Dish ID"	int
// @Success		200	{object}	res.SingleDishRes	"Successfully fetched dish"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/dishes/{id} [get]
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

// @Summary		Search dishes
// @Description	Search for dishes based on a query
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			q	query		string			true	"Search query"
// @Success		200	{object}	res.DishListRes	"Successfully fetched dishes"
// @Failure		401	{object}	res.CommonRes	"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes	"Internal Server Error"
// @Router			/search/dishes [get]
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

// @Summary		Get paginated list of sellers
// @Description	Retrieve a paginated list of sellers for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			p	query		string				false	"Page number (default: 1)"
// @Param			l	query		string				false	"Number of items per page"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/sellers [get]
func (h *UserHandler) GetSellersPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l")

	sellerList, err := h.usecase.GetSellersPage(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch sellers",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched seller",
			SellerList: *sellerList,
		})

}

// @Summary		Get a seller
// @Description	Retrieve a specific seller by ID for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Seller ID"	int
// @Success		200	{object}	res.SingleSellerRes	"Successfully fetched seller"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/sellers/{id} [get]
func (h *UserHandler) GetSeller(c *fiber.Ctx) error {
	sellerId := c.Params("id")

	seller, err := h.usecase.GetSeller(sellerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch seller",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SingleSellerRes{
			Status:  "success",
			Message: "successfully fetched seller",
			Seller:  *seller,
		})
}

// @Summary		Search sellers
// @Description	Search for sellers based on a query
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			q	query		string				true	"Search query"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/search/sellers [get]
func (h *UserHandler) SearchSeller(c *fiber.Ctx) error {
	searchQuery := c.Query("q")

	sellersList, err := h.usecase.SearchSeller(searchQuery)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch sellers",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.SellerListRes{
			Status:     "success",
			Message:    "successfully fetched sellers",
			SellerList: *sellersList,
		})
}
