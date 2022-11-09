package dto

import "backend/internal/product-service/entity"

type CreateImageDTO struct {
	Image *entity.ProductImage
}

func NewCreateImageDTO(image *entity.ProductImage) *CreateImageDTO {
	return &CreateImageDTO{Image: image}
}

type ImageDTO struct {
	ImageId string
}
type UpdateImageDTO struct {
	Image   *entity.ProductImage
	ImageId string
}

func NewUpdateImageDTO(image *entity.ProductImage) *UpdateImageDTO {
	return &UpdateImageDTO{Image: image}
}
