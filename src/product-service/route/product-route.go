package route

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/product-service/controller"
	admin_server_controller "chilindo/src/product-service/controller/admin-grpc-controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	GetRouter()
}

type ProductRoute struct {
	ProductController  controller.ProductController
	Router             *gin.Engine
	AdminSrvController admin_server_controller.IAdminServiceController
	AdminClient        admin.AdminServiceClient
}

func NewProductRoute(productController controller.ProductController, router *gin.Engine, adminSrvController admin_server_controller.IAdminServiceController, adminClient admin.AdminServiceClient) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router, AdminSrvController: adminSrvController, AdminClient: adminClient}
}

func (p ProductRoute) GetRouter() {
	productRoutes := p.Router.Group("chilindo/product")
	productRoutes.Use(middleware.Logger())
	{
		productRoutes.POST("/create", p.AdminSrvController.CheckIsAuth(p.AdminClient), p.ProductController.Insert)
		productRoutes.PUT("/:productId", p.ProductController.Update)
		productRoutes.DELETE("/:productId", p.ProductController.Delete)
		productRoutes.GET("/:productId", p.ProductController.FindByID)
		productRoutes.GET("/", p.ProductController.All)

	}
}
