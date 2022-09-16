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

func (a *AuctionRoute) GetRouter() {
	auctionRoute := a.Router.Group("/auctee")
	{
		auctionRoute.POST("/user/auction", a.AccountSrvController.MiddlewareCheckIsAuth(), a.AuctionController.CreateAuction)
		auctionRoute.PUT("/user/auction/detail/id=:auctionId", a.AccountSrvController.MiddlewareCheckIsAuth(), a.AuctionController.UpdateAuctionByAuctionId)
		auctionRoute.DELETE("/user/auction/detail/id=:auctionId", a.AccountSrvController.MiddlewareCheckIsAuth(), a.AuctionController.DeleteAuctionByAuctionId)
		auctionRoute.GET("/auctions", a.AuctionController.GetAllAuctions)
		auctionRoute.GET("/auctions/product_name=:productName", a.AuctionController.GetAllAuctionsByProductName)
		auctionRoute.GET("/auction/detail/id=:auctionId", a.AuctionController.GetAuctionByAuctionId)
	}
}
