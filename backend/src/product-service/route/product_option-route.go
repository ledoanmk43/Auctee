package route

import (
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	"backend/src/product-service/controller"
	"github.com/gin-gonic/gin"
)

type IProductOptionRoute interface {
	GetRouter()
}

type OptionRoute struct {
	ProductOptionController controller.IProductOptionController
	Router                  *gin.Engine
	AccountSrvController    account_server_controller.IAccountServiceController
	//AccountClient           account.AccountServiceClient
}

func NewOptionRoute(productOptionController controller.IProductOptionController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController) *OptionRoute {
	return &OptionRoute{ProductOptionController: productOptionController, Router: router, AccountSrvController: accountSrvController}
}

func (o *OptionRoute) GetRouter() {
	optionRoutes := o.Router.Group("/auctee/user/:productId")
	{
		optionRoutes.POST("/option", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.CreateOption)
		optionRoutes.PUT("/option", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.UpdateOption)
		optionRoutes.DELETE("/option", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.DeleteOption)
	}
}
