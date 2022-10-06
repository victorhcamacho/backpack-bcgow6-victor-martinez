package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/cmd/server/handler"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/internal/products"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/pkg/store"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("No se logro cargar el archivo .env")
	}

	db := store.New(store.FileType, os.Getenv("PATH_JSON"))

	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	ph := handler.NewProduct(serv)

	server := gin.Default()

	p := server.Group("/products")
	p.POST("/", ph.Save())
	p.GET("/", ph.GetAll())
	p.GET("/:id", ph.GetByID())
	p.PUT("/:id", ph.Update())
	p.PATCH("/:id", ph.UpdatePatch())
	p.DELETE("/:id", ph.Delete())

	server.Run(":9000")
}
