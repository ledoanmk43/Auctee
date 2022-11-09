package service

import (
	"backend/internal/account-service/dto"
	"backend/internal/account-service/entity"
	"backend/internal/account-service/repository"
	"backend/pkg/pb/account"
	"backend/pkg/token"
	"errors"
	//"github.com/mashingan/smapping"

	//"github.com/mashingan/smapping"
	"log"
)

type IAccountService interface {
	UpdatePassword(dto *dto.PasswordToUpdate, userId uint) error
	VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Account, error)
	CreateUser(user *entity.Account) (*entity.Account, error)
	IsDuplicateUsername(username string) (bool, error)
	CheckIsAuth(req *account.CheckIsAuthRequest) (*account.CheckIsAuthResponse, error)
	GetUserByUserId(userId uint) (*entity.Account, error)
	UpdateProfileByUserId(userId uint, updateBody *dto.UpdateProfileDTO) error
}

type AccountService struct {
	AccountRepository repository.IAccountRepository
}

func NewAccountServiceDefault(accountRepository repository.IAccountRepository) *AccountService {
	return &AccountService{AccountRepository: accountRepository}
}

func (a *AccountService) VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Account, error) {
	admin, err := a.AccountRepository.VerifyCredential(loginDTO)

	if err != nil {
		log.Println("SignIn: Error VerifyCredential in package service: ", err.Error())
		return nil, err
	}
	return admin, nil
}

func (a *AccountService) CreateUser(user *entity.Account) (*entity.Account, error) {
	if len(user.Password) < 6 {
		log.Println("Create Password: Error empty field in package repository: empty input")
		return nil, errors.New("password too short")
	}
	newUser, err := a.AccountRepository.InsertUser(user)
	if err != nil {
		log.Println("Error: Error in package service: ", err.Error())
		return nil, err
	}
	return newUser, nil
}

func (a *AccountService) IsDuplicateUsername(username string) (bool, error) {
	user, err := a.AccountRepository.IsDuplicateUsername(username)
	if user != nil {
		return true, nil
	}
	return false, err
}

func (a *AccountService) UpdatePassword(dto *dto.PasswordToUpdate, userId uint) error {
	if len(dto.OldPassword) == 0 || len(dto.NewPassword) == 0 {
		log.Println("Update Password: Error empty field in package repository: empty input")
		return errors.New("password field must not be empty")
	}
	if len(dto.NewPassword) < 6 {
		log.Println("Update Password: Error empty field in package repository: empty input")
		return errors.New("password too short")
	}
	err := a.AccountRepository.UpdatePassword(dto, userId)
	if err != nil {
		log.Println("Error: Error in package service: ", err.Error())
		return err
	}
	return nil
}

func (a *AccountService) UpdateProfileByUserId(userId uint, updateBody *dto.UpdateProfileDTO) error {
	err := a.AccountRepository.UpdateProfileByUserId(userId, updateBody)
	if err != nil {
		log.Println("Error: Error in package service: ", err.Error())
		return err
	}
	return nil
}

func (a *AccountService) GetUserByUserId(userId uint) (*entity.Account, error) {
	user, err := a.AccountRepository.GetUserByUserId(userId)
	if err != nil {
		log.Println("Error: Error in package service: ", err.Error())
		return nil, err
	}
	return user, nil
}

func (u AccountService) CheckIsAuth(req *account.CheckIsAuthRequest) (*account.CheckIsAuthResponse, error) {
	isAuth := false
	tokenString := req.Token

	claims, err := token.ExtractToken(tokenString)
	if err != nil {
		log.Println("CheckIsAuth: ", err)
		return nil, err
	}

	if claims != nil {
		isAuth = true
	}

	return &account.CheckIsAuthResponse{
		IsAuth: isAuth,
	}, nil
}
