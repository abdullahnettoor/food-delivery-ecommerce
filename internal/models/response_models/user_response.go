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

type ViewCartRes struct {
	Status  string        `json:"status,omitempty"`
	Cart    entities.Cart `json:"cart,omitempty"`
	Message string        `json:"message,omitempty"`
}
type ViewAddressRes struct {
	Status  string           `json:"status,omitempty"`
	Address entities.Address `json:"address,omitempty"`
	Message string           `json:"message,omitempty"`
}

type ViewAddressListRes struct {
	Status  string           `json:"status,omitempty"`
	AddressList []entities.Address `json:"addressList,omitempty"`
	Message string           `json:"message,omitempty"`
}
