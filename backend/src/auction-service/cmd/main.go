package main

import (
	"backend/pkg/utils"
	rpcClientAuction "backend/src/auction-service/cmd/grpc-auction"
	"backend/src/auction-service/config"
	"backend/src/auction-service/controller"
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	"backend/src/auction-service/repository"
	"backend/src/auction-service/route"
	"backend/src/auction-service/service"
	"fmt"
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
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService, productClient)
	accountSrvCtrl := account_server_controller.NewAccountServiceController()
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter, accountSrvCtrl, adminClient)
	auctionRouter.GetRouter()

	bidRepository := repository.NewBidRepositoryDefault(db, auctionRepository)
	bidService := service.NewBidServiceDefault(bidRepository, auctionRepository)
	bidController := controller.NewBidController(bidService)
	bidRouter := route.NewBidRoute(bidController, newRouter, accountSrvCtrl, adminClient)
	bidRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 1009")
}
