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

type AdminUpdateDTO struct {
	Id       uint   `json:"id" form:"id"`
	Username string `json:"email" from:"username" binding:"required,username"`
	Password string `json:"password" from:"password" binding:"required"`
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
