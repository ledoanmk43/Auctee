package main

import (
	"backend/pkg/utils"
	"backend/src/account-service/config"
	grpc_auction "backend/src/auction-service/cmd/grpc-auction"
	rpcClientAuction "backend/src/auction-service/cmd/grpc-auction"
	config_auction "backend/src/auction-service/config"
	"backend/src/auction-service/controller"
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	"backend/src/auction-service/repository"
	"backend/src/auction-service/route"
	"backend/src/auction-service/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"log"
	"net"
)

const (
	ginPort               = ":1009"
	grpcServerPortAccount = "account:50051"
	grpcServerPortProduct = "product:50052"
	grpcServerPort        = ":50053"
)

func main() {
	//Create new gRPC Client from Product Server
	grpcClientFromProductServer := rpcClientAuction.NewRPCClient()
	productClient := grpcClientFromProductServer.SetUpProductClient(grpcServerPortProduct)

	//Create new gRPC Client from Account Server
	grpcClientFromAccountServer := rpcClientAuction.NewRPCClient()
	accountClient := grpcClientFromAccountServer.SetUpAccountClient(grpcServerPortAccount)

	//Auction service DB
	db := config_auction.GetDB()
	defer config_auction.CloseDatabase(db)
	newRouter := utils.Router()

	//Cookie
	newRouter.Use(sessions.Sessions(config.CookieAuth, config.CookieStore))

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService, productClient)
	accountSrvCtrl := account_server_controller.NewAccountServiceController(accountClient)
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter, accountSrvCtrl, accountClient)
	auctionRouter.GetRouter()

	bidRepository := repository.NewBidRepositoryDefault(db, auctionRepository)
	bidService := service.NewBidServiceDefault(bidRepository, auctionRepository)
	bidController := controller.NewBidController(bidService, auctionService, productClient)
	bidRouter := route.NewBidRoute(bidController, newRouter, accountSrvCtrl)
	bidRouter.GetRouter()

	go func() {
		if err := newRouter.Run(ginPort); err != nil {

			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 1009")
	}()
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen from auction service: %v", err)
	}

	if err = grpc_auction.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server auction is running")
}
