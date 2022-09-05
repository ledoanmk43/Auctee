package grpc_product

import (
	"chilindo/pkg/pb/admin"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	adminClientPort = ":50051"
)

type IRPCClient interface {
	SetUpProductClient() admin.AdminServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpAdminClient() admin.AdminServiceClient {
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

	adminClient := admin.NewAdminServiceClient(conn)
	fmt.Println("Listen to AdminService on port", adminClientPort)
	return adminClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
