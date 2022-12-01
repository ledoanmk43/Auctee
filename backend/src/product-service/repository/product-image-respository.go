package repository

import (
	"backend/src/product-service/dto"
	"backend/src/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

const (
	inactive = false
)

type IProductImageRepository interface {
	CreateImage(image *entity.ProductImage, userId uint) error
	UpdateImage(image *entity.ProductImage, userId uint) error
	DeleteImage(image *entity.ProductImage, userId uint) error
	DeleteAllImages(proId string) error
	GetDefaultImageByProductId(productId string) (*entity.ProductImage, error)
	ProductImageByID(productId string) (int64, error)
	GetAllImages(productId string) (*[]entity.ProductImage, error)
	GetImageByID(b *dto.ImageDTO) (*entity.ProductImage, error)
}

type ProductImageRepository struct {
	connection *gorm.DB
}

func NewProductImageRepository(dbConn *gorm.DB) *ProductImageRepository {
	return &ProductImageRepository{
		connection: dbConn,
	}
}

func (p *ProductImageRepository) GetDefaultImageByProductId(productId string) (*entity.ProductImage, error) {
	var image *entity.ProductImage
	var countImg int64
	res := p.connection.Where("product_id = ? AND is_default = ?", productId, 1).Find(&image).Count(&countImg)

	if res.Error != nil || countImg == 0 {
		log.Println("line 44")
		return nil, errors.New("image not found")
	}

	return image, nil
}
func (p *ProductImageRepository) UpdateImage(imageBody *entity.ProductImage, userId uint) error {

	//var countImg int64
	//res := p.connection.Where("product_id = ? AND id = ?", imageBody.ProductId, imageBody.ID).Find(&image).Count(&countImg)
	//if res.Error != nil || countImg == 0 {
	//	log.Println("line 62: ", res.Error)
	//}
	//
	//image.Path = imageBody.Path
	//image.IsDefault = imageBody.IsDefault

	//var imageCheckDefault *entity.ProductImage
	//var countDefault int64
	//_ = p.connection.Where("product_id = ? AND id != ? AND is_default = ?", imageBody.ProductId, imageBody.ID, true).Find(&imageCheckDefault).Count(&countDefault)
	//if imageCheckDefault != nil {
	//	imageCheckDefault.IsDefault = utils.BoolAddr(false)
	//	_ = p.connection.Where("product_id = ? AND id != ? AND is_default = ?", imageBody.ProductId, imageBody.ID, true).Updates(&imageCheckDefault)
	//}
	resUpdate := p.connection.Create(&imageBody)
	if resUpdate.Error != nil {
		return errors.New("image not found")
	}

	return nil
}

func (p *ProductImageRepository) GetImageByID(b *dto.ImageDTO) (*entity.ProductImage, error) {
	var image *entity.ProductImage
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

func (p *ProductImageRepository) DeleteImage(image *entity.ProductImage, userId uint) error {
	var countPro int64
	var product *entity.Product
	record := p.connection.Where("id = ? AND user_id = ? ", image.ProductId, userId).Find(&product).Count(&countPro)
	if record.Error != nil || countPro == 0 {
		return errors.New("image not found")
	}

	var countImg int64
	res := p.connection.Where("product_id = ? AND id = ?", image.ProductId, image.ID).Find(&image).Count(&countImg)
	if res.Error != nil || countImg == 0 {
		return errors.New("image not found")
	}

	resDel := p.connection.Delete(&image)
	if resDel.Error != nil {
		return errors.New("image not found")
	}
	return nil
}

func (p *ProductImageRepository) DeleteAllImages(proId string) error {
	var images *[]entity.ProductImage
	res := p.connection.Where("product_id = ?", proId).Delete(&images)
	if res.Error != nil {
		return errors.New("no image to delete")
	}
	return nil
}

func (p *ProductImageRepository) GetAllImages(productId string) (*[]entity.ProductImage, error) {
	var images *[]entity.ProductImage
	var count int64
	record := p.connection.Where("product_id = ?", productId).Find(&images).Count(&count)
	if record.Error != nil {
		log.Println("GetOptions : Error to get all images", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		log.Println("GetImages : Not found images", count)
		return nil, nil
	}
	return images, nil
}

func (p *ProductImageRepository) ProductImageByID(productId string) (int64, error) {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ?", productId).Count(&count)
	if record.Error != nil {
		log.Println("CountProductById: Get product by ID", record.Error)
		return count, record.Error
	}
	return count, nil
}

func (p *ProductImageRepository) CreateImage(image *entity.ProductImage, userId uint) error {
	var count int64
	record := p.connection.Model(&entity.Product{}).Where("id = ? AND user_id = ? ", image.ProductId, userId).Count(&count)
	if record.Error != nil {
		return errors.New("product not found")
	}

	if count == 0 {
		log.Println("image not found")
		return errors.New("product not found")
	}

	record = p.connection.Create(&image)
	if record.Error != nil {
		log.Println("Create image: Error to create repository")
		return record.Error
	}
	return nil
}
