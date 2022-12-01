package repository

import (
	"backend/pkg/utils"
	"backend/src/account-service/dto"
	"backend/src/account-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type IAccountRepository interface {
	VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Account, error)
	InsertUser(user *entity.Account) (*entity.Account, error)
	UpdatePassword(dto *dto.PasswordToUpdate, userId uint) error
	IsDuplicateUsername(username string) (*entity.Account, error)
	GetUserByUserId(userId uint) (*entity.Account, error)
	UpdateProfileByUserId(userId uint, updateBody *dto.UpdateProfileDTO) error
	UpdateHonorPoint(userId, caseId uint) error
	UpdateInCome(userId, caseId uint, value float64) error
}

type AccountRepositoryDefault struct {
	db *gorm.DB
}

func NewAccountRepositoryDefault(db *gorm.DB) *AccountRepositoryDefault {
	return &AccountRepositoryDefault{db: db}
}

func (a *AccountRepositoryDefault) InsertUser(user *entity.Account) (*entity.Account, error) {
	if errCheckEmptyField := user.Validate("register"); errCheckEmptyField != nil {
		log.Println("Verify Credential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	if errHashPassword := user.HashPassword(user.Password); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return nil, errHashPassword
	}
	user.Email = user.Username
	result := a.db.Create(&user)
	if result.Error != nil {
		log.Println("Create User: Error in package repository", result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (a *AccountRepositoryDefault) UpdatePassword(dto *dto.PasswordToUpdate, userId uint) error {

	var userToUpdate *entity.Account
	result := a.db.Where("id = ?", userId).Find(&userToUpdate)
	if result.Error != nil {
		log.Println("Update Password: Error in package repository: ", result.Error)
		return errors.New("Unauthorized")
	}
	if userToUpdate.CheckPassword(dto.OldPassword) != nil { //compare
		return errors.New("wrong password")
	}
	if userToUpdate.CheckPassword(dto.NewPassword) == nil { //compare
		return errors.New("new password must not be the same as old password")
	}
	if errHashPassword := userToUpdate.HashPassword(dto.NewPassword); errHashPassword != nil {
		log.Println("CreateUser: Error in package repository", errHashPassword)
		return errHashPassword
	}

	res := a.db.Updates(&userToUpdate)
	if res.Error != nil {
		log.Println("CreateUser: Error in package repository", res.Error)
		return result.Error
	}
	return nil
}

func (a *AccountRepositoryDefault) UpdateProfileByUserId(userId uint, updateBody *dto.UpdateProfileDTO) error {
	var userToUpdate *entity.Account
	result := a.db.Where("id = ?", userId).Find(&userToUpdate)
	if result.Error != nil {
		log.Println("Update Password: Error in package repository: ", result.Error)
		return errors.New("Unauthorized")
	}

	userToUpdate.Firstname = updateBody.Firstname
	userToUpdate.Lastname = updateBody.Lastname
	userToUpdate.Phone = updateBody.Phone
	userToUpdate.Email = updateBody.Email
	userToUpdate.Birthday = updateBody.Birthday
	if updateBody.Gender != nil {
		userToUpdate.Gender = utils.BoolAddr(*updateBody.Gender)
	}
	userToUpdate.Country = updateBody.Country
	userToUpdate.Language = updateBody.Language
	userToUpdate.Shopname = updateBody.Shopname
	userToUpdate.Avatar = updateBody.Avatar
	userToUpdate.Nickname = updateBody.Nickname
	id, _ := strconv.Atoi(updateBody.PresentAuction)
	userToUpdate.PresentAuction = uint(id)

	res := a.db.Where("id = ?", userId).Updates(&userToUpdate)
	if res.Error != nil {
		log.Println("Update User: Error in package repository", res.Error)
		return result.Error
	}
	userToUpdate.Password = ""
	return nil
}

func (a *AccountRepositoryDefault) UpdateHonorPoint(userId, caseId uint) error {
	var userToUpdate *entity.Account
	result := a.db.Where("id = ?", userId).Find(&userToUpdate)
	if result.Error != nil {
		log.Println("Update Password: Error in package repository: ", result.Error)
		return errors.New("Unauthorized")
	}
	// 1: +2	2:-5
	switch caseId {
	case 1:
		userToUpdate.HonorPoint += 2
		res := a.db.Where("id = ?", userId).Updates(&userToUpdate)
		if res.Error != nil {
			log.Println("Update User: Error in package repository", res.Error)
			return result.Error
		}
	case 2:
		userToUpdate.HonorPoint -= 5
		res := a.db.Where("id = ?", userId).Updates(&userToUpdate)
		if res.Error != nil {
			log.Println("Update User: Error in package repository", res.Error)
			return result.Error
		}
	default:
		break
	}

	return nil
}

func (a *AccountRepositoryDefault) UpdateInCome(userId, caseId uint, value float64) error {
	//var userToUpdate *entity.Account
	//result := a.db.Where("id = ?", userId).Find(&userToUpdate)
	//if result.Error != nil {
	//	log.Println("Update Password: Error in package repository: ", result.Error)
	//	return errors.New("Unauthorized")
	//}
	//// 1: +2	2:-5
	//switch caseId {
	//case 1:
	//	userToUpdate.HonorPoint += 2
	//	res := a.db.Where("id = ?", userId).Updates(&userToUpdate)
	//	if res.Error != nil {
	//		log.Println("Update User: Error in package repository", res.Error)
	//		return result.Error
	//	}
	//case 2:
	//	userToUpdate.HonorPoint -= 5
	//	res := a.db.Where("id = ?", userId).Updates(&userToUpdate)
	//	if res.Error != nil {
	//		log.Println("Update User: Error in package repository", res.Error)
	//		return result.Error
	//	}
	//default:
	//	break
	//}

	return nil
}

func (a *AccountRepositoryDefault) IsDuplicateUsername(username string) (*entity.Account, error) {
	var user *entity.Account
	result := a.db.Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if user.Id == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (a *AccountRepositoryDefault) GetUserByUserId(userId uint) (*entity.Account, error) {
	var user *entity.Account
	result := a.db.Where("id = ?", userId).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if user.Id == 0 {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (a *AccountRepositoryDefault) VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Account, error) {
	if errCheckEmptyField := loginDTO.Validate("login"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	var user *entity.Account
	res := a.db.Where("username = ?", loginDTO.Username).Find(&user)
	if res.Error != nil {
		log.Println("VerifyCredential: Error find username in package repository: ", res.Error)

		return nil, res.Error
	}

	if len(user.Username) == 0 {
		err := errors.New("username doesn't exist")
		return nil, err
	}
	if err := user.CheckPassword(loginDTO.Password); err != nil {
		log.Println("VerifyCredential: Error in check password package repository: ", err.Error())
		err = errors.New("wrong password")
		return nil, err
	}
	return user, nil
}
