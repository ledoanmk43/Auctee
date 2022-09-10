package controller

import (
	"backend/pkg/token"
	"backend/pkg/utils"
	"backend/src/account-service/config"
	"backend/src/account-service/dto"
	"backend/src/account-service/entity"
	"backend/src/account-service/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IAccountController interface {
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
	SignOut(c *gin.Context)
	UpdatePassword(c *gin.Context)
	GetUserByUserId(c *gin.Context)
	UpdateProfileByUserId(c *gin.Context)
}

type AccountController struct {
	AccountService service.IAccountService
	token          *token.Claims
}

func NewAccountControllerDefault(accountService service.IAccountService) *AccountController {
	return &AccountController{AccountService: accountService}
}

func (a *AccountController) SignUp(ctx *gin.Context) {
	var newUser *entity.Account
	errDTO := ctx.ShouldBindJSON(&newUser)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	isDuplicated, errChecking := a.AccountService.IsDuplicateUsername(newUser.Username)
	if isDuplicated {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "username existed",
		})
		log.Println("SignUp: ", errChecking)
		ctx.Abort()
		return
	}

	createdUser, errCreate := a.AccountService.CreateUser(newUser)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreate.Error(),
		})
		log.Println("SignUp: Error in package controller", errCreate)
		ctx.Abort()
		return
	}

	tokenString, errGenToken := a.token.GenerateJWT("", createdUser.Id, createdUser.Username)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	//Create Session with token
	createdUser.Token = tokenString
	newSession := sessions.DefaultMany(ctx, config.CookieAuth)
	newSession.Set(config.CookieAuth, tokenString)
	if errSave := newSession.Save(); errSave != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "register successfully"})
}

func (a *AccountController) SignIn(ctx *gin.Context) {
	var loginDTO *dto.AdminLoginDTO

	errDTO := ctx.ShouldBindJSON(&loginDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	user, errVerify := a.AccountService.VerifyCredential(loginDTO)
	if errVerify != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": errVerify.Error(),
		})
		log.Println("SignIn: Error in UserService.SignIn in package controller")
		ctx.Abort()
		return
	}

	tokenString, errGenToken := a.token.GenerateJWT("", user.Id, user.Username)
	if errGenToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}

	//Create Session with token
	newSession := sessions.DefaultMany(ctx, config.CookieAuth)
	newSession.Set(config.CookieAuth, tokenString)
	if errSave := newSession.Save(); errSave != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "log in successfully",
	})

}

func (a *AccountController) SignOut(ctx *gin.Context) {
	newSession := sessions.DefaultMany(ctx, config.CookieAuth)
	tokenFromCookie := newSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	newSession.Delete(config.CookieAuth)
	if err := newSession.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "log out successfully"})
}

func (a *AccountController) UpdatePassword(ctx *gin.Context) {
	newSession := sessions.DefaultMany(ctx, config.CookieAuth)
	tokenFromCookie := newSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, errGetId := strconv.Atoi(ctx.Param(config.UserId))
	if errGetId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	var passwordToUpdate *dto.PasswordToUpdate
	err := ctx.ShouldBindJSON(&passwordToUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie.(string))
	if errExtract != nil || len(tokenFromCookie.(string)) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	if uint(userId) != claims.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	errUpdate := a.AccountService.UpdatePassword(passwordToUpdate, claims.UserId)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("Update Password: Error in package controller: ", errUpdate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "password updated",
	})
}

func (a *AccountController) GetUserByUserId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	userId, errGetId := strconv.Atoi(ctx.Param(config.UserId))
	if errGetId != nil {
		log.Println("error in get user by userId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	if uint(userId) != claims.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	user, errGet := a.AccountService.GetUserByUserId(uint(userId))
	if errGet != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errGet.Error(),
		})
		log.Println("Get User: Error in package controller", errGet)
		ctx.Abort()
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (a *AccountController) UpdateProfileByUserId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	userId, errGetId := strconv.Atoi(ctx.Param(config.UserId))
	if errGetId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	var updateBody *dto.UpdateProfileDTO
	err := ctx.ShouldBindJSON(&updateBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	if uint(userId) != claims.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	errUpdate := a.AccountService.UpdateProfileByUserId(uint(userId), updateBody)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("Update User: Error in package controller: ", errUpdate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "profile updated",
	})
}
