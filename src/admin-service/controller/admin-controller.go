package controller

import (
	"chilindo/pkg/token"
	"chilindo/src/admin-service/dto"
	"chilindo/src/admin-service/entity"
	"chilindo/src/admin-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAdminController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	//Update(c *gin.Context)
}

type AdminController struct {
	AdminService service.IAdminService
	token        *token.Claims
}

func (a AdminController) Update(c *gin.Context) {

}

func NewAdminControllerDefault(adminService service.IAdminService) *AdminController {
	return &AdminController{AdminService: adminService}
}

func (a AdminController) SignUp(ctx *gin.Context) {
	var newAdmin *entity.Admin
	errDTO := ctx.ShouldBindJSON(&newAdmin)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}
	if a.AdminService.IsDuplicateUsername(newAdmin.Username) {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "username existed",
		})
		log.Println("SignUp: username existed", errDTO)
		ctx.Abort()
		return
	}

	createdAdmin, errCreate := a.AdminService.CreateAdmin(newAdmin)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreate.Error(),
		})
		log.Println("SignUp: Error in package controller", errDTO)
		ctx.Abort()
		return
	}

	tokenString, errGenToken := a.token.GenerateJWT("", createdAdmin.Id, createdAdmin.Username)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	createdAdmin.Token = tokenString
	ctx.JSON(http.StatusCreated, gin.H{"token": createdAdmin.Token})
}

func (a *AdminController) SignIn(ctx *gin.Context) {
	var loginDTO *dto.AdminLoginDTO

	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	admin, errVerify := a.AdminService.VerifyCredential(loginDTO)
	if errVerify != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": errVerify.Error(),
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		ctx.Abort()
		return
	}

	tokenString, errGenToken := a.token.GenerateJWT("", admin.Id, admin.Username)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
