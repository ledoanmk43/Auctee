package main

import (
	"backend/pkg/utils"
	config_account "backend/src/account-service/config"
	rpcClientAuction "backend/src/auction-service/cmd/grpc-auction"
	config_auction "backend/src/auction-service/config"
	"backend/src/auction-service/controller"
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	"backend/src/auction-service/repository"
	"backend/src/auction-service/route"
	"backend/src/auction-service/service"
	"fmt"
	"github.com/gin-contrib/sessions"
)

const (
	ginPort               = ":1009"
	grpcServerPortAdmin   = "localhost:50051"
	grpcServerPortProduct = "localhost:50052"
)

func main() {
	//Create new gRPC Client
	grpcClientFromProductServer := rpcClientAuction.NewRPCClient()
	productClient := grpcClientFromProductServer.SetUpProductClient(grpcServerPortProduct)

	//Create new gRPC Client
	grpcClientFromAdminServer := rpcClientAuction.NewRPCClient()
	adminClient := grpcClientFromAdminServer.SetUpAccountClient(grpcServerPortAdmin)

	//Product service DB
	db := config_auction.GetDB()
	defer config_auction.CloseDatabase(db)
	newRouter := utils.Router()

	//Cookie
	newRouter.Use(sessions.SessionsMany(config_account.NewSessions, config_account.CookieStore))

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService, productClient)
	accountSrvCtrl := account_server_controller.NewAccountServiceController(adminClient)
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter, accountSrvCtrl, adminClient)
	auctionRouter.GetRouter()

	bidRepository := repository.NewBidRepositoryDefault(db, auctionRepository)
	bidService := service.NewBidServiceDefault(bidRepository, auctionRepository)
	bidController := controller.NewBidController(bidService)
	bidRouter := route.NewBidRoute(bidController, newRouter, accountSrvCtrl)
	bidRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 1009")
}
