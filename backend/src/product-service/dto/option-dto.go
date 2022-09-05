package dto

import "chilindo/src/product-service/entity"

type CreateOptionDTO struct {
	Option *entity.ProductOption
}
type UpdateOptionDTO struct {
	Option   *entity.ProductOption
	OptionId string
}

func NewUpdateOptionDTO(option *entity.ProductOption) *UpdateOptionDTO {
	return &UpdateOptionDTO{Option: option}
}
func NewCreateOptionDTO(option *entity.ProductOption) *CreateOptionDTO {
	return &CreateOptionDTO{Option: option}
}

type OptionIdDTO struct {
	OptionId string
}

type ProductIdDTO struct {
	ProductId string
}
