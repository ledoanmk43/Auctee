package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ProductRepository interface {
	InsertProduct(b *entity.Product) (*entity.Product, error)
	UpdateProduct(b *dto.ProductUpdateDTO) (*entity.Product, error)
	DeleteProduct(proId string, adminId uint) (*entity.Product, error)
	AllProduct() (*[]entity.Product, error)
	FindProductByID(b *dto.ProductDTO) (*entity.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func (db *productConnection) DeleteProduct(proId string, adminId uint) (*entity.Product, error) {
	var product *entity.Product
	var count int64
	resultProduct := db.connection.Where("id = ? AND admin_id = ?", proId, adminId).Find(&product).Count(&count)
	if resultProduct.Error != nil {
		log.Println("DeleteProduct: Error in find product to delete in package repository", resultProduct.Error)
		return nil, resultProduct.Error
	}
	if count == 0 {
		return nil, errors.New("product not found")
	}
	db.connection.Delete(&product)

	var productOptions *entity.ProductOption
	resultOption := db.connection.Where("product_id = ?", proId).Limit(99).Delete(&productOptions)
	if resultOption.Error != nil {
		log.Println("DeleteProduct: Error in find option to delete in package repository", resultOption.Error)
		return nil, resultOption.Error
	}
	var productImages *entity.ProductImages
	resultImage := db.connection.Where("product_id = ?", proId).Limit(99).Delete(&productImages)
	if resultOption.Error != nil {
		log.Println("DeleteProduct: Error in find option to delete in package repository", resultImage.Error)
		return nil, resultImage.Error
	}
	return product, nil
}

func (db *productConnection) InsertProduct(pro *entity.Product) (*entity.Product, error) {
	if errCheckEmptyField := pro.Validate("insert"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}

	record := db.connection.Create(&pro)
	if record.Error != nil {
		log.Println("Error to create product in repo")
		return nil, record.Error
	}
	return pro, nil
}

func (db *productConnection) UpdateProduct(b *dto.ProductUpdateDTO) (*entity.Product, error) {
	var updateProduct *entity.Product
	var count int64
	record := db.connection.Where("id = ?", b.ProductId).Find(&updateProduct).Count(&count)

	if record.Error != nil {
		log.Println("Error to find product in repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("product not found")
	}
	//b.Product.Id = b.ProductId
	updateProduct = b.Product
	recordSave := db.connection.Updates(&updateProduct)
	if recordSave.Error != nil {
		log.Println("Error to update produce repo", recordSave.Error)
		return nil, recordSave.Error
	}
	return updateProduct, nil
}

func (db *productConnection) AllProduct() (*[]entity.Product, error) {
	var products *[]entity.Product
	record := db.connection.Find(&products)
	if record.Error != nil {
		log.Println("GetProducts: Error get all in repo", record.Error)
		return nil, record.Error
	}
	return products, nil
}

func (db *productConnection) FindProductByID(pid *dto.ProductDTO) (*entity.Product, error) {
	var product *entity.Product
	var count int64
	record := db.connection.Where("id = ?", pid.ProductId).Find(&product).Count(&count)
	if record.Error != nil {
		log.Println("Get product by ID", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetProductById: product not found", count)
		return nil, errors.New("error: product not found")
	}
	return product, nil
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}
