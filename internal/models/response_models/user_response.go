package res

import "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"

type UserLoginRes struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type UserListRes struct {
	Status   string          `json:"status,omitempty"`
	UserList []entities.User `json:"userList,omitempty"`
	Error    string          `json:"error,omitempty"`
	Message  string          `json:"message,omitempty"`
}
