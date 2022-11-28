package controller

import (
	account "backend/internal/account-service/config"
	auction "backend/internal/auction-service/config"
	"backend/internal/auction-service/entity"
	"backend/internal/auction-service/service"
	"backend/pkg/pb/product"
	"backend/pkg/token"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IBidController interface {
	CreateBid(ctx *gin.Context)
	AutoBid(ctx *gin.Context)
	GetAllBidsByAuctionId(ctx *gin.Context)
	GetAllBidsByUserId(ctx *gin.Context)
}

type BidController struct {
	BidService     service.IBidService
	AuctionService service.IAuctionService
	ProductClient  product.ProductServiceClient
}

func NewBidController(bidService service.IBidService, auctionService service.IAuctionService, productClient product.ProductServiceClient) *BidController {
	return &BidController{BidService: bidService, AuctionService: auctionService, ProductClient: productClient}
}

func (b *BidController) GetAllBidsByAuctionId(ctx *gin.Context) {
	auctionId, errGetId := strconv.Atoi(ctx.Query(auction.Id))
	if errGetId != nil {
		log.Println("error when get auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}
	bids, err := b.BidService.GetAllBidsByAuctionId(uint(auctionId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to get all bids",
		})
		log.Println("Get bids: Error get all bids in package controller", err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, bids)
}

func (b *BidController) GetAllBidsByUserId(ctx *gin.Context) {
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

	auctions, err := b.BidService.GetAllBidsByUserId(claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to get all bids",
		})
		log.Println("Get bids: Error get all bids in package controller", err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, auctions)
}

func (b *BidController) AutoBid(ctx *gin.Context) {
	//var bidBody entity.Bid
	//
	//tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
	//if errGetToken != nil {
	//	log.Println("Error when get token in controller: ", errGetToken)
	//	ctx.Abort()
	//	return
	//}
	//
	//claims, errExtract := token.ExtractToken(tokenFromCookie)
	//if errExtract != nil || claims == nil {
	//	log.Println("Error: Error when extracting token in controller: ", errExtract)
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "Unauthorized",
	//	})
	//	ctx.Abort()
	//	return
	//}
	//
	//auctionId, errGetId := strconv.Atoi(ctx.Query(auction.AuctionId))
	//if errGetId != nil {
	//	log.Println("error in get auctionId: ", errGetId)
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Error when get id of auction in url",
	//	})
	//	ctx.Abort()
	//	return
	//}
	//
	//productId := ctx.Query(auction.ProductId)
	//if len(productId) == 0 {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "Error when get id of product in url",
	//	})
	//	ctx.Abort()
	//	return
	//}
	//
	//in := product.GetProductByIdRequest{ProductId: productId}
	//res, errRes := b.ProductClient.GetProductById(ctx, &in)
	//if errRes != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": errRes.Error(),
	//	})
	//	log.Println("Create Bid: Error to call productService rpc server", errRes)
	//	ctx.Abort()
	//	return
	//}
	//
	//if res == nil {
	//	ctx.JSON(http.StatusNotFound, gin.H{
	//		"message": "no product found",
	//	})
	//	log.Println("Create Bid: product not found")
	//	ctx.Abort()
	//	return
	//}
	//
	//auctionData, errGetAuction := b.AuctionService.GetAuctionById(uint(auctionId))
	//if errGetAuction != nil || auctionData.ProductId != productId {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "auction does not exist",
	//	})
	//	log.Println("GetAuctionById: Error in package controller", errGetAuction)
	//	ctx.Abort()
	//	return
	//}
	////check is active
	//if *auctionData.IsActive == false {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		//check bid_value
	//		"message": "auction has not run yet",
	//	})
	//	ctx.Abort()
	//	return
	//}
	//
	//if auctionData.CurrentBid < float64(res.ExpectPrice) {
	//	bidBody.BidValue = auctionData.CurrentBid + auctionData.PricePerStep
	//} else {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "price meet expected number",
	//	})
	//	ctx.Abort()
	//	return
	//}
	//
	//bidBody.AuctionId = auctionData.ID
	//errCreateBid := b.BidService.CreateBid(&bidBody)
	//if errCreateBid != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"error": errCreateBid.Error(),
	//	})
	//	log.Println("CreateBid: Error auto create new Bid in package controller")
	//	ctx.Abort()
	//	return
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "bid is automatically created",
	})
}

func (b *BidController) CreateBid(ctx *gin.Context) {
	var bidBody *entity.Bid
	if err := ctx.ShouldBindJSON(&bidBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "err.Error()",
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

	auctionId, errGetId := strconv.Atoi(ctx.Query(auction.AuctionId))
	if errGetId != nil {
		log.Println("error in get auctionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id of auction in url",
		})
		ctx.Abort()
		return
	}

	productId := ctx.Query(auction.ProductId)
	if len(productId) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error when get id of product in url",
		})
		ctx.Abort()
		return
	}

	auctionData, errGetAuction := b.AuctionService.GetAuctionById(uint(auctionId))
	if errGetAuction != nil || auctionData.ProductId != productId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Phiên đấu giá không tồn tại",
		})
		log.Println("GetAuctionById: Error in package controller", errGetAuction)
		ctx.Abort()
		return
	}

	//check is active
	if *auctionData.IsActive == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Phiên đấu giá chưa diễn ra",
		})
		ctx.Abort()
		return
	}

	in := product.GetProductByIdRequest{ProductId: productId}
	res, errRes := b.ProductClient.GetProductById(ctx, &in)
	if errRes != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errRes.Error(),
		})
		log.Println("Create Bid: Error to call productService rpc server", errRes)
		ctx.Abort()
		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no product found",
		})
		log.Println("Create Bid: product not found")
		ctx.Abort()
		return
	}
	//
	//if bidBody.BidValue <= float64(res.MinPrice) {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		//check bid_value
	//		"message": "Số tiền không hợp lệ",
	//	})
	//	ctx.Abort()
	//	return
	//}

	bidBody.UserId = claims.UserId
	bidBody.AuctionId = uint(auctionId)
	errCreateBid := b.BidService.CreateBid(bidBody, res.ExpectPrice)
	if errCreateBid != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreateBid.Error(),
		})
		log.Println("CreateBid: Error create new Bid in package controller")
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "bid created",
	})
}
