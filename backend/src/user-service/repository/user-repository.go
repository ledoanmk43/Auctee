package repository

import (
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IUserRepository interface {
	VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error)
	InsertUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) *entity.User
	IsDuplicateEmail(email string) bool
	FindByEmail(email string) *entity.User
	ProfileUser(userID string) *entity.User
}

type UserRepositoryDefault struct {
	db *gorm.DB
}

func NewUserRepositoryDefault(db *gorm.DB) *UserRepositoryDefault {
	return &UserRepositoryDefault{db: db}
}

func (u UserRepositoryDefault) InsertUser(user *entity.User) (*entity.User, error) {
	if errCheckEmptyField := user.Validate("register"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	if errHashPassword := user.HashPassword(user.Password); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return nil, errHashPassword
	}

	result := u.db.Create(&user)
	if result.Error != nil {
		log.Println("CreateUser: Error in package repository", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (u UserRepositoryDefault) UpdateUser(user *entity.User) *entity.User {
	//if user.Password != "" {
	//	user.Password, _ = user.HashPassword(user.Password)
	//} else {
	//	var tempUser entity.User
	//	u.db.Find(&tempUser, user.ID)
	//	user.Password = tempUser.Password
	//}
	//
	//u.db.Save(&user)
	return user
}

func (u UserRepositoryDefault) IsDuplicateEmail(email string) bool {
	var user *entity.User
	result := u.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return true
	}
	return false
}

func (u UserRepositoryDefault) FindByEmail(email string) *entity.User {
	var user *entity.User
	u.db.Where("email = ?", email).Find(&user)

	return user
}

func (u UserRepositoryDefault) ProfileUser(userID string) *entity.User {
	var user *entity.User
	u.db.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func (u UserRepositoryDefault) VerifyCredential(loginDTO *dto.UserLoginDTO) (*entity.User, error) {
	if errCheckEmptyField := loginDTO.Validate("login"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	var user *entity.User
	res := u.db.Where("email = ?", loginDTO.Email).Find(&user)
	if res.Error != nil {
		log.Println("VerifyCredential: Error find username in package repository: ", res.Error)

		return nil, res.Error
	}

	if len(user.Email) == 0 {
		err := errors.New("email doesn't exist")
		return nil, err
	}

	if err := user.CheckPassword(loginDTO.Password); err != nil {
		log.Println("VerifyCredential: Error in check password package repository: ", err.Error())
		err = errors.New("wrong password")
		return nil, err
	}
	return user, nil
}
