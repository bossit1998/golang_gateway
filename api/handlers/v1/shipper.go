package v1

import (
	"context"
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

// @Router /v1/shippers [post]
// @Summary Create Shipper
// @Description API for creating shipper
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param shipper body models.CreateShipperModel true "shipper"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateShipper(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		shipper        pbu.Shipper
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &shipper)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	accessToken, err := jwt.GenerateJWT(id.String(), "Shipper", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	shipper.Id = id.String()
	shipper.AccessToken = accessToken

	res, err := h.grpcClient.ShipperService().CreateShipper(
		context.Background(), &pbu.CreateShipperRequest{
			Shipper: &shipper,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating Shipper") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Shipper)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/shippers [put]
// @Summary Update Shipper
// @Description API for updating shipper
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param shipper body models.UpdateShipperModel true "shipper"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateShipper(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		shipper        pbu.Shipper
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &shipper)
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

	res, err := h.grpcClient.ShipperService().UpdateShipper(
		context.Background(),
		&pbu.UpdateShipperRequest{
			Shipper: &shipper,
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
		h.log.Error("Error while updating Shipper", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating Shipper, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetShipper())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags shipper
// @Router /v1/shippers/{shipper_id} [delete]
// @Summary Delete Shipper
// @Description API for deleting shipper
// @Accept  json
// @Produce  json
// @Param shipper_id path string true "shipper_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteShipper(c *gin.Context) {

	_, err := h.grpcClient.ShipperService().DeleteShipper(
		context.Background(),
		&pbu.DeleteShipperRequest{
			Id: c.Param("shipper_id"),
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
		h.log.Error("Error while deleting Shipper", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting shipper, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting shipper, service unavailable", logger.Error(err))
		return
	}
	c.Status(http.StatusOK)
}

// @Tags shipper
// @Router /v1/shippers/{shipper_id} [get]
// @Summary Get Shipper
// @Description API for getting shipper info
// @Accept  json
// @Produce json
// @Param shipper_id path string true "shipper_id"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetShipper(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.ShipperService().GetShipper(
		context.Background(), &pbu.GetShipperRequest{
			Id: c.Param("shipper_id"),
		},
	)
	st, ok := status.FromError(err)
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Shipper Not Found",
			},
		})
		h.log.Error("Error while getting Shipper, Shipper Not Found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Server unavailable",
			},
		})
		h.log.Error("Error while getting Shipper, service unavailable", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting Shipper", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetShipper())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/shippers [get]
// @Summary Get All shippers
// @Description API for getting shippers
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllShippersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllShippers(c *gin.Context) {
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

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.ShipperService().GetAllShippers(
		context.Background(),
		&pbu.GetAllShippersRequest{
			Page:  uint64(page),
			Limit: uint64(limit),
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

// @Router /v1/shippers/check-login/ [POST]
// @Summary Check Shipper Login
// @Description API that checks whether shipper exists
// @Description and if exists sends sms to their number
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckShipperLoginRequest true "check login"
// @Success 200 {object} models.CheckShipperLoginResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckShipperLogin(c *gin.Context) {
	var (
		checkShipperLoginModel models.CheckShipperLoginRequest
		code                  string
	)

	err := c.ShouldBindJSON(&checkShipperLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	checkShipperLoginModel.Phone = strings.TrimSpace(checkShipperLoginModel.Phone)

	resp, err := h.grpcClient.ShipperService().ExistsShipper(
		context.Background(), &pbu.ExistsShipperRequest{
			Phone: checkShipperLoginModel.Phone,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking shipper") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Shipper not found",
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
				Recipients: []string{checkShipperLoginModel.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(checkShipperLoginModel.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.JSON(http.StatusOK, models.CheckCustomerLoginResponse{
		Code:  code,
		Phone: checkShipperLoginModel.Phone,
	})
}

// @Router /v1/shippers/confirm-login/ [POST]
// @Summary Confirm shipper Login
// @Description API that checks whether - Shipper entered
// @Description valid token
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param confirm_phone body models.ConfirmShipperLoginRequest true "confirm login"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmShipperLogin(c *gin.Context) {
	var (
		cb models.ConfirmShipperLoginRequest
	)

	err := c.ShouldBindJSON(&cb)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	cb.Code = strings.TrimSpace(cb.Code)

	//Getting code from redis
	key := cb.Phone
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
	if cb.Code != s {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "Code is invalid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	_, err = h.grpcClient.ShipperService().GetShipper(
		context.Background(), &pbu.GetShipperRequest{
			Id: cb.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting Shipper") {
		return
	}

	c.Status(http.StatusOK)
}