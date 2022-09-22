package entity

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"not null"`
	Path       string `json:"path" gorm:"type:varchar(100)"`
	IsDefault  *bool  `json:"is_default"gorm:"type:boolean;default:false"`
	ProductId  string `gorm:"type:varchar(20);not null" json:"-"`
}
