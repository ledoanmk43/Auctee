package main

import (
	"chilindo/pkg/utils"
	rpcClientAuction "chilindo/src/auction-service/cmd/grpc-auction"
	"chilindo/src/auction-service/config"
	"chilindo/src/auction-service/controller"
	admin_server_controller "chilindo/src/auction-service/controller/admin-grpc-controller"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/route"
	"chilindo/src/auction-service/service"
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
	adminClient := grpcClientFromAdminServer.SetUpAdminClient(grpcServerPortAdmin)

	//Product service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService, productClient)
	adminSrvCtrl := admin_server_controller.NewAdminServiceController()
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter, adminSrvCtrl, adminClient)
	auctionRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 1009")
}
