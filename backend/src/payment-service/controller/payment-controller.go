package controller

import (
	"backend/pkg/pb/account"
	"backend/pkg/pb/auction"
	"backend/pkg/token"
	"backend/pkg/utils"
	account_config "backend/src/account-service/config"
	auction_config "backend/src/auction-service/config"
	payment_config "backend/src/payment-service/config"
	"backend/src/payment-service/entity"
	"backend/src/payment-service/service"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sony/sonyflake"
	"log"
	"net/http"
	"os"
	"strconv"
)

type IPaymentController interface {
	CreatePayment(ctx *gin.Context)
	UpdateAddressPayment(ctx *gin.Context)
	DeletePayment(ctx *gin.Context)
	GetAllPaymentsForWinner(ctx *gin.Context)
	GetAllPaymentsForOwner(ctx *gin.Context)
	GetPaymentByPaymentId(ctx *gin.Context)
	CheckoutMoMo(ctx *gin.Context)
	CheckoutCOD(ctx *gin.Context)
	SetShippingStatusCompleted(ctx *gin.Context)
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

func (p *PaymentController) SetShippingStatusCompleted(ctx *gin.Context) {
	var paymentBody entity.Payment
	if err := ctx.ShouldBindJSON(&paymentBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
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

	paymentId := ctx.Query(payment_config.Id)

	//Address
	paymentBody.Id = paymentId
	paymentBody.WinnerId = claims.UserId
	paymentBody.ShippingStatus = utils.BoolAddr(true)
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

func (p *PaymentController) CheckoutMoMo(ctx *gin.Context) {
	var paymentBody entity.Payment
	if err := ctx.ShouldBindJSON(&paymentBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	// Get user id
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

	// Checking detail payment before Checkout with MoMo
	paymentDetail, err := p.PaymentService.GetPaymentByPaymentId(paymentBody.Id, claims.UserId)
	if err != nil || paymentDetail == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "payment not found",
		})
		ctx.Abort()
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env in auction file")
	}
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	reqID, _ := flake.NextID()
	//total := fmt.Sprintf("%f", paymentBody.Total)

	redirectToThisURL := fmt.Sprintf("http://localhost:3000/auctee/user/order/?id=%s",
		paymentBody.Id,
	)

	var endpoint = os.Getenv("MOMO_EP")
	var accessKey = os.Getenv("MOMO_ACCESS_KEY")
	var secretKey = os.Getenv("MOMO_SECRET_KEY")
	var partnerCode = "MOMO"
	var partnerName = "Đấu giá trực tuyến"
	var storeId = "MoMoTestStore"
	var requestId = strconv.FormatUint(reqID, 16)
	var amount = int64(paymentBody.Total)
	var orderId = paymentBody.Id
	var orderInfo = fmt.Sprintf("Đơn hàng: %s",
		paymentBody.ProductName,
	)
	var redirectUrl = redirectToThisURL
	var ipnUrl = "http://localhost:3000/auctee/user/checkout/shipping-status-payment"
	var requestType = "payWithMethod"
	var extraData = ""
	var orderGroupId = ""
	var autoCapture = true
	var lang = "vi"

	//build raw signature
	var rawSignature bytes.Buffer
	rawSignature.WriteString("accessKey=")
	rawSignature.WriteString(accessKey)
	rawSignature.WriteString("&amount=")
	rawSignature.WriteString(strconv.FormatInt(amount, 10))
	rawSignature.WriteString("&extraData=")
	rawSignature.WriteString(extraData)
	rawSignature.WriteString("&ipnUrl=")
	rawSignature.WriteString(ipnUrl)
	rawSignature.WriteString("&orderId=")
	rawSignature.WriteString(orderId)
	rawSignature.WriteString("&orderInfo=")
	rawSignature.WriteString(orderInfo)
	rawSignature.WriteString("&partnerCode=")
	rawSignature.WriteString(partnerCode)
	rawSignature.WriteString("&redirectUrl=")
	rawSignature.WriteString(redirectUrl)
	rawSignature.WriteString("&requestId=")
	rawSignature.WriteString(requestId)
	rawSignature.WriteString("&requestType=")
	rawSignature.WriteString(requestType)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hmac := hmac.New(sha256.New, []byte(secretKey))
	// Write Data to it
	hmac.Write(rawSignature.Bytes())

	// Get result and encode as hexadecimal string
	signature := hex.EncodeToString(hmac.Sum(nil))

	var payload = entity.Payload{
		PartnerCode:  partnerCode,
		AccessKey:    accessKey,
		RequestID:    requestId,
		Amount:       int64(paymentBody.Total),
		RequestType:  requestType,
		RedirectUrl:  redirectUrl,
		IpnUrl:       ipnUrl,
		OrderID:      orderId,
		StoreId:      storeId,
		PartnerName:  partnerName,
		OrderGroupId: orderGroupId,
		AutoCapture:  autoCapture,
		Lang:         lang,
		OrderInfo:    orderInfo,
		ExtraData:    extraData,
		Signature:    signature,
	}
	log.Println(payload)
	var jsonPayload []byte

	jsonPayload, err = json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Payload: " + string(jsonPayload))

	//send HTTP to momo endpoint
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}

	//result
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}
	fmt.Println("Response from Momo: ", result)
	ctx.JSON(http.StatusOK, gin.H{
		"redirectURL": result["shortLink"],
	})

}

func (p *PaymentController) CheckoutCOD(ctx *gin.Context) {
	var paymentBody entity.Payment
	if err := ctx.ShouldBindJSON(&paymentBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	// Get user id
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

	addressId, errGetId := strconv.Atoi(ctx.Query("id"))
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

	//Address

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
	paymentBody.CheckoutStatus = 3
	paymentBody.PaymentMethod = "COD"
	paymentBody.Total = 0
	paymentBody.ShippingStatus = utils.BoolAddr(false)
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
		"message": "ship cod confirmed",
	})

}

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

	if err := ctx.ShouldBindJSON(&paymentBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("Error to ShouldBindJSON controller", err)
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

	paymentId := ctx.Query(payment_config.Id)

	errDeletePayment := p.PaymentService.DeletePayment(paymentId, claims.UserId)
	if errDeletePayment != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errDeletePayment.Error(),
		})
		log.Println("CreatePayment: Error create new payment in package controller")
		ctx.Abort()
		return
	}
	inAccount := account.UpdateHonorPointRequest{
		UserId: uint32(claims.UserId),
		CaseId: 2,
	}
	_, err := p.AccountClient.UpdateHonorPoint(ctx, &inAccount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("CreatePayment: Error to call productService rpc server", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "payment deleted",
	})
}

func (p *PaymentController) GetPaymentByPaymentId(ctx *gin.Context) {
	paymentId := ctx.Query(payment_config.Id)

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

	paymentDetail, err := p.PaymentService.GetPaymentByPaymentId(paymentId, claims.UserId)
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
