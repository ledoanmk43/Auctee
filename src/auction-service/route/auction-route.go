package route

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/auction-service/controller"
	admin_server_controller "chilindo/src/auction-service/controller/admin-grpc-controller"
	"github.com/gin-gonic/gin"
)

type IAuctionRoute interface {
	GetRouter()
}

type AuctionRoute struct {
	AuctionController  controller.IAuctionController
	Router             *gin.Engine
	AdminSrvController admin_server_controller.IAdminServiceController
	AdminClient        admin.AdminServiceClient
}

func NewAuctionRoute(auctionController controller.IAuctionController, router *gin.Engine, adminSrvController admin_server_controller.IAdminServiceController, adminClient admin.AdminServiceClient) *AuctionRoute {
	return &AuctionRoute{AuctionController: auctionController, Router: router, AdminSrvController: adminSrvController, AdminClient: adminClient}
}

func (a AuctionRoute) GetRouter() {
	auctionRoute := a.Router.Group("/chilindo/auction/")
	{
		auctionRoute.POST("/create", a.AdminSrvController.CheckIsAuth(a.AdminClient), a.AuctionController.CreateAuction)
	}
}
