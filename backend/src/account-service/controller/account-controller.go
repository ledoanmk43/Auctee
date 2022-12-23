package controller

import (
	"backend/pkg/token"
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
	GetUserAsGuestByUserId(c *gin.Context)
	UpdateProfileByUserId(c *gin.Context)
	RefreshToken(c *gin.Context)
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
			"message": "email has already been taken",
		})
		log.Println("SignUp: ", errChecking)
		ctx.Abort()
		return
	}

	createdUser, errCreate := a.AccountService.CreateUser(newUser)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreate.Error(),
		})
		log.Println("SignUp: Error in package controller", errCreate)
		ctx.Abort()
		return
	}

	tokenString, errGenToken := a.token.GenerateJWT("", createdUser.Id, createdUser.Username)
	if errGenToken != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": errGenToken.Error(),
		})
		log.Println("SignIn: Error in GenerateJWT in package controller")
		ctx.Abort()
		return
	}
	//Create Session with token
	createdUser.Token = tokenString
	newSession := sessions.Default(ctx)
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
	newSession := sessions.Default(ctx)
	newSession.Set(config.CookieAuth, tokenString)
	if errSave := newSession.Save(); errSave != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "log in successfully",
	})

}

func (a *AccountController) RefreshToken(ctx *gin.Context) {
	//Checking session whether use is logged in
	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
		})
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

	user, errGet := a.AccountService.GetUserByUserId(claims.UserId)
	if errGet != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		log.Println("Get User: Error in package controller", errGet)
		ctx.Abort()
		return
	}
	user.Password = ""
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
	newSession := sessions.Default(ctx)
	newSession.Set(config.CookieAuth, tokenString)
	if errSave := newSession.Save(); errSave != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "session extended",
	})
}

func (a *AccountController) SignOut(ctx *gin.Context) {
	authSession := sessions.Default(ctx)
	log.Println(authSession.Get(config.CookieAuth))
	authSession.Set(config.CookieAuth, "")
	authSession.Clear()
	log.Println(authSession.Get(config.CookieAuth))
	authSession.Options(sessions.Options{Path: "/", MaxAge: -1, SameSite: http.SameSiteNoneMode, Secure: true, HttpOnly: true})
	authSession.Delete(config.CookieAuth)
	if err := authSession.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "log out successfully"})
	ctx.Redirect(301, "/auctee/login")
}

func (a *AccountController) UpdatePassword(ctx *gin.Context) {
	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
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

	errUpdate := a.AccountService.UpdatePassword(passwordToUpdate, claims.UserId)

	if errUpdate != nil {
		if errUpdate.Error() == "wrong password" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": errUpdate.Error(),
			})

		}
		if errUpdate.Error() == "new password must not be the same as old password" {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": errUpdate.Error(),
			})
		}
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "password updated",
	})
}

func (a *AccountController) GetUserByUserId(ctx *gin.Context) {
	//tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	//if errGetToken != nil {
	//	log.Println("Error when get token in controller: ", errGetToken)
	//	ctx.Abort()
	//	return
	//}
	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
		})
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

	user, errGet := a.AccountService.GetUserByUserId(claims.UserId)
	if errGet != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		log.Println("Get User: Error in package controller", errGet)
		ctx.Abort()
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (a *AccountController) GetUserAsGuestByUserId(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("id"))
	user, errGet := a.AccountService.GetUserByUserId(uint(userId))
	if errGet != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errGet.Error(),
		})
		log.Println("Get User: Error in package controller", errGet)
		ctx.Abort()
		return
	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (a *AccountController) UpdateProfileByUserId(ctx *gin.Context) {
	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(config.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
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

	claims, errExtract := token.ExtractToken(tokenFromCookie.(string))
	if errExtract != nil || len(tokenFromCookie.(string)) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	errUpdate := a.AccountService.UpdateProfileByUserId(claims.UserId, updateBody)
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
