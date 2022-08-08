package route

import (
	"chilindo/src/user-service/controller"
	"github.com/gin-gonic/gin"
)

type IUserRoute interface {
	GetRouter()
}

type UserRouteDefault struct {
	UserController controller.IUserController
	Router         *gin.Engine
}

func (u *UserRouteDefault) GetRouter() {
	newUserRoute(u.UserController, u.Router)
}

func newUserRoute(controller controller.IUserController, group *gin.Engine) {
	userRoute := group.Group("/chilindo/user")
	{
		userRoute.POST("/sign-up", controller.SignUp)
		userRoute.POST("/sign-in", controller.SignIn)
	}
	//userAuthRoute := group.Group("/chilindo/user").Use(middleware.AuthorizeJWT())
	//{
	//	userAuthRoute.PUT("/update", controller.Update)
	//}
}

func NewUserRouteDefault(userController controller.IUserController, router *gin.Engine) *UserRouteDefault {
	return &UserRouteDefault{UserController: userController, Router: router}
}
