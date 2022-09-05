package controller

import (
	"chilindo/pkg/utils"
	"chilindo/src/product-service/dto"
	"chilindo/src/product-service/entity"
	"chilindo/src/product-service/service"
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

func (p *productController) Insert(ctx *gin.Context) {
	var productBody *entity.Product
	if err := ctx.ShouldBindJSON(&productBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create product",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		ctx.Abort()
		return
	}

	adminId := utils.GetIdFromToken(ctx)

	productBody.AdminId = adminId
	//dto := dto.NewProductCreatedDTO(productBody)
	createdProduct, err := p.productService.Insert(productBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		log.Println("Error to create product  controller", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, createdProduct)
}

func (p *productController) Update(context *gin.Context) {
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

func (p *productController) Delete(ctx *gin.Context) {

	ProductId := ctx.Param(productId)

	adminId := utils.GetIdFromToken(ctx)
	_, err := p.productService.Delete(ProductId, adminId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		log.Println("DeleteProduct: Error to get id product in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "product deleted",
	})
}

func NewProductController(productServ service.ProductService) ProductController {
	return &productController{
		productService: productServ,
	}
}
