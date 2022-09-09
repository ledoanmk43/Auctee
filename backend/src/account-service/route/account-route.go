package route

import (
	"chilindo/src/account-service/controller"
	"chilindo/src/account-service/middleware"
	"github.com/gin-gonic/gin"
)

type IAccountRoute interface {
	GetRouter()
}

type AccountRouteDefault struct {
	AccountController controller.IAccountController
	Router            *gin.Engine
	JWTMiddleware     *middleware.SMiddleWare
}

func (a *AccountRouteDefault) GetRouter() {
	userRoute := a.Router.Group("/auctee")
	{
		userRoute.POST("/register", a.AccountController.SignUp)
		userRoute.POST("/login", a.AccountController.SignIn)
		userRoute.POST("/logout", a.AccountController.SignOut)
		userRoute.PUT("/user/profile/id=:id", a.AccountController.UpdatePassword)
		userRoute.GET("/user/profile/id=:id", a.AccountController.GetUserByUserId)
		userRoute.PUT("/user/profile/update/id=:id", a.AccountController.UpdateProfileByUserId)

	}
}

func NewAccountRouteDefault(accountController controller.IAccountController, router *gin.Engine) *AccountRouteDefault {
	return &AccountRouteDefault{AccountController: accountController, Router: router}
}
