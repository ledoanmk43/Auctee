package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"primaryKey"`
	Firstname  string `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname   string `json:"description" gorm:"type:nvarchar(100);not null"`
	//Username  string `json:"username" gorm:"type:nvarchar(100)"`
	Password string `json:"password" gorm:"type:nvarchar(100);not null"`
	Birthday string `json:"birthday" gorm:"type:nvarchar(100)"`
	Phone    string `json:"phone" gorm:"type:nvarchar(100)"`
	Email    string `json:"email" gorm:"type:nvarchar(100);not null"`
	Gender   bool   `json:"gender" gorm:"type:boolean"`
	Country  string `json:"country" gorm:"type:nvarchar(100)"`
	Language string `json:"language" gorm:"type:nvarchar(100)"`
	Token    string `gorm:"-" json:"token,omitempty"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Validate(action string) error {
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
