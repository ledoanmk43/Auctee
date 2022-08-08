package repository

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ProductOptionRepository interface {
	CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error)
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error)
	ProductOptionByID(b *dto.ProductDTO) (int64, error)
	UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error)
}

func (p productOptionRepository) UpdateOption(b *dto.UpdateOptionDTO) (*entity.ProductOption, error) {
	//TODO implement me
	var updateOption *entity.ProductOption
	var count int64
	record := p.connection.Where("id = ?", b.Option.ID).Find(&updateOption).Count(&count)

	if record.Error != nil {
		log.Println("Error to find product repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("option not found")
	}
	updateOption = b.Option
	recordSave := p.connection.Updates(&updateOption)
	if recordSave.Error != nil {
		log.Println("Error to update produce repo", recordSave.Error)
		return nil, recordSave.Error
	}
	return updateOption, nil
}

func (p productOptionRepository) ProductOptionByID(b *dto.ProductDTO) (int64, error) {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ?", b.ProductId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
}

func (p productOptionRepository) CreateOption(b *dto.CreateOptionDTO) (*entity.ProductOption, error) {
	record := p.connection.Create(&b.Option)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return nil, record.Error
	}
	return b.Option, nil

}

func (p productOptionRepository) GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error) {
	//TODO implement me
	var options *[]entity.ProductOption
	var count int64
	record := p.connection.Where("product_id = ?", b.ProductId).Find(&options).Count(&count)
	if record.Error != nil {
		log.Println("GetOptions : Error to get all option", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptions : Not found Options", count)
		return nil, nil
	}
	return options, nil
}

func (p productOptionRepository) GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	//TODO implement me
	var option *entity.ProductOption
	var count int64
	record := p.connection.Where("id = ?", b.OptionId).Find(&option).Count(&count)
	if record.Error != nil {
		log.Println("GetOptionById: Error to get option in repo pkg", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetOptionById: Not found option", count)
		return nil, nil
	}
	return option, nil
}

func (p productOptionRepository) DeleteOption(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
	//TODO implement me
	var option *entity.ProductOption
	record := p.connection.Where("id = ?", b.OptionId).Find(&option)
	if record.Error != nil {
		log.Println("DeleteOption: Error to find option", record.Error)
		return nil, record.Error
	}
	recordDelete := p.connection.Delete(&option)
	if recordDelete.Error != nil {
		log.Println("DeleteOption: Error to delete option", record.Error)
		return nil, recordDelete.Error
	}
	return option, nil
}

type productOptionRepository struct {
	connection *gorm.DB
}

func NewProductOptionRepository(connection *gorm.DB) *productOptionRepository {
	return &productOptionRepository{connection: connection}
}
