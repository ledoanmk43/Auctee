package account_grpc_controller

import (
	"backend/pkg/pb/account"
	"backend/src/account-service/config"
	"github.com/gin-contrib/sessions"
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

func (a AccountServiceController) MiddlewareCheckIsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authSession := sessions.Default(ctx)
		tokenFromCookie := authSession.Get(config.CookieAuth)
		if tokenFromCookie == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "no cookie",
			})
			ctx.Abort()
			return
		}
		res, err := a.AccountClient.CheckIsAuth(ctx, &account.CheckIsAuthRequest{
			Token: tokenFromCookie.(string),
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
