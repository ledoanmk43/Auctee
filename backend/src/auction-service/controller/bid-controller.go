package controller

import (
	"chilindo/pkg/utils"
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	auctionId = "auctionId"
)

type IBidController interface {
	CreateBid(c *gin.Context)
}

type BidController struct {
	BidService service.IBidService
	//ProductClient product.ProductServiceClient
	//UserClient     user.UserServiceClient
}

func NewBidController(bidService service.IBidService) *BidController {
	return &BidController{BidService: bidService}
}

func (b *BidController) CreateBid(ctx *gin.Context) {
	var bidBody *entity.Bid
	if err := ctx.ShouldBindJSON(&bidBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error when binding JSON",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		ctx.Abort()
		return
	}

	UserIdFromToken := utils.GetIdFromToken(ctx)
	auctionID := ctx.Param(auctionId)
	rawId, _ := strconv.ParseUint(auctionID, 10, 64)
	bidBody.AuctionId = uint(rawId)
	bidBody.UserId = UserIdFromToken

	createdBid, errCreateBid := b.BidService.CreateBid(bidBody)
	if errCreateBid != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateBid.Error(),
		})
		log.Println("CreateBid: Error create new bid in package controller")
		return
	}
	ctx.JSON(http.StatusOK, createdBid)
}
