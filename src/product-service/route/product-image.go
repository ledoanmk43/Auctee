package route

import (
	"chilindo/src/product-service/controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IImageRoute interface {
	GetRouter()
}

type ImageRoute struct {
	ProductImageController controller.ProductImageController
	Router                 *gin.Engine
}

func NewImageRoute(productImageController controller.ProductImageController, router *gin.Engine) *ImageRoute {
	return &ImageRoute{ProductImageController: productImageController, Router: router}
}

func (i ImageRoute) GetRouter() {
	imageRoutes := i.Router.Group("chilindo/image")
	imageRoutes.Use(middleware.Logger())
	{
		imageRoutes.POST("/id=:productId", i.ProductImageController.CreateImage)
		imageRoutes.GET("/id=:productId", i.ProductImageController.GetImage)
		imageRoutes.DELETE("/:imageId", i.ProductImageController.DeleteImage)
		imageRoutes.GET("/imageId=:imageId", i.ProductImageController.GetImageByID)
		imageRoutes.PUT("/:imageId", i.ProductImageController.UpdateImage)
	}
}
