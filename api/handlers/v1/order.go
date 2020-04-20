package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"context"
	"fmt"
	pbo "genproto/order_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

// @Router /v1/order/ [post]
// @Summary Create Order
// @Description API for creating order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.CreateOrder true "order"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order pbo.Order
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}
	fromLocation := models.Location{
		Long:float64(order.FromLocation.Long),
		Lat:float64(order.FromLocation.Lat),
	}
	toLocation := models.Location{
		Long:float64(order.ToLocation.Long),
		Lat:float64(order.ToLocation.Lat),
	}
	deliveryTotalPrice, err := calcDeliveryPriceWithFare(c, h, order.FareId, fromLocation, toLocation)

	if handleGrpcErrWithMessage(c, h.log, err, "error while calculating delivery price by fare") {
		return
	}

	order.DeliverTotalPrice = float32(deliveryTotalPrice)

	_, err = h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "order successfully created",
	})
}

// @Router /v1/order/{order_id} [get]
// @Summary Get Order
// @Description API for getting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.GetOrder
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrder(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		orderID string
	)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	orderID = c.Param("order_id")

	order, err := h.grpcClient.OrderService().Get(context.Background(), &pbo.GetRequest{
		Id: orderID,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/order [get]
// @Summary Get Orders
// @Description API for getting orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetOrders
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
	)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err := ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	order, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.GetAllRequest{
		Page:page,
		Limit:limit,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/order/{order_id}/change-status [patch]
// @Summary Change Order Status
// @Description API for changing order status
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "ORDER ID"
// @Param status body models.ChangeStatusRequest true "status"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h handlerV1) ChangeOrderStatus(c *gin.Context) {
	var (
		orderID string
		changeStatusModel models.ChangeStatusRequest
	)
	orderID = c.Param("order_id")

	err := c.ShouldBindJSON(&changeStatusModel)

	if handleBadRequestErrWithMessage(c, h.log, err,"error while binding to json") {
		return
	}
	fmt.Println(changeStatusModel)

	_, err = h.grpcClient.OrderService().ChangeStatus(
		context.Background(),
		&pbo.ChangeStatusRequest{
			Id: orderID,
			StatusId: changeStatusModel.StatusID,
		})
	fmt.Println(err)

	if handleGrpcErrWithMessage(c, h.log, err, "error while changing order status") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "changing order status successfully",
	})
}

// @Router /v1/order-statuses [get]
// @Summary Get All Possible Order Statuses
// @Description API for getting order statuses
// @Tags order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetStatuses
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetStatuses(c *gin.Context) {
	var (
		model models.GetStatuses
		status models.Status
	)

	status = models.Status{"986a0d09-7b4d-4ca9-8567-aa1c6d770505", "New"}
	model.Statuses = append(model.Statuses, status)

	status = models.Status{"6ba783a3-1c2e-479c-9626-25526b3d9d36", "Cancelled"}
	model.Statuses = append(model.Statuses, status)

	status = models.Status{"8781af8e-f74d-4fb6-ae23-fd997f4a2ee0", "Accepted"}
	model.Statuses = append(model.Statuses, status)

	status = models.Status{"84be5a2f-3a92-4469-8283-220ca34a0de4", "Picked up"}
	model.Statuses = append(model.Statuses, status)

	status = models.Status{"79413606-a56f-45ed-97c3-f3f18e645972", "Delivered"}
	model.Statuses = append(model.Statuses, status)

	status = models.Status{"e665273d-5415-4243-a329-aee410e39465", "Finished"}
	model.Statuses = append(model.Statuses, status)

	c.JSON(http.StatusOK, model)
}