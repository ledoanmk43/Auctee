package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"unicode"
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

func (user *Account) Validate(pass string) error {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	var message error
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++

		case unicode.IsNumber(char):
			num = true
			tot++

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++

		default:
			return message
		}
	}

	if !upp || !low || !num || !sym || tot < 6 {
		if upp == false {
			message = errors.New("cần ít nhất một ký tự in hoa")
		}
		if low == false {
			message = errors.New("cần ít nhất một ký tự viết thường")
		}
		if num == false {
			message = errors.New("cần ít nhất một số tự nhiên")
		}
		if sym == false {
			message = errors.New("cần ít nhất một ký tự đặc biệt")
		}
		return message
	}
	return nil
}
