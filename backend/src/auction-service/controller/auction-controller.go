package controller

import (
	product "backend/pkg/pb/product"
	"backend/src/auction-service/entity"
	"backend/src/auction-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAuctionController interface {
	CreateAuction(c *gin.Context)
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

	in := product.GetProductRequest{ProductId: auctionBody.ProductId}
	res, errRes := a.ProductClient.GetProduct(ctx, &in)

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

	auctionBody.CurrentBid = float64(res.MinPrice)
	createdAuction, errCreateAuction := a.AuctionService.CreateAuction(auctionBody)
	if errCreateAuction != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAuction.Error(),
		})
		log.Println("CreateAuction: Error create new auction in package controller")
		return
	}
	ctx.JSON(http.StatusOK, createdAuction)
}
