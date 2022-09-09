package entity

import (
	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model   `json:"-"`
	ProductId    string  `json:"product_id" gorm:"type:varchar(20);not null"`
	StartTime    string  `json:"start_time" gorm:"type:datetime;not null"`
	EndTime      string  `json:"end_time" gorm:"type:datetime;not null"`
	CurrentBid   float64 `json:"current_bid" gorm:"type:double;not null"`
	IsActive     bool    `json:"is_active" gorm:"default:false"`
	Quantity     int     `json:"quantity" gorm:"type:nvarchar(100);not null"`
	PricePerStep float64 `json:"price_per_step" gorm:"type:double;not null"`
	WinnerId     uint    `json:"winner_id"`
}
