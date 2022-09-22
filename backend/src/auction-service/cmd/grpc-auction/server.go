package grpc_auction

import (
	"backend/pkg/pb/auction"
	"backend/src/auction-service/config"
	"backend/src/auction-service/repository"
	"backend/src/auction-service/service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	grpcServerPort = ":50053"
	certFile       = "pkg/ssl/server.crt"
	keyFile        = "pkg/ssl/server.pem"
)

type AuctionServer struct {
	auction.AuctionServiceServer
	auctionService service.IAuctionService
	//imageService   service.IProductImageService
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
	db := config.GetDB()
	auctionRepo := repository.NewAuctionRepositoryDefault(db)
	AuctionService := service.NewAuctionServiceDefault(auctionRepo)

	auction.RegisterAuctionServiceServer(s, &AuctionServer{
		auctionService: AuctionService,
	})

	log.Printf("  Auction Server is on port  %s\n", grpcServerPort)
	return s.Serve(lis)
}

func (a *AuctionServer) GetAuctionById(ctx context.Context, in *auction.GetAuctionByIdRequest) (*auction.GetAuctionByIdResponse, error) {
	auctionId := in.GetAuctionId()
	if auctionId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument productId= %v", auctionId)
	}

	res, err := a.auctionService.GetAuctionById(uint(auctionId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not found: %v", err)
	}

	response := &auction.GetAuctionByIdResponse{
		ProductId:   res.ProductId,
		ProductName: res.ProductName,
		EndTime:     res.EndTime,
		Quantity:    int32(res.Quantity),
		UserId:      uint32(res.UserId),
		WinnerId:    uint32(res.WinnerId),
		CurrentBid:  float32(res.CurrentBid),
	}
	return response, nil
}
