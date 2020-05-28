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

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage/redis"
)

// @Router /v1/vendors [post]
// @Summary Create Vendor
// @Description API for creating vendor
// @Tags vendor
// @Accept  json
// @Produce  json
// @Param vendor body models.CreateVendorModel true "vendor"
// @Success 200 {object} models.GetVendorModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateVendor(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		vendor        pbu.Vendor
	)
	jspbMarshal.OrigName = true
	err := jspbUnmarshal.Unmarshal(c.Request.Body, &vendor)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	accessToken, err := jwt.GenerateJWT(id.String(), "user", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	vendor.Id = id.String()
	vendor.AccessToken = accessToken

	res, err := h.grpcClient.VendorService().CreateVendor(
		context.Background(), &pbu.CreateVendorRequest{
			Vendor: &vendor,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating vendor") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Vendor)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/vendors [put]
// @Summary Update Vendor
// @Description API for updating vendor
// @Tags vendor
// @Accept  json
// @Produce  json
// @Param vendor body models.UpdateVendorModel true "vendor"
// @Success 200 {object} models.GetVendorModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateVendor(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		vendor        pbu.Vendor
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &vendor)
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

	res, err := h.grpcClient.VendorService().UpdateVendor(
		context.Background(),
		&pbu.UpdateVendorRequest{
			Vendor: &vendor,
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
		h.log.Error("Error while updating vendor", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating vendor, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetVendor())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags vendor
// @Router /v1/vendors/{vendor_id} [delete]
// @Summary Delete Vendor
// @Description API for deleting vendor
// @Accept  json
// @Produce  json
// @Param vendor_id path string true "vendor_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteVendor(c *gin.Context) {

	_, err := h.grpcClient.VendorService().DeleteVendor(
		context.Background(),
		&pbu.DeleteVendorRequest{
			Id: c.Param("vendor_id"),
		},
	)
	fmt.Println(1111111111)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting vendor", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting vendor, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting vendor, service unavailable", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Tags vendor
// @Router /v1/vendors/{vendor_id} [get]
// @Summary Get Vendor
// @Description API for getting vendor info
// @Accept  json
// @Produce json
// @Param vendor_id path string true "vendor_id"
// @Success 200 {object} models.GetVendorModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetVendor(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.VendorService().GetVendor(
		context.Background(), &pbu.GetVendorRequest{
			Id: c.Param("vendor_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "User Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetVendor())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/vendors [get]
// @Summary Get All Vendors
// @Description API for getting vendors
// @Tags vendor
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllVendorsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllVendors(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	pageSize, err := ParsePageSizeQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.VendorService().GetAllVendors(
		context.Background(),
		&pbu.GetAllVendorsRequest{
			Page:  uint64(page),
			Limit: uint64(pageSize),
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

// @Router /v1/vendors/check-login/ [POST]
// @Summary Check Vendor Login
// @Description API that checks whether vendor exists
// @Description and if exists sends sms to their number
// @Tags vendor
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckVendorLoginRequest true "check login"
// @Success 200 {object} models.CheckVendorLoginResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckVendorLogin(c *gin.Context) {
	var (
		checkVendorLoginModel models.CheckVendorLoginRequest
		code            string
	)
	
	err := c.ShouldBindJSON(&checkVendorLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	checkVendorLoginModel.Phone = strings.TrimSpace(checkVendorLoginModel.Phone)

	resp, err := h.grpcClient.VendorService().ExistsVendor(
		context.Background(), &pbu.ExistsVendorRequest{
			Phone: checkVendorLoginModel.Phone,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking vendor") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "User not found",
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
				Text:       code,
				Recipients: []string{checkVendorLoginModel.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(checkVendorLoginModel.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.JSON(http.StatusOK, models.CheckUserLoginResponse{
		Code:  code,
		Phone: checkVendorLoginModel.Phone,
	})
}

// @Router /v1/vendors/confirm-login/ [POST]
// @Summary Confirm Vendor Login
// @Description API that checks whether - vendor entered
// @Description valid token
// @Tags vendor
// @Accept  json
// @Produce  json
// @Param confirm_phone body models.ConfirmVendorLoginRequest true "confirm login"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmVendorLogin(c *gin.Context) {
	var (
		cm models.ConfirmUserLoginRequest
	)

	err := c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}
//ConfirmVendorLoginResponse ...
type ConfirmVendorLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}
	cm.Code = strings.TrimSpace(cm.Code)

	//Getting code from redis
	key := cm.Phone
	s, err := redis.String(h.inMemoryStorage.Get(key))
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
	if cm.Code != s {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "Code is invalid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	user, err := h.grpcClient.UserService().GetClient(
		context.Background(), &pbu.GetClientRequest{
			Id: cm.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting client") {
		return
	}

	c.JSON(http.StatusOK, &models.ConfirmVendorLoginResponse{
		ID:          user.Client.Id,
		AccessToken: user.Client.AccessToken,
	})
}