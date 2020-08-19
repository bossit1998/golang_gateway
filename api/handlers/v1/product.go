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

// @Security ApiKeyAuth
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
		product   pb.Product
		userInfo  models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	err = unmarshal.Unmarshal(c.Request.Body, &product)

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
	product.ShipperId = userInfo.ShipperID

	resp, err := h.grpcClient.ProductService().Create(
		context.Background(),
		&product,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating product") {
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// @Security ApiKeyAuth
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
		model      models.GetAllProductsModel
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
	fmt.Println(shipperId)
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

	resp, err := h.grpcClient.ProductService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			ShipperId: shipperId,
			Page:      int64(page),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all products") {
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

	for i, p := range model.Products {
		if p.Image != "" {
			model.Products[i].Image = fmt.Sprintf("https://sdn.delever.uz/delever/%s", p.Image)
		}
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/product/{product_id} [get]
// @Summary Get Product
// @Description API for getting a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.GetProductModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetProduct(c *gin.Context) {
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

	resp, err := h.grpcClient.ProductService().Get(
		context.Background(),
		&pb.GetRequest{
			ShipperId: shipperId,
			Id:        c.Param("product_id"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting a product") {
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

// @Security ApiKeyAuth
// @Router /v1/product/{product_id} [put]
// @Summary Update Product
// @Description API for updating product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Param product body models.CreateProductModel true "product"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		unmarshal jsonpb.Unmarshaler
		product   pb.Product
		userInfo  models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	productID := c.Param("product_id")

	err = unmarshal.Unmarshal(c.Request.Body, &product)

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
	product.Id = productID
	product.ShipperId = userInfo.ShipperID

	_, err = h.grpcClient.ProductService().Update(
		context.Background(),
		&product,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while updating product") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{Message: "product updated successfully"})
}

// @Security ApiKeyAuth
// @Router /v1/product/{product_id} [delete]
// @Summary Delete Product
// @Description API for deleting product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	var (
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	productID := c.Param("product_id")

	_, err = h.grpcClient.ProductService().Delete(
		context.Background(),
		&pb.DeleteRequest{
			ShipperId: userInfo.ShipperID,
			Id:        productID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while deleting product") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "product deleted successfully",
	})
}
