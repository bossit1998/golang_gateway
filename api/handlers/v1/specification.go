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

func (h *handlerV1) CreateSpecification (c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		specification pb.Specification
	)
	err := unmarshal.Unmarshal(c.Request.Body, &specification)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.SpecificationService().Create(context.Background(), &specification)

	fmt.Println(err)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating specification") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *handlerV1) GetAllSpecification(c *gin.Context) {
	page, err := ParsePageQueryParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while parsing page",
			},
		})
		return
	}

	resp, err := h.grpcClient.SpecificationService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page: int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all specifications") {
		return
	}

	c.JSON(http.StatusOK, resp)
}
