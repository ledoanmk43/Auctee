package main

import (
	grpc_server "backend/src/account-service/cmd/grpc-account"
	"backend/src/account-service/config"
	"backend/src/account-service/controller"
	"backend/src/account-service/repository"
	"backend/src/account-service/route"
	"backend/src/account-service/service"
	"backend/src/account-service/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"log"
	"net"
)

const (
	accountPortForClientsGRPC = ":50051"
)

func main() {
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	//Cookie
	newRouter.Use(sessions.SessionsMany(config.NewSessions, config.CookieStore))

	accountRepo := repository.NewAccountRepositoryDefault(db)
	accountService := service.NewAccountServiceDefault(accountRepo)
	accountController := controller.NewAccountControllerDefault(accountService)
	accountRouter := route.NewAccountRouteDefault(accountController, newRouter)
	accountRouter.GetRouter()

	addressRepo := repository.NewAddressRepositoryDefault(db)
	addressService := service.NewAddressServiceDefault(addressRepo)
	addressController := controller.NewAddressControllerDefault(addressService)
	addressRouter := route.NewAddressRouteDefault(addressController, newRouter)
	addressRouter.GetRouter()

	go func() {
		if err := newRouter.Run(":1001"); err != nil {
			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 1001")

	}()
	lis, err := net.Listen("tcp", accountPortForClientsGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpc_server.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server account is running")

}
