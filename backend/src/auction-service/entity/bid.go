package entity

import (
	"gorm.io/gorm"
	"time"
)

type Bid struct {
	gorm.Model `json:"-"`
	UserId     uint      `json:"user_id"`
	AuctionId  uint      `json:"auction_id"`
	BidValue   float64   `json:"bid_value" gorm:"type:double;not null"`
	BidTime    time.Time `json:"bid_time" gorm:"type:datetime;not null"`
}
