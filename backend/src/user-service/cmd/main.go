package main

import (
	"chilindo/pkg/utils"
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/controller"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/route"
	"chilindo/src/user-service/service"
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
	addressService := service.NewAddressServiceDefault(addressRepo)
	addressController := controller.NewAddressControllerDefault(addressService)
	addressRouter := route.NewAddressRouteDefault(addressController, newRouter)
	addressRouter.GetRouter()

	if err := newRouter.Run(":8080"); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")
}
