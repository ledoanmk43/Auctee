package route

import (
	"chilindo/src/admin-service/controller"
	"github.com/gin-gonic/gin"
)

type IAdminRoute interface {
	GetRouter()
}

type AdminRouteDefault struct {
	AdminController controller.IAdminController
	Router          *gin.Engine
}

func (a *AdminRouteDefault) GetRouter() {
	newAdminRoute(a.AdminController, a.Router)
}

func newAdminRoute(controller controller.IAdminController, group *gin.Engine) {
	userRoute := group.Group("/chilindo/admin")
	{
		userRoute.POST("/sign-up", controller.SignUp)
		userRoute.POST("/sign-in", controller.SignIn)
	}
	//userAuthRoute := group.Group("/chilindo/user").Use(middleware.AuthorizeJWT())
	//{
	//	userAuthRoute.PUT("/update", controller.Update)
	//}
}

func NewAdminRouteDefault(adminController controller.IAdminController, router *gin.Engine) *AdminRouteDefault {
	return &AdminRouteDefault{AdminController: adminController, Router: router}
}
