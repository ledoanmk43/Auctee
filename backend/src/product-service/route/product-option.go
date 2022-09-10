package route

import (
	"chilindo/pkg/pb/account"
	account_server_controller "chilindo/src/auction-service/controller/account-grpc-controller"
	"chilindo/src/product-service/controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IProductOptionRoute interface {
	GetRouter()
}

type OptionRoute struct {
	ProductOptionController controller.ProductOptionController
	Router                  *gin.Engine
	AccountSrvController    account_server_controller.IAccountServiceController
	AccountClient           account.AccountServiceClient
}

func NewOptionRoute(productOptionController controller.ProductOptionController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController, accountClient account.AccountServiceClient) *OptionRoute {
	return &OptionRoute{ProductOptionController: productOptionController, Router: router, AccountSrvController: accountSrvController, AccountClient: accountClient}
}

func (o OptionRoute) GetRouter() {
	optionRoutes := o.Router.Group("auctee/option")
	optionRoutes.Use(middleware.Logger())
	{
		optionRoutes.POST("/id=:productId", o.AccountSrvController.MiddlewareCheckIsAuth(o.AccountClient), o.ProductOptionController.CreateOption)
		optionRoutes.GET("/id=:productId", o.ProductOptionController.GetOptions)
		optionRoutes.DELETE("/:optionId", o.ProductOptionController.DeleteOption)
		optionRoutes.GET("/optionId=:optionId", o.ProductOptionController.GetOptionByID)
		optionRoutes.PUT("/:optionId", o.ProductOptionController.UpdateOption)

	}
}
