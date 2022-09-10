package account_grpc_controller

import (
	"chilindo/pkg/pb/account"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAccountServiceController interface {
	MiddlewareCheckIsAuth(accountClient account.AccountServiceClient) gin.HandlerFunc
}

type AccountServiceController struct {
}

func NewAccountServiceController() *AccountServiceController {
	return &AccountServiceController{}
}

func (a AccountServiceController) MiddlewareCheckIsAuth(accountClient account.AccountServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			c.Abort()
			return
		}
		res, err := accountClient.CheckIsAuth(c, &account.CheckIsAuthRequest{
			Token: token,
		})
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
			c.Abort()
			return
		}
		if !(res.IsAuth) {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
