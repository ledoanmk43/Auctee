package main

import (
	"backend/pkg/utils"
	grpc_server "backend/src/account-service/cmd/grpc-account"
	"backend/src/account-service/config"
	"backend/src/account-service/controller"
	"backend/src/account-service/repository"
	"backend/src/account-service/route"
	"backend/src/account-service/service"
	"github.com/gin-contrib/sessions"
	"log"
	"net"
)

const (
	ginPort        = ":1001"
	grpcServerPort = ":50051"
)

func main() {
	//Account service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	newRouter.Use(sessions.Sessions(config.CookieAuth, config.CookieStore))
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
		if err := newRouter.Run(ginPort); err != nil {
			log.Println("Open port is fail")
			return
		}
		log.Println("Server is opened on port 1001")

	}()
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpc_server.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server account is running")

}
