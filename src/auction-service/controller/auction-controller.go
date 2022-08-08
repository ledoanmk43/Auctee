package controller

import (
	product2 "chilindo/pkg/pb/product"
	"chilindo/src/auction-service/entity"
	"chilindo/src/auction-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAuctionController interface {
	CreateAuction(c *gin.Context)
}

type AuctionController struct {
	AuctionService service.IAuctionService
	ProductClient  product2.ProductServiceClient
}

func NewAuctionController(auctionService service.IAuctionService, productClient product2.ProductServiceClient) *AuctionController {
	return &AuctionController{AuctionService: auctionService, ProductClient: productClient}
}

func (a AuctionController) CreateAuction(c *gin.Context) {
	var auctionBody *entity.Auction
	if err := c.ShouldBindJSON(&auctionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error when binding JSON",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		c.Abort()
		return
	}
	in := product2.GetProductRequest{ProductId: auctionBody.ProductId}
	res, errRes := a.ProductClient.GetProduct(c, &in)

	if errRes != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create auction",
		})
		log.Println("CreateAuction: Error to call productService rpc server", errRes)
		c.Abort()
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found product",
		})
		log.Println("CreateAuction: product not found")
		c.Abort()
		return
	}

	createdAuction, errCreateAuction := a.AuctionService.CreateAuction(auctionBody)
	if errCreateAuction != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAuction.Error(),
		})
		log.Println("CreateAuction: Error create new auction in package controller")
		return
	}
	c.JSON(http.StatusOK, createdAuction)
}
