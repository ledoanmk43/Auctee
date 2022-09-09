package controller

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

const (
	imageId = "imageId"
)

type ProductImageController interface {
	CreateImage(c *gin.Context)
	GetImage(c *gin.Context)
	GetImageByID(c *gin.Context)
	DeleteImage(c *gin.Context)
	UpdateImage(c *gin.Context)
}

func (p productImageController) UpdateImage(c *gin.Context) {
	var imageUpdateBody *entity.ProductImages
	if err := c.ShouldBindJSON(&imageUpdateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	oid, errCv := strconv.Atoi(c.Param(imageId))
	if errCv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error update image",
		})
		log.Println("UpdateImage: Error parse param", errCv)
		c.Abort()
		return
	}
	dtoUpdate := dto.NewUpdateImageDTO(imageUpdateBody)
	dtoUpdate.ImageId = imageId
	dtoUpdate.Image.ID = uint(oid)
	product, err := p.productImageService.UpdateImage(dtoUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to update product",
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
}

func (p productImageController) GetImageByID(c *gin.Context) {

	var dto dto.ImageDTO
	dto.ImageId = c.Param(imageId)
	c.Set(imageId, dto.ImageId)
	image, err := p.productImageService.GetImageByID(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error get image",
		})
		log.Println("GetImageById: Error call service in pkg controller", err)
		c.Abort()
		return
	}
	if image == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no image found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, image)
}
func (p productImageController) DeleteImage(c *gin.Context) {
	oId := c.Param(imageId)
	var dto dto.ImageDTO
	dto.ImageId = oId
	image, err := p.productImageService.DeleteImage(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to delete image",
		})
		log.Println("DeleteImage: Error to parse oId", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, image)
}
func (p productImageController) GetImage(c *gin.Context) {
	id := c.Param(productId)
	var dto dto.ProductIdDTO
	dto.ProductId = id
	images, err := p.productImageService.GetImage(&dto)
	if err != nil {
		log.Println("GetImages: error in controller package", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to get images",
		})
		c.Abort()
		return
	}
	if images == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "no image found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, images)
}

func (p productImageController) CreateImage(c *gin.Context) {
	var imageBody *entity.ProductImages
	if err := c.ShouldBindJSON(&imageBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to create image",
		})
		log.Println("CreateImage: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dtoImage := dto.NewCreateImageDTO(imageBody)
	dtoImage.Image.ProductId = c.Param(productId)
	image, err := p.productImageService.CreateImage(dtoImage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create image",
		})
		log.Println("CreateImage: Error to CreateImage in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, image)
}

type productImageController struct {
	productImageService service.ProductImageService
}

func NewProductImageController(productImageService service.ProductImageService) *productImageController {
	return &productImageController{productImageService: productImageService}
}
