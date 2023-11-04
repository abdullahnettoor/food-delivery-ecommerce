package entities

type Admin struct {
	ID       uint   `json:"adminId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
