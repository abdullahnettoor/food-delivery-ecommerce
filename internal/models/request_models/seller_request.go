package req

type SellerSignUpReq struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,gte=3"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password,gte=3"`
	Name            string `json:"name" validate:"required,min=2"`
	Description     string `json:"description"`
	PinCode         string `json:"pinCode" validate:"required,len=6"`
}

type SellerLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=3"`
}
