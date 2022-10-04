package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC2/practica_tt/cmd/server/handler"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC2/practica_tt/internal/products"
)

func main() {

	repo := products.NewRepository()
	serv := products.NewService(repo)
	ph := handler.NewProduct(serv)

	server := gin.Default()

	p := server.Group("/products")
	p.POST("/", ph.Save())
	p.GET("/", ph.GetAll())
	p.GET("/:id", ph.GetByID())

	server.Run(":9000")

}
