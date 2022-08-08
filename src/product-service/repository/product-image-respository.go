package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ProductImageRepository interface {
	CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error)
	GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error)
	ProductImageByID(b *dto.ProductDTO) (int64, error)
	GetImageByID(b *dto.ImageDTO) (*entity.ProductImages, error)
	DeleteImage(b *dto.ImageDTO) (*entity.ProductImages, error)
	UpdateImage(b *dto.UpdateImageDTO) (*entity.ProductImages, error)
}

func (p productImageRepository) UpdateImage(b *dto.UpdateImageDTO) (*entity.ProductImages, error) {
	//TODO implement me
	var count int64
	var updateImage *entity.ProductImages
	record := p.connection.Where("id = ?", b.Image.ID).Find(&updateImage).Count(&count)

	if record.Error != nil {
		log.Println("Error to find product repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("image not found")
	}
	updateImage = b.Image
	recordSave := p.connection.Updates(&updateImage)
	if recordSave.Error != nil {
		log.Println("Error to update produce repo", recordSave.Error)
		return nil, recordSave.Error
	}
	return updateImage, nil
}

func (p productImageRepository) GetImageByID(b *dto.ImageDTO) (*entity.ProductImages, error) {
	//TODO implement me
	var image *entity.ProductImages
	var count int64
	record := p.connection.Where("id = ?", b.ImageId).Find(&image).Count(&count)
	if record.Error != nil {
		log.Println("GetOptionById: Error to get option in repo pkg", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptionById: Not found option", count)
		return nil, nil
	}
	return image, nil
}

func (p productImageRepository) DeleteImage(b *dto.ImageDTO) (*entity.ProductImages, error) {
	//TODO implement me
	var images *entity.ProductImages
	record := p.connection.Where("id = ?", b.ImageId).Find(&images)
	if record.Error != nil {
		log.Println("DeleteOption: Error to find option", record.Error)
		return nil, record.Error
	}
	recordDelete := p.connection.Delete(&images)
	if recordDelete.Error != nil {
		log.Println("DeleteOption: Error to delete option", record.Error)
		return nil, recordDelete.Error
	}
	return images, nil
}

func (p productImageRepository) GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error) {
	//TODO implement me
	var images *[]entity.ProductImages
	var count int64
	record := p.connection.Where("product_id = ?", b.ProductId).Find(&images).Count(&count)
	if record.Error != nil {
		log.Println("GetOptions : Error to get all option", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptions : Not found Options", count)
		return nil, nil
	}
	return images, nil
}

func (p productImageRepository) ProductImageByID(b *dto.ProductDTO) (int64, error) {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ?", b.ProductId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
}

func (p productImageRepository) CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error) {
	//TODO implement me
	record := p.connection.Create(&b.Image)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return nil, record.Error
	}
	return b.Image, nil
}

type productImageRepository struct {
	connection *gorm.DB
}

func NewProductImageRepository(connection *gorm.DB) *productImageRepository {
	return &productImageRepository{connection: connection}
}
