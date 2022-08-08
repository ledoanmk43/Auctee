package dto

import (
	"errors"
	"strings"
)

type AdminLoginDTO struct {
	Username string
	Password string
}

type AdminUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"email" from:"username" binding:"required,username"`
	Password string `json:"password" from:"password" binding:"required"`
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
