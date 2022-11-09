package controller

import (
	account "backend/internal/account-service/config"
	product "backend/internal/product-service/config"
	"backend/internal/product-service/entity"
	"backend/internal/product-service/service"
	"backend/pkg/token"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IProductController interface {
	CreateProduct(c *gin.Context)
	UpdateProductByProductId(c *gin.Context)
	GetProductByProductId(c *gin.Context)
	GetAllProducts(c *gin.Context)
	DeleteProductByProductId(c *gin.Context)
}
type ProductController struct {
	ProductService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (p *ProductController) GetAllProducts(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	products, err := p.ProductService.GetAllProducts(claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "fail to get all products",
		})
		log.Println("GetProducts: Error get all products in package controller", err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) GetProductByProductId(ctx *gin.Context) {
	productId := ctx.Query(product.Id)
	if len(productId) == 0 {
		log.Println("error in get product by productId: nil productId")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	productDetail, err := p.ProductService.GetProductDetailByProductId(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "product not found",
		})
		log.Println("GetProductById: Error in package controller", err)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, productDetail)

}

func (p *ProductController) CreateProduct(ctx *gin.Context) {
	var productBody *entity.Product
	if err := ctx.ShouldBindJSON(&productBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding JSON",
		})
		log.Println("Error to ShouldBindJSON controller", err)
		ctx.Abort()
		return
	}

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	productBody.UserId = claims.UserId
	log.Println("nenenene: ", productBody)
	errCreate := p.ProductService.Insert(productBody)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreate.Error(),
		})
		log.Println("Error to create product  controller", errCreate)
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "product created",
	})
}

func (p *ProductController) UpdateProductByProductId(ctx *gin.Context) {
	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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

	productId := ctx.Query(product.Id)
	if len(productId) == 0 {
		log.Println("error in update product by productId: nil productId")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	var updateBody *entity.Product
	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}
	updateBody.Id = productId
	updateBody.UserId = claims.UserId
	err := p.ProductService.Update(updateBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product updated",
	})
}

func (p *ProductController) DeleteProductByProductId(ctx *gin.Context) {
	productId := ctx.Query(product.Id)
	if len(productId) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "missing param",
		})
		ctx.Abort()
		return
	}

	tokenFromCookie, errGetToken := utils.GetTokenFromCookie(ctx, account.CookieAuth)
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
	err := p.ProductService.Delete(productId, claims.UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Println("DeleteProduct: Error to get id product in package controller", err)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted",
	})
}
