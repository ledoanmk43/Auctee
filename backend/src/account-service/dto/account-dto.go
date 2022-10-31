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
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Birthday       string `json:"birthday"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Gender         *bool  `json:"gender"`
	Country        string `json:""`
	Language       string `json:"language"`
	Shopname       string `json:"shopname"`
	Avatar         string `json:"avatar"`
	Nickname       string `json:"nickname"`
	PresentAuction string `json:"present_auction"`
}

type UpdateAddressDTO struct {
	Id          uint
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Province    string `json:"province"`
	District    string `json:"district"`
	SubDistrict string `json:"sub_district"`
	Address     string `json:"address"`
	TypeAddress string `json:"type_address"`
	UserId      uint   `json:"user_id"`
	IsDefault   *bool  `json:"is_default"`
}
type PasswordToUpdate struct {
	OldPassword string `json:"old_password" gorm:"type:nvarchar(100);not null"`
	NewPassword string `json:"new_password" gorm:"type:nvarchar(100);not null"`
	KeepLogin   bool
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
