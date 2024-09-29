package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		reponse := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		reponse := model.Response{
			Message: "Id do produto deve ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, reponse)
		return

	}

	product, err := p.productUseCase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		reponse := model.Response{
			Message: "Produto nao foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, reponse)
		return

	}

	ctx.JSON(http.StatusOK, product)
}
