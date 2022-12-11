package route

import (
	"backend/pkg/pb/account"
	"backend/src/payment-service/controller"
	account_server_controller "backend/src/product-service/controller/account-grpc-controller"
	"github.com/gin-gonic/gin"
)

type IPaymentRoute interface {
	GetRouter()
}

type PaymentRoute struct {
	PaymentController    controller.IPaymentController
	Router               *gin.Engine
	AccountSrvController account_server_controller.IAccountServiceController
	AccountClient        account.AccountServiceClient
}

func NewPaymentRoute(paymentController controller.IPaymentController, router *gin.Engine, accountSrvController account_server_controller.IAccountServiceController, accountClient account.AccountServiceClient) *PaymentRoute {
	return &PaymentRoute{PaymentController: paymentController, Router: router, AccountSrvController: accountSrvController, AccountClient: accountClient}
}

func (p *PaymentRoute) GetRouter() {
	paymentRoute := p.Router.Group("/auctee")
	{
		paymentRoute.POST("/user/checkout/auction", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.CreatePayment)
		paymentRoute.GET("/user/checkout/payment-history", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.GetAllPaymentsForWinner)
		paymentRoute.GET("/user/checkout/all-bills", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.GetAllPaymentsForOwner)
		paymentRoute.GET("/user/checkout/payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.GetPaymentByPaymentId)
		paymentRoute.PUT("/user/checkout/shipping-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.UpdateAddressPayment)
		paymentRoute.PUT("/user/checkout/cancel-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.CancelPayment)
		paymentRoute.POST("/user/checkout/momo-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.CheckoutMoMo)
		//paymentRoute.POST("/user/ipn/momo-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.MoMoIPNResult)
		paymentRoute.PUT("/user/update/momo-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.UpdateMoMoCheckOut)
		paymentRoute.PUT("/user/checkout/cod-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.CheckoutCOD)
		paymentRoute.PUT("/user/checkout/shipping-confirm", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.SetShippingStatusDelivering)
		paymentRoute.PUT("/user/checkout/shipping-status-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.SetShippingStatusCompleted)
		paymentRoute.PUT("/user/checkout/checkout-status-done", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.SetCheckOutStatusDone)
	}
}
