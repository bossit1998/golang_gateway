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

// @Security ApiKeyAuth
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
		userInfo  models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	err = unmarshal.Unmarshal(c.Request.Body, &category)

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
	category.ShipperId = userInfo.ShipperID

	resp, err := h.grpcClient.CategoryService().Create(
		context.Background(),
		&category,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating category") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// @Router /v1/category [get]
// @Summary Get All Category
// @Description API for getting all category
// @Tags category
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param parent_id query integer false "parent_id"
// @Success 200 {object} models.GetAllCategoriesModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCategory(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model      models.GetAllCategoriesModel
		userInfo   models.UserInfo
		shipperId  string
	)

	if c.GetHeader("Authorization") == "" && c.GetHeader("Shipper") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: Authorization or shipper id have to be on the header")
		return
	}

	if c.GetHeader("Authorization") != "" {
		err := getUserInfo(h, c, &userInfo)

		if err != nil {
			return
		}
		shipperId = userInfo.ShipperID
	} else if c.GetHeader("Shipper") != "" {
		shipperId = c.GetHeader("Shipper")
	}

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
			ShipperId: shipperId,
			Page:      int64(page),
			ParentId:  c.Param("parent_id"),
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

// @Security ApiKeyAuth
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
		userInfo  models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	categoryID := c.Param("category_id")

	err = unmarshal.Unmarshal(c.Request.Body, &category)

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
	category.ShipperId = userInfo.ShipperID

	_, err = h.grpcClient.CategoryService().Update(
		context.Background(),
		&category,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while updating category") {
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{Message: "category updated successfully"})
}

// @Security ApiKeyAuth
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
	var (
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	categoryID := c.Param("category_id")

	_, err = h.grpcClient.CategoryService().Delete(
		context.Background(),
		&pb.DeleteRequest{
			ShipperId: userInfo.ShipperID,
			Id:        categoryID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while deleting category") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "category deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/category/{category_id} [get]
// @Summary Get Category
// @Description API for getting a category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.GetCategoryModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCategory(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		userInfo   models.UserInfo
		shipperId  string
	)

	if c.GetHeader("Authorization") == "" && c.GetHeader("Shipper") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: Authorization or shipper id have to be on the header")
		return
	}

	if c.GetHeader("Authorization") != "" {
		err := getUserInfo(h, c, &userInfo)

		if err != nil {
			return
		}
		shipperId = userInfo.ShipperID
	} else if c.GetHeader("Shipper") != "" {
		shipperId = c.GetHeader("Shipper")
	}

	marshaller.OrigName = true
	marshaller.EmitDefaults = true

	resp, err := h.grpcClient.CategoryService().Get(
		context.Background(),
		&pb.GetRequest{
			ShipperId: shipperId,
			Id:        c.Param("category_id"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting a category") {
		return
	}

	js, err := marshaller.MarshalToString(resp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "err while marshaling",
			},
		})
		h.log.Error("error while parsing proto to struct", logger.Error(err))
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
