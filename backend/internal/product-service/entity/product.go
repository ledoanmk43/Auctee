package entity

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Product struct {
	//Deleted     gorm.DeletedAt
	gorm.Model    `json:"-"`
	Id            string          `json:"id" gorm:"primary_key;type:varchar(20);not null;unique"`
	Name          string          `json:"name" gorm:"type:nvarchar(100);not null"`
	MinPrice      float64         `json:"min_price" gorm:"type:double;not null"`
	Description   string          `json:"description" gorm:"type:nvarchar(1500);not null"`
	Quantity      int             `json:"quantity" gorm:"type:bigint;not null"`
	ExpectPrice   float64         `json:"expect_price" gorm:"type:double;not null"`
	UserId        uint            `gorm:"not null" json:"-"`
	ProductImage  []ProductImage  `json:"product_images"`
	ProductOption []ProductOption `json:"product_options"`
}

type ProductResponse struct {
	IdList      []string
	ProductName []string
}

//type ProductDetail struct {
//	Product Product
//}

func (product *Product) Validate(action string) error {
	switch strings.ToLower(action) {
	case "insert":
		if len(strings.TrimSpace(product.Name)) == 0 {
			return errors.New("required product name")
		}
		if len(strings.TrimSpace(strconv.Itoa(int(product.MinPrice)))) == 0 {
			return errors.New("required min price")
		}
		if product.Quantity == 0 {
			return errors.New("required quantity of product")
		}
		if len(strings.TrimSpace(strconv.Itoa(int(product.ExpectPrice)))) == 0 {
			return errors.New("required expect price")
		}
		return nil
	default:
		if len(strings.TrimSpace(product.Name)) == 0 {
			return errors.New("required product name")
		}
		if len(strings.TrimSpace(strconv.Itoa(int(product.MinPrice)))) == 0 {
			return errors.New("required min price")
		}
		if len(strings.TrimSpace(string(product.Quantity))) == 0 {
			return errors.New("required quantity of product")
		}
		return nil
	}
}
