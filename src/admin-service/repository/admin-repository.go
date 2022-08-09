package repository

import (
	"chilindo/src/admin-service/dto"
	"chilindo/src/admin-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IAdminRepository interface {
	VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Admin, error)
	InsertAdmin(admin *entity.Admin) (*entity.Admin, error)
	UpdatePassword(password string, adminId uint) error
	IsDuplicateUsername(username string) bool
}

type AdminRepositoryDefault struct {
	db *gorm.DB
}

func NewAdminRepositoryDefault(db *gorm.DB) *AdminRepositoryDefault {
	return &AdminRepositoryDefault{db: db}
}

func (a *AdminRepositoryDefault) InsertAdmin(admin *entity.Admin) (*entity.Admin, error) {
	if errCheckEmptyField := admin.Validate("register"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	if errHashPassword := admin.HashPassword(admin.Password); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return nil, errHashPassword
	}

	result := a.db.Create(&admin)
	if result.Error != nil {
		log.Println("CreateUser: Error in package repository", result.Error)
		return nil, result.Error
	}
	return admin, nil
}

func (a *AdminRepositoryDefault) UpdatePassword(password string, adminId uint) error {
	if len(password) == 0 {
		log.Println("Change Password: Error empty field in package repository: empty input")
		return errors.New("password field must not be empty")
	}

	var adminToUpdate *entity.Admin
	result := a.db.Where("id = ?", adminId).Find(&adminToUpdate)
	if result.Error != nil {
		log.Println("Update Password: Error in package repository: ", result.Error)
		return result.Error
	}
	if adminToUpdate.CheckPassword(password) == nil { //compare
		return errors.New("new password must not be the same as old password")
	}
	if errHashPassword := adminToUpdate.HashPassword(password); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return errHashPassword
	}

	res := a.db.Updates(&adminToUpdate)
	if res.Error != nil {
		log.Println("CreateUser: Error in package repository", res.Error)
		return result.Error
	}
	return nil
}

func (a *AdminRepositoryDefault) IsDuplicateUsername(username string) bool {
	var admin *entity.Admin
	result := a.db.Where("username = ?", username).Find(&admin)
	if result.Error != nil {
		return true
	}
	return false
}

func (a *AdminRepositoryDefault) VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Admin, error) {
	if errCheckEmptyField := loginDTO.Validate("login"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	var admin *entity.Admin
	res := a.db.Where("username = ?", loginDTO.Username).Find(&admin)
	if res.Error != nil {
		log.Println("VerifyCredential: Error find username in package repository: ", res.Error)

		return nil, res.Error
	}

	if len(admin.Username) == 0 {
		err := errors.New("username doesn't exist")
		return nil, err
	}
	if err := admin.CheckPassword(loginDTO.Password); err != nil {
		log.Println("VerifyCredential: Error in check password package repository: ", err.Error())
		err = errors.New("wrong password")
		return nil, err
	}
	return admin, nil
}
