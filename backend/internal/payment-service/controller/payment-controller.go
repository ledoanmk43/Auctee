package controller

import (
	account_config "backend/internal/account-service/config"
	auction_config "backend/internal/auction-service/config"
	payment_config "backend/internal/payment-service/config"
	"backend/pkg/pb/account"
	"backend/pkg/pb/auction"
	"backend/pkg/token"
	"backend/pkg/utils"
	"strconv"

	"backend/internal/payment-service/entity"
	"backend/internal/payment-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IPaymentController interface {
	CreatePayment(ctx *gin.Context)
	UpdateAddressPayment(ctx *gin.Context)
	DeletePayment(ctx *gin.Context)
	GetAllPaymentsForWinner(ctx *gin.Context)
	GetAllPaymentsForOwner(ctx *gin.Context)
	GetPaymentByPaymentId(ctx *gin.Context)
	//CheckoutMoMo(ctx *gin.Context)
	//SetShippingStatusCompleted(ctx *gin.Context)
}

type PaymentController struct {
	PaymentService service.IPaymentService
	AuctionClient  auction.AuctionServiceClient
	AccountClient  account.AccountServiceClient
	//UserClient     user.UserServiceClient
}

func NewPaymentController(paymentService service.IPaymentService, auctionClient auction.AuctionServiceClient, accountClient account.AccountServiceClient) *PaymentController {
	return &PaymentController{PaymentService: paymentService, AuctionClient: auctionClient, AccountClient: accountClient}
}

//func (p *PaymentController) CheckoutMoMo(ctx *gin.Context) {
//
//
//}

//func (p *PaymentController) SetShippingStatusCompleted(ctx *gin.Context) {
//
//
//}

func (p *PaymentController) CreatePayment(ctx *gin.Context) {
	var paymentBody entity.Payment
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
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

	auctionId, errGetAuctionId := strconv.Atoi(ctx.Query(payment_config.Id))
	if errGetAuctionId != nil {
		log.Println("error when get auctionId: ", errGetAuctionId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	inAuction := auction.GetAuctionByIdRequest{AuctionId: uint32(auctionId)}
	resAuction, errResAuction := p.AuctionClient.GetAuctionById(ctx, &inAuction)
	if errResAuction != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errResAuction.Error(),
		})
		log.Println("CreatePayment: Error to call Auction Service rpc server", errResAuction)
		ctx.Abort()
		return
	}

	if resAuction == nil || claims.UserId != uint(resAuction.WinnerId) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no auction found",
		})
		log.Println("FindAuction: auction not found")
		ctx.Abort()
		return
	}

	inAccount := account.GetUserByUserIdRequest{
		UserId: resAuction.UserId,
	}
	resAccount, errResAccount := p.AccountClient.GetUserByUserId(ctx, &inAccount)
	if errResAccount != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errResAccount.Error(),
		})
		log.Println("CreatePayment: Error to call productService rpc server", errResAccount)
		ctx.Abort()
		return
	}

	if resAccount == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no account found",
		})
		log.Println("CreatePayment: account not found")
		ctx.Abort()
		return
	}
	//Auction
	paymentBody.ImagePath = resAuction.ImagePath
	paymentBody.Shopname = resAccount.Shopname
	paymentBody.AuctionId = uint(auctionId)
	paymentBody.WinnerId = claims.UserId          //Winner of the auction
	paymentBody.OwnerId = uint(resAuction.UserId) //Owner of the auction
	paymentBody.ProductId = resAuction.ProductId
	paymentBody.ProductName = resAuction.ProductName
	paymentBody.EndTime = resAuction.EndTime
	paymentBody.Quantity = int(resAuction.Quantity)
	paymentBody.BeforeDiscount = float64(resAuction.CurrentBid)
	paymentBody.CheckoutStatus = 1

	id, errCreatePayment := p.PaymentService.CreatePayment(&paymentBody)
	if errCreatePayment != nil {
		if errCreatePayment.Error() == "order is pending" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errCreatePayment.Error(),
				"id":      id,
			})
			log.Println("CreatePayment: Error create new payment in package controller")
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreatePayment.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (p *PaymentController) UpdateAddressPayment(ctx *gin.Context) {
	var paymentBody entity.Payment
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
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

	addressId, errGetId := strconv.Atoi(ctx.Query(payment_config.AddressId))
	if errGetId != nil {
		log.Println("error when get addressId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	inAccount := account.GetAddressByUserIdRequest{
		UserId:    uint32(claims.UserId),
		AddressId: uint32(addressId),
	}
	resAccount, errResAccount := p.AccountClient.GetAddressByUserId(ctx, &inAccount)
	if errResAccount != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errResAccount.Error(),
		})
		log.Println("CreatePayment: Error to call productService rpc server", errResAccount)
		ctx.Abort()
		return
	}

	if resAccount == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no auction found",
		})
		log.Println("CreatePayment: auction not found")
		ctx.Abort()
		return
	}
	paymentId, errGetPaymentId := strconv.Atoi(ctx.Query(payment_config.Id))
	if errGetPaymentId != nil {
		log.Println("error when get auctionId: ", errGetPaymentId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	//Address
	paymentBody.ID = uint(paymentId)
	paymentBody.Firstname = resAccount.Firstname
	paymentBody.Lastname = resAccount.Lastname
	paymentBody.Phone = resAccount.Phone
	paymentBody.Email = resAccount.Email
	paymentBody.Province = resAccount.Province
	paymentBody.District = resAccount.District
	paymentBody.SubDistrict = resAccount.SubDistrict
	paymentBody.Address = resAccount.Address
	paymentBody.TypeAddress = resAccount.TypeAddress
	paymentBody.WinnerId = claims.UserId

	errUpdatePayment := p.PaymentService.UpdateAddressPayment(&paymentBody)
	if errUpdatePayment != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdatePayment.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "payment updated",
	})
}

func (p *PaymentController) DeletePayment(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
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

	paymentId, errGetPaymentId := strconv.Atoi(ctx.Query(payment_config.Id))
	if errGetPaymentId != nil {
		log.Println("error when get auctionId: ", errGetPaymentId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	errDeletePayment := p.PaymentService.DeletePayment(uint(paymentId), claims.UserId)
	if errDeletePayment != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errDeletePayment.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "payment deleted",
	})
}

func (p *PaymentController) GetPaymentByPaymentId(ctx *gin.Context) {
	paymentId, errGetId := strconv.Atoi(ctx.Query(payment_config.Id))
	if errGetId != nil {
		log.Println("error when get auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 || claims == nil {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	paymentDetail, err := p.PaymentService.GetPaymentByPaymentId(uint(paymentId), claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "payment not found",
		})
		log.Println("GetPaymentById: Error in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, paymentDetail)

}

func (p *PaymentController) GetAllPaymentsForOwner(ctx *gin.Context) {
	page, errGetPage := strconv.Atoi(ctx.Query(auction_config.Page))
	if errGetPage != nil {
		log.Println("error when get auctionId: ", errGetPage)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
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

	ownerId := claims.UserId
	payments, err := p.PaymentService.GetAllPaymentsForOwner(page, ownerId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, payments)
}

func (p *PaymentController) GetAllPaymentsForWinner(ctx *gin.Context) {
	page, errGetPage := strconv.Atoi(ctx.Query(auction_config.Page))
	if errGetPage != nil {
		log.Println("error when get auctionId: ", errGetPage)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account_config.CookieAuth)
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

	payments, err := p.PaymentService.GetAllPaymentsForWinner(page, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, payments)
}
