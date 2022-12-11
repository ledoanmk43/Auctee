package controller

import (
	"backend/pkg/token"
	account "backend/src/account-service/config"
	product "backend/src/product-service/config"
	"backend/src/product-service/entity"
	"backend/src/product-service/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IProductOptionController interface {
	CreateOption(ctx *gin.Context)
	GetOptions(ctx *gin.Context)
	GetOptionByID(ctx *gin.Context)
	DeleteOption(ctx *gin.Context)
	UpdateOption(ctx *gin.Context)
}

type ProductOptionController struct {
	ProductOptionService service.IProductOptionService
}

func NewProductOptionController(productOptionService service.IProductOptionService) *ProductOptionController {
	return &ProductOptionController{ProductOptionService: productOptionService}
}

func (p *ProductOptionController) UpdateOption(ctx *gin.Context) {
	var optionBody entity.ProductOption
	if err := ctx.ShouldBindJSON(&optionBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("Create option: Error to ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}
	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(account.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
		})
		ctx.Abort()
		return
	}
	claims, errExtract := token.ExtractToken(tokenFromCookie.(string))
	if errExtract != nil || len(tokenFromCookie.(string)) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	optionId, errGetId := strconv.Atoi(ctx.Query(product.Id))
	if errGetId != nil {
		log.Println("error in get image by optionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	optionBody.ProductId = ctx.Param(product.ProductId)
	optionBody.ID = uint(optionId)
	errDelete := p.ProductOptionService.UpdateOption(&optionBody, claims.UserId)
	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errDelete.Error(),
		})
		log.Println("Create option: Error to Create option in package controller", errDelete)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "option updated",
	})
}
func (p *ProductOptionController) CreateOption(ctx *gin.Context) {
	var optionBody *entity.ProductOption
	if err := ctx.ShouldBindJSON(&optionBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error Binding JSON",
		})
		log.Println("Create option: Error to ShouldBindJSON in package controller", err)
		ctx.Abort()
		return
	}

	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(account.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
		})
		ctx.Abort()
		return
	}
	claims, errExtract := token.ExtractToken(tokenFromCookie.(string))
	if errExtract != nil || len(tokenFromCookie.(string)) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	optionBody.ProductId = ctx.Param(product.ProductId)
	errCreate := p.ProductOptionService.CreateOption(optionBody, claims.UserId)
	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errCreate.Error(),
		})
		log.Println("Create option: Error to Create option in package controller", errCreate)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "option created",
	})
}
func (p *ProductOptionController) GetOptions(c *gin.Context) {
	//id := c.Param(product.ProductId)
	//var dto dto.ProductIdDTO
	//dto.ProductId = id
	//options, err := p.productOptionService.GetOptions(&dto)
	//if err != nil {
	//	log.Println("GetOptions: error in controller package", err)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Fail to get options",
	//	})
	//	c.Abort()
	//	return
	//}
	//if options == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Not found options",
	//	})
	//	c.Abort()
	//	return
	//}
	//c.JSON(http.StatusOK, options)
}

func (p *ProductOptionController) GetOptionByID(c *gin.Context) {
	//var dto dto.OptionIdDTO
	//dto.OptionId = c.Param(optionId)
	//c.Set(optionId, dto.OptionId)
	//option, err := p.productOptionService.GetOptionByID(&dto)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Error get option",
	//	})
	//	log.Println("GetOptionById: Error call service in pkg controller", err)
	//	c.Abort()
	//	return
	//}
	//if option == nil {
	//	c.JSON(http.StatusNotFound, gin.H{
	//		"message": "Option not found",
	//	})
	//	c.Abort()
	//	return
	//}
	//c.JSON(http.StatusOK, option)
}

func (p *ProductOptionController) DeleteOption(ctx *gin.Context) {
	var optionBody entity.ProductOption

	authSession := sessions.Default(ctx)
	tokenFromCookie := authSession.Get(account.CookieAuth)
	if tokenFromCookie == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "no cookie",
		})
		ctx.Abort()
		return
	}
	claims, errExtract := token.ExtractToken(tokenFromCookie.(string))
	if errExtract != nil || len(tokenFromCookie.(string)) == 0 {
		log.Println("Error: Error when extracting token in controller: ", errExtract)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	optionId, errGetId := strconv.Atoi(ctx.Query(product.Id))
	if errGetId != nil {
		log.Println("error in get option by optionId: ", errGetId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error when get id in url",
		})
		ctx.Abort()
		return
	}

	optionBody.ProductId = ctx.Param(product.ProductId)
	optionBody.ID = uint(optionId)
	errDelete := p.ProductOptionService.DeleteOption(&optionBody, claims.UserId)
	if errDelete != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errDelete.Error(),
		})
		log.Println("Delete Option: Error to delete option in package controller", errDelete)
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "option deleted",
	})
}
