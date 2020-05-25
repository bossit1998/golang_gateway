package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	pb "genproto/catalog_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

// @Router /v1/product-kind [post]
// @Summary Create Product Kind
// @Description API for creating product kind
// @Tags product-kind
// @Accept  json
// @Produce  json
// @Param product_kind body models.CreateProductKindModel true "product_kind"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateProductKind(c *gin.Context) {
	var (
		unmarshal   jsonpb.Unmarshaler
		productKind pb.ProductKind
	)

	err := unmarshal.Unmarshal(c.Request.Body, &productKind)

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

	resp, err := h.grpcClient.ProductKindService().Create(context.Background(), &productKind)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating product kind") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/product-kind [get]
// @Summary Get All Product Kind
// @Description API for getting all product kind
// @Tags product-kind
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Success 200 {object} models.GetAllProductKindsModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllProductKind(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model models.GetAllProductKindsModel
	)
	marshaller.OrigName = true
	marshaller.EmitDefaults = true
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

	resp, err := h.grpcClient.ProductKindService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page:int64(page),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all product kind") {
		return
	}

	js, _ := marshaller.MarshalToString(resp)
	fmt.Println(js)

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