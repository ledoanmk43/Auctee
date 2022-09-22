package controller

import (
	"backend/pkg/pb/product"
	"backend/pkg/token"
	"backend/pkg/utils"
	account "backend/src/account-service/config"
	auction "backend/src/auction-service/config"
	"backend/src/auction-service/entity"
	"backend/src/auction-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IAuctionController interface {
	CreateAuction(ctx *gin.Context)
	UpdateAuctionByAuctionId(ctx *gin.Context)
	DeleteAuctionByAuctionId(ctx *gin.Context)
	GetAuctionByAuctionId(ctx *gin.Context)
	GetAllAuctions(ctx *gin.Context)
	GetAllAuctionsByProductName(ctx *gin.Context)
}

type AuctionController struct {
	AuctionService service.IAuctionService
	ProductClient  product.ProductServiceClient
	//UserClient     user.UserServiceClient
}

func NewAuctionController(auctionService service.IAuctionService, productClient product.ProductServiceClient) *AuctionController {
	return &AuctionController{AuctionService: auctionService, ProductClient: productClient}
}

func (a *AuctionController) CreateAuction(ctx *gin.Context) {
	var auctionBody *entity.Auction
	if err := ctx.ShouldBindJSON(&auctionBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("Error to ShouldBindJSON controller", err)
		ctx.Abort()
		return
	}

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	in := product.GetProductByIdRequest{ProductId: auctionBody.ProductId}
	res, errRes := a.ProductClient.GetProductById(ctx, &in)
	if errRes != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errRes.Error(),
		})
		log.Println("CreateAuction: Error to call productService rpc server", errRes)
		ctx.Abort()
		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no product found",
		})
		log.Println("CreateAuction: product not found")
		ctx.Abort()
		return
	}
	if res.UserId != uint32(claims.UserId) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "product not found",
		})
		ctx.Abort()
		return
	}

	auctionBody.UserId = claims.UserId
	if auctionBody.CurrentBid <= float64(res.MinPrice) {
		auctionBody.CurrentBid = float64(res.MinPrice)
	}

	if len(res.Path) == 0 && len(auctionBody.ImagePath) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "no image as default, set a default image for this product first or choose an image",
		})
		ctx.Abort()
		return
	}

	auctionBody.ProductName = res.Name
	errCreateAuction := a.AuctionService.CreateAuction(auctionBody)
	if errCreateAuction != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAuction.Error(),
		})
		log.Println("CreateAuction: Error create new auction in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "auction created",
	})
}

func (a *AuctionController) UpdateAuctionByAuctionId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	auctionId, errGetId := strconv.Atoi(ctx.Query(auction.Id))
	if errGetId != nil {
		log.Println("error in get auction by auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	var updateBody *entity.Auction
	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}

	updateBody.ID = uint(auctionId)
	updateBody.UserId = claims.UserId
	err := a.AuctionService.Update(updateBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("Update auction: Error Update in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "auction updated",
	})
}

func (a *AuctionController) DeleteAuctionByAuctionId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	auctionId, errGetId := strconv.Atoi(ctx.Query(auction.Id))
	if errGetId != nil {
		log.Println("error when get auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	err := a.AuctionService.Delete(uint(auctionId), claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("Update auction: Error Update in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "auction deleted",
	})
}
func (a *AuctionController) GetAuctionByAuctionId(ctx *gin.Context) {
	auctionId, errGetId := strconv.Atoi(ctx.Query(auction.Id))
	if errGetId != nil {
		log.Println("error when get auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	auctionDetail, err := a.AuctionService.GetAuctionById(uint(auctionId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "auction not found",
		})
		log.Println("GetAuctionById: Error in package controller", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, auctionDetail)

}
func (a *AuctionController) GetAllAuctions(ctx *gin.Context) {
	page, errGetPage := strconv.Atoi(ctx.Query(auction.Page))
	if errGetPage != nil {
		log.Println("error when get auctionId: ", errGetPage)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	auctions, err := a.AuctionService.GetAllAuctions(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to get all auctions",
		})
		log.Println("Get auctions: Error get all auctions in package controller", err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, auctions)
}

func (a *AuctionController) GetAllAuctionsByProductName(ctx *gin.Context) { //Search auction by productName
	productName := ctx.Query(auction.ProductName)
	in := product.GetProductByProductNameRequest{ProductName: productName}
	res, errRes := a.ProductClient.GetProductByProductName(ctx, &in)
	if errRes != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errRes.Error(),
		})
		log.Println("CreateAuction: Error to call productService rpc server", errRes)
		ctx.Abort()
		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no product found",
		})
		log.Println("CreateAuction: product not found")
		ctx.Abort()
		return
	}

	log.Println("list ne: ", res.ProductName)
	auctions, err := a.AuctionService.GetAllAuctionsByProductName(res.ProductName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, auctions)

}
