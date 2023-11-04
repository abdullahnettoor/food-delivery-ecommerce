package requestmodels

type AdminLoginReq struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"gte=3,required"`
}
