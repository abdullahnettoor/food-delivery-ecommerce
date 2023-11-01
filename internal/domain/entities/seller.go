package entities

type Restaurant struct {
	ID          uint   `json:"restaurantId"`
	Name        string `json:"restaurantName"`
	Description string `json:"restaurantDescription"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Status      string `json:"status"`
}
