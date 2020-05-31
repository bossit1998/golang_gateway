package v1

import (
	"context"
	"encoding/json"
	"fmt"
	pb "genproto/catalog_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Router /v1/product [post]
// @Summary Create Product
// @Description API for creating product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.CreateProductModel true "product"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		product pb.Product
	)
	err := unmarshal.Unmarshal(c.Request.Body, &product)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while parsing json to proto",
			},
		})
		h.log.Error("error while parsing json to proto", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.ProductService().Create(
		context.Background(),
		&product,
		)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating product") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/product [get]
// @Summary Get All Product
// @Description API for getting all product
// @Tags product
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Success 200 {object} models.GetAllProductsModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model models.GetAllProductsModel
	)
	marshaller.OrigName = true

	page, err := ParsePageQueryParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while parsing page",
			},
		})
		h.log.Error("error while parsing page", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.ProductService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page:int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all products") {
		return
	}

	js, _ := marshaller.MarshalToString(resp)

	err = json.Unmarshal([]byte(js), &model)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError {
				Code:ErrorCodeInternal,
				Message:"error while parsing proto to struct",
			},
		})
		h.log.Error("error while parsing proto to struct", logger.Error(err))
		return
	}

	for i, p := range model.Products {
		model.Products[i].Image = fmt.Sprintf("http://api.delever.uz:9001/delever/%s", p.Image)
	}

	c.JSON(http.StatusOK, model)
}