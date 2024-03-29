package main

import (
	"backend/pkg/utils"
	config_account "backend/src/account-service/config"
	account_server_controller "backend/src/auction-service/controller/account-grpc-controller"
	rpcClientPayment "backend/src/payment-service/cmd/grpc-payment"
	config_payment "backend/src/payment-service/config"
	"backend/src/payment-service/controller"
	"backend/src/payment-service/repository"
	"backend/src/payment-service/route"
	"backend/src/payment-service/service"
	"fmt"
	"github.com/gin-contrib/sessions"
)

const (
	ginPort               = ":1003"
	grpcServerPortAccount = "account:50051"
	grpcServerPortAuction = "auction:50053"
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
	newRouter.Use(sessions.Sessions(config_account.CookieAuth, config_account.CookieStore))

	paymentRepository := repository.NewPaymentRepositoryDefault(db)
	paymentService := service.NewPaymentServiceDefault(paymentRepository)
	paymentController := controller.NewPaymentController(paymentService, auctionClient, accountClient)
	accountSrvCtrl := account_server_controller.NewAccountServiceController(accountClient)
	paymentRouter := route.NewPaymentRoute(paymentController, newRouter, accountSrvCtrl, accountClient)
	paymentRouter.GetRouter()

	//go func() {
	if err := newRouter.Run(ginPort); err != nil {
		fmt.Println("Open port is fail: ", err)
		return
	}
	//}()

	//_, err := tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	//if err != nil {
	//	log.Fatalf("failed to load server key pairs: %v", err)
	//}

	//if err := newRouter.RunTLS(ginPort, "localhost.pem", "localhost-key.pem"); err != nil {
	//	fmt.Println("Open port is fail: ", err)
	//	return
	//}
	fmt.Println("Server is opened on port 1003")
}
