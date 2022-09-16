package route

import (
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	"backend/src/product-service/controller"
	"github.com/gin-gonic/gin"
)

type IProductImageRoute interface {
	GetRouter()
}

type ImageRoute struct {
	ProductImageController controller.IProductImageController
	Router                 *gin.Engine
	AccountSrvController   account_server_controller.IAccountServiceController
}

func NewImageRoute(productImageController controller.IProductImageController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController) *ImageRoute {
	return &ImageRoute{ProductImageController: productImageController, Router: router, AccountSrvController: accountSrvController}
}

func (i *ImageRoute) GetRouter() {
	imageRoutes := i.Router.Group("/auctee") //http://localhost:1002/auctee/shirt-1/image/id=1
	{
		imageRoutes.POST("/:productId/image", i.AccountSrvController.MiddlewareCheckIsAuth(), i.ProductImageController.CreateImage)
		imageRoutes.PUT("/:productId/image/id=:imageId", i.AccountSrvController.MiddlewareCheckIsAuth(), i.ProductImageController.UpdateImage)
		imageRoutes.DELETE("/:productId/image/id=:imageId", i.AccountSrvController.MiddlewareCheckIsAuth(), i.ProductImageController.DeleteImage)
		//imageRoutes.GET("/id=:productId", i.ProductImageController.GetImage)
		//imageRoutes.GET("/imageId=:imageId", i.ProductImageController.GetImageByID)
	}
}
