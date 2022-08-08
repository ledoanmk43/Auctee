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
	//Update(admin *dto.AdminUpdateDTO) *entity.Admin
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
	return newAdmin, err
}

func (a *AdminService) IsDuplicateUsername(username string) bool {
	res := a.AdminRepository.IsDuplicateUsername(username)
	return res
}

func (a *AdminService) Update(admin *dto.AdminUpdateDTO) *entity.Admin {
	//var adminToUpdate *entity.Admin
	//err := smapping.FillStruct(&adminToUpdate, smapping.MapFields(&admin))
	//if err != nil {
	//	log.Fatalf("Failed map %v:", err)
	//}
	//updatedAdmin := a.AdminRepository.UpdateAdmin(adminToUpdate)
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
