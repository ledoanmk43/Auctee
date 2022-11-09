package account_grpc_controller

import (
	"backend/internal/account-service/config"
	"backend/pkg/pb/account"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAccountServiceController interface {
	MiddlewareCheckIsAuth() gin.HandlerFunc
}

type AccountServiceController struct {
	AccountClient account.AccountServiceClient
}

func NewAccountServiceController(accountClient account.AccountServiceClient) *AccountServiceController {
	return &AccountServiceController{AccountClient: accountClient}
}

func (a *AccountServiceController) MiddlewareCheckIsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
		if errGetToken != nil {
			log.Println("Error when get token in account-rpc-controller: ", errGetToken)
			//ctx.JSON(http.StatusUnauthorized, gin.H{
			//	"message": "Unauthorized",
			//})
			ctx.Abort()
			return
		}
		res, err := a.AccountClient.CheckIsAuth(ctx, &account.CheckIsAuthRequest{
			Token: tokenFromCookie,
		})
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
			ctx.Abort()
			return
		}
		if !(res.IsAuth) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
