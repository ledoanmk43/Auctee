package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Account struct {
	gorm.Model     `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Id             uint      `json:"id" gorm:"primaryKey"`
	Username       string    `json:"username" gorm:"type:nvarchar(100);not null"`
	Password       string    `json:"password,omitempty" gorm:"type:nvarchar(100);not null"`
	Firstname      string    `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname       string    `json:"lastname" gorm:"type:nvarchar(100);not null"`
	Birthday       string    `json:"birthday" gorm:"type:nvarchar(100)"`
	Phone          string    `json:"phone" gorm:"type:nvarchar(100)"`
	Email          string    `json:"email" gorm:"type:nvarchar(100)"`
	Gender         *bool     `json:"gender" gorm:"type:boolean;default:true"`
	Country        string    `json:"country" gorm:"type:nvarchar(100)"`
	Language       string    `json:"language" gorm:"type:nvarchar(100)"`
	Token          string    `gorm:"-" json:"token,omitempty"`
	Shopname       string    `json:"shopname" gorm:"type:nvarchar(100);not null"`
	TotalIncome    float64   `json:"total_income" gorm:"type:double;not null;default:500000"`
	Avatar         string    `json:"avatar" gorm:"type:mediumtext"`
	Nickname       string    `json:"nickname" gorm:"type:nvarchar(100);not null"`
	PresentAuction uint      `json:"present_auction"`
	HonorPoint     uint      `json:"honor_point" gorm:"default:90"`
	Role           uint      `json:"role" gorm:"default:0"` // 0: normal user 1: admin user
	// Additional data for admin
	SystemBalance float64   `json:"system_balance" gorm:"-"`
	TotalUser     uint      `json:"total_user" gorm:"-"`
	Users         []Account `json:"users_list" gorm:"-"`
}

func (user *Account) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *Account) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *Account) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *Account) Validate(action string) error {
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
