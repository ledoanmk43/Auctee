package route

import (
	"chilindo/pkg/pb/admin"
	"chilindo/src/auction-service/controller"
	admin_server_controller "chilindo/src/auction-service/controller/admin-grpc-controller"
	"github.com/gin-gonic/gin"
)

type IBidRoute interface {
	GetRouter()
}

type BidRoute struct {
	BidController      controller.IBidController
	Router             *gin.Engine
	AdminSrvController admin_server_controller.IAdminServiceController
	AdminClient        admin.AdminServiceClient
}

func NewBidRoute(bidController controller.IBidController, router *gin.Engine, adminSrvController admin_server_controller.IAdminServiceController, adminClient admin.AdminServiceClient) *BidRoute {
	return &BidRoute{BidController: bidController, Router: router, AdminSrvController: adminSrvController, AdminClient: adminClient}
}

func (b BidRoute) GetRouter() {
	bidRoute := b.Router.Group("/chilindo/bid/productId=:productId/auctionId=:auctionId")
	{
		bidRoute.POST("/create", b.AdminSrvController.CheckIsAuth(b.AdminClient), b.BidController.CreateBid)
	}
}
