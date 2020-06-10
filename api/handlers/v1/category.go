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

// @Router /v1/category [post]
// @Summary Create Category
// @Description API for creating category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category body models.CreateCategoryModel true "category"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		category  pb.Category
	)
	err := unmarshal.Unmarshal(c.Request.Body, &category)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing json to proto",
			},
		})
		h.log.Error("error while parsing json to proto", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.CategoryService().Create(
		context.Background(),
		&category,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating category") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Router /v1/category [get]
// @Summary Get All Category
// @Description API for getting all category
// @Tags category
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Success 200 {object} models.GetAllCategoriesModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCategory(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model models.GetAllCategoriesModel
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

	resp, err := h.grpcClient.CategoryService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			Page: int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all categories") {
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

// @Router /v1/category/{category_id} [put]
// @Summary Update Category
// @Description API for updating category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category_id path string true "category_id"
// @Param category body models.CreateCategoryModel true "category"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		category  pb.Category
	)
	categoryID := c.Param("category_id")

	err := unmarshal.Unmarshal(c.Request.Body, &category)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing json to proto",
			},
		})
		h.log.Error("error while parsing json to proto", logger.Error(err))
		return
	}
	category.Id = categoryID

	_, err = h.grpcClient.CategoryService().Update(
		context.Background(),
		&category,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while updating category") {
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{Message:"category updated successfully"})
}

// @Router /v1/category/{category_id} [delete]
// @Summary Delete Category
// @Description API for deleting category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("category_id")

	_, err := h.grpcClient.CategoryService().Delete(
		context.Background(),
		&pb.DeleteRequest{
			Id: categoryID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while deleting category") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "category deleted successfully",
	})
}