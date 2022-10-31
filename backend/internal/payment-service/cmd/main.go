package main

import (
	config_account "backend/internal/account-service/config"
	account_server_controller "backend/internal/auction-service/controller/account-grpc-controller"
	rpcClientPayment "backend/internal/payment-service/cmd/grpc-payment"
	config_payment "backend/internal/payment-service/config"
	"backend/internal/payment-service/controller"
	"backend/internal/payment-service/repository"
	"backend/internal/payment-service/route"
	"backend/internal/payment-service/service"
	"backend/pkg/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
)

const (
	ginPort               = ":1003"
	grpcServerPortAccount = "localhost:50051"
	grpcServerPortAuction = "localhost:50053"
)

func main() {
	//Create new gRPC Client from Auction Server
	grpcClientFromAuctionServer := rpcClientPayment.NewRPCClient()
	auctionClient := grpcClientFromAuctionServer.SetUpAuctionClient(grpcServerPortAuction)

	//Create new gRPC Client from Account Server
	grpcClientFromAccountServer := rpcClientPayment.NewRPCClient()
	accountClient := grpcClientFromAccountServer.SetUpAccountClient(grpcServerPortAccount)

	//Product service DB
	db := config_payment.GetDB()
	defer config_payment.CloseDatabase(db)
	newRouter := utils.Router()

	//Cookie
	newRouter.Use(sessions.SessionsMany(config_account.NewSessions, config_account.CookieStore))

	paymentRepository := repository.NewPaymentRepositoryDefault(db)
	paymentService := service.NewPaymentServiceDefault(paymentRepository)
	paymentController := controller.NewPaymentController(paymentService, auctionClient, accountClient)
	accountSrvCtrl := account_server_controller.NewAccountServiceController(accountClient)
	paymentRouter := route.NewPaymentRoute(paymentController, newRouter, accountSrvCtrl, accountClient)
	paymentRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 1003")
}
