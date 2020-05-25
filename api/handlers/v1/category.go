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

func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		category pb.Category
	)
	err := unmarshal.Unmarshal(c.Request.Body, &category)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while parsing json to proto",
			},
		})
		h.log.Error("error while parsing json to proto", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.CategortService().Create(
		context.Background(),
		&category,
		)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating category") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *handlerV1) GetAllCategory(c *gin.Context) {
	var (
		marshaler jsonpb.Marshaler
		model models.GetAllCategory
	)
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

	resp, err := h.grpcClient.CategortService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page: int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all categories") {
		return
	}

	js, _ := marshaler.MarshalToString(resp)
	fmt.Println(js)

	err = json.Unmarshal([]byte(js), &model)
	//model.Count = resp.Count

	c.JSON(http.StatusOK, model)
}
