package entity

import "gorm.io/gorm"

type ProductImages struct {
	gorm.Model `jon:"-"`
	ProductId  string  `json:"productId" gorm:"type:varchar(20);not null"`
	Link       string  `json:"link" gorm:"type:varchar(100)"`
	Product    Product `json:"-" gorm:"foreignKey:ProductId"`
}
