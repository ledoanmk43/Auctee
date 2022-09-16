package main

import (
	"backend/pkg/utils"
	"backend/src/account-service/config"
	grpc_product "backend/src/product-service/cmd/grpc-product"
	rpcClient "backend/src/product-service/cmd/grpc-product"
	"backend/src/product-service/controller"
	account_server_controller "backend/src/product-service/controller/account-grpc-controller"
	"backend/src/product-service/repository"
	"backend/src/product-service/route"
	"backend/src/product-service/service"
	"github.com/gin-contrib/sessions"
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

	//Cookie
	newRouter.Use(sessions.SessionsMany(config.NewSessions, config.CookieStore))

	productRepository := repository.NewProductRepositoryDefault(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	adminSrvCtrl := account_server_controller.NewAccountServiceController(accountClient)
	productRouter := route.NewProductRoute(productController, newRouter, adminSrvCtrl)
	productRouter.GetRouter()

	productOptionRepository := repository.NewProductOptionRepository(db)
	productOptionService := service.NewProductOptionService(productOptionRepository)
	productOptionController := controller.NewProductOptionController(productOptionService)
	optionRouter := route.NewOptionRoute(productOptionController, newRouter, adminSrvCtrl)
	optionRouter.GetRouter()

	productImageRepository := repository.NewProductImageRepository(db)
	productImageService := service.NewProductImageService(productImageRepository)
	productImageController := controller.NewProductImageController(productImageService)
	imageRouter := route.NewImageRoute(productImageController, newRouter, adminSrvCtrl)
	imageRouter.GetRouter()

	go func() {
		if err := newRouter.Run(ginPort); err != nil {

			log.Println("Open port is fail")
			return
		}
		log.Println("Server product is opened on port 1002")
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
