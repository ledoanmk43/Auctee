package main

import (
	"chilindo/pkg/utils"
	grpc_product "chilindo/src/product-service/cmd/grpc-product"
	rpcClient "chilindo/src/product-service/cmd/grpc-product"
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/controller"
	account_server_controller "chilindo/src/product-service/controller/account-grpc-controller"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/route"
	"chilindo/src/product-service/service"
	"fmt"
	"log"
	"net"
)

const (
	ginPort        = ":1002"
	grpcServerPort = "localhost:50052"
)

func main() {
	//Create new gRPC Client
	grpcClient := rpcClient.NewRPCClient()
	accountClient := grpcClient.SetUpAccountClient()

	//adminClient := admin.NewAdminServiceClient(conn)

	//Product service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	adminSrvCtrl := account_server_controller.NewAccountServiceController()
	productRouter := route.NewProductRoute(productController, newRouter, adminSrvCtrl, accountClient)
	productRouter.GetRouter()

	productOptionRepository := repository.NewProductOptionRepository(db)
	productOptionService := service.NewProductOptionService(productOptionRepository)
	productOptionController := controller.NewProductOptionController(productOptionService)
	optionRouter := route.NewOptionRoute(productOptionController, newRouter, adminSrvCtrl, accountClient)
	optionRouter.GetRouter()

	productImageRepository := repository.NewProductImageRepository(db)
	productImageService := service.NewProductImageService(productImageRepository)
	productImageController := controller.NewProductImageController(productImageService)
	imageRouter := route.NewImageRoute(productImageController, newRouter)
	imageRouter.GetRouter()

	go func() {
		if err := newRouter.Run(ginPort); err != nil {

			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server product is opened on port 1002")
	}()
	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen from product service: %v", err)
	}

	if err = grpc_product.RunGRPCServer(false, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server product is running")

}
