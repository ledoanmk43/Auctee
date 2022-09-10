package route

import (
	"backend/src/account-service/controller"
	"backend/src/account-service/middleware"
	"github.com/gin-gonic/gin"
)

type IAddressRoute interface {
	GetRouter()
}

type AddressRoute struct {
	AddressController controller.IAddressController
	Router            *gin.Engine
	JWTMiddleware     *middleware.SMiddleWare
}

func NewAddressRouteDefault(addressController controller.IAddressController, router *gin.Engine) *AddressRoute {
	return &AddressRoute{AddressController: addressController, Router: router}
}

func (a AddressRoute) GetRouter() {
	addressRoute := a.Router.Group("/auctee")
	{
		addressRoute.POST("/user/address", a.AddressController.CreateAddress)
		addressRoute.GET("/user/address/id=:id", a.AddressController.GetAddressByAddressId)
		addressRoute.GET("/user/address", a.AddressController.GetAllAddresses)
		addressRoute.PUT("/user/address/id=:id", a.AddressController.UpdateAddressByAddressId)
		addressRoute.DELETE("/user/address/id=:id", a.AddressController.DeleteAddressByAddressId)
	}
}
