package v1

import (
	"context"
	"fmt"
	pbo "genproto/order_service"
	"net/http"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
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
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
	)
	userInfo, err := userInfo(h, c)

	if err != nil {
		return
	}
	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}
	order.DeliveryPrice = order.CoDeliveryPrice
	order.CoId = userInfo.ID
	order.UserId = userInfo.ID
	order.CreatorTypeId = userInfo.ID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"

	_, err = h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "order successfully created",
	})
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id} [put]
// @Summary Update Order
// @Description API for updating order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param order body models.UpdateOrder true "order"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
	)
	userInfo, err := userInfo(h, c)
	orderID := c.Param("order_id")

	if err != nil {
		return
	}
	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}
	order.Id = orderID
	order.DeliveryPrice = order.CoDeliveryPrice
	order.CoId = userInfo.ID
	order.UserId = userInfo.ID
	order.CreatorTypeId = userInfo.ID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"

	_, err = h.grpcClient.OrderService().Update(context.Background(), &order)
	fmt.Println(err)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "order successfully updated",
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
		orderID     string
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

	order, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		Page:  page,
		Limit: limit,
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
		orderID           string
		changeStatusModel models.ChangeStatusRequest
	)
	orderID = c.Param("order_id")

	err := c.ShouldBindJSON(&changeStatusModel)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}
	fmt.Println(changeStatusModel)

	_, err = h.grpcClient.OrderService().ChangeStatus(
		context.Background(),
		&pbo.ChangeStatusRequest{
			Id:       orderID,
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
		model  models.GetStatuses
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

// @Router /v1/order/{order_id}/add-courier [patch]
// @Summary Add Order Courier
// @Description API for adding order courier
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "ORDER ID"
// @Param courier body models.AddCourierRequest true "courier"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) AddCourier(c *gin.Context) {
	var (
		orderID         string
		addCourierModel models.AddCourierRequest
	)
	orderID = c.Param("order_id")
	err := c.ShouldBindJSON(&addCourierModel)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	_, err = h.grpcClient.OrderService().AddCourier(
		context.Background(),
		&pbo.AddCourierRequest{
			OrderId:   orderID,
			CourierId: addCourierModel.CourierID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while adding order courier") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "courier added successfully",
	})
}

// @Router /v1/order/{order_id}/remove-courier [patch]
// @Summary Remove Order Courier
// @Description API for changing order courier
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "ORDER ID"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) RemoveCourier(c *gin.Context) {
	var (
		orderID string
	)
	orderID = c.Param("order_id")

	_, err := h.grpcClient.OrderService().RemoveCourier(
		context.Background(),
		&pbo.RemoveCourierRequest{
			OrderId: orderID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while removing order courier") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "courier removed successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/courier/order [get]
// @Summary Get Courier Orders
// @Description API for getting courier orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param courier_id query string false "courier_id"
// @Success 200 {object} models.GetOrders
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourierOrders(c *gin.Context) {
	var (
		courierID string
	)
	userInfo, err := userInfo(h, c)

	if err != nil {
		return
	}

	if userInfo.Role == config.RoleCourier {
		courierID = userInfo.ID
	} else {
		courierID = c.Query("courier_id")

		_, err := uuid.Parse(courierID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "courier id is not valid",
			})
			return
		}
	}

	// page, err := ParsePageQueryParam(c)

	// if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
	// 	return
	// }

	// limit, err := ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	orders, err := h.grpcClient.OrderService().GetCourierOrders(
		context.Background(),
		&pbo.GetCourierOrdersRequest{
			CourierId: courierID,
			Page:      1,
			Limit:     100,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting courier orders") {
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *handlerV1) GetCOOrders(c *gin.Context) {
	var (
		coID string
	)
	userInfo, err := userInfo(h, c)

	if err != nil {
		return
	}

	if userInfo.Role == config.RoleCargoOwnerAdmin {
		coID = userInfo.ID
	} else {
		coID = c.Query("co_id")

		_, err := uuid.Parse(coID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "cargo owner id is not valid",
			})
			return
		}
	}

	page, err := ParsePageQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err := ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	orders, err := h.grpcClient.OrderService().GetCOOrders(
		context.Background(),
		&pbo.GetCOOrdersRequest{
			CoId:  coID,
			Page:  page,
			Limit: limit,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting courier orders") {
		return
	}

	c.JSON(http.StatusOK, orders)
}

// @Router /v1/new-order [get]
// @Summary Get New Orders
// @Description API for getting new orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetOrders
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) NewOrders(c *gin.Context) {
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

	order, err := h.grpcClient.OrderService().GetNewOrders(context.Background(), &pbo.OrdersRequest{
		Page:  page,
		Limit: limit,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting new orders") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
