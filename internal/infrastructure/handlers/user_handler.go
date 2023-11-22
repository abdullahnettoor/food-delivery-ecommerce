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
		return c.Status(fiber.StatusInternalServerError).
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
		return c.Status(fiber.StatusInternalServerError).
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

// @Summary		View User Details
// @Description Get user details
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.UserDetailsRes	"Successfully fetched user details"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile [get]
func (h *UserHandler) ViewUserDetails(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	userDetails, err := h.usecase.GetUserDetails(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to fetch user details",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserDetailsRes{
			Status:  "success",
			Message: "successfully fetched user details",
			User:    *userDetails,
		})
}

// @Summary		Update User Details
// @Description Update user name, email details
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.UpdateUserDetailsReq true	"Edit user details request"
// @Success		200	{object}	res.UserDetailsRes	"Successfully updated user details"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile/edit [patch]
func (h *UserHandler) EditUserDetails(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	var req req.UpdateUserDetailsReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	updatedUser, err := h.usecase.UpdateUserDetails(userId, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to edit details",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.UserDetailsRes{
			Status:  "success",
			Message: "successfully updated user details",
			User:    *updatedUser,
		})
}

// @Summary		Change Password
// @Description Change existing password of user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.ChangePasswordReq true	"Change user password request"
// @Success		200	{object}	res.CommonRes	"Successfully changed user password details"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile/changePassword [post]
func (h *UserHandler) ChangePassword(c *fiber.Ctx) error {
	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	var req req.ChangePasswordReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	if err := h.usecase.ChangePassword(userId, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to change password",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "password changed successfully",
		})
}

// @Summary		Forgot Password
// @Description Forgot user password
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.ForgotPasswordReq true	"Forgot password request"
// @Success		200	{object}	res.CommonRes	"Successfully sent otp to number related to the given email"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/forgotPassword [post]
func (h *UserHandler) ForgotPassword(c *fiber.Ctx) error {

	var req req.ForgotPasswordReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	if err := h.usecase.ForgotPassword(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to process",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "otp sent to given phone",
		})
}

// @Summary		Reset Password
// @Description Reset user password
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.ResetPasswordReq true	"Reset password request"
// @Success		200	{object}	res.CommonRes	"Reset Password Successfull"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/resetPassword [post]
func (h *UserHandler) ResetPassword(c *fiber.Ctx) error {

	var req req.ResetPasswordReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	if err := h.usecase.ResetPassword(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(res.CommonRes{
				Status:  "failed",
				Message: "failed to reset password",
				Error:   err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "reset password successfull",
		})
}

// @Summary		Add an address
// @Description	Add a new address for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.NewAddressReq	true	"New address request"
// @Success		200	{object}	res.CommonRes		"Successfully added address"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile/addAddress [post]
func (h *UserHandler) AddAddress(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	var req req.NewAddressReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	if err := h.usecase.AddAddress(userId, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to add address",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully added address",
		})
}

// @Summary		Update address
// @Description	Update existing address for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			req	body		req.UpdateAddressReq	true	"Update address request"
// @Param			id	path		string				true	"Address ID"
// @Success		200	{object}	res.CommonRes		"Successfully updated address"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile/address/{id} [put]
func (h *UserHandler) EditAddress(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])
	addressId := c.Params("id")

	var req req.UpdateAddressReq

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
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

	if err := h.usecase.UpdateAddress(userId, addressId, &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to update address",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.CommonRes{
			Status:  "success",
			Message: "successfully updated address",
		})
}

// @Summary		View a specific address
// @Description	View details of a specific address for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Address ID"
// @Success		200	{object}	res.ViewAddressRes	"Successfully fetched address"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		401	{object}	res.CommonRes		"Unauthorized Access"
// @Failure		404	{object}	res.CommonRes		"Address not found"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/profile/address/{id} [get]
func (h *UserHandler) ViewAddress(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])
	addressId := c.Params("id")

	address, err := h.usecase.ViewAddress(addressId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch address",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewAddressRes{
			Status:  "success",
			Message: "successfully added address",
			Address: *address,
		})
}

// @Summary		View all addresses
// @Description	View details of all addresses for the user
// @Security		Bearer
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	res.ViewAddressListRes	"Successfully fetched addresses"
// @Failure		401	{object}	res.CommonRes			"Unauthorized Access"
// @Failure		500	{object}	res.CommonRes			"Internal Server Error"
// @Router			/profile/address [get]
func (h *UserHandler) ViewAllAddress(c *fiber.Ctx) error {

	user := c.Locals("UserModel").(map[string]any)
	userId := fmt.Sprint(user["userId"])

	addressList, err := h.usecase.ViewAllAddresses(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(res.CommonRes{
				Status:  "failed",
				Error:   err.Error(),
				Message: "failed to fetch address",
			})
	}

	return c.Status(fiber.StatusOK).
		JSON(res.ViewAddressListRes{
			Status:      "success",
			Message:     "successfully added address",
			AddressList: *addressList,
		})
}

// @Summary		Get paginated list of sellers
// @Description	Retrieve a paginated list of sellers for the user
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			p	query		string				false	"Page number (default: 1)"
// @Param			l	query		string				false	"Number of items per page"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/user/sellers [get]
func (h *UserHandler) GetSellersPage(c *fiber.Ctx) error {
	page := c.Query("p", "1")
	limit := c.Query("l", "10")

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
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string				true	"Seller ID"	int
// @Success		200	{object}	res.SingleSellerRes	"Successfully fetched seller"
// @Failure		400	{object}	res.CommonRes		"Bad Request"
// @Failure		500	{object}	res.CommonRes		"Internal Server Error"
// @Router			/user/sellers/{id} [get]
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
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			q	query		string				true	"Search query"
// @Success		200	{object}	res.SellerListRes	"Successfully fetched sellers"
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
