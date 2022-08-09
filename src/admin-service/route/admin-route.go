package route

import (
	"chilindo/src/admin-service/controller"
	"chilindo/src/admin-service/middleware"
	"github.com/gin-gonic/gin"
)

type IAdminRoute interface {
	GetRouter()
}

type AdminRouteDefault struct {
	AdminController controller.IAdminController
	Router          *gin.Engine
	JWTMiddleware   *middleware.SMiddleWare
}

func (a *AdminRouteDefault) GetRouter() {
	userRoute := a.Router.Group("/chilindo/admin")
	{
		userRoute.POST("/sign-up", a.AdminController.SignUp)
		userRoute.POST("/sign-in", a.AdminController.SignIn)
		userRoute.PATCH("/id=:id/password-setting", a.AdminController.UpdatePassword).Use(a.JWTMiddleware.IsAuthenticated())
	}
}

func NewAdminRouteDefault(adminController controller.IAdminController, router *gin.Engine) *AdminRouteDefault {
	return &AdminRouteDefault{AdminController: adminController, Router: router}
}
