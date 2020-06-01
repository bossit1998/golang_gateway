package v1

import (
	"context"
	"fmt"
	pbo "genproto/order_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
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

	resp, err := h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, resp)
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
// @Param status_id query string false "status_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetOrders
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		order       *pbo.OrdersResponse
		statusID    string
		err         error
		page        uint64
		limit       uint64
	)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	statusID = c.Query("status_id")

	page, err = ParsePageQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err = ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	if statusID == "" {
		order, err = h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
			Page:  page,
			Limit: limit,
		})
	} else {
		_, err = uuid.Parse(statusID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "status_id is invalid",
			})
			return
		}

		order, err = h.grpcClient.OrderService().GetOrdersByStatus(context.Background(), &pbo.GetOrdersByStatusRequest{
			StatusId: statusID,
			Page:     page,
			Limit:    limit,
		})
	}

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
	var ()
	m := make(map[string]string)
	m["new"] = config.NewStatusId
	m["courier_accepted"] = config.CourierAcceptedStatusId
	m["courier_picked_up"] = config.CourierPickedUpStatusId
	m["delivered"] = config.DeliveredStatusId
	m["finished"] = config.FinishedStatusId
	m["courier_cancelled"] = config.CourierCancelledStatusId
	m["vendor_accepted"] = config.VendorAcceptedStatusId
	m["vendor_cancelled"] = config.VendorCancelledStatusId
	m["vendor_ready"] = config.VendorReadyStatusId
	m[config.NewStatusId] = "New"
	m[config.CourierAcceptedStatusId] = "Courier Accepted"
	m[config.CourierPickedUpStatusId] = "Courier Picked up"
	m[config.DeliveredStatusId] = "Delivered"
	m[config.FinishedStatusId] = "Finished"
	m[config.CourierCancelledStatusId] = "Courier Cancelled"
	m[config.VendorAcceptedStatusId] = "Vendor Accepted"
	m[config.VendorCancelledStatusId] = "Vendor Cancelled"
	m[config.VendorReadyStatusId] = "Vendor Ready"

	//status = models.Status{ID: config.NEW_STATUS_ID, Name: "New"}
	//model.Statuses = append(model.Statuses, status)
	//
	//status = models.Status{ID: config.CANCELLED_STATUS_ID, Name: "Cancelled"}
	//model.Statuses = append(model.Statuses, status)
	//
	//status = models.Status{ID: config.ACCEPTED_STATUS_ID, Name: "Accepted"}
	//model.Statuses = append(model.Statuses, status)
	//
	//status = models.Status{ID: "84be5a2f-3a92-4469-8283-220ca34a0de4", Name: "Picked up"}
	//model.Statuses = append(model.Statuses, status)
	//
	//status = models.Status{ID: config.DELIVERED_STATUS_ID, Name: "Delivered"}
	//model.Statuses = append(model.Statuses, status)
	//
	//status = models.Status{ID: config.FINISHED_STATUS_ID, Name: "Finished"}
	//model.Statuses = append(model.Statuses, status)
	//
	//var a int
	//fmt.Scan(a)

	c.JSON(http.StatusOK, m)
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

	order, err := h.grpcClient.OrderService().GetOrdersByStatus(context.Background(), &pbo.GetOrdersByStatusRequest{
		StatusId: config.NewStatusId,
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

// @Security ApiKeyAuth
// @Router /v1/order-step/{step_id}/take [patch]
// @Summary Take Order Steps
// @Description API for taking order step
// @Tags order
// @Accept  json
// @Produce  json
// @Param step_id path string true "step_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) TakeOrderStep(c *gin.Context) {
	userInfo, err := userInfo(h, c)

	if err != nil {
		return
	}

	if userInfo.Role != config.RoleCourier {
		c.JSON(http.StatusForbidden, "")
		return
	}

	stepID := c.Param("step_id")

	_, err = uuid.Parse(stepID)

	if err != nil {
		c.JSON(http.StatusOK, models.ResponseError{
			Error: "invalid uuid format in param",
		})
	}

	_, err = h.grpcClient.OrderService().ChangeStatusStep(
		context.Background(),
		&pbo.ChangeStatusStepRequest{
			StepId: stepID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while taking order step") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "order step took",
	})
}

// @Security ApiKeyAuth
// @Router /v1/customer-addresses/{phone} [get]
// @Summary Get Customer Order Addresses
// @Description API for taking all order addresses
// @Tags order
// @Accept  json
// @Produce  json
// @Param phone path string true "phone"
// @Success 200 {object} models.CustomerAddressesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerAddresses(c *gin.Context) {
	phone := c.Param("phone")

	res, _ := h.grpcClient.OrderService().GetCustomerAddresses(
		context.Background(),
		&pbo.GetCustomerAddressesRequest{
			Phone:phone,
		})

	c.JSON(http.StatusOK, res)
}