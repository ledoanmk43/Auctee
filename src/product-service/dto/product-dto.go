package dto

import "chilindo/src/product-service/entity"

type ProductCreatedDTO struct {
	Product *entity.Product
}

func NewProductCreatedDTO(product *entity.Product) *ProductCreatedDTO {
	return &ProductCreatedDTO{Product: product}
}

type ProductUpdateDTO struct {
	ProductId string
	Product   *entity.Product
}

func NewProductUpdateDTO(product *entity.Product) *ProductUpdateDTO {
	return &ProductUpdateDTO{Product: product}
}

type ProductDTO struct {
	ProductId string
}
