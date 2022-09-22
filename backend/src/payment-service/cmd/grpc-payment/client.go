package grpc_payment

import (
	"backend/pkg/pb/account"
	"backend/pkg/pb/auction"
	"google.golang.org/grpc"
	"log"
)

type IRPCClient interface {
	SetUpClient(port string) auction.AuctionServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpAuctionClient(auctionClientPort string) auction.AuctionServiceClient {
	//var opts []grpc.DialOption
	//creds, tlsErr := ssl.LoadTLSCredentials()
	//
	//if tlsErr != nil {
	//	log.Fatalf("Failed to load credentials: %v", tlsErr)
	//}
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial(productClientPort, opts...
	conn, err := grpc.Dial(auctionClientPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", auctionClientPort)
	auctionClient := auction.NewAuctionServiceClient(conn)
	return auctionClient
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
