package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type UserLoginRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type UserListRes struct {
	Status   string          `json:"status,omitempty"`
	UserList []entities.User `json:"userList,omitempty"`
	Message  string          `json:"message,omitempty"`
}
