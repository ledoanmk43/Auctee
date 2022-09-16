package route

import (
	"backend/src/auction-service/controller"
	account_server_controller "backend/src/product-service/controller/account-grpc-controller"
	"github.com/gin-gonic/gin"
)

type IBidRoute interface {
	GetRouter()
}

type BidRoute struct {
	BidController        controller.IBidController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
}

func NewBidRoute(bidController controller.IBidController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController) *BidRoute {
	return &BidRoute{BidController: bidController, Router: router, AccountSrvController: accountSrvController}
}

func (b *BidRoute) GetRouter() {
	bidRoute := b.Router.Group("/auctee")
	{
		bidRoute.POST("/product/productId=:productId&auctionId=auctionId", b.AccountSrvController.MiddlewareCheckIsAuth(), b.BidController.CreateBid)
	}
}
