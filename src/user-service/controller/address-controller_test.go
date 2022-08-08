package controller

import (
	"bytes"
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	service "chilindo/src/user-service/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateTestAddress(t *testing.T) (*service.MockIAddressService, *AddressController) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockSvc := service.NewMockIAddressService(ctl)
	userCtl := NewAddressControllerDefault(mockSvc)
	return mockSvc, userCtl
}

func TestAddressController_CreateAddress(t *testing.T) {
	mockSrv, userCtr := CreateTestAddress(t)
	//Mock
	mockSrv.EXPECT().CreateAddress(gomock.Any()).Return(&entity.Address{
		Model:       gorm.Model{},
		Firstname:   "",
		Lastname:    "",
		Phone:       "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		UserId:      0,
		User:        entity.User{},
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("POST", "/chilindo/user/address/create", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Set(config.UserId, uint(1))

	c.Request = req

	userCtr.CreateAddress(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
}

func TestAddressController_GetAddress(t *testing.T) {
	mockSvr, userCtr := CreateTestAddress(t)
	mockSvr.EXPECT().GetAddress(gomock.Any()).Return(&[]entity.Address{{
		Model:       gorm.Model{},
		Firstname:   "",
		Lastname:    "",
		Phone:       "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		UserId:      0,
		User:        entity.User{},
	}}, nil).Times(1)

	req, err := http.NewRequest("GET", "chilindo/user/address/getaddress", nil)

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set(config.UserId, uint(1))

	userCtr.GetAddress(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
}

func TestAddressController_DeleteAddress(t *testing.T) {
	mockSvc, addressCtl := CreateTestAddress(t)
	mockSvc.EXPECT().DeleteAddress(gomock.Any()).Return(nil).Times(1)

	req, err := http.NewRequest("DELETE", "chilindo/user/address/delete/:id", nil)
	if err != nil {
		t.Fatal("Error")
	}

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
	c.Set(config.UserId, uint(1))
	addressCtl.DeleteAddress(c)
	if rr.Code != http.StatusOK {
		t.Fatalf("Status is expected 200 but %v", rr.Code)
	}
}

func TestAddressController_UpdateAddress(t *testing.T) {
	mockSvc, addressCtl := CreateTestAddress(t)
	//Mock svc
	mockSvc.EXPECT().UpdateAddress(gomock.Any()).Return(&entity.Address{
		Model:       gorm.Model{},
		Firstname:   "",
		Lastname:    "",
		Phone:       "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		UserId:      0,
		User:        entity.User{},
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("PUT", "chilindo/user/address/update/:id", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal("Error")
	}

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
	c.Set(config.UserId, uint(1))
	c.Request = req
	addressCtl.UpdateAddress(c)

	if rr.Code != http.StatusOK {
		t.Fatalf("Status expected is 200 but %v", rr.Code)
	}
}

func TestAddressController_GetAddressById(t *testing.T) {
	mockSvr, userCtr := CreateTestAddress(t)
	mockSvr.EXPECT().GetAddressById(gomock.Any()).Return(&entity.Address{
		Model:       gorm.Model{},
		Firstname:   "",
		Lastname:    "",
		Phone:       "",
		Province:    "",
		District:    "",
		SubDistrict: "",
		Address:     "",
		TypeAddress: "",
		UserId:      0,
		User:        entity.User{},
	}, nil).Times(1)

	req, err := http.NewRequest("GET", "chilindo/user/address/getaddress/:id", nil)

	if err != nil {
		t.Fatalf("Error")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Params = []gin.Param{gin.Param{Key: "id", Value: "1"}}
	c.Set(config.UserId, uint(1))
	userCtr.GetAddressById(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
}
