package route

import (
	"backend/internal/account-service/controller"
	"github.com/gin-gonic/gin"
)

type IAddressRoute interface {
	GetRouter()
}

type AddressRoute struct {
	AddressController controller.IAddressController
	Router            *gin.Engine
}

func NewAddressRouteDefault(addressController controller.IAddressController, router *gin.Engine) *AddressRoute {
	return &AddressRoute{AddressController: addressController, Router: router}
}

func (a *AddressRoute) GetRouter() {
	addressRoute := a.Router.Group("/auctee")
	{
		addressRoute.POST("/user/address", a.AddressController.CreateAddress)
		addressRoute.GET("/user/address", a.AddressController.GetAddressByAddressId)
		addressRoute.GET("/user/addresses", a.AddressController.GetAllAddresses)
		addressRoute.PUT("/user/address", a.AddressController.UpdateAddressByAddressId)
		addressRoute.DELETE("/user/address", a.AddressController.DeleteAddressByAddressId)
	}
}
