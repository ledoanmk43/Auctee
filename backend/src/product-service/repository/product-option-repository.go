package repository

import (
	"backend/src/product-service/dto"
	"backend/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IProductOptionRepository interface {
	CreateOption(option *entity.ProductOption, userId uint) error
	UpdateOption(option *entity.ProductOption, userId uint) error
	DeleteOption(option *entity.ProductOption, userId uint) error
	ProductOptionByID(productId string) (int64, error)
	GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error)
	GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error)
}

type ProductOptionRepository struct {
	connection *gorm.DB
}

func NewProductOptionRepository(dbConn *gorm.DB) *ProductOptionRepository {
	return &ProductOptionRepository{
		connection: dbConn,
	}
}

func (p *ProductOptionRepository) UpdateOption(optionBody *entity.ProductOption, userId uint) error {
	var countPro int64
	var product *entity.Product
	var option *entity.ProductOption
	record := p.connection.Where("id = ? AND user_id = ? ", optionBody.ProductId, userId).Find(&product).Count(&countPro)
	if record.Error != nil || countPro == 0 {
		log.Println("line 36")
		return errors.New("option not found")
	}

	var countOption int64
	res := p.connection.Where("product_id = ? AND id = ?", optionBody.ProductId, optionBody.ID).Find(&option).Count(&countOption)

	if res.Error != nil || countOption == 0 {
		log.Println("line 44")
		return errors.New("option not found")
	}

	option.Model = optionBody.Model
	option.Color = optionBody.Color
	option.Size = optionBody.Size
	option.Quantity = optionBody.Quantity
	resUpdate := p.connection.Updates(&option)
	if resUpdate.Error != nil {
		return errors.New("option not found")
	}

	return nil
}

func (p *ProductOptionRepository) ProductOptionByID(productId string) (int64, error) {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ?", productId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
}

func (p *ProductOptionRepository) CreateOption(option *entity.ProductOption, userId uint) error {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ? AND user_id = ? ", option.ProductId, userId).Count(&count)
	if record.Error != nil {
		return errors.New("product not found")
	}

	if count == 0 {
		log.Println("option not found")
		return errors.New("product not found")
	}

	record = p.connection.Create(&option)
	if record.Error != nil {
		log.Println("CreateOption: Error to create repository")
		return record.Error
	}
	return nil
}

func (p *ProductOptionRepository) GetOptions(b *dto.ProductIdDTO) (*[]entity.ProductOption, error) {
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

func (p *ProductOptionRepository) GetOptionByID(b *dto.OptionIdDTO) (*entity.ProductOption, error) {
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

func (p *ProductOptionRepository) DeleteOption(option *entity.ProductOption, userId uint) error {
	var countPro int64
	var product *entity.Product
	record := p.connection.Where("id = ? AND user_id = ? ", option.ProductId, userId).Find(&product).Count(&countPro)
	if record.Error != nil || countPro == 0 {
		return errors.New("option not found")
	}

	var countImg int64
	res := p.connection.Where("product_id = ? AND id = ?", option.ProductId, option.ID).Find(&option).Count(&countImg)
	if res.Error != nil || countImg == 0 {
		return errors.New("option not found")
	}

	resDel := p.connection.Delete(&option)
	if resDel.Error != nil {
		return errors.New("option not found")
	}
	return nil
}
