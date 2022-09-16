package service

import (
	"backend/src/product-service/dto"
	"backend/src/product-service/entity"
	"backend/src/product-service/repository"
	"errors"
	"log"
)

type IProductOptionService interface {
	CreateOption(option *entity.ProductOption, userId uint) error
	UpdateOption(option *entity.ProductOption, userId uint) error
	DeleteOption(option *entity.ProductOption, userId uint) error
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error)
}

type ProductOptionService struct {
	ProductOptionRepository repository.IProductOptionRepository
}

func NewProductOptionService(productOptionRepository repository.IProductOptionRepository) *ProductOptionService {
	return &ProductOptionService{ProductOptionRepository: productOptionRepository}
}

func (p *ProductOptionService) UpdateOption(option *entity.ProductOption, userId uint) error {
	err := p.ProductOptionRepository.UpdateOption(option, userId)
	if err != nil {
		log.Println("Update Option: ", err)
		return err
	}
	return nil
}

func (p *ProductOptionService) CreateOption(option *entity.ProductOption, userId uint) error {
	productId := option.ProductId
	countProd, prodErr := p.ProductOptionRepository.ProductOptionByID(productId)
	if prodErr != nil {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return prodErr
	}
	if countProd == 0 {
		log.Println("CreateOption: Error not found product to create option", prodErr)
		return errors.New("product not found")
	}
	err := p.ProductOptionRepository.CreateOption(option, userId)
	if err != nil {
		log.Println("CreateOption: Error to create option", err)
		return err
	}
	return nil
}
func (p *ProductOptionService) GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error) {
	options, err := p.ProductOptionRepository.GetOptions(b)
	if err != nil {
		log.Println("GetOptions: Error get options", err)
		return nil, err
	}
	return options, nil
}
func (p *ProductOptionService) GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	option, err := p.ProductOptionRepository.GetOptionByID(b)
	if err != nil {
		log.Println("GetOptionById: Error get option", err)
		return nil, err
	}
	return option, nil
}
func (p *ProductOptionService) DeleteOption(option *entity.ProductOption, userId uint) error {
	err := p.ProductOptionRepository.DeleteOption(option, userId)
	if err != nil {
		log.Println("DeleteOption: Error delete option", err)
		return err
	}
	return nil
}
