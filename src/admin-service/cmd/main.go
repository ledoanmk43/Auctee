package main

import (
	grpc_server "chilindo/src/admin-service/cmd/grpc-admin"
	"chilindo/src/admin-service/config"
	"chilindo/src/admin-service/controller"
	"chilindo/src/admin-service/repository"
	"chilindo/src/admin-service/route"
	"chilindo/src/admin-service/service"
	"chilindo/src/admin-service/utils"
	"fmt"
	"log"
	"net"
)

const (
	adminPortForClients = ":50051"
	//adminClientPortForAuction = ":50053"
)

func main() {

	db := config.GetDB()
	newRouter := utils.Router()

	adminRepo := repository.NewAdminRepositoryDefault(db)
	adminService := service.NewAdminServiceDefault(adminRepo)
	adminController := controller.NewAdminControllerDefault(adminService)
	adminRouter := route.NewAdminRouteDefault(adminController, newRouter)
	adminRouter.GetRouter()

	go func() {
		if err := newRouter.Run(":1001"); err != nil {
			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 1001")

	}()
	lis1, err1 := net.Listen("tcp", adminPortForClients)
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}

	if err1 = grpc_server.RunGRPCServer(false, lis1); err1 != nil {
		log.Fatalf("failed to serve: %v", err1)
	}
	log.Println("gRPC server admin is running")

}
