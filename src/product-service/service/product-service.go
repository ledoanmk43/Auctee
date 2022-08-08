package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/repository"
	"log"
)

type ProductService interface {
	Insert(b *dto.ProductCreatedDTO) (*entity.Product, error)
	Update(b *dto.ProductUpdateDTO) (*entity.Product, error)
	Delete(b *dto.ProductDTO) (*entity.Product, error)
	All() (*[]entity.Product, error)
	FindProductByID(b *dto.ProductDTO) (*entity.Product, error)
}

type productService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{ProductRepository: productRepository}
}

func (t productService) Insert(b *dto.ProductCreatedDTO) (*entity.Product, error) {
	createProduct, err := t.ProductRepository.InsertProduct(b)
	if err != nil {
		log.Println("Error in create product in service", err)
		return nil, err
	}
	return createProduct, nil
}

func (t productService) Update(b *dto.ProductUpdateDTO) (*entity.Product, error) {
	updateProduct, err := t.ProductRepository.UpdateProduct(b)
	if err != nil {
		log.Println("Error in package service", err)
		return nil, err
	}
	return updateProduct, nil
}

func (t productService) Delete(b *dto.ProductDTO) (*entity.Product, error) {
	product, err := t.ProductRepository.DeleteProduct(b)
	if err != nil {
		log.Println("Error in package service", err)
		return nil, err
	}
	return product, nil
}

//
func (t productService) All() (*[]entity.Product, error) {
	products, err := t.ProductRepository.AllProduct()
	if err != nil {
		log.Println("GetProducts : Error get products in package service", err)
	}
	return products, nil
}

func (t productService) FindProductByID(b *dto.ProductDTO) (*entity.Product, error) {
	product, err := t.ProductRepository.FindProductByID(b)
	if err != nil {
		log.Println("GetProductById: Error in get product by Id", err)
		return nil, err
	}
	return product, nil
}
