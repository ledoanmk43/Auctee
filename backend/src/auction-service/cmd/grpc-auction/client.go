package grpc_auction

import (
	"chilindo/pkg/pb/account"
	"chilindo/pkg/pb/product"
	"google.golang.org/grpc"
	"log"
)

type IRPCClient interface {
	SetUpClient(port string) product.ProductServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpProductClient(productClientPort string) product.ProductServiceClient {
	//var opts []grpc.DialOption
	//creds, tlsErr := ssl.LoadTLSCredentials()
	//
	//if tlsErr != nil {
	//	log.Fatalf("Failed to load credentials: %v", tlsErr)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial(productClientPort, opts...
	conn, err := grpc.Dial(productClientPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", productClientPort)
	productClient := product.NewProductServiceClient(conn)
	return productClient
}
func (r RPCClient) SetUpAccountClient(accountClientPort string) account.AccountServiceClient {
	//var opts []grpc.DialOption
	//creds, tlsErr := ssl.LoadTLSCredentials()
	//
	//if tlsErr != nil {
	//	log.Fatalf("Failed to load credentials: %v", tlsErr)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(accountClientPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", accountClientPort)
	accountClient := account.NewAccountServiceClient(conn)
	return accountClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
