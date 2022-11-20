package websocket

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(c *gin.Context) (*websocket.Conn, error) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}
