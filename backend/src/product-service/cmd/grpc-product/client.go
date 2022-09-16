package grpc_product

import (
	"backend/pkg/pb/account"

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
	addr := os.Getenv("ACCOUNT_SRV_HOST")

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
