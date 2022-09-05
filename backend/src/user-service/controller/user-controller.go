package controller

import (
	"chilindo/pkg/token"
	"chilindo/src/user-service/dto"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IUserController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	Update(c *gin.Context)
}

type UserController struct {
	UserService service.IUserService
	token       *token.Claims
}

func (u UserController) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewUserControllerDefault(userService service.IUserService) *UserController {
	return &UserController{UserService: userService}
}

func (u *UserController) SignUp(ctx *gin.Context) {
	var newUser *entity.User
	errDTO := ctx.ShouldBindJSON(&newUser)

	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	if u.UserService.IsDuplicateEmail(newUser.Email) {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "email existed",
		})
		log.Println("SignUp: email existed", errDTO)
		ctx.Abort()
		return
	}

	createdUser, errCreate := u.UserService.CreateUser(newUser)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreate.Error(),
		})
		log.Println("SignUp: Error in package controller", errDTO)
		ctx.Abort()
		return
	}

	tokenString, errGenToken := u.token.GenerateJWT(createdUser.Email, createdUser.Id, "")
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	createdUser.Token = tokenString
	ctx.JSON(http.StatusCreated, gin.H{"token": createdUser.Token})
} //done

func (u *UserController) SignIn(ctx *gin.Context) {
	var loginDTO *dto.UserLoginDTO

	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	user, errVerify := u.UserService.VerifyCredential(loginDTO)
	if errVerify != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Message": errVerify.Error(),
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		ctx.Abort()
		return
	}

	tokenString, errGenToken := u.token.GenerateJWT(user.Email, user.Id, "")
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
