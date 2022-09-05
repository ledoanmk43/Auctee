package entity

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Product struct {
	Deleted     gorm.DeletedAt
	Id          string  `json:"id" gorm:"primary_key;type:varchar(20);not null;unique"`
	Name        string  `json:"name" gorm:"type:nvarchar(100);not null"`
	MinPrice    float64 `json:"min_price" gorm:"type:double;not null"`
	Description string  `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int     `json:"quantity" gorm:"type:nvarchar(500);not null"`
	AdminId     uint    `json:"admin_id"`
}

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
