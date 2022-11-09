package grpc_account

import (
	"backend/internal/account-service/config"
	"backend/internal/account-service/repository"
	"backend/internal/account-service/service"
	"backend/pkg/pb/account"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	addrAccountServerGRPC = ":50051"
	certFile              = "pkg/ssl/server.crt"
	keyFile               = "pkg/ssl/server.pem"
)

type AccountServer struct {
	account.AccountServiceServer
	accountService service.IAccountService
	addressService service.IAddressService
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

	addressRepo := repository.NewAddressRepositoryDefault(config.DB)
	AddressService := service.NewAddressServiceDefault(addressRepo)
	account.RegisterAccountServiceServer(s, &AccountServer{
		accountService: AccountService,
		addressService: AddressService,
	})

	log.Printf("Account Server is on port %s\n", addrAccountServerGRPC)
	return s.Serve(lis)
}

func (a *AccountServer) CheckIsAuth(ctx context.Context, in *account.CheckIsAuthRequest) (*account.CheckIsAuthResponse, error) {
	res, err := a.accountService.CheckIsAuth(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	if res == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return res, nil
}

func (a *AccountServer) GetAddressByUserId(ctx context.Context, in *account.GetAddressByUserIdRequest) (*account.GetAddressByUserIdResponse, error) {
	userId := in.GetUserId()
	addressId := in.GetAddressId()
	if userId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument productId= %v", userId)
	}

	res, err := a.addressService.GetAddressByAddressId(uint(addressId), uint(userId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not found: %v", err)
	}

	response := &account.GetAddressByUserIdResponse{
		Firstname:   res.Firstname,
		Lastname:    res.Lastname,
		Phone:       res.Phone,
		Email:       res.Email,
		Province:    res.Province,
		District:    res.District,
		SubDistrict: res.SubDistrict,
		Address:     res.Address,
		TypeAddress: res.TypeAddress,
	}
	return response, nil
}