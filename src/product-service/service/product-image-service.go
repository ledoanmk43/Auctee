package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/repository"
	"errors"
	"log"
)

type ProductImageService interface {
	CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error)
	GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error)
	GetImageByID(b *dto.ImageDTO) (*entity.ProductImages, error)
	DeleteImage(b *dto.ImageDTO) (*entity.ProductImages, error)
	UpdateImage(b *dto.UpdateImageDTO) (*entity.ProductImages, error)
}

func (p productImageService) UpdateImage(b *dto.UpdateImageDTO) (*entity.ProductImages, error) {
	image, err := p.ProductImageRepository.UpdateImage(b)
	if err != nil {
		log.Println("UpdateOption: Error call repo")
		return nil, err
	}
	return image, nil

}

func (p productImageService) GetImageByID(b *dto.ImageDTO) (*entity.ProductImages, error) {
	image, err := p.ProductImageRepository.GetImageByID(b)
	if err != nil {
		log.Println("GetOptionById: Error get option", err)
		return nil, err
	}
	return image, nil
}

func (p productImageService) DeleteImage(b *dto.ImageDTO) (*entity.ProductImages, error) {
	image, err := p.ProductImageRepository.DeleteImage(b)
	if err != nil {
		log.Println("DeleteOption: Error delete option", err)
		return nil, err
	}
	return image, nil
}

func (p productImageService) GetImage(b *dto.ProductIdDTO) (*[]entity.ProductImages, error) {
	options, err := p.ProductImageRepository.GetImage(b)
	if err != nil {
		log.Println("GetOptions: Error get options", err)
		return nil, err
	}
	return options, nil
}

func (p productImageService) CreateImage(b *dto.CreateImageDTO) (*entity.ProductImages, error) {
	var proDTO dto.ProductDTO
	proDTO.ProductId = b.Image.ProductId
	countProd, prodErr := p.ProductImageRepository.ProductImageByID(&proDTO)
	if prodErr != nil {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, prodErr
	}
	if countProd == 0 {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, errors.New("not found product")
	}
	image, err := p.ProductImageRepository.CreateImage(b)
	if err != nil {
		log.Println("CreateOption: Error to create option", err)
		return nil, err
	}
	return image, nil
}

type productImageService struct {
	ProductImageRepository repository.ProductImageRepository
}

func NewProductImageService(productImageRepository repository.ProductImageRepository) *productImageService {
	return &productImageService{ProductImageRepository: productImageRepository}
}
