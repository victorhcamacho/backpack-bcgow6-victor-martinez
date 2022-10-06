package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/cmd/server/handler"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/docs"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/internal/products"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/pkg/store"
)

// @title BootcampGoWeb - Products API
// @version 1.0
// @description This API Handle Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	if err := godotenv.Load(); err != nil {
		panic("No se logro cargar el archivo .env")
	}

	host := os.Getenv("HOST")
	basePath := os.Getenv("BASE_PATH")

	db := store.New(store.FileType, os.Getenv("PATH_JSON"))

	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	ph := handler.NewProduct(serv)

	server := gin.Default()

	api := server.Group(basePath)

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = basePath

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := api.Group("/products")
	{
		p.POST("/", ph.Save())
		p.GET("/", ph.GetAll())
		p.GET("/:id", ph.GetByID())
		p.PUT("/:id", ph.Update())
		p.PATCH("/:id", ph.UpdatePatch())
		p.DELETE("/:id", ph.Delete())
	}

	server.Run(":9000")
}
