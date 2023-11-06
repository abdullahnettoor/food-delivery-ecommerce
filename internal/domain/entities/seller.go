package entities

type Seller struct {
	ID          uint   `json:"restaurantId" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"-"`
	PinCode     string `json:"pinCode"`
	Status      string `json:"status" gorm:"default:Pending"`
}
