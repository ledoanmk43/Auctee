package route

import (
	"backend/internal/payment-service/controller"
	account_server_controller "backend/internal/product-service/controller/account-grpc-controller"
	"backend/pkg/pb/account"
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
		paymentRoute.DELETE("/user/checkout/payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.DeletePayment)
		//paymentRoute.PUT("/user/checkout/shipping-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.CheckoutMoMo) //set checkout_status to true
		//paymentRoute.PUT("/user/checkout/shipping-status-payment", p.AccountSrvController.MiddlewareCheckIsAuth(), p.PaymentController.SetShippingStatusCompleted) //set checkout_status to true
	}
}
