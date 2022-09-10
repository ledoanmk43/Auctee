package route

import (
	"backend/pkg/pb/account"
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
	AccountClient        account.AccountServiceClient
}

func NewBidRoute(bidController controller.IBidController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController, accountClient account.AccountServiceClient) *BidRoute {
	return &BidRoute{BidController: bidController, Router: router, AccountSrvController: accountSrvController, AccountClient: accountClient}
}

func (b BidRoute) GetRouter() {
	bidRoute := b.Router.Group("auctee/bid/productId=:productId/auctionId=:auctionId")
	{
		bidRoute.POST("/create", b.AccountSrvController.MiddlewareCheckIsAuth(b.AccountClient), b.BidController.CreateBid)
	}
}
