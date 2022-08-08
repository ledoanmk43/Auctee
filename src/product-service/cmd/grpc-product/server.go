package grpc_product

import (
	"chilindo/pkg/pb/product"
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/service"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	grpcServerPort = ":50052"
	certFile       = "pkg/ssl/server.crt"
	keyFile        = "pkg/ssl/server.pem"
)

type ProductServer struct {
	product.ProductServiceServer
	productService service.ProductService
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
	repoProduct := repository.NewProductRepository(db)
	serviceProduct := service.NewProductService(repoProduct)

	product.RegisterProductServiceServer(s, &ProductServer{
		productService: serviceProduct,
	})

	log.Printf("  Product Server is on port  %s\n", grpcServerPort)
	return s.Serve(lis)
}

func (p *ProductServer) GetProduct(ctx context.Context, in *product.GetProductRequest) (*product.GetProductResponse, error) {
	log.Printf("Login request: %v\n", in)

	pid := in.GetProductId()
	if pid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "InvalidArgument productId= %v", pid)
	}
	var dto dto.ProductDTO
	dto.ProductId = in.GetProductId()

	prod, err := p.productService.FindProductByID(&dto)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not found: %v", err)
	}

	response := &product.GetProductResponse{
		Id:          prod.Id,
		Name:        prod.Name,
		MinPrice:    prod.MinPrice,
		Description: prod.Description,
		Quantity:    int32(prod.Quantity),
	}
	return response, nil
}
