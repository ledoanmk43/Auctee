package route

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/product-service/controller"
	admin_server_controller "chilindo/src/product-service/controller/admin-grpc-controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IProductOptionRoute interface {
	GetRouter()
}

type OptionRoute struct {
	ProductOptionController controller.ProductOptionController
	Router                  *gin.Engine
	AdminSrvController      admin_server_controller.IAdminServiceController
	AdminClient             admin.AdminServiceClient
}

func NewOptionRoute(productOptionController controller.ProductOptionController, router *gin.Engine, adminSrvController admin_server_controller.IAdminServiceController, adminClient admin.AdminServiceClient) *OptionRoute {
	return &OptionRoute{ProductOptionController: productOptionController, Router: router, AdminSrvController: adminSrvController, AdminClient: adminClient}
}

func (o OptionRoute) GetRouter() {
	optionRoutes := o.Router.Group("chilindo/option")
	optionRoutes.Use(middleware.Logger())
	{
		optionRoutes.POST("/Id=:productId", o.AdminSrvController.MiddlewareCheckIsAuth(o.AdminClient), o.ProductOptionController.CreateOption)
		optionRoutes.GET("/Id=:productId", o.ProductOptionController.GetOptions)
		optionRoutes.DELETE("/:optionId", o.ProductOptionController.DeleteOption)
		optionRoutes.GET("/optionId=:optionId", o.ProductOptionController.GetOptionByID)
		optionRoutes.PUT("/:optionId", o.ProductOptionController.UpdateOption)

	}
}
