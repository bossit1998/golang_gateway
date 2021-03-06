package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	pbo "genproto/order_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Security ApiKeyAuth
// @Router /v1/demand-order/ [post]
// @Summary Create Demand Order
// @Description API for creating demand order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.CreateDemandOrderModel true "order"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateDemandOrder(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
		userInfo      models.UserInfo
	)

	accessToken := c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		ErrUnauthorized := errors.New("unauthorized")
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return
	}

	userInfo, err := ReturnUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: err.Error(),
		})
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
	order.ShipperId = userInfo.ShipperID
	order.CreatorId = userInfo.UserID
	order.CreatorTypeId = userInfo.UserTypeID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"
	order.StatusId = config.VendorReadyStatusId

	resp, err := h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router /v1/ondemand-order/ [post]
// @Summary Create On Demand Order
// @Description API for creating on demand order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.CreateOnDemandOrderModel true "order"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateOnDemandOrder(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
		userInfo      models.UserInfo
	)

	// err := getUserInfo(h, c, &userInfo)
	// if err != nil {
	// 	return
	// }
	accessToken := c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		ErrUnauthorized := errors.New("unauthorized")
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return
	}

	userInfo, err := ReturnUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: err.Error(),
		})
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

	if order.ToLocation == nil || order.ToLocation.Lat == 0 || order.ToLocation.Long == 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("Location is not valid", logger.Error(err))
		return
	}

	if order.PaymentType != "cash" && order.PaymentType != "payme" && order.PaymentType != "click" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("payment type is not valid", logger.Error(err))
		return
	}

	if order.Source != "admin_panel" && order.Source != "website" &&
		order.Source != "bot" && order.Source != "android" && order.Source != "ios" {

		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("source type is not valid", logger.Error(err))
		return
	}

	order.DeliveryPrice = order.CoDeliveryPrice
	order.ShipperId = userInfo.ShipperID
	order.CreatorId = userInfo.UserID
	order.CreatorTypeId = userInfo.UserTypeID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"

	if order.Steps[0].BranchId.GetValue() == "" {
		order.StatusId = config.NewStatusId
	} else {
		order.StatusId = config.OperatorAcceptedStatusId
	}

	resp, err := h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	go func() {
		if order.Steps[0].BranchId.GetValue() != "" {
			values, err := json.Marshal(map[string]string{
				"order_id": resp.OrderId,
			})
			if err != nil {
				h.log.Error("Error while marshaling", logger.Error(err))
				return
			}

			_, err = http.Post(config.TelegramBotURL+"/send-order/", "application/json", bytes.NewBuffer(values))
			if err != nil {
				h.log.Error("Error while sending push to vendor bot", logger.Error(err))
			}
		}
	}()

	c.JSON(http.StatusOK, resp)
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
		userInfo      models.UserInfo
	)

	// err := getUserInfo(h, c, &userInfo)
	// if err != nil {
	// 	return
	// }
	accessToken := c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		ErrUnauthorized := errors.New("unauthorized")
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return
	}

	userInfo, err := ReturnUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	orderID := c.Param("order_id")

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	if order.PaymentType != "cash" && order.PaymentType != "payme" && order.PaymentType != "click" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("payment type is not valid", logger.Error(err))
		return
	}

	if order.Source != "admin_panel" && order.Source != "website" &&
		order.Source != "bot" && order.Source != "android" && order.Source != "ios" {

		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("source type is not valid", logger.Error(err))
		return
	}

	order.Id = orderID
	order.DeliveryPrice = order.CoDeliveryPrice
	order.ShipperId = userInfo.ShipperID
	order.CreatorId = userInfo.UserID
	order.CreatorTypeId = userInfo.UserTypeID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"

	if order.Steps[0].BranchId.GetValue() == "" {
		order.StatusId = config.NewStatusId
	} else {
		order.StatusId = config.OperatorAcceptedStatusId
	}

	_, err = h.grpcClient.OrderService().Update(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	go func() {
		if order.Steps[0].BranchId.GetValue() != "" {
			values, err := json.Marshal(map[string]string{
				"order_id": orderID,
			})
			if err != nil {
				h.log.Error("Error while marshaling", logger.Error(err))
				return
			}

			_, err = http.Post(config.TelegramBotURL+"/send-order/", "application/json", bytes.NewBuffer(values))
			if err != nil {
				h.log.Error("Error while sending push to vendor bot", logger.Error(err))
			}
		}
	}()

	c.JSON(200, models.ResponseOK{
		Message: "order successfully updated",
	})
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id} [get]
// @Summary Get Order
// @Description API for getting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.GetOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrder(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		orderID     string
		userInfo    models.UserInfo
		//model models.GetOrderModel
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	orderID = c.Param("order_id")

	order, err := h.grpcClient.OrderService().Get(context.Background(), &pbo.GetRequest{
		ShipperId: userInfo.ShipperID,
		Id:        orderID,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}
	//
	//err = json.Unmarshal([]byte(js), &model)
	//
	//if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
	//	return
	//}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/order [get]
// @Summary Get Orders
// @Description API for getting orders
// @Tags order
// @Accept json
// @Produce json
// @Param status_id query string false "status_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param branch_ids query []string false "branch_ids"
// @Param customer_phone query string false "customer_phone"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Success 200 {object} models.GetAllOrderModel
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
		model       models.GetAllOrderModel
		userInfo    models.UserInfo
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	//jspbMarshal.EmitDefaults = true

	statusID = c.Query("status_id")

	if statusID != "" {
		_, err = uuid.Parse(statusID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "status_id is invalid",
			})
			return
		}
	}

	page, err = ParsePageQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err = ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	order, err = h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		ShipperId:     userInfo.ShipperID,
		StatusId:      statusID,
		Page:          page,
		Limit:         limit,
		CustomerPhone: c.Query("customer_phone"),
		BranchIds:     c.QueryArray("branch_ids[]"),
		StartDate:     c.Query("start_date"),
		EndDate:       c.Query("end_date"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
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
func (h *handlerV1) ChangeOrderStatus(c *gin.Context) {
	var (
		jspbUnmarshal jsonpb.Unmarshaler
		statusNote    pbo.StatusNote
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &statusNote)
	statusNote.OrderId = c.Param("order_id")
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	_, err = h.grpcClient.OrderService().ChangeStatus(
		context.Background(),
		&pbo.ChangeStatusRequest{
			ShipperId:  userInfo.ShipperID,
			StatusNote: &statusNote,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while changing order status") {
		return
	}

	// send push for aliftech
	go func() {
		if userInfo.ShipperID == config.Load().AliftechShipperId && statusNote.StatusId == config.CourierPickedUpStatusId {
			helpers.SendPush(c.Param("order_id"), statusNote.StatusId, h.log)
		}
	}()

	c.JSON(200, models.ResponseOK{
		Message: "changing order status successfully",
	})
}

// @Security ApiKeyAuth
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
	m["operator_accepted"] = config.OperatorAcceptedStatusId
	m["operator_cancelled"] = config.OperatorCancelledStatusId
	m["vendor_accepted"] = config.VendorAcceptedStatusId
	m["vendor_cancelled"] = config.VendorCancelledStatusId
	m["vendor_ready"] = config.VendorReadyStatusId
	m["courier_accepted"] = config.CourierAcceptedStatusId
	m["courier_cancelled"] = config.CourierCancelledStatusId
	m["courier_picked_up"] = config.CourierPickedUpStatusId
	m["delivered"] = config.DeliveredStatusId
	m["finished"] = config.FinishedStatusId
	m["server_cancelled"] = config.ServerCancelledStatusId
	m[config.NewStatusId] = "New"
	m[config.OperatorAcceptedStatusId] = "Operator Accepted"
	m[config.OperatorCancelledStatusId] = "Operator Cancelled"
	m[config.VendorAcceptedStatusId] = "Vendor Accepted"
	m[config.VendorCancelledStatusId] = "Vendor Cancelled"
	m[config.VendorReadyStatusId] = "Vendor Ready"
	m[config.CourierAcceptedStatusId] = "Courier Accepted"
	m[config.CourierPickedUpStatusId] = "Courier Picked up"
	m[config.CourierCancelledStatusId] = "Courier Cancelled"
	m[config.DeliveredStatusId] = "Delivered"
	m[config.FinishedStatusId] = "Finished"
	m[config.ServerCancelledStatusId] = "Server Cancelled"

	c.JSON(http.StatusOK, m)
}

// @Security ApiKeyAuth
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
		userInfo        models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	orderID = c.Param("order_id")
	order, err := h.grpcClient.OrderService().Get(context.Background(), &pbo.GetRequest{
		ShipperId: userInfo.ShipperID,
		Id:        orderID,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting order") {
		return
	}

	if order.CourierId.GetValue() != "" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("Courier have been already assigned", logger.Error(err))
		return
	}

	err = c.ShouldBindJSON(&addCourierModel)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	_, err = h.grpcClient.OrderService().AddCourier(
		context.Background(),
		&pbo.AddCourierRequest{
			ShipperId: userInfo.ShipperID,
			OrderId:   orderID,
			CourierId: addCourierModel.CourierID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while adding order courier") {
		return
	}

	go func() {
		// send push for aliftech
		if userInfo.ShipperID == config.Load().AliftechShipperId {
			go helpers.SendPush(c.Param("order_id"), config.CourierAcceptedStatusId, h.log)
		}

		values, err := json.Marshal(map[string]string{
			"order_id":   orderID,
			"courier_id": addCourierModel.CourierID,
		})
		if err != nil {
			h.log.Error("Error while marshaling", logger.Error(err))
			return
		}

		_, err = http.Post(config.TelegramBotURL+"/send-courier-order/", "application/json", bytes.NewBuffer(values))
		if err != nil {
			h.log.Error("Error while sending push to vendor bot", logger.Error(err))
		}
	}()

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "courier added successfully",
	})
}

// @Security ApiKeyAuth
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
		orderID  string
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	orderID = c.Param("order_id")

	_, err = h.grpcClient.OrderService().RemoveCourier(
		context.Background(),
		&pbo.RemoveCourierRequest{
			ShipperId: userInfo.ShipperID,
			OrderId:   orderID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while removing order courier") {
		return
	}

	// send push for aliftech
	go func() {
		if userInfo.ShipperID == config.Load().AliftechShipperId {
			helpers.SendPush(c.Param("order_id"), config.VendorReadyStatusId, h.log)
		}
	}()

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
// @Success 200 {object} models.GetCourierOrdersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourierOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		courierID   string
		model       models.GetCourierOrdersModel
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	if userInfo.UserType == config.RoleCourier {
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
			ShipperId: userInfo.ShipperID,
			CourierId: courierID,
			Page:      1,
			Limit:     100,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting courier orders") {
		return
	}

	js, err := jspbMarshal.MarshalToString(orders)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/new-order [get]
// @Summary Get Courier New Orders
// @Description API for getting courier new orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param courier_id query string false "courier_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetOrders
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CourierNewOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		courierID   string
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	if userInfo.UserType == config.RoleCourier {
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

	page, err := ParsePageQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err := ParseLimitQueryParam(c)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	order, err := h.grpcClient.OrderService().GetCourierNewOrders(context.Background(), &pbo.GetCourierNewOrdersRequest{
		ShipperId: userInfo.ShipperID,
		CourierId: courierID,
		StatusId:  config.VendorAcceptedStatusId,
		Page:      page,
		Limit:     limit,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting courier new orders") {
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
// @Param courier_pickedup body models.CourierPickedUpRequest true "courier_pickedup"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) TakeOrderStep(c *gin.Context) {
	var (
		userInfo models.UserInfo
		model    models.CourierPickedUpRequest
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	if userInfo.UserType != config.RoleCourier {
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

	err = c.ShouldBindJSON(&model)
	if err != nil {
		h.log.Error("order_id does not exists in body", logger.Error(err))
		// c.JSON(http.StatusBadRequest, models.ResponseError{
		// 	Error: models.InternalServerError{
		// 		Code: ErrorBadRequest,
		// 	},
		// })
		// return
	}

	_, err = h.grpcClient.OrderService().ChangeStatusStep(
		context.Background(),
		&pbo.ChangeStatusStepRequest{
			StepId:    stepID,
			ShipperId: userInfo.ShipperID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while taking order step") {
		return
	}

	// send push for aliftech
	go func() {
		if userInfo.ShipperID == config.Load().AliftechShipperId && model.OrderID != "" {
			helpers.SendPush(model.OrderID, config.CourierPickedUpStatusId, h.log)
		}
	}()

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
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	phone := c.Param("phone")

	res, err := h.grpcClient.OrderService().GetCustomerAddresses(
		context.Background(),
		&pbo.GetCustomerAddressesRequest{
			ShipperId: userInfo.ShipperID,
			Phone:     phone,
		})

	if handleInternalWithMessage(c, h.log, err, "error while getting customer addresses") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id}/add-branch [patch]
// @Summary Add Branch ID to orders
// @Description API for adding branch_id
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param branch body models.AddBranchIDModel true "branch"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) AddBranchID(c *gin.Context) {
	var (
		model    models.AddBranchIDModel
		userInfo models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	orderID := c.Param("order_id")

	err = c.ShouldBindJSON(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code: ErrorBadRequest,
			},
		})
		return
	}

	_, err = h.grpcClient.OrderService().AddBranchID(
		context.Background(),
		&pbo.AddBranchIDRequest{
			OrderId:   orderID,
			ShipperId: userInfo.ShipperID,
			BranchId:  model.BranchID,
		})
	if handleInternalWithMessage(c, h.log, err, "error while adding branch_id") {
		return
	}

	go func() {
		values, err := json.Marshal(map[string]string{
			"order_id": orderID,
		})
		if err != nil {
			h.log.Error("Error while marshaling", logger.Error(err))
			return
		}

		_, err = http.Post(config.TelegramBotURL+"/send-order/", "application/json", bytes.NewBuffer(values))
		if err != nil {
			h.log.Error("Error while sending push to vendor bot", logger.Error(err))
		}
	}()

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "branch_id added successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/customers/{customer_id}/orders [get]
// @Summary Get Customer Orders
// @Description API for getting customer orders
// @Tags customer
// @Accept json
// @Produce json
// @Param customer_id path string true "customer_id"
// @Param status_id query string false "status_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		order       *pbo.OrdersResponse
		statusID    string
		err         error
		page        uint64
		limit       uint64
		model       models.GetAllOrderModel
		userInfo    models.UserInfo
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	customerID := c.Param("customer_id")

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	//jspbMarshal.EmitDefaults = true

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
		order, err = h.grpcClient.OrderService().GetCustomerOrders(context.Background(), &pbo.GetCustomerOrdersRequest{
			ShipperId:  userInfo.ShipperID,
			CustomerId: customerID,
			Page:       page,
			Limit:      limit,
		})
	} else {
		_, err = uuid.Parse(statusID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "status_id is invalid",
			})
			return
		}

		order, err = h.grpcClient.OrderService().GetCustomerOrders(context.Background(), &pbo.GetCustomerOrdersRequest{
			ShipperId:  userInfo.ShipperID,
			CustomerId: customerID,
			StatusId:   statusID,
			Page:       page,
			Limit:      limit,
		})
	}

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all customer order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/branches/{branch_id}/orders [get]
// @Summary Get Branch Orders
// @Description API for getting branch orders
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string false "branch_id"
// @Param status_id query string false "status_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBranchOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		order       *pbo.OrdersResponse
		statusID    string
		err         error
		page        uint64
		limit       uint64
		model       models.GetAllOrderModel
		userInfo    models.UserInfo
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	branchID := c.Param("branch_id")

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	//jspbMarshal.EmitDefaults = true

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
		order, err = h.grpcClient.OrderService().GetBranchOrders(context.Background(), &pbo.GetBranchOrdersRequest{
			ShipperId: userInfo.ShipperID,
			BranchId:  branchID,
			Page:      page,
			Limit:     limit,
		})
	} else {
		_, err = uuid.Parse(statusID)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: "status_id is invalid",
			})
			return
		}

		order, err = h.grpcClient.OrderService().GetBranchOrders(context.Background(), &pbo.GetBranchOrdersRequest{
			ShipperId: userInfo.ShipperID,
			BranchId:  branchID,
			StatusId:  statusID,
			Page:      page,
			Limit:     limit,
		})
	}

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all branch order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id}/review [patch]
// @Summary Create Review For An Order
// @Description API for creating review for order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param order body models.OrderReview true "order_review"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateReview(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	orderID := c.Param("order_id")

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
	_, err = h.grpcClient.OrderService().CreateReview(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order review") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "review created successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id}/finish-otp [POST]
// @Summary Finish Order OTP
// @Description API that finish order otp
// @Tags order
// @Accept json
// @Produce json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) FinishOTP(c *gin.Context) {
	var userInfo models.UserInfo

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	orderID := c.Param("order_id")

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		config.Load().AliftechURL+orderID+"/request-complete",
		nil)
	if err != nil {
		h.log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Token", config.Load().AliftechAccessToken)
	resp, err := client.Do(request)
	if err != nil {
		h.log.Error("Error while sending push", logger.Error(err))
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("code has not been sent", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Code has been sent",
	})
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id}/confirm-finish-otp [POST]
// @Summary Confirm Finish Order OTP
// @Description API that confirm finish order otp
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param confirm_finish body models.ConfirmFinishOrderModel true "confirm finish"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmFinishOTP(c *gin.Context) {
	var (
		cm       models.ConfirmFinishOrderModel
		userInfo models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	err = c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	cm.Code = strings.TrimSpace(cm.Code)

	orderID := c.Param("order_id")
	values, err := json.Marshal(map[string]string{
		"otp": cm.Code,
	})
	if err != nil {
		h.log.Error("Error while marshaling", logger.Error(err))
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		config.Load().AliftechURL+orderID+"/complete",
		bytes.NewBuffer(values))
	if err != nil {
		h.log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Token", config.Load().AliftechAccessToken)
	resp, err := client.Do(request)
	if err != nil {
		h.log.Error("Error while sending push", logger.Error(err))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("code is invalid", logger.Error(err))
		return
	}

	// change order status to finish
	_, err = h.grpcClient.OrderService().ChangeStatus(
		context.Background(),
		&pbo.ChangeStatusRequest{
			ShipperId: userInfo.ShipperID,
			StatusNote: &pbo.StatusNote{
				OrderId:  orderID,
				StatusId: config.FinishedStatusId,
			},
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while changing order status") {
		return
	}

	c.JSON(http.StatusOK, &models.ResponseOK{
		Message: "Order finished successfully",
	})
}
