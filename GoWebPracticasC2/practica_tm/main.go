package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type product struct {
	ID         int
	NAME       string  `validate:"required"`
	UNIT_PRICE float64 `validate:"required,min=1"`
	STOCK      int     `validate:"required,min=1"`
}

var products []product

// var validate *validator.Validate

const TOKEN string = "QWERTY12345"

func main() {

	server := gin.Default()

	rp := server.Group("/products")
	{
		rp.POST("/", StoreProduct())
		rp.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, products)
		})
	}

	server.Run(":9000")
}

func StoreProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(400, gin.H{
				"headersError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != TOKEN {
			ctx.JSON(401, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		var request product

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{
				"parseError": err.Error(),
			})
			return
		}

		validate := validator.New()

		errValidate := validate.Struct(&request)

		if errValidate != nil {

			var strError string
			var validationErrors []string

			for _, verr := range errValidate.(validator.ValidationErrors) {

				tag := verr.ActualTag()
				param := verr.StructField()

				switch tag {
				case "required":
					strError = fmt.Sprintf("El parametro %s es requerido", param)
				case "min":
					strError = fmt.Sprintf("El valor del parametro %s debe ser mayor a 0", param)
				}

				validationErrors = append(validationErrors, strError)
			}

			ctx.JSON(400, gin.H{
				"validationErrors": validationErrors,
			})

			return
		}

		request.ID = len(products) + 1
		products = append(products, request)

		ctx.JSON(http.StatusOK, request)
	}
}
