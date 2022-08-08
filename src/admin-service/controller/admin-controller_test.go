package controller

import (
	"bytes"
	"chilindo/src/admin-service/entity"
	service "chilindo/src/admin-service/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTestAdmin(t *testing.T) (*service.MockIAdminService, *AdminController) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockSvc := service.NewMockIAdminService(ctr)
	adminCtr := NewAdminControllerDefault(mockSvc)
	return mockSvc, adminCtr
}

func TestAdminController_SignIn(t *testing.T) {
	mockSvc, authCtr := createTestAdmin(t)

	mockSvc.EXPECT().VerifyCredential(gomock.Any()).Return(&entity.Admin{
		Model:    gorm.Model{},
		Id:       1,
		Username: "",
		Password: "",
		Token:    "",
	}, nil).Times(1)

	body := []byte("{}")

	req, err := http.NewRequest("POST", "chilindo/admin/sign-in", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Error")
	}
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	authCtr.SignIn(c)
	if w.Code != http.StatusOK {
		t.Fatalf("200 but got %v", w.Code)
	}
}

func TestAdminController_SignUp(t *testing.T) {
	mockSvc, adminCtr := createTestAdmin(t)

	//mock service
	mockSvc.EXPECT().CreateAdmin(gomock.Any()).Return(&entity.Admin{
		Model:    gorm.Model{},
		Username: "",
		Password: "",
		Token:    "",
	}, nil).Times(1)

	mockSvc.EXPECT().IsDuplicateUsername(gomock.Any()).Return(false)

	bodyRequest := `{"username":""}`

	req, err := http.NewRequest("POST", "chilindo/admin/sign-up", strings.NewReader(bodyRequest))
	if err != nil {
		t.Fatalf("error %v", err)
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	adminCtr.SignUp(c)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status is 201 but got %v", w.Code)
	}

}
