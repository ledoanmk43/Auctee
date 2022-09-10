package dto

import (
	"errors"
	"strings"
)

type AdminLoginDTO struct {
	Username string
	Password string
}

type UpdateProfileDTO struct {
	Firstname string
	Lastname  string
	Birthday  string
	Phone     string
	Email     string
	Gender    bool
	Country   string
	Language  string
}

type UpdateAddressDTO struct {
	Id          uint
	Firstname   string
	Lastname    string
	Phone       string
	Email       string
	Province    string
	District    string
	SubDistrict string `json:"sub_district" gorm:"type:nvarchar(100);not null"`
	Address     string
	TypeAddress string `json:"type_address" gorm:"type:nvarchar(100)"`
	UserId      uint
}
type PasswordToUpdate struct {
	Password  string
	KeepLogin bool
}

func (admin *AdminLoginDTO) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if len(strings.TrimSpace(admin.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(admin.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	case "register":
		if len(strings.TrimSpace(admin.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(admin.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	default:
		if len(strings.TrimSpace(admin.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(admin.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	}
}
func (address *UpdateAddressDTO) Validate() error {

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
