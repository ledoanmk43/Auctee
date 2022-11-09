package service

import (
	"backend/internal/product-service/entity"
	"backend/internal/product-service/repository"
	"log"
)

type IProductService interface {
	Insert(b *entity.Product) error
	Update(b *entity.Product) error
	Delete(proId string, userId uint) error
	GetAllProducts(userId uint) (*[]entity.Product, error)
	GetProductByProductId(productId string) (*entity.Product, error)
	GetProductsByProductName(productName string) (*entity.ProductResponse, error)
	GetProductDetailByProductId(productId string) (*entity.Product, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (p *ProductService) Insert(b *entity.Product) error {
	err := p.ProductRepository.InsertProduct(b)
	if err != nil {
		log.Println("Error in create product in service", err)
		return err
	}
	return nil
}

func (p *ProductService) Update(b *entity.Product) error {
	err := p.ProductRepository.UpdateProduct(b)
	if err != nil {
		log.Println("Error in package service", err)
		return err
	}
	return nil
}

func (p *ProductService) Delete(proId string, userId uint) error {
	err := p.ProductRepository.DeleteProduct(proId, userId)
	if err != nil {
		log.Println("Error in package service", err)
		return err
	}
	return nil
}

func (p *ProductService) GetAllProducts(userId uint) (*[]entity.Product, error) {
	products, err := p.ProductRepository.GetAllProducts(userId)
	if err != nil {
		log.Println("GetProducts : Error get products in package service", err)
	}
	return products, nil
}

func (p *ProductService) GetProductDetailByProductId(productId string) (*entity.Product, error) {
	Product, err := p.ProductRepository.GetProductDetailByProductId(productId)
	if err != nil {
		log.Println("GetProductById: Error in get product by Id", err)
		return nil, err
	}
	return Product, nil
}

func (p *ProductService) GetProductByProductId(productId string) (*entity.Product, error) {
	product, err := p.ProductRepository.GetProductByProductId(productId)
	if err != nil {
		log.Println("GetProductById: Error in get product by Id", err)
		return nil, err
	}
	return product, nil
}

func (p *ProductService) GetProductsByProductName(productName string) (*entity.ProductResponse, error) {
	idList, err := p.ProductRepository.GetProductsByProductName(productName)
	if err != nil {
		log.Println("GetProductById: Error in get product by Id", err)
		return nil, err
	}
	return idList, nil
}
