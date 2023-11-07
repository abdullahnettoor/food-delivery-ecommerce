package req

type AdminLoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=3"`
}

type CreateCategoryReq struct {
	Name string `json:"name" validate:"required,gte=3"`
}
type UpdateCategoryReq struct {
	Name string `json:"name" validate:"required,gte=3"`
}
