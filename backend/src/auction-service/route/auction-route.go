package route

import (
	"backend/pkg/pb/account"
	"backend/src/auction-service/controller"
	account_server_controller "backend/src/product-service/controller/account-grpc-controller"
	"github.com/gin-gonic/gin"
)

type IAuctionRoute interface {
	GetRouter()
}

type AuctionRoute struct {
	AuctionController    controller.IAuctionController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
	AccountClient        account.AccountServiceClient
}

func NewAuctionRoute(auctionController controller.IAuctionController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController, accountClient account.AccountServiceClient) *AuctionRoute {
	return &AuctionRoute{AuctionController: auctionController, Router: router, AccountSrvController: accountSrvController, AccountClient: accountClient}
}

func (a AuctionRoute) GetRouter() {
	auctionRoute := a.Router.Group("/backend/auction/")
	{
		auctionRoute.POST("/create", a.AccountSrvController.MiddlewareCheckIsAuth(a.AccountClient), a.AuctionController.CreateAuction)
	}
}
