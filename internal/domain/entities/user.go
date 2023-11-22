package entities

type User struct {
	ID        uint   `json:"userId" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"notNull"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique;notNull"`
	Phone     string `json:"phone" gorm:"notNull"`
	Password  string `json:"-"`
	Status    string `json:"status" gorm:"default:Pending"`
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

	FkUser  User `json:"-" gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`

}
