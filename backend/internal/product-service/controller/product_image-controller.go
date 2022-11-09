package controller

import (
	"backend/internal/account-service/config"
	product "backend/internal/product-service/config"
	"backend/internal/product-service/entity"
	"backend/internal/product-service/service"
	"backend/pkg/token"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IProductImageController interface {
	CreateImage(ctx *gin.Context)
	UpdateImage(ctx *gin.Context)
	DeleteImage(ctx *gin.Context)
	GetImage(ctx *gin.Context)
	GetImageByID(ctx *gin.Context)
}
type ProductImageController struct {
	ProductImageService service.IProductImageService
}

func NewProductImageController(productImageService service.IProductImageService) *ProductImageController {
	return &ProductImageController{ProductImageService: productImageService}
}

func (p *ProductImageController) UpdateImage(ctx *gin.Context) {
	var imageBody entity.ProductImage
	if err := ctx.ShouldBindJSON(&imageBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("CreateImage: Error to ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	imageId, errGetId := strconv.Atoi(ctx.Query(product.Id))
	if errGetId != nil {
		log.Println("error in get image by imageId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	imageBody.ProductId = ctx.Param(product.ProductId)
	imageBody.ID = uint(imageId)
	errUpdate := p.ProductImageService.UpdateImage(&imageBody, claims.UserId)
	if errUpdate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errUpdate.Error(),
		})
		log.Println("CreateImage: Error to CreateImage in package controller", errUpdate)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "image updated",
	})
}

func (p *ProductImageController) GetImageByID(c *gin.Context) {
	//
	//var dto dto.ImageDTO
	//dto.ImageId = c.Param(imageId)
	//c.Set(imageId, dto.ImageId)
	//image, err := p.ProductImageService.GetImageByID(&dto)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Error get image",
	//	})
	//	log.Println("GetImageById: Error call service in pkg controller", err)
	//	c.Abort()
	//	return
	//}
	//if image == nil {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "no image found",
	//	})
	//	c.Abort()
	//	return
	//}
	//c.JSON(http.StatusOK, image)
}

func (p *ProductImageController) GetImage(c *gin.Context) {
	//id := c.Param(product.ProductId)
	//var dto dto.ProductIdDTO
	//dto.ProductId = id
	//images, err := p.ProductImageService.GetImage(&dto)
	//if err != nil {
	//	log.Println("GetImages: error in controller package", err)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Fail to get images",
	//	})
	//	c.Abort()
	//	return
	//}
	//if images == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "no image found",
	//	})
	//	c.Abort()
	//	return
	//}
	//c.JSON(http.StatusOK, images)
}

func (p *ProductImageController) CreateImage(ctx *gin.Context) {
	var imageBody *entity.ProductImage
	if err := ctx.ShouldBindJSON(&imageBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("CreateImage: Error to ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	imageBody.ProductId = ctx.Param(product.ProductId)
	errCreate := p.ProductImageService.CreateImage(imageBody, claims.UserId)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreate.Error(),
		})
		log.Println("CreateImage: Error to CreateImage in package controller", errCreate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "image created",
	})
}

func (p *ProductImageController) DeleteImage(ctx *gin.Context) {
	var imageBody entity.ProductImage

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, config.CookieAuth)
	if errGetToken != nil {
		log.Println("Error when get token in controller: ", errGetToken)
		ctx.Abort()
		return
	}

	claims, errExtract := token.ExtractToken(tokenFromCookie)
	if errExtract != nil || len(tokenFromCookie) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	imageId, errGetId := strconv.Atoi(ctx.Query(product.Id))
	if errGetId != nil {
		log.Println("error in get image by imageId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	imageBody.ProductId = ctx.Param(product.ProductId)
	imageBody.ID = uint(imageId)
	errDelete := p.ProductImageService.DeleteImage(&imageBody, claims.UserId)
	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errDelete.Error(),
		})
		log.Println("Delete Image: Error to delete image in package controller", errDelete)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "image deleted",
	})
}
