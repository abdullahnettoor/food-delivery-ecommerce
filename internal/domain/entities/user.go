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
