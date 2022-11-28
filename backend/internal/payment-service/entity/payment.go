package entity

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model `json:"-"`
	Id         uint `json:"id"`
	//User & auction fields
	AuctionId   uint   `json:"auction_id"`
	ProductId   string `json:"product_id" gorm:"type:varchar(20);not null"`
	ProductName string `json:"product_name" gorm:"type:nvarchar(100);not null"`
	Shopname    string `json:"shopname" gorm:"type:nvarchar(100);not null"`
	EndTime     string `json:"end_time" gorm:"type:datetime;not null"`
	OwnerId     uint   `gorm:"not null" json:"-"`
	WinnerId    uint   `gorm:"not null" json:"-"`
	ImagePath   string `json:"image_path" gorm:"type:mediumtext"`
	Quantity    int    `json:"quantity" gorm:"type:nvarchar(100);not null"`
	//Address
	Firstname   string `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname    string `json:"lastname" gorm:"type:nvarchar(100);not null"`
	Phone       string `json:"phone" gorm:"type:nvarchar(100);not null"`
	Email       string `json:"email" gorm:"type:nvarchar(100); not null"`
	Province    string `json:"province" gorm:"type:nvarchar(100); not null"`
	District    string `json:"district" gorm:"type:nvarchar(100); not null"`
	SubDistrict string `json:"sub_district" gorm:"type:nvarchar(100);not null"`
	Address     string `json:"address" gorm:"type:nvarchar(200); not null"`
	TypeAddress string `json:"type_address" gorm:"type:nvarchar(100); not null"`
	//Payment fields
	PaymentMethod  string    `json:"payment_method" gorm:"type:nvarchar(100)"`
	BeforeDiscount float64   `json:"before_discount" gorm:"type:double;not null"`
	ShippingValue  float64   `json:"shipping_value" gorm:"type:double;not null"`
	DiscountValue  float64   `json:"discount_value" gorm:"type:double;not null"`
	Total          float64   `json:"total" gorm:"type:double;not null"`
	CheckoutStatus uint      `json:"checkout_status" gorm:"default:0"`
	CheckoutTime   time.Time `json:"checkout_time" gorm:"type:datetime;not null"`
	ShippingStatus *bool     `json:"shipping_status" gorm:"default:false"`
}
