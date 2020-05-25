package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"context"
	"fmt"
	pb "genproto/catalog_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

func (h *handlerV1) CreateMeasure(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		measure pb.Measure
	)
	err := unmarshal.Unmarshal(c.Request.Body, &measure)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while creating measure",
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

func (h *handlerV1) GetAllMeasure(c *gin.Context) {
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

	resp, err := h.grpcClient.MeasureService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page:int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all measures") {
		return
	}
	fmt.Println(resp)

	c.JSON(http.StatusOK, resp)
}
