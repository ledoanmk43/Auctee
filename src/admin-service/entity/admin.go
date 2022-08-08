package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type Admin struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"primaryKey"`
	Username   string `json:"username" gorm:"type:nvarchar(100);unique"`
	Password   string `json:"password" gorm:"type:nvarchar(100);not null"`
	Token      string `gorm:"-" json:"token,omitempty"`
}

func (user *Admin) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *Admin) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *Admin) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *Admin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if len(strings.TrimSpace(user.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	case "register":
		if len(strings.TrimSpace(user.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	default:
		if len(strings.TrimSpace(user.Username)) == 0 {
			return errors.New("required username")
		}
		if len(strings.TrimSpace(user.Password)) == 0 {
			return errors.New("required password")
		}
		return nil
	}
}
