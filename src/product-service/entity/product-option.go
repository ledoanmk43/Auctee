package entity

import "gorm.io/gorm"

type ProductOption struct {
	gorm.Model `jon:"-"`
	ProductId  string  `json:"productId" gorm:"type:varchar(20);not null"`
	Color      string  `json:"color" gorm:"type:nvarchar(100)"`
	Size       string  `json:"size" gorm:"type:nvarchar(100)"`
	Models     string  `json:"model" gorm:"type:nvarchar(100)"`
	Product    Product `json:"-" gorm:"foreignKey:ProductId"`
}
