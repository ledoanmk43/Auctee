package entity

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model `json:"-"`
	Id         string `json:"id" gorm:"type:varchar(40);not null"`
	//User & auction fields
	AuctionId   uint   `json:"auction_id"`
	ProductId   string `json:"product_id" gorm:"type:varchar(20);not null"`
	ProductName string `json:"product_name" gorm:"type:nvarchar(100);not null"`
	Shopname    string `json:"shopname" gorm:"type:nvarchar(100);not null"`
	EndTime     string `json:"end_time" gorm:"type:datetime;not null"`
	OwnerId     uint   `gorm:"not null" json:"-"`
	WinnerId    uint   `gorm:"not null" json:"winner_id"`
	ImagePath   string `json:"image_path" gorm:"type:mediumtext"`
	Quantity    int    `json:"quantity" gorm:"type:nvarchar(100);not null"`
	//Address
	AddressId   uint   `json:"address_id"`
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
	Note           string  `json:"note" gorm:"type:nvarchar(1000)"`
	PaymentMethod  string  `json:"payment_method" gorm:"type:nvarchar(100)"`
	BeforeDiscount float64 `json:"before_discount" gorm:"type:double;not null"`
	ShippingValue  float64 `json:"shipping_value" gorm:"type:double;not null"`
	DiscountValue  float64 `json:"discount_value" gorm:"type:double;not null"`
	Total          float64 `json:"total" gorm:"type:double;not null"`
	// 1: chờ xác nhận 	2: đang giao 	3:đã nhận 	4: đã huỷ	5: hoàn thành
	CheckoutStatus uint      `json:"checkout_status" gorm:"default:0"`
	CheckoutTime   time.Time `json:"checkout_time" gorm:"type:datetime;not null"`
	// 1: chờ shipper 	2: đang giao 	3:đã nhận 	4: đã huỷ	5: hoàn thành	0: không nhận hàng
	ShippingStatus uint `json:"shipping_status" gorm:"default:0"`
}

type Payload struct {
	PartnerCode  string `json:"partnerCode"`
	AccessKey    string `json:"accessKey"`
	RequestID    string `json:"requestId"`
	Amount       int64  `json:"amount"`
	OrderID      string `json:"orderId"`
	OrderInfo    string `json:"orderInfo"`
	PartnerName  string `json:"partnerName"`
	StoreId      string `json:"storeId"`
	OrderGroupId string `json:"orderGroupId"`
	Lang         string `json:"lang"`
	AutoCapture  bool   `json:"autoCapture"`
	RedirectUrl  string `json:"redirectUrl"`
	IpnUrl       string `json:"ipnUrl"`
	ExtraData    string `json:"extraData"`
	RequestType  string `json:"requestType"`
	Signature    string `json:"signature"`
}
