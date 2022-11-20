package route

import (
	"backend/internal/auction-service/config"
	"backend/internal/auction-service/controller"
	account_server_controller "backend/internal/product-service/controller/account-grpc-controller"
	"backend/pkg/websocket"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IBidRoute interface {
	GetRouter()
}

type BidRoute struct {
	BidController        controller.IBidController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
}

func NewBidRoute(bidController controller.IBidController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController) *BidRoute {
	return &BidRoute{BidController: bidController, Router: router, AccountSrvController: accountSrvController}
}

func serveWs(c *gin.Context) {
	conn, err := websocket.Upgrade(c)
	if err != nil {
		fmt.Fprintf(c.Writer, "%+V\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: config.Pool,
	}
	config.Pool.Register <- client
	client.Read()
}

func (b *BidRoute) GetRouter() {
	bidRoute := b.Router.Group("/auctee")
	{
		bidRoute.GET("/ws", serveWs)
		bidRoute.POST("/auction", b.BidController.CreateBid)
		bidRoute.POST("/bot/auction", b.BidController.AutoBid)
		bidRoute.GET("/all-bids/auction", b.BidController.GetAllBidsByAuctionId)
	}
}
