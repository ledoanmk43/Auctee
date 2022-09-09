package main

import (
	grpc_server "chilindo/src/account-service/cmd/grpc-admin"
	"chilindo/src/account-service/config"
	"chilindo/src/account-service/controller"
	"chilindo/src/account-service/repository"
	"chilindo/src/account-service/route"
	"chilindo/src/account-service/service"
	"chilindo/src/account-service/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"log"
	"net"
)

const (
	userPortForClients = ":50051"
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

	go func() {
		if err := newRouter.Run(":1001"); err != nil {
			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 1001")

	}()
	lis, err := net.Listen("tcp", userPortForClients)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpc_server.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server account is running")

}
