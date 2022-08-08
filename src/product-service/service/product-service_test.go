package service

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	service "chilindo/src/product-service/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func Init(t *testing.T) (*service.MockProductRepository, *productService) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockProductRepo := service.NewMockProductRepository(ctr)
	return mockProductRepo, NewProductService(mockProductRepo)
}
func TestNewProductService(t *testing.T) {
	mockProductRepo, testProductService := Init(t)
	mockProductRepo.EXPECT().InsertProduct(gomock.Any()).Return(&entity.Product{
		Id:          "",
		Name:        "",
		MinPrice:    "",
		Description: "",
		Quantity:    0,
	}, nil)
	var b *dto.ProductCreatedDTO
	_, err := testProductService.Insert(b)
	if err != nil {
		t.Fail()
	}
}
func TestProductService_Update(t *testing.T) {
	mockProductRepo, testProductService := Init(t)
	mockProductRepo.EXPECT().UpdateProduct(gomock.Any()).Return(&entity.Product{
		Id:          "",
		Name:        "",
		MinPrice:    "",
		Description: "",
		Quantity:    0,
	}, nil)
	var b *dto.ProductUpdateDTO

	_, err := testProductService.Update(b)
	if err != nil {
		t.Fail()
	}
}

func TestProductService_FindProductByID(t *testing.T) {
	mockProductRepo, testProductService := Init(t)
	mockProductRepo.EXPECT().FindProductByID(gomock.Any()).Return(&entity.Product{
		Id:          "",
		Name:        "",
		MinPrice:    "",
		Description: "",
		Quantity:    0,
	}, nil)
	var b *dto.ProductDTO
	_, err := testProductService.FindProductByID(b)
	if err != nil {
		t.Fail()
	}
}
func TestProductService_All(t *testing.T) {
	mockProductRepo, testProductService := Init(t)
	mockProductRepo.EXPECT().AllProduct().Return(&[]entity.Product{}, nil)
	_, err := testProductService.All()
	if err != nil {
		t.Fail()
	}

}
func TestProductService_Delete(t *testing.T) {
	mockProductRepo, testProductService := Init(t)
	mockProductRepo.EXPECT().DeleteProduct(gomock.Any()).Return(&entity.Product{
		Id:          "",
		Name:        "",
		MinPrice:    "",
		Description: "",
		Quantity:    0,
	}, nil)
	var b *dto.ProductDTO
	_, err := testProductService.Delete(b)
	if err != nil {
		t.Fail()
	}
}
