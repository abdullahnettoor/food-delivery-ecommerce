package entities

type Seller struct {
	ID          uint   `json:"restaurantId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Status      string `json:"status"`
}
