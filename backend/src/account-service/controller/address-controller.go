package controller

import (
	"backend/pkg/token"
	"backend/pkg/utils"
	"backend/src/account-service/config"
	"backend/src/account-service/entity"
	"backend/src/account-service/service"
	"backend/src/user-service-mock/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IAddressController interface {
	CreateAddress(c *gin.Context)
	UpdateAddressByAddressId(c *gin.Context)
	GetAddressByAddressId(c *gin.Context)
	GetAllAddresses(c *gin.Context)
	DeleteAddressByAddressId(c *gin.Context)
}

type AddressController struct {
	AddressService service.IAddressService
}

func NewAddressControllerDefault(addressService service.IAddressService) *AddressController {
	return &AddressController{AddressService: addressService}
}

func (a *AddressController) CreateAddress(ctx *gin.Context) {
	var newAddress *entity.Address
	errDTO := ctx.ShouldBindJSON(&newAddress)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("CreateAddress: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	newAddress.UserId = claims.UserId
	errCreateAddress := a.AddressService.CreateAddress(newAddress)
	if errCreateAddress != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAddress.Error(),
		})
		log.Println("CreateAddress: Error create new address in package controller")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "address created",
	})
}

func (a *AddressController) UpdateAddressByAddressId(c *gin.Context) {
	var updateAddress *entity.Address
	if err := c.ShouldBindJSON(&updateAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("UpdateAddress: Error ShouldBindJSON in package controller", err)
		return
	}
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	addressId, err := strconv.Atoi(c.Param(config.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	updateAddress.ID = uint(addressId)
	updateAddress.UserId = userId.(uint)

	updatedAddress, errUpdate := a.AddressService.UpdateAddress(updateAddress)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("UpdateAddress: Error update address in package controller")
		c.Abort()
		return
	}
	fmt.Println("check here")

	if updatedAddress == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found address",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, updatedAddress)
}

func (a *AddressController) GetAllAddresses(c *gin.Context) {
	var dTo dto.GetAddressDTO
	userId, oke := c.Get(config.UserId)
	dTo.UserId = userId.(uint)
	if !oke {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address is fail",
		})
		log.Println("GetAddress: Error Get Address in package controller")
		c.Abort()
		return
	}
	address, err := a.AddressService.GetAllAddresses(&dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address is fail",
		})
		log.Println("GetAddress: Error Get Address in package controller")
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, address)
}

func (a *AddressController) DeleteAddressByAddressId(c *gin.Context) {
	var dTo dto.GetAddressByIdDTO
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	addressId, err := strconv.Atoi(c.Param(config.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	dTo.AddressId = uint(addressId)
	dTo.UserId = userId.(uint)
	errDelete := a.AddressService.DeleteAddress(&dTo)
	if errDelete != nil {
		c.JSONP(http.StatusUnauthorized, gin.H{
			"Message": "DeleteAddress: not exist id address to delete",
		})
		log.Println("DeleteAddress: Error to delete Address in package controller")
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"Message": "success",
	})
}

func (a *AddressController) GetAddressByAddressId(c *gin.Context) {
	var dTo dto.GetAddressByIdDTO
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	addressId, err := strconv.Atoi(c.Param(config.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	dTo.AddressId = uint(addressId)
	dTo.UserId = userId.(uint)
	address, err := a.AddressService.GetAddressById(&dTo)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Get Address by ID fail",
		})
		log.Println("GetAddressById: Error in package controllers", err)
		c.Abort()
		return
	}
	c.JSONP(http.StatusOK, address)
}
