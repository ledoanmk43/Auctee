package grpc_product

import (
	"chilindo/pkg/pb/account"

	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	accountClientPort = ":50051"
)

type IRPCClient interface {
	SetUpProductClient() account.AccountServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpAccountClient() account.AccountServiceClient {
	//var opts []grpc.DialOption
	//creds, tlsErr := ssl.LoadTLSCredentials()
	//
	//if tlsErr != nil {
	//	log.Fatalf("Failed to load credentials: %v", tlsErr)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial(adminClientPort, opts...)
	//if err != nil {
	//	log.Fatalf("failed to connect: %v", err)
	//}
	//log.Println("Listening from port :", adminClientPort)
	//adminClient := admin.NewAdminServiceClient(conn)
	//return adminClient
	addr := os.Getenv("ADMIN_SRV_HOST")

	conn, dialErr := grpc.Dial(addr, grpc.WithInsecure())
	if dialErr != nil {
		log.Fatalf("failed to connect: %v", dialErr)
	}

	accountClient := account.NewAccountServiceClient(conn)
	fmt.Println("Listen to Account-service on port", accountClientPort)
	return accountClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
