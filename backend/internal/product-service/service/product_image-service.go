package service

import (
	"backend/internal/product-service/dto"
	"backend/internal/product-service/entity"
	"backend/internal/product-service/repository"
	"errors"
	"log"
)

type IProductImageService interface {
	CreateImage(image *entity.ProductImage, userId uint) error
	GetDefaultImageByProductId(productId string) (*entity.ProductImage, error)
	GetImageByID(b *dto.ImageDTO) (*entity.ProductImage, error)
	DeleteImage(image *entity.ProductImage, userId uint) error
	UpdateImage(image *entity.ProductImage, userId uint) error
}

type ProductImageService struct {
	ProductImageRepository repository.IProductImageRepository
}

func NewProductImageService(productImageRepository repository.IProductImageRepository) *ProductImageService {
	return &ProductImageService{ProductImageRepository: productImageRepository}
}

func (p *ProductImageService) GetDefaultImageByProductId(productId string) (*entity.ProductImage, error) {
	image, err := p.ProductImageRepository.GetDefaultImageByProductId(productId)
	if err != nil {
		log.Println("GetOptionById: Error get image", err)
		return nil, err
	}
	return image, nil
}

func (p *ProductImageService) GetImageByID(b *dto.ImageDTO) (*entity.ProductImage, error) {
	image, err := p.ProductImageRepository.GetImageByID(b)
	if err != nil {
		log.Println("GetOptionById: Error get option", err)
		return nil, err
	}
	return image, nil
}

func (p *ProductImageService) GetAllImagesOfAProduct(productId string) (*[]entity.ProductImage, error) {
	images, err := p.ProductImageRepository.GetAllImages(productId)
	if err != nil {
		log.Println("GetOptions: Error get images", err)
		return nil, err
	}
	return images, nil
}

func (p *ProductImageService) DeleteImage(image *entity.ProductImage, userId uint) error {
	err := p.ProductImageRepository.DeleteImage(image, userId)
	if err != nil {
		log.Println("Delete image: Error delete option", err)
		return err
	}
	return nil
}

func (p *ProductImageService) UpdateImage(image *entity.ProductImage, userId uint) error {
	err := p.ProductImageRepository.UpdateImage(image, userId)
	if err != nil {
		log.Println("Update image: error in service ", err)
		return err
	}
	return nil

}

func (p *ProductImageService) CreateImage(image *entity.ProductImage, userId uint) error {
	if len(image.Path) == 0 {
		return errors.New("required path")
	}
	productId := image.ProductId
	countProd, prodErr := p.ProductImageRepository.ProductImageByID(productId)
	if prodErr != nil {
		log.Println("Create image: Error not found product to create option", prodErr)
		return prodErr
	}
	if countProd == 0 {
		log.Println("Create image: Error not found product to create option", prodErr)
		return errors.New("product not found")
	}
	err := p.ProductImageRepository.CreateImage(image, userId)
	if err != nil {
		log.Println("CreateOption: Error to create option", err)
		return err
	}
	return nil
}
