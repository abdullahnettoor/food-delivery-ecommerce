package entities

type Favourite struct {
	ID     uint `json:"-"`
	UserID uint `json:"-"`
	DishID uint `json:"-"`

	FkUser User `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	FkDish Dish `json:"-" gorm:"foreignkey:DishID;constraint:OnDelete:CASCADE"`
}
