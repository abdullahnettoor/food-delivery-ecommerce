package entities

type Dish struct {
	ID           uint    `json:"dishId"`
	RestaurantID uint    `json:"sellerId"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Quantity     uint    `json:"quantity"`
	Category     uint    `json:"category"`
	IsVeg        bool    `json:"isVeg"`
	Availability bool    `json:"isAvailable" grom:"type:boolean;default:true"`
	Deleted      bool    `json:"-" gorm:"type:boolean;default:false"`
}
