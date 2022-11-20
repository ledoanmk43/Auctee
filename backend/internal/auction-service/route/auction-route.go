package route

import (
	"backend/internal/auction-service/controller"
	account_server_controller "backend/internal/product-service/controller/account-grpc-controller"
	"backend/pkg/pb/account"
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
		auctionRoute.PUT("/user/auction/detail", a.AccountSrvController.MiddlewareCheckIsAuth(), a.AuctionController.UpdateAuctionByAuctionId)
		auctionRoute.DELETE("/user/auction/detail", a.AccountSrvController.MiddlewareCheckIsAuth(), a.AuctionController.DeleteAuctionByAuctionId)
		auctionRoute.GET("/user/auctions", a.AuctionController.GetAllAuctionsByUserId)
		auctionRoute.GET("/auctions", a.AuctionController.GetAllAuctions)
		auctionRoute.GET("/auctions/products", a.AuctionController.GetAllAuctionsByProductName)
		auctionRoute.GET("/auction/detail", a.AuctionController.GetAuctionByAuctionId)
	}
}
