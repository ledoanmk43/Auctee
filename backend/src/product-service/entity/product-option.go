package entity

import "gorm.io/gorm"

type ProductOption struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"not null"`
	Color      string `json:"color" gorm:"type:nvarchar(100)"`
	Size       string `json:"size" gorm:"type:nvarchar(100)"`
	Models     string `json:"model" gorm:"type:nvarchar(100)"`
	Quantity   int    `json:"quantity" gorm:"type:bigint;not null"`
	ProductId  string `gorm:"type:varchar(20);not null" json:"-"`
}
