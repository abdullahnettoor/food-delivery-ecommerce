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

type CreateDishReq struct {
	Name         string  `form:"name" validate:"required,gte=3"`
	Description  string  `form:"description"`
	Price        float64 `form:"price" validate:"required,gte=0"`
	SalePrice        float64 `form:"salePrice" validate:"required,gte=0"`
	Quantity     uint    `form:"quantity" validate:"required,gte=0"`
	CategoryID   uint    `form:"categoryId" validate:"required,number"`
	IsVeg        bool    `form:"isVeg" validate:"boolean"`
	Availability bool    `form:"isAvailable" validate:"boolean"`
	ImageUrl     string  `swaggerignore:"true"`
}

type UpdateDishReq struct {
	Name         string  `json:"name" validate:"required,gte=3"`
	Description  string  `json:"description"`
	Price        float64 `json:"price" validate:"required,gte=0"`
	SalePrice        float64 `json:"salePrice" validate:"gte=0"`
	Quantity     uint    `json:"quantity" validate:"required,gte=0"`
	CategoryID   uint    `json:"categoryId" validate:"required,number"`
	IsVeg        bool    `json:"isVeg" validate:"boolean"`
	Availability bool    `json:"isAvailable" validate:"boolean"`
}

type UpdateOrderStatusReq struct {
	OrderStatus string `json:"orderStatus" validate:"required,oneof='Cooking' 'Food Ready' 'Delivered'"`
}