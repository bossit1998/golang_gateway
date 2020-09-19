package v1

import (
	"context"
	"encoding/json"
	pb "genproto/catalog_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Router /v1/measure [post]
// @Summary Create Measure
// @Description API for creating measure
// @Tags measure
// @Accept  json
// @Produce  json
// @Param measure body models.CreateMeasureModel true "measure"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateMeasure(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		measure   pb.Measure
	)
	err := unmarshal.Unmarshal(c.Request.Body, &measure)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while creating measure",
			},
		})
		h.log.Error("error while parsing json to proto", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.MeasureService().Create(
		context.Background(),
		&measure,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating measure") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/measure [get]
// @Summary Get All Measure
// @Description API for getting all measure
// @Tags measure
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Success 200 {object} models.GetAllMeasuresModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllMeasure(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model      models.GetAllMeasuresModel
	)
	marshaller.OrigName = true

	page, err := ParsePageQueryParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing page",
			},
		})
		h.log.Error("error while parsing page", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.MeasureService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page: int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all measures") {
		return
	}

	js, _ := marshaller.MarshalToString(resp)

	err = json.Unmarshal([]byte(js), &model)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "error while parsing proto to struct",
			},
		})
		h.log.Error("error while parsing proto to struct", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, model)
}
