package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	service "chilindo/src/product-service/mock"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func testoption(t *testing.T) (*service.MockProductOptionRepository, *productOptionService) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockOptionRepo := service.NewMockProductOptionRepository(ctr)
	return mockOptionRepo, NewProductOptionService(mockOptionRepo)
}
func TestProductOptionService_GetOptions(t *testing.T) {
	mockOptionRepo, testOptionService := testoption(t)
	mockOptionRepo.EXPECT().GetOptions(gomock.Any()).Return(&[]entity.ProductOption{}, nil)
	var b *dto.ProductIdDTO
	_, err := testOptionService.GetOptions(b)
	if err != nil {
		t.Fail()
	}
}
func TestProductOptionService_UpdateOption(t *testing.T) {
	mockOptionRepo, testOptionService := testoption(t)
	mockOptionRepo.EXPECT().UpdateOption(gomock.Any()).Return(&entity.ProductOption{
		Model:     gorm.Model{},
		ProductId: "",
		Color:     "",
		Size:      "",
		Models:    "",
		Product:   entity.Product{},
	}, nil)
	var b *dto.UpdateOptionDTO
	_, err := testOptionService.UpdateOption(b)
	if err != nil {
		t.Fail()
	}
}
func TestProductOptionService_DeleteOption(t *testing.T) {
	mockOptionRepo, testOptionService := testoption(t)
	mockOptionRepo.EXPECT().DeleteOption(gomock.Any()).Return(&entity.ProductOption{
		Model:     gorm.Model{},
		ProductId: "",
		Color:     "",
		Size:      "",
		Models:    "",
		Product:   entity.Product{},
	}, nil)
	var b *dto.OptionIdDTO
	_, err := testOptionService.DeleteOption(b)
	if err != nil {
		t.Fail()
	}
}
func TestProductOptionService_GetOptionByID(t *testing.T) {
	mockOptionRepo, testOptionService := testoption(t)
	mockOptionRepo.EXPECT().GetOptionByID(gomock.Any()).Return(&entity.ProductOption{
		Model:     gorm.Model{},
		ProductId: "",
		Color:     "",
		Size:      "",
		Models:    "",
		Product:   entity.Product{},
	}, nil)
	var b *dto.OptionIdDTO
	_, err := testOptionService.GetOptionByID(b)
	if err != nil {
		t.Fail()
	}
}
