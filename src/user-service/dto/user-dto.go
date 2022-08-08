package dto

import (
	"errors"
	"strings"
)

type UserLoginDTO struct {
	Email    string
	Password string
}

type GetAddressDTO struct {
	UserId uint
}

type GetAddressByIdDTO struct {
	AddressId uint
	UserId    uint
}

type UserUpdateDTO struct {
	ID        uint64 `json:"id" form:"id"`
	Firstname string `json:"firstname" form:"name" binding:"required"`
	Lastname  string `json:"lastname" form:"name" binding:"required"`
	Email     string `json:"email" from:"email" binding:"required,email"`
	Password  string `json:"password" from:"password" binding:"required"`
}

func (user *UserLoginDTO) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if len(strings.TrimSpace(user.Email)) == 0 {
			return errors.New("required email")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	case "register":
		if len(strings.TrimSpace(user.Email)) == 0 {
			return errors.New("required email")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	default:
		if len(strings.TrimSpace(user.Email)) == 0 {
			return errors.New("required email")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	}
}
