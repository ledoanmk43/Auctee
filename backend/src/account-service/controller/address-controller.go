package controller

import (
	"backend/pkg/token"
	"backend/pkg/utils"
	"backend/src/account-service/config"
	"backend/src/account-service/dto"
	"backend/src/account-service/entity"
	"backend/src/account-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IAddressController interface {
	CreateAddress(ctx *gin.Context)
	UpdateAddressByAddressId(ctx *gin.Context)
	GetAddressByAddressId(ctx *gin.Context)
	GetAllAddresses(ctx *gin.Context)
	DeleteAddressByAddressId(ctx *gin.Context)
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

func (a *AddressController) UpdateAddressByAddressId(ctx *gin.Context) {
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

	addressId, errGetId := strconv.Atoi(ctx.Query(config.Id))
	if errGetId != nil {
		log.Println("error in get address by addressId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	var updateBody *dto.UpdateAddressDTO
	err := ctx.ShouldBindJSON(&updateBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("SignIn: Error ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}

	updateBody.Id = uint(addressId)
	errUpdate := a.AddressService.UpdateAddress(claims.UserId, updateBody)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("Update User: Error in package controller: ", errUpdate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "address updated",
	})
}

func (a *AddressController) GetAllAddresses(ctx *gin.Context) {
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
	userId := claims.UserId

	addresses, err := a.AddressService.GetAllAddresses(userId)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{
			"message": err,
		})
		log.Println("GetAddress: Error Get Address in package controller")
		ctx.Abort()
		return
	}
	ctx.JSONP(http.StatusOK, addresses)
}

func (a *AddressController) DeleteAddressByAddressId(ctx *gin.Context) {
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

	addressId, errGetId := strconv.Atoi(ctx.Query(config.Id))
	if errGetId != nil {
		log.Println("error in get address by addressId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	errUpdate := a.AddressService.DeleteAddress(claims.UserId, uint(addressId))
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("Update User: Error in package controller: ", errUpdate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "address deleted",
	})
}

func (a *AddressController) GetAddressByAddressId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	addressId, errGetId := strconv.Atoi(ctx.Query(config.Id))
	if errGetId != nil {
		log.Println("error in get address by addressId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
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

	address, errGet := a.AddressService.GetAddressByAddressId(uint(addressId), claims.UserId)
	if errGet != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errGet.Error(),
		})
		log.Println("Get User: Error in package controller", errGet)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, address)
}
