package models

import validator "github.com/go-playground/validator/v10"

//Product alhjolhoh
type (
	Product struct {
		//gorm.Model
		ProductID uint   `json:"productId" gorm:"primary_key"`
		Code      string `json:"code" validate:"required"`
		Name      string `json:"name" validate:"required"`
		Price     uint   `json:"price" validate:"required,numeric"`
	}
	CustomValidator struct {
		Validator *validator.Validate
	}
)

//Validate alhoh
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
