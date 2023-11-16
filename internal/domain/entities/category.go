package entities

type Category struct {
	ID      uint   `json:"categoryId"`
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}
