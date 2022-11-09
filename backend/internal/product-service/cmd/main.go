package main

import (
	config_account "backend/internal/account-service/config"
	grpc_product "backend/internal/product-service/cmd/grpc-product"
	rpcClient "backend/internal/product-service/cmd/grpc-product"
	config_product "backend/internal/product-service/config"
	"backend/internal/product-service/controller"
	account_server_controller "backend/internal/product-service/controller/account-grpc-controller"
	"backend/internal/product-service/repository"
	"backend/internal/product-service/route"
	"backend/internal/product-service/service"
	"backend/pkg/utils"
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
	db := config_product.GetDB()
	defer config_product.CloseDatabase(db)
	newRouter := utils.Router()

	//Cookie
	newRouter.Use(sessions.SessionsMany(config_account.NewSessions, config_account.CookieStore))

	productOptionRepository := repository.NewProductOptionRepository(db)
	productImageRepository := repository.NewProductImageRepository(db)
	productRepository := repository.NewProductRepositoryDefault(db, productOptionRepository, productImageRepository)
	
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	adminSrvCtrl := account_server_controller.NewAccountServiceController(accountClient)
	productRouter := route.NewProductRoute(productController, newRouter, adminSrvCtrl)
	productRouter.GetRouter()

	productOptionService := service.NewProductOptionService(productOptionRepository)
	productOptionController := controller.NewProductOptionController(productOptionService)
	optionRouter := route.NewOptionRoute(productOptionController, newRouter, adminSrvCtrl)
	optionRouter.GetRouter()

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
