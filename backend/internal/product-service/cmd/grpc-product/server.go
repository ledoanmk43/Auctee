package grpc_product

import (
	"backend/internal/product-service/config"
	"backend/internal/product-service/repository"
	"backend/internal/product-service/service"
	"backend/pkg/pb/product"
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
	productService service.IProductService
	imageService   service.IProductImageService
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
	productRepo := repository.NewProductRepositoryDefault(db, nil, nil)
	ProductService := service.NewProductService(productRepo)

	imageRepo := repository.NewProductImageRepository(db)
	ProductImageService := service.NewProductImageService(imageRepo)
	product.RegisterProductServiceServer(s, &ProductServer{
		productService: ProductService,
		imageService:   ProductImageService,
	})

	log.Printf("  Product Server is on port  %s\n", grpcServerPort)
	return s.Serve(lis)
}

func (p *ProductServer) GetProductById(ctx context.Context, in *product.GetProductByIdRequest) (*product.GetProductByIdResponse, error) {
	productId := in.GetProductId()
	if productId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument productId= %v", productId)
	}

	prod, err := p.productService.GetProductDetailByProductId(productId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not found: %v", err)
	}
	var imagePath string
	image, _ := p.imageService.GetDefaultImageByProductId(productId)
	if image == nil {
		imagePath = ""
	} else {
		imagePath = image.Path
	}

	response := &product.GetProductByIdResponse{
		Id:          prod.Id,
		Name:        prod.Name,
		MinPrice:    float32(prod.MinPrice),
		Description: prod.Description,
		Quantity:    int32(prod.Quantity),
		ExpectPrice: float32(prod.ExpectPrice),
		UserId:      uint32(prod.UserId),
		Path:        imagePath, //default image
	}
	return response, nil
}

func (p *ProductServer) GetProductByProductName(ctx context.Context, in *product.GetProductByProductNameRequest) (*product.GetProductByProductNameResponse, error) {
	productName := in.GetProductName()
	if productName == "" {
		return nil, status.Errorf(codes.InvalidArgument, "InvalidArgument productId= %v", productName)
	}

	productList, err := p.productService.GetProductsByProductName(productName)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Not found: %v", err)
	}
	response := &product.GetProductByProductNameResponse{
		IdList:      productList.IdList,
		ProductName: productList.ProductName,
	}
	return response, nil
}
