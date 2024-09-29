package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// conexao com o banco
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8000")
}
