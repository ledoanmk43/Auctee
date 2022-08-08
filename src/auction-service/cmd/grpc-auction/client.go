package grpc_auction

import (
	"chilindo/pkg/pb/admin"
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
func (r RPCClient) SetUpAdminClient(adminClientPort string) admin.AdminServiceClient {
	//var opts []grpc.DialOption
	//creds, tlsErr := ssl.LoadTLSCredentials()
	//
	//if tlsErr != nil {
	//	log.Fatalf("Failed to load credentials: %v", tlsErr)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(adminClientPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", adminClientPort)
	adminClient := admin.NewAdminServiceClient(conn)
	return adminClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
