package route

import (
	"chilindo/pkg/pb/account"
	account_server_controller "chilindo/src/auction-service/controller/account-grpc-controller"
	"chilindo/src/product-service/controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	GetRouter()
}

type ProductRoute struct {
	ProductController    controller.ProductController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
	AccountClient        account.AccountServiceClient
}

func NewProductRoute(productController controller.ProductController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController, accountClient account.AccountServiceClient) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router, AccountSrvController: accountSrvController, AccountClient: accountClient}
}

func (p ProductRoute) GetRouter() {
	productRoutes := p.Router.Group("auctee/product")
	productRoutes.Use(middleware.Logger())
	{
		productRoutes.POST("/create", p.AccountSrvController.MiddlewareCheckIsAuth(p.AccountClient), p.ProductController.Insert)
		productRoutes.PUT("/:productId", p.ProductController.Update)
		productRoutes.DELETE("/:productId", p.AccountSrvController.MiddlewareCheckIsAuth(p.AccountClient), p.ProductController.Delete)
		productRoutes.GET("/:productId", p.ProductController.FindByID)
		productRoutes.GET("/", p.ProductController.All)

	}
}
