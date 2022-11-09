package route

import (
	account_server_controller "backend/internal/auction-service/controller/account-grpc-controller"
	"backend/internal/product-service/controller"
	"github.com/gin-gonic/gin"
)

type IProductRoute interface {
	GetRouter()
}

type ProductRoute struct {
	ProductController    controller.IProductController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
	//AccountClient        account.AccountServiceClient
}

func NewProductRoute(productController controller.IProductController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController) *ProductRoute {
	return &ProductRoute{ProductController: productController, Router: router, AccountSrvController: accountSrvController}
}

func (p *ProductRoute) GetRouter() {
	productRoutes := p.Router.Group("/auctee")
	{
		productRoutes.POST("/user/product", p.AccountSrvController.MiddlewareCheckIsAuth(), p.ProductController.CreateProduct)
		productRoutes.GET("/products", p.ProductController.GetAllProducts)
		productRoutes.GET("/product/detail", p.ProductController.GetProductByProductId)
		productRoutes.PUT("/user/product/detail", p.AccountSrvController.MiddlewareCheckIsAuth(), p.ProductController.UpdateProductByProductId)
		productRoutes.DELETE("/user/product/detail", p.AccountSrvController.MiddlewareCheckIsAuth(), p.ProductController.DeleteProductByProductId)

	}
}
