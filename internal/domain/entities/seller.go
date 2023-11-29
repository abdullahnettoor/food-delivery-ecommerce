package entities

type Seller struct {
	ID          uint   `json:"sellerId" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"-"`
	PinCode     string `json:"pinCode"`
	Status      string `json:"status" gorm:"default:Pending"`
}

type Sales struct {
	Count    uint    `json:"saleCount"`
	TotalAmt float64 `json:"totalSales"`
}
