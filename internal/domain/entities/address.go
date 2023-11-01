package entities

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
