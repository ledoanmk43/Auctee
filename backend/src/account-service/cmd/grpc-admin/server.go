package grpc_admin

import (
	admin2 "chilindo/pkg/pb/admin"
	"chilindo/src/account-service/config"
	"chilindo/src/account-service/repository"
	"chilindo/src/account-service/service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	addrAdminServer = ":50051"
	certFile        = "pkg/ssl/server.crt"
	keyFile         = "pkg/ssl/server.pem"
)

type AccountServer struct {
	admin2.AdminServiceServer
	AccountService service.IAccountService
}

func RunGRPCServer(enabledTLS bool, lis net.Listener) error {
	var opts []grpc.ServerOption
	if enabledTLS {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	accountRepo := repository.NewAccountRepositoryDefault(config.DB)
	AccountService := service.NewAccountServiceDefault(accountRepo)

	admin2.RegisterAdminServiceServer(s, &AccountServer{
		AccountService: AccountService,
	})

	log.Printf("Admin Server is on port %s\n", addrAdminServer)
	return s.Serve(lis)
}

func (a *AccountServer) CheckIsAuth(ctx context.Context, in *admin2.CheckIsAuthRequest) (*admin2.CheckIsAuthResponse, error) {
	res, err := a.AccountService.CheckIsAuth(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	if res == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return res, nil
}
