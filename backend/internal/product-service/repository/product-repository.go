package repository

import (
	"backend/internal/product-service/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type IProductRepository interface {
	InsertProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(productId string, userIdId uint) error
	GetAllProducts(userId uint) (*[]entity.Product, error)
	GetProductByProductId(productId string) (*entity.Product, error)
	GetProductsByProductName(productName string) (*entity.ProductResponse, error)
	GetProductDetailByProductId(productId string) (*entity.Product, error)
}

type ProductRepositoryDefault struct {
	connection        *gorm.DB
	ProductOptionRepo IProductOptionRepository
	ProductImageRepo  IProductImageRepository
}

func NewProductRepositoryDefault(dbConn *gorm.DB, optionRepo IProductOptionRepository, imageRepo IProductImageRepository) *ProductRepositoryDefault {
	return &ProductRepositoryDefault{
		connection:        dbConn,
		ProductOptionRepo: optionRepo,
		ProductImageRepo:  imageRepo,
	}
}

func (p *ProductRepositoryDefault) DeleteProduct(proId string, userId uint) error {
	var product *entity.Product
	var count int64
	resultProduct := p.connection.Where("id = ? AND user_id = ?", proId, userId).Find(&product).Count(&count)
	if resultProduct.Error != nil {
		log.Println("DeleteProduct: Error in find product to delete in package repository", resultProduct.Error)
		return resultProduct.Error
	}
	if count == 0 {
		return errors.New("product not found")
	}
	p.connection.Delete(&product)

	var productOptions *entity.ProductOption
	resultOption := p.connection.Where("product_id = ?", proId).Limit(99).Delete(&productOptions)
	if resultOption.Error != nil {
		log.Println("DeleteProduct: Error in find option to delete in package repository", resultOption.Error)
		return resultOption.Error
	}
	var productImages *entity.ProductImage
	resultImage := p.connection.Where("product_id = ?", proId).Limit(99).Delete(&productImages)
	if resultOption.Error != nil {
		log.Println("DeleteProduct: Error in find option to delete in package repository", resultImage.Error)
		return resultImage.Error
	}
	return nil
}

func (p *ProductRepositoryDefault) InsertProduct(pro *entity.Product) error {
	if errCheckEmptyField := pro.Validate("insert"); errCheckEmptyField != nil {
		log.Println("VerifyCredential: Error empty field in package repository", errCheckEmptyField)
		return errCheckEmptyField
	}

	record := p.connection.Create(&pro)
	if record.Error != nil {
		log.Println("Error to create product in repo")
		return record.Error
	}

	//for i := 0; i < len(pro.ProductOption); i++ {
	//	pro.ProductOption[i].ProductId = pro.Id
	//	errOption := p.ProductOptionRepo.CreateOption(&pro.ProductOption[i], pro.UserId)
	//	if errOption != nil {
	//		log.Println("Error to create product in repo")
	//		return errOption
	//	}
	//}
	//for i := 0; i < len(pro.ProductImage); i++ {
	//	pro.ProductImage[i].ProductId = pro.Id
	//	errImage := p.ProductImageRepo.CreateImage(&pro.ProductImage[i], pro.UserId)
	//	if errImage != nil {
	//		log.Println("Error to create product in repo")
	//		return errImage
	//	}
	//}
	return nil
}

func (p *ProductRepositoryDefault) UpdateProduct(productBody *entity.Product) error {
	var productToUpdate *entity.Product
	var count int64
	record := p.connection.Where("id = ? AND user_id = ?", productBody.Id, productBody.UserId).Find(&productToUpdate).Count(&count)

	if record.Error != nil {
		log.Println("Error to find product in repo", record.Error)
		return record.Error
	}
	if count == 0 {
		return errors.New("product not found")
	}
	if productBody.ExpectPrice < productToUpdate.MinPrice || productBody.MinPrice >= productBody.ExpectPrice {
		return errors.New("expect price should be larger than minimum price")
	}

	errImage := p.ProductImageRepo.DeleteAllImages(productBody.Id)
	if errImage != nil {
		log.Println("Error to delete image in product repo")
	}

	for i := 0; i < len(productBody.ProductImage); i++ {
		productBody.ProductImage[i].ProductId = productBody.Id
		if productBody.ProductImage[i].Id > 0 {
			productBody.ProductImage[i].Id = 0
		}
		errOption := p.ProductImageRepo.UpdateImage(&productBody.ProductImage[i], productBody.UserId)
		if errOption != nil {
			log.Println("Error to create product in repo")
		}
	}

	errOption := p.ProductOptionRepo.DeleteAllOption(productBody.Id)
	if errOption != nil {
		log.Println("Error to delete option in product repo")
	}
	for i := 0; i < len(productBody.ProductOption); i++ {
		productBody.ProductOption[i].ProductId = productBody.Id
		if productBody.ProductOption[i].Id > 0 {
			productBody.ProductOption[i].Id = 0
		}
		errOption := p.ProductOptionRepo.UpdateOption(&productBody.ProductOption[i], productBody.UserId)
		if errOption != nil {
			log.Println("Error to create product in repo")
		}
	}

	productToUpdate.Name = productBody.Name
	productToUpdate.MinPrice = productBody.MinPrice
	productToUpdate.Description = productBody.Description
	productToUpdate.Quantity = productBody.Quantity
	productToUpdate.ExpectPrice = productBody.ExpectPrice

	recordSave := p.connection.Updates(&productToUpdate)
	if recordSave.Error != nil {
		log.Println("Error to update product repo", recordSave.Error)
		return recordSave.Error
	}
	return nil
}

func (p *ProductRepositoryDefault) GetAllProducts(userId uint) (*[]entity.Product, error) {
	var products *[]entity.Product
	var count int64
	record := p.connection.Where("user_id = ?", userId).Order("created_at desc").Find(&products).Count(&count)
	if record.Error != nil {
		log.Println("GetProducts: Error get all products in repo", record.Error)
		return nil, record.Error
	}
	if count == 0 {
		return nil, errors.New("product not found")
	}

	for i := 0; i < len(*products); i++ {
		options, _ := p.ProductOptionRepo.GetOptions((*products)[i].Id)
		if options == nil {
			continue
		}
		(*products)[i].ProductOption = *options
	}

	for j := 0; j < len(*products); j++ {
		images, _ := p.ProductImageRepo.GetAllImages((*products)[j].Id)
		if images == nil {
			continue
		}
		(*products)[j].ProductImage = *images
	}
	return products, nil
}
func (p *ProductRepositoryDefault) GetProductByProductId(productId string) (*entity.Product, error) {
	var product *entity.Product
	var count int64
	record := p.connection.Where("id = ?", productId).Find(&product).Count(&count)
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

func (p *ProductRepositoryDefault) GetProductsByProductName(productName string) (*entity.ProductResponse, error) {
	var products []entity.Product
	record := p.connection.Where("name LIKE ?", "%"+productName+"%").Find(&products)
	if record.Error != nil {
		log.Println("Get product by name", record.Error)
		return nil, record.Error
	}
	var productList entity.ProductResponse
	for i, _ := range products {
		productList.IdList = append(productList.IdList, products[i].Id)
		productList.ProductName = append(productList.ProductName, products[i].Name)
	}
	return &productList, nil
}

func (p *ProductRepositoryDefault) GetProductDetailByProductId(productId string) (*entity.Product, error) {
	var productDetail *entity.Product
	var countProd int64
	resultProduct := p.connection.Where("id = ? ", productId).Find(&productDetail).Count(&countProd)
	if resultProduct.Error != nil {
		log.Println("Get Product: Error in find product to get in package repository", resultProduct.Error)
		return nil, resultProduct.Error
	}
	if countProd == 0 {
		return nil, errors.New("product not found")
	}

	//Options

	resultOption := p.connection.Where("product_id = ?", productId).Limit(99).Find(&productDetail.ProductOption)
	if resultOption.Error != nil {
		log.Println("Get Product: Error in find option to get in package repository", resultOption.Error)
		return nil, resultOption.Error
	}

	//Images

	resultImage := p.connection.Where("product_id = ?", productId).Limit(99).Find(&productDetail.ProductImage)
	if resultOption.Error != nil {
		log.Println("Get Product: Error in find image to get in package repository", resultImage.Error)
		return nil, resultImage.Error
	}

	// return all properties of the product
	return productDetail, nil
}
