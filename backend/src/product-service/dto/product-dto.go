package dto

import (
	"backend/src/product-service/entity"
	"gorm.io/gorm"
)

type ProductCreatedDTO struct {
	Product *entity.Product
}

func NewProductCreatedDTO(product *entity.Product) *ProductCreatedDTO {
	return &ProductCreatedDTO{Product: product}
}

type ProductUpdateDTO struct {
	ProductId string
	Product   *entity.Product
}

type ProductDTO struct {
	gorm.Model  `json:"-"`
	Id          string  `json:"id" gorm:"primary_key;type:varchar(20);not null;unique"`
	Name        string  `json:"name" gorm:"type:nvarchar(100);not null"`
	MinPrice    float64 `json:"min_price" gorm:"type:double;not null"`
	Description string  `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int     `json:"quantity" gorm:"type:nvarchar(500);not null"`
	ExpectPrice float64 `json:"expect_price" gorm:"type:double;not null"`
	UserId      uint    `gorm:"not null" json:"-"`
}
