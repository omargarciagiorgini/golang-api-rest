package models

//Product alhjolhoh
type Product struct {
	//gorm.Model
	ProductID uint   `json:"productId" gorm:"primary_key"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
}
