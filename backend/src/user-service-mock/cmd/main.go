package main

import (
	"chilindo/pkg/utils"
	controller2 "chilindo/src/account-service/controller"
	route2 "chilindo/src/account-service/route"
	service2 "chilindo/src/account-service/service"
	"chilindo/src/user-service-mock/config"
	"chilindo/src/user-service-mock/controller"
	"chilindo/src/user-service-mock/repository"
	"chilindo/src/user-service-mock/route"
	"chilindo/src/user-service-mock/service"
	"fmt"
)

func main() {
	db := config.GetDB()
	newRouter := utils.Router()

	userRepo := repository.NewUserRepositoryDefault(db)
	userService := service.NewUserServiceDefault(userRepo)
	userController := controller.NewUserControllerDefault(userService)
	userRouter := route.NewUserRouteDefault(userController, newRouter)
	userRouter.GetRouter()

	addressRepo := repository.NewAddressRepositoryDefault(db)
	addressService := service2.NewAddressServiceDefault(addressRepo)
	addressController := controller2.NewAddressControllerDefault(addressService)
	addressRouter := route2.NewAddressRouteDefault(addressController, newRouter)
	addressRouter.GetRouter()

	if err := newRouter.Run(":8080"); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")
}
