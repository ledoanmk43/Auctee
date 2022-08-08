package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	InsertProduct(b *dto.ProductCreatedDTO) (*entity.Product, error)
	UpdateProduct(b *dto.ProductUpdateDTO) (*entity.Product, error)
	DeleteProduct(b *dto.ProductDTO) (*entity.Product, error)
	AllProduct() (*[]entity.Product, error)
	FindProductByID(b *dto.ProductDTO) (*entity.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func (t productConnection) DeleteProduct(b *dto.ProductDTO) (*entity.Product, error) {
	var product *entity.Product
	recordFind := t.connection.Where("id = ?", b.ProductId).Delete(&product)
	if recordFind.Error != nil {
		log.Println("DeleteProduct: Error in find product to delete in package repository", recordFind.Error)
		return nil, recordFind.Error
	}
	return product, nil
}

func (t productConnection) InsertProduct(b *dto.ProductCreatedDTO) (*entity.Product, error) {
	record := t.connection.Create(&b.Product)
	if record.Error != nil {
		log.Println("Error to create product repo")
		return nil, record.Error
	}
	return b.Product, nil
}

func (t productConnection) UpdateProduct(b *dto.ProductUpdateDTO) (*entity.Product, error) {
	var updateProduct *entity.Product
	var count int64
	record := t.connection.Where("id = ?", b.ProductId).Find(&updateProduct).Count(&count)

	if record.Error != nil {
		log.Println("Error to find product repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("product not found")
	}
	//b.Product.Id = b.ProductId
	updateProduct = b.Product
	recordSave := t.connection.Updates(&updateProduct)
	if recordSave.Error != nil {
		log.Println("Error to update produce repo", recordSave.Error)
		return nil, recordSave.Error
	}
	return updateProduct, nil
}

func (t productConnection) AllProduct() (*[]entity.Product, error) {
	var products *[]entity.Product
	record := t.connection.Find(&products)
	if record.Error != nil {
		log.Println("GetProducts: Error get all in package", record.Error)
		return nil, record.Error
	}
	return products, nil
}

func (t productConnection) FindProductByID(pid *dto.ProductDTO) (*entity.Product, error) {
	var product *entity.Product
	var count int64
	record := t.connection.Where("id = ?", pid.ProductId).Find(&product).Count(&count)
	if record.Error != nil {
		log.Println("Get product by ID", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetProductById: Product not found", count)
		return nil, errors.New("error: Product not found")
	}
	return product, nil
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}
