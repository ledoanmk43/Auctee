package entity

import (
	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model `json:"-"`
	ProductId  string  `json:"product_id" gorm:"type:varchar(20);not null"`
	StartTime  string  `json:"start-time" gorm:"type:nvarchar(100);not null"`
	EndTime    string  `json:"end-ime" gorm:"type:nvarchar(100);not null"`
	CurrentBid float64 `json:"current-bid" gorm:"type:nvarchar(100);not null"`
	isActive   bool    `json:"isActive"gorm:"default:false"`
	Quantity   int     `json:"quantity" gorm:"type:nvarchar(100);not null"`
}
