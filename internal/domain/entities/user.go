package entities

type User struct {
	ID        uint   `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	Status    string `json:"status"`
}

type Address struct {
	ID        uint   `json:"addressId"`
	UserID    uint   `json:"userId"`
	Name      string `json:"name"`
	HouseName string `json:"houseName"`
	Street    string `json:"street"`
	District  string `json:"district"`
	State     string `json:"state"`
	PinCode   string `json:"pinCode"`
	Phone     string `json:"phone"`
}
