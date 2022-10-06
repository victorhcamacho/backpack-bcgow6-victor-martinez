package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoWebPracticasC3/internal/products"
)

type requestDTO struct {
	ID            int     `json:"id"`
	Name          string  `json:"nombre" validate:"required"`
	UnitPrice     float64 `json:"precioUnitario" validate:"required,min=1"`
	StockQuantity int     `json:"cantidadExistencias" validate:"required,min=1"`
}

type productHandler struct {
	service products.Service
}

func NewProduct(s products.Service) *productHandler {
	return &productHandler{service: s}
}

func (ph *productHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		products, err := ph.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}

func (ph *productHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fue posible obtener el id del producto"})
			return
		}

		product, err := ph.service.GetByID(idProduct)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (ph *productHandler) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"parseError": err.Error()})
			return
		}

		errorList := validateRequest(&request)

		if len(errorList) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"validationErrors": errorList})
			return
		}

		product, err := ph.service.Save(request.Name, request.UnitPrice, request.StockQuantity)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"saveError": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, product)

	}
}

func (ph *productHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fue posible obtener el id del producto"})
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"parseError": err.Error()})
			return
		}

		errorList := validateRequest(&request)

		if len(errorList) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"validationErrors": errorList})
			return
		}

		product, err := ph.service.Update(idProduct, request.Name, request.UnitPrice, request.StockQuantity)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (ph *productHandler) UpdatePatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fue posible obtener el id del producto"})
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"parseError": err.Error()})
			return
		}

		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre del producto es requerido"})
			return
		} else if request.UnitPrice <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El precio unitario del producto es requerido"})
			return
		}

		product, err := ph.service.UpdateNameAndPrice(idProduct, request.Name, request.UnitPrice)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, product)
	}
}

func (ph *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"authError": "Es necesario enviar el cabecero 'token' para verificar la petición",
			})
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"authError": "No cuenta con los permisos necesarios para procesar la petición solicitada",
			})
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fue posible obtener el id del producto"})
			return
		}

		errDel := ph.service.Delete(idProduct)

		if errDel != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": errDel.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("El producto %d ha sido eliminado", idProduct)})
	}
}

func validateRequest(req *requestDTO) (errors []string) {

	validate := validator.New()

	err := validate.Struct(req)

	if err != nil {

		var strError string

		for _, verr := range err.(validator.ValidationErrors) {

			tag := verr.ActualTag()
			param := verr.StructField()

			switch tag {
			case "required":
				strError = fmt.Sprintf("El parametro %s es requerido", param)
			case "min":
				strError = fmt.Sprintf("El valor del parametro %s debe ser mayor a 0", param)
			}

			errors = append(errors, strError)
		}
	}

	return
}
