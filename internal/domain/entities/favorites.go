package entities

type Favourite struct {
	ID     uint `json:"-"`
	UserID uint `json:"-"`
	DishID uint `json:"-"`
	User User  `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	Dish  Dish `json:"dish" gorm:"foreignkey:DishID;constraint:OnDelete:CASCADE"`
}
