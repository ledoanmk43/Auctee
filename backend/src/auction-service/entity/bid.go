package entity

import "time"

type Bid struct {
	BidId     uint      `json:"bid_id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	AuctionId uint      `json:"auction_id"`
	BidValue  float64   `json:"bid_value" gorm:"type:double;not null"`
	BidTime   time.Time `json:"bid_time" gorm:"type:datetime;not null"`
	Auction   Auction   `json:"-" gorm:"foreignKey:AuctionId"`
	//User      entity.User `json:"-" gorm:"foreignKey:UserId"`
}
