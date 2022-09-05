package service

import (
	"chilindo/pkg/pb/admin"
	"chilindo/pkg/token"
	"chilindo/src/admin-service/dto"
	"chilindo/src/admin-service/entity"
	"chilindo/src/admin-service/repository"
	"strings"

	//"github.com/mashingan/smapping"

	//"github.com/mashingan/smapping"
	"log"
)

type IAdminService interface {
	UpdatePassword(in *dto.PasswordToUpdate, adminId uint) error
	VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Admin, error)
	CreateAdmin(admin *entity.Admin) (*entity.Admin, error)
	IsDuplicateUsername(username string) bool
	CheckIsAuth(req *admin.CheckIsAuthRequest) (*admin.CheckIsAuthResponse, error)
}

type AdminService struct {
	AdminRepository repository.IAdminRepository
}

func NewAdminServiceDefault(adminRepository repository.IAdminRepository) *AdminService {
	return &AdminService{AdminRepository: adminRepository}
}

func (a *AdminService) VerifyCredential(loginDTO *dto.AdminLoginDTO) (*entity.Admin, error) {
	admin, err := a.AdminRepository.VerifyCredential(loginDTO)

	if err != nil {
		log.Println("SignIn: Error VerifyCredential in package service: ", err.Error())
		return nil, err
	}
	return admin, nil
}

func (a *AdminService) CreateAdmin(admin *entity.Admin) (*entity.Admin, error) {
	newAdmin, err := a.AdminRepository.InsertAdmin(admin)
	if err != nil {
		log.Println("Error: Error in package repository: ", err.Error())
		return nil, err
	}
	return newAdmin, nil
}

func (a *AdminService) IsDuplicateUsername(username string) bool {
	res := a.AdminRepository.IsDuplicateUsername(username)
	return res
}

func (a *AdminService) UpdatePassword(in *dto.PasswordToUpdate, adminId uint) error {
	err := a.AdminRepository.UpdatePassword(in.Password, adminId)
	if err != nil {
		log.Println("Error: Error in package repository: ", err.Error())
		return err
	}
	return nil
}

func (u AdminService) CheckIsAuth(req *admin.CheckIsAuthRequest) (*admin.CheckIsAuthResponse, error) {
	isAuth := false
	tokenString := req.Token

	tokenResult := strings.TrimPrefix(tokenString, "Bearer ")
	claims, err := token.ExtractToken(tokenResult)

	if err != nil {
		log.Println("CheckIsAuth: ", err)
		return nil, err
	}

	if claims != nil {
		isAuth = true
	}

	return &admin.CheckIsAuthResponse{
		IsAuth: isAuth,
	}, nil
}
