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
	optionId = "optionId"
)

type ProductOptionController interface {
	CreateOption(c *gin.Context)
	GetOptions(c *gin.Context)
	GetOptionByID(c *gin.Context)
	DeleteOption(c *gin.Context)
	UpdateOption(c *gin.Context)
}

func (p productOptionController) UpdateOption(c *gin.Context) {
	var optionUpdateBody *entity.ProductOption
	if err := c.ShouldBindJSON(&optionUpdateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	oid, errCv := strconv.Atoi(c.Param(optionId))
	if errCv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error update option",
		})
		log.Println("UpdateOption: Error parse param", errCv)
		c.Abort()
		return
	}
	dtoUpdate := dto.NewUpdateOptionDTO(optionUpdateBody)
	dtoUpdate.OptionId = optionId
	dtoUpdate.Option.ID = uint(oid)
	product, err := p.productOptionService.UpdateOption(dtoUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to update product",
		})
		log.Println("UpdateProduct: Error Update in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, product)
}
func (p productOptionController) CreateOption(c *gin.Context) {
	var optionBody *entity.ProductOption
	if err := c.ShouldBindJSON(&optionBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to ShouldBindJSON in package controller", err)
		c.Abort()
		return
	}
	dtoOption := dto.NewCreateOptionDTO(optionBody)
	dtoOption.Option.ProductId = c.Param(productId)
	option, err := p.productOptionService.CreateOption(dtoOption)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to create option",
		})
		log.Println("CreateOption: Error to CreateOption in package controller", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, option)
}
func (p productOptionController) GetOptions(c *gin.Context) {
	id := c.Param(productId)
	var dto dto.ProductIdDTO
	dto.ProductId = id
	options, err := p.productOptionService.GetOptions(&dto)
	if err != nil {
		log.Println("GetOptions: error in controller package", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Fail to get options",
		})
		c.Abort()
		return
	}
	if options == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found options",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, options)
}

func (p productOptionController) GetOptionByID(c *gin.Context) {
	var dto dto.OptionIdDTO
	dto.OptionId = c.Param(optionId)
	c.Set(optionId, dto.OptionId)
	option, err := p.productOptionService.GetOptionByID(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error get option",
		})
		log.Println("GetOptionById: Error call service in pkg controller", err)
		c.Abort()
		return
	}
	if option == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Option not found",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, option)
}

func (p productOptionController) DeleteOption(c *gin.Context) {
	oId := c.Param(optionId)
	var dto dto.OptionIdDTO
	dto.OptionId = oId
	option, err := p.productOptionService.DeleteOption(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error to delete option",
		})
		log.Println("DeleteOption: Error to parse oId", err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, option)
}

type productOptionController struct {
	productOptionService service.ProductOptionService
}

func NewProductOptionController(productOptionService service.ProductOptionService) *productOptionController {
	return &productOptionController{productOptionService: productOptionService}
}
