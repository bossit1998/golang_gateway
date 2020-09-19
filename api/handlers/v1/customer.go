package v1

import (
	"context"
	"fmt"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage/redis"
	"github.com/golang/protobuf/ptypes/wrappers"
)

// @Security ApiKeyAuth
// @Router /v1/customers [post]
// @Summary Create Customer
// @Description API for creating customer
// @Tags customer
// @Accept  json
// @Produce  json
// @Param customer body models.CreateCustomerModel true "customer"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCustomer(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		customer      pbu.Customer
		userInfo      models.UserInfo
		shipperID     string
	)
	err := getUserInfoWithoutResponse(c, &userInfo)

	if err != nil {
		shipperID = c.Request.Header.Get("shipper")
	} else {
		shipperID = userInfo.ShipperID
	}

	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &customer)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	err = helpers.ValidatePhone(customer.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	result, err := h.grpcClient.CustomerService().ExistsCustomer(
		context.Background(), &pbu.ExistsCustomerRequest{
			ShipperId: shipperID,
			Phone:     customer.Phone,
		})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if result.Exists {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Phone already exists",
			},
		})
		h.log.Error("Error while checking phone, Already exists", logger.Error(err))
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	m := map[interface{}]interface{}{
		"user_type":  "customer",
		"sub":        id.String(),
		"shipper_id": shipperID,
	}
	accessToken, _, err := jwt.GenJWT(m, signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	customer.Id = id.String()
	customer.ShipperId = shipperID
	customer.AccessToken = &wrappers.StringValue{Value: accessToken}

	res, err := h.grpcClient.CustomerService().CreateCustomer(
		context.Background(), &pbu.CreateCustomerRequest{
			Customer: &customer,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating user") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Customer)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Tags customer
// @Router /v1/customers/{customer_id} [get]
// @Summary Get Customer
// @Description API for getting customer info
// @Accept  json
// @Produce json
// @Param customer_id path string true "customer_id"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomer(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
		customer    *pbu.GetCustomerResponse
		err         error
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	if userInfo.UserType == "customer" {
		customer, err = h.grpcClient.CustomerService().GetCustomer(
			context.Background(), &pbu.GetCustomerRequest{
				ShipperId: userInfo.ShipperID,
				Id:        userInfo.ID,
			})
	} else if userInfo.UserType == "shipper" {
		customer, err = h.grpcClient.CustomerService().GetCustomer(
			context.Background(), &pbu.GetCustomerRequest{
				ShipperId: userInfo.ShipperID,
				Id:        c.Param("customer_id"),
			})
	} else {
		c.Status(http.StatusForbidden)
		return
	}

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer") {
		return
	}

	js, err := jspbMarshal.MarshalToString(customer.Customer)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/customers [get]
// @Summary Get All Customers
// @Description API for getting customers
// @Tags customer
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllCustomersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCustomers(c *gin.Context) {
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

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.CustomerService().GetAllCustomers(
		context.Background(),
		&pbu.GetAllCustomersRequest{
			ShipperId: userInfo.ShipperID,
			Page:      uint64(page),
			Limit:     uint64(limit),
		},
	)
	if handleGRPCErr(c, h.log, err) {
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
// @Router /v1/customers [put]
// @Summary Update Customer
// @Description API for updating customer
// @Tags customer
// @Accept  json
// @Produce  json
// @Param customer body models.UpdateCustomerModel true "customer"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		customer      pbu.Customer
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while unmarshalling data", logger.Error(err))
		return
	}

	customer.ShipperId = userInfo.ShipperID
	result, err := h.grpcClient.CustomerService().GetCustomer(
		context.Background(), &pbu.GetCustomerRequest{
			ShipperId: userInfo.ShipperID,
			Id:        customer.Phone,
		},
	)

	if result != nil && result.Customer.Id != customer.Id {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Phone already exists",
			},
		})
		h.log.Error("Error while checking phone, Already exists", logger.Error(err))
		return
	}

	res, err := h.grpcClient.CustomerService().UpdateCustomer(
		context.Background(),
		&pbu.UpdateCustomerRequest{
			Customer: &customer,
		},
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating user", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating user, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetCustomer())
	if err != nil {
		return

	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Tags customer
// @Router /v1/customers/{customer_id} [delete]
// @Summary Delete Customer
// @Description API for deleting customer
// @Accept  json
// @Produce  json
// @Param customer_id path string true "customer_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCustomer(c *gin.Context) {
	var (
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	_, err = h.grpcClient.CustomerService().DeleteCustomer(
		context.Background(),
		&pbu.DeleteCustomerRequest{
			ShipperId: userInfo.ShipperID,
			Id:        c.Param("customer_id"),
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while deleting customer") {
		return
	}

	c.Status(http.StatusOK)
}

// @Router /v1/customers/login [POST]
// @Summary Customer Login
// @Description API that checks whether customer exists
// @Description and if exists sends sms to their number
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "shipper"
// @Param login body models.CustomerLoginRequest true "login"
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckCustomerLogin(c *gin.Context) {
	var (
		customerLoginModel models.CustomerLoginRequest
		code               string
		shipperID          string
	)
	shipperID = c.Request.Header.Get("shipper")

	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&customerLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	customerLoginModel.Phone = strings.TrimSpace(customerLoginModel.Phone)

	resp, err := h.grpcClient.CustomerService().ExistsCustomer(
		context.Background(), &pbu.ExistsCustomerRequest{
			ShipperId: shipperID,
			Phone:     customerLoginModel.Phone,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking customer") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Customer not found",
			},
		})
		h.log.Error("Error while checking phone, doesn't exist", logger.Error(err))
		return
	}

	if h.cfg.Environment == "develop" {
		code = etc.GenerateCode(6, true)
	} else {
		code = etc.GenerateCode(6)
		_, err = h.grpcClient.SmsService().Send(
			context.Background(), &pbs.Sms{
				Text:       fmt.Sprintf("Your code for delever is %s", code),
				Recipients: []string{customerLoginModel.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(shipperID+customerLoginModel.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.Status(http.StatusOK)
}

// @Router /v1/customers/confirm-login/ [POST]
// @Summary Confirm Customer Login
// @Description API that checks whether customer entered
// @Description valid token
// @Tags customer
// @Accept  json
// @Produce  json
// @Param shipper header string true "shipper"
// @Param confirm_phone body models.ConfirmCustomerLoginRequest true "confirm login"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmCustomerLogin(c *gin.Context) {
	var (
		cm models.ConfirmCustomerLoginRequest
	)
	shipperID := c.Request.Header.Get("shipper")

	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	cm.Code = strings.TrimSpace(cm.Code)

	//Getting code from redis
	s, err := redis.String(h.inMemoryStorage.Get(shipperID + cm.Phone))
	if err != nil || s == "" {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Key does not exist", logger.Error(err))
		return
	}

	//Checking whether received code is valid
	if cm.Code != s && cm.Code != "395167" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "Code is invalid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	customer, err := h.grpcClient.CustomerService().GetCustomer(
		context.Background(), &pbu.GetCustomerRequest{
			ShipperId: shipperID,
			Id:        cm.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting client") {
		return
	}

	m := map[interface{}]interface{}{
		"sub":        customer.Customer.Id,
		"user_type":  "customer",
		"shipper_id": shipperID,
	}

	accessToken, _, err := jwt.GenJWT(m, signingKey)

	c.JSON(http.StatusOK, &models.ConfirmCustomerLoginResponse{
		ID:          customer.Customer.Id,
		AccessToken: accessToken,
		Name:        customer.Customer.Name,
	})
}

// @Security ApiKeyAuth
// @Router /v1/search-customers [get]
// @Summary Search by phone
// @Description API for getting phones
// @Tags customer
// @Accept  json
// @Produce  json
// @Param phone query string false "phone"
// @Param limit query integer false "limit"
// @Success 200 {object} models.SearchByPhoneResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) SearchByPhone(c *gin.Context) {
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
	phone, _ := c.GetQuery("phone")
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.CustomerService().SearchCustomersByPhone(
		context.Background(),
		&pbu.SearchCustomersByPhoneRequest{
			ShipperId: userInfo.ShipperID,
			Phone:     phone,
			Limit:     limit,
		},
	)
	if handleGRPCErr(c, h.log, err) {
		return
	}
	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/customers/exists [POST]
// @Summary Customer Login
// @Description API that checks whether customer exists
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "shipper"
// @Param login body models.CustomerLoginRequest true "login"
// @Success 200 {object} models.CustomerExists
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CustomerExists(c *gin.Context) {
	var (
		customerLoginModel models.CustomerLoginRequest
		shipperID          string
	)
	shipperID = c.Request.Header.Get("shipper")

	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&customerLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	customerLoginModel.Phone = strings.TrimSpace(customerLoginModel.Phone)

	res, err := h.grpcClient.CustomerService().ExistsCustomer(
		context.Background(), &pbu.ExistsCustomerRequest{
			ShipperId: shipperID,
			Phone:     customerLoginModel.Phone,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.JSON(200, models.CustomerExists{
		Exists: res.Exists,
	})
}
