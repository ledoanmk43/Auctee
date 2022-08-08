package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/repository"
	"errors"
	"log"
)

type ProductOptionService interface {
	CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error)
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error)
}

func (p productOptionService) UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error) {
	option, err := p.ProductOptionRepository.UpdateOption(b)
	if err != nil {
		log.Println("UpdateOption: Error call repo")
		return nil, err
	}
	return option, nil
}

func (p productOptionService) CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error) {
	var proDTO dto.ProductDTO
	proDTO.ProductId = b.Option.ProductId
	countProd, prodErr := p.ProductOptionRepository.ProductOptionByID(&proDTO)
	if prodErr != nil {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, prodErr
	}
	if countProd == 0 {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return nil, errors.New("not found product")
	}
	option, err := p.ProductOptionRepository.CreateOption(b)
	if err != nil {
		log.Println("CreateOption: Error to create option", err)
		return nil, err
	}
	return option, nil
}
func (p productOptionService) GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error) {
	options, err := p.ProductOptionRepository.GetOptions(b)
	if err != nil {
		log.Println("GetOptions: Error get options", err)
		return nil, err
	}
	return options, nil
}
func (p productOptionService) GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	option, err := p.ProductOptionRepository.GetOptionByID(b)
	if err != nil {
		log.Println("GetOptionById: Error get option", err)
		return nil, err
	}
	return option, nil
}
func (p productOptionService) DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	option, err := p.ProductOptionRepository.DeleteOption(b)
	if err != nil {
		log.Println("DeleteOption: Error delete option", err)
		return nil, err
	}
	return option, nil
}

type productOptionService struct {
	ProductOptionRepository repository.ProductOptionRepository
}

func NewProductOptionService(productOptionRepository repository.ProductOptionRepository) *productOptionService {
	return &productOptionService{ProductOptionRepository: productOptionRepository}
}
