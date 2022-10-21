package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/products"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/pkg/web"
)

type requestDTO struct {
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

// List All Products
// @Summary List products
// @Tags Products
// @Description get all products
// @Produce  json
// @Param token header string true "Access Token"
// @Success  200    {object}  web.Response  "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/ [GET]
func (ph *productHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		products, err := ph.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, products, ""))
	}
}

// List One Product
// @Summary List a product
// @Tags Products
// @Description get only one product
// @Produce  json
// @Param token header string true "Access Token"
// @Param    id       path      int             true   "ID Product"
// @Success  200    {object}  web.Response  "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/{id} [GET]
func (ph *productHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "No fue posible obtener el id del producto"))
			return
		}

		product, err := ph.service.GetByID(idProduct)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}

// Store Product
// @Summary Store a product
// @Tags Products
// @Description create a new product
// @Accept   json
// @Produce  json
// @Param token header string true "Access Token"
// @Param    product  body    requestDTO	true	"Product to Store"
// @Success  201    {object}  web.Response  "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/ [POST]
func (ph *productHandler) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		errorParam := validateRequest(request)

		if errorParam != "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, errorParam))
			return
		}

		product, err := ph.service.Save(request.Name, request.UnitPrice, request.StockQuantity)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, web.NewResponse(201, product, ""))

	}
}

// Update Product
// @Summary Update a product
// @Tags Products
// @Description update an existing product
// @Accept   json
// @Produce  json
// @Param token header string true "Access Token"
// @Param    id       path      int         true   "ID Product"
// @Param    product  body      requestDTO  true  "Product to Update"
// @Success  200    {object}  web.Response  "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/{id} [PUT]
func (ph *productHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "No fue posible obtener el id del producto"),
			)
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		errorParam := validateRequest(request)

		if errorParam != "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, errorParam))
			return
		}

		product, err := ph.service.Update(idProduct, request.Name, request.UnitPrice, request.StockQuantity)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}

// Update Product Name and Unit Price
// @Summary Update product name and unit price
// @Tags Products
// @Description update an existing product name and unit price
// @Accept   json
// @Produce  json
// @Param token header string true "Access Token"
// @Param    id       path      int         true   "ID Product"
// @Param    product  body      requestDTO  true  "Params to Update"
// @Success  200    {object}  web.Response  "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/{id} [PATCH]
func (ph *productHandler) UpdatePatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "No fue posible obtener el id del producto"),
			)
			return
		}

		var request requestDTO

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			return
		}

		if request.Name == "" {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "El nombre del producto es requerido"),
			)
			return
		} else if request.UnitPrice <= 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "El precio unitario del producto es requerido"),
			)
			return
		}

		product, err := ph.service.UpdateNameAndPrice(idProduct, request.Name, request.UnitPrice)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, product, ""))
	}
}

// Delete One Product
// @Summary Delete a product
// @Tags Products
// @Description delete an existing product
// @Produce  json
// @Param token header string true "Access Token"
// @Param id	path   int    true "ID Product"
// @Success  204    "Success"
// @Failure  400    {object}  web.Response  "Bad Request"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  404    {object}  web.Response  "Not Found"
// @Failure  500    {object}  web.Response  "Internal Server Error"
// @Router /products/{id} [DELETE]
func (ph *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenHeader := ctx.GetHeader("token")

		if len(tokenHeader) == 0 {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "Es necesario enviar el cabecero 'token' para verificar la petición"),
			)
			return
		} else if tokenHeader != os.Getenv("TOKEN_AUTH") {
			ctx.JSON(
				http.StatusUnauthorized,
				web.NewResponse(401, nil, "No cuenta con los permisos necesarios para procesar la petición solicitada"),
			)
			return
		}

		idProduct, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				web.NewResponse(400, nil, "No fue posible obtener el id del producto"),
			)
			return
		}

		errDel := ph.service.Delete(idProduct)

		if errDel != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, errDel.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func validateRequest(req requestDTO) (strError string) {

	err := validator.New().Struct(req)

	if err != nil {
		for _, verr := range err.(validator.ValidationErrors) {

			tag := verr.ActualTag()
			param := verr.StructField()

			switch tag {
			case "required":
				strError = fmt.Sprintf("El parametro %s es requerido", param)
				return
			case "min":
				strError = fmt.Sprintf("El valor del parametro %s debe ser mayor a 0", param)
				return
			}
		}
	}

	return
}
