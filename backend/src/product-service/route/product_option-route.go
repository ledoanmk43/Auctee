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
	optionRoutes := o.Router.Group("/auctee")
	{
		optionRoutes.POST("/:productId/option", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.CreateOption)
		optionRoutes.PUT("/:productId/option/id=:optionId", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.UpdateOption)
		optionRoutes.DELETE("/:productId/option/id=:optionId", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.DeleteOption)
		//optionRoutes.POST("/id=:productId", o.AccountSrvController.MiddlewareCheckIsAuth(), o.ProductOptionController.CreateOption)
		//optionRoutes.GET("/id=:productId", o.ProductOptionController.GetOptions)
		//optionRoutes.DELETE("/:optionId", o.ProductOptionController.DeleteOption)
		//optionRoutes.GET("/optionId=:optionId", o.ProductOptionController.GetOptionByID)
		//optionRoutes.PUT("/:optionId", o.ProductOptionController.UpdateOption)
	}
}
