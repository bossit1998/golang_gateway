package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"context"
	"encoding/json"
	pb "genproto/catalog_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

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

	c.JSON(http.StatusOK, model)
}