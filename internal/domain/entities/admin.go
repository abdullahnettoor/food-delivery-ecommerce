package entities

type Admin struct {
	ID       uint   `json:"adminId"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
