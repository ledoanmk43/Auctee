package controller

import (
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	productId = "productId"
)

type ProductController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
type productController struct {
	productService service.ProductService
}

func (p *productController) All(context *gin.Context) {
	products, err := p.productService.All()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to get all products",
		})
		log.Println("GetProducts: Error get all product in package controller", err)
		context.Abort()
	}
	context.JSON(http.StatusOK, products)

}

func (p *productController) FindByID(context *gin.Context) {
	var dto dto.ProductDTO
	dto.ProductId = context.Param(productId)
	context.Set(productId, dto.ProductId)
	product, err := p.productService.FindProductByID(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to get Product ",
		})
		log.Println("GetProductById: Error in package controller", err)
		context.Abort()
		return
	}
	if product == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"Message": "Not found product",
		})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, product)

}

func (p productController) Insert(context *gin.Context) {
	var productBody *entity.Product
	if err := context.ShouldBindJSON(&productBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		context.Abort()
		return
	}
	dto := dto.NewProductCreatedDTO(productBody)
	product, err := p.productService.Insert(dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
		log.Println("Error to create product  controller", err)
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, product)
}

func (p productController) Update(context *gin.Context) {
	productId := context.Param(productId)
	var productUpdateBody *entity.Product
	if err := context.ShouldBindJSON(&productUpdateBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		context.Abort()
		return
	}
	dtoUpdate := dto.NewProductUpdateDTO(productUpdateBody)
	dtoUpdate.ProductId = productId
	dtoUpdate.Product.Id = productId
	product, err := p.productService.Update(dtoUpdate)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, product)

}

func (p productController) Delete(context *gin.Context) {
	var dto dto.ProductDTO
	fmt.Println(dto.ProductId)
	dto.ProductId = context.Param(productId)
	fmt.Println(dto.ProductId)
	product, err := p.productService.Delete(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete product",
		})
		log.Println("DeleteProduct: Error to get id product in package controller", err)
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, product)

}

func NewProductController(productServ service.ProductService) ProductController {
	return &productController{
		productService: productServ,
	}
}
