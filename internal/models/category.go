package models

type Category struct {
	ID   uint   `json:"categoryId" gorm:"type:bigInt;primaryKey"`
	Name string `json:"categoryName"`
}
