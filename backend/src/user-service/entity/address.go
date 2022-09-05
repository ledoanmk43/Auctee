package entity

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Address struct {
	gorm.Model
	Firstname   string `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname    string `json:"lastname" gorm:"type:nvarchar(100);not null"`
	Phone       string `json:"phone" gorm:"type:nvarchar(100);not null"`
	Email       string `json:"email" gorm:"type:nvarchar(100); not null"`
	Province    string `json:"province" gorm:"type:nvarchar(100); not null"`
	District    string `json:"district" gorm:"type:nvarchar(100); not null"`
	SubDistrict string `json:"sub_district" gorm:"type:nvarchar(100);not null"`
	Address     string `json:"address" gorm:"type:nvarchar(200); not null"`
	TypeAddress string `gorm:"type_address" gorm:"type:nvarchar(100); not null"`
	UserId      uint   `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE"json:"-"`
}

func (address *Address) Validate() error {

	if len(strings.TrimSpace(address.Firstname)) == 0 {
		return errors.New("required firstname")
	}
	if len(strings.TrimSpace(address.Lastname)) == 0 {
		return errors.New("required lastname")
	}
	if len(strings.TrimSpace(address.Phone)) == 0 {
		return errors.New("required phone")
	}
	if len(strings.TrimSpace(address.Email)) == 0 {
		return errors.New("required email")
	}
	if len(strings.TrimSpace(address.Province)) == 0 {
		return errors.New("required province")
	}
	if len(strings.TrimSpace(address.District)) == 0 {
		return errors.New("required district")
	}
	if len(strings.TrimSpace(address.Address)) == 0 {
		return errors.New("required address")
	}
	return nil
}
