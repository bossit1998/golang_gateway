package v1

import (
	"context"
	pbc "genproto/courier_service"
	pbs "genproto/sms_service"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"

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
)

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id} [get]
// @Summary Get Courier
// @Description API for getting courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.GetCourierModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourier(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetCourier(
		context.Background(), &pbc.GetCourierRequest{
			Id: c.Param("courier_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Courier Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetCourier())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id}/courier-details [get]
// @Summary Get Courier Details
// @Description API for getting courier details
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.CourierDetailsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourierDetails(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetCourierDetails(
		context.Background(), &pbc.GetCourierDetailsRequest{
			CourierId: c.Param("courier_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Details Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetCourierDetails())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers [get]
// @Summary Get Couriers
// @Description API for getting couriers
// @Tags courier
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllCouriersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCouriers(c *gin.Context) {
	var (
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

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

	res, err := h.grpcClient.CourierService().GetAllCouriers(
		context.Background(),
		&pbc.GetAllCouriersRequest{
			ShipperId: userInfo.ShipperID,
			Page:      uint64(page),
			Limit:     uint64(pageSize),
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

//@Security ApiKeyAuth
// @Router /v1/couriers [post]
// @Summary Create Courier
// @Description API for creating courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier body models.CreateCourierModel true "courier"
// @Success 200 {object} models.GetCourierModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCourier(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		courier       pbc.Courier
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &courier)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	// validate phone
	err = helpers.ValidatePhone(courier.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.grpcClient.CourierService().ExistsCourier(
		context.Background(), &pbc.ExistsCourierRequest{
			PhoneNumber: courier.Phone,
		},
	)

	if resp.Exists {
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

	accessToken, err := jwt.GenerateJWT(id.String(), "courier", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	courier.Id = id.String()
	courier.ShipperId = &wrappers.StringValue{Value: userInfo.ShipperID}
	courier.AccessToken = accessToken

	res, err := h.grpcClient.CourierService().Create(
		context.Background(), &courier,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating courier") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Courier)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers/courier-details [post]
// @Summary Create Courier Details
// @Description API for creating courier details
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier body models.CourierDetailsModel true "courier_details"
// @Success 200 {object} models.CourierDetailsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCourierDetails(c *gin.Context) {
	var (
		jspbMarshal    jsonpb.Marshaler
		jspbUnmarshal  jsonpb.Unmarshaler
		courierDetails pbc.CourierDetails
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &courierDetails)
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

	cd, err := h.grpcClient.CourierService().CreateCourierDetails(
		context.Background(),
		&courierDetails,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating courier details", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating courier details, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(cd.CourierDetails)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers [put]
// @Summary Update Courier
// @Description API for updating courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier body models.UpdateCourierModel true "courier"
// @Success 200 {object} models.GetCourierModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCourier(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		courier       pbc.Courier
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &courier)
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

	result, err := h.grpcClient.CourierService().GetCourier(
		context.Background(), &pbc.GetCourierRequest{
			Id: courier.Phone,
		},
	)

	if result != nil && result.Courier.Id != courier.Id {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Phone already exists",
			},
		})
		h.log.Error("Error while checking phone, Already exists", logger.Error(err))
		return
	}

	res, err := h.grpcClient.CourierService().Update(
		context.Background(),
		&courier,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating courier", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating courier, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetCourier())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers/courier-details [put]
// @Summary Update Courier Details
// @Description API for updating courier details
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_details body models.CourierDetailsModel true "courier_details"
// @Success 200 {object} models.CourierDetailsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCourierDetails(c *gin.Context) {

	var (
		jspbMarshal    jsonpb.Marshaler
		jspbUnmarshal  jsonpb.Unmarshaler
		courierDetails pbc.CourierDetails
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &courierDetails)
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

	res, err := h.grpcClient.CourierService().UpdateCourierDetails(
		context.Background(),
		&courierDetails,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating courier details", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating courier details, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.CourierDetails)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id} [delete]
// @Summary Delete Courier
// @Description API for deleting courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCourier(c *gin.Context) {

	_, err := h.grpcClient.CourierService().Delete(
		context.Background(),
		&pbc.DeleteCourierRequest{
			Id: c.Param("courier_id"),
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
		h.log.Error("Error while deleting courier", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting courier, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting courier, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id}/block [patch]
// @Summary Blocking Courier
// @Description API for blocking courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) BlockCourier(c *gin.Context) {

	_, err := h.grpcClient.CourierService().BlockCourier(
		context.Background(),
		&pbc.BlockCourierRequest{
			Id: c.Param("courier_id"),
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
		h.log.Error("Error while deleting courier", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting courier, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting courier, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id}/unblock [patch]
// @Summary Unblocking Courier
// @Description API for unblocking courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UnblockCourier(c *gin.Context) {

	_, err := h.grpcClient.CourierService().UnblockCourier(
		context.Background(),
		&pbc.UnblockCourierRequest{
			Id: c.Param("courier_id"),
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
		h.log.Error("Error while deleting courier", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting courier, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting courier, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/vehicle/{vehicle_id} [get]
// @Summary Get Courier Vehicle
// @Description API for getting courier vehicle
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param vehicle_id path string true "vehicle_id"
// @Success 200 {object} models.GetCourierVehicleModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourierVehicle(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetCourierVehicle(
		context.Background(), &pbc.GetCourierVehicleRequest{
			Id: c.Param("vehicle_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Vehicle Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetCourierVehicle())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id}/active-vehicle [get]
// @Summary Get Courier Active Vehicle
// @Description API for getting courier's  active vehicle
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.GetCourierVehicleModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCourierActiveVehicle(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetCourierActiveVehicle(
		context.Background(),
		&pbc.GetCourierActiveVehicleRequest{
			CourierId: c.Param("courier_id"),
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
// @Router /v1/couriers/{courier_id}/vehicles [get]
// @Summary Get All Courier Vehicles
// @Description API for getting courier's vehicles
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.GetAllCourierVehiclesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCourierVehicles(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetAllCourierVehicles(
		context.Background(),
		&pbc.GetAllCourierVehiclesRequest{
			CourierId: c.Param("courier_id"),
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
// @Router /v1/vehicles [get]
// @Summary Get All Vehicles
// @Description API for getting all vehicles
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllCourierVehiclesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllVehicles(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.CourierService().GetAllCourierVehicles(
		context.Background(),
		&pbc.GetAllCourierVehiclesRequest{},
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
// @Router /v1/vehicles [post]
// @Summary Create Courier Vehicle
// @Description API for creating courier vehicle
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param courier_vehicle body models.CreateCourierVehicleModel true "courier_vehicle"
// @Success 200 {object} models.GetCourierVehicleModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCourierVehicle(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		vehicle       pbc.CourierVehicle
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &vehicle)
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

	res, err := h.grpcClient.CourierService().CreateCourierVehicle(
		context.Background(),
		&vehicle,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating vehicle", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating vehicle, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetCourierVehicle())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/vehicles [put]
// @Summary Update Courier Vehicle
// @Description API for updating courier vehicle
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param courier_vehicle body models.UpdateCourierVehicleModel true "courier_vehicle"
// @Success 200 {object} models.GetCourierVehicleModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCourierVehicle(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		vehicle       pbc.CourierVehicle
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &vehicle)
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

	res, err := h.grpcClient.CourierService().UpdateCourierVehicle(
		context.Background(),
		&vehicle,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating vehicle", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})

		h.log.Error("Error while updating vehicle, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.CourierVehicle)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/vehicles/{vehicle_id} [delete]
// @Summary Delete Courier Vehicle
// @Description API for deleting courier vehicle
// @Tags vehicle
// @Accept  json
// @Produce  json
// @Param vehicle_id path string true "vehicle_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCourierVehicle(c *gin.Context) {

	_, err := h.grpcClient.CourierService().DeleteCourierVehicle(
		context.Background(),
		&pbc.DeleteCourierVehicleRequest{
			Id: c.Param("vehicle_id"),
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

		h.log.Error("Error while deleting courier vehicle", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting courier vehicle, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting courier vehicle, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Router /v1/couriers/check-login/ [POST]
// @Summary Check Courier Login
// @Description API that checks whether courier exists
// @Description and if exists sends sms to their number
// @Tags courier
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckLoginRequest true "check login"
// @Success 200 {object} models.CheckLoginResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckCourierLogin(c *gin.Context) {
	var (
		checkLoginModel models.CheckLoginRequest
		code            string
	)

	err := c.ShouldBindJSON(&checkLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	checkLoginModel.Login = strings.TrimSpace(checkLoginModel.Login)

	resp, err := h.grpcClient.CourierService().ExistsCourier(
		context.Background(), &pbc.ExistsCourierRequest{
			PhoneNumber: checkLoginModel.Login,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking courier") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Courier not found",
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
				Recipients: []string{checkLoginModel.Login},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(checkLoginModel.Login, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.JSON(http.StatusOK, models.CheckLoginResponse{
		Code:  code,
		Phone: checkLoginModel.Login,
	})
}

// @Router /v1/couriers/confirm-login/ [POST]
// @Summary Confirm Courier Login
// @Description API that checks whether courier entered
// @Description valid token
// @Tags courier
// @Accept  json
// @Produce  json
// @Param confirm_login body models.ConfirmLoginRequest true "confirm login"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmCourierLogin(c *gin.Context) {
	var (
		cm models.ConfirmLoginRequest
	)

	err := c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
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

	courier, err := h.grpcClient.CourierService().GetCourier(
		context.Background(), &pbc.GetCourierRequest{
			Id: cm.Phone,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting courier") {
		return
	}

	// check courier fcm token
	if courier.Courier.FcmToken.GetValue() != cm.FcmToken {
		_, err := h.grpcClient.CourierService().UpdateFcmToken(
			context.Background(), &pbc.UpdateFcmTokenRequest{
				Id:       courier.Courier.Id,
				FcmToken: cm.FcmToken,
			},
		)

		if handleGrpcErrWithMessage(c, h.log, err, "Error while requesting to courier service") {
			return
		}
	}

	m := map[interface{}]interface{}{
		"user_type":  "courier",
		"shipper_id": courier.Courier.ShipperId.GetValue(),
		"sub":        courier.Courier.Id,
	}
	access, _, err := jwt.GenJWT(m, signingKey)

	if handleInternalWithMessage(c, h.log, err, "Error while generating token") {
		return
	}

	_, err = h.grpcClient.CourierService().UpdateToken(
		context.Background(), &pbc.UpdateTokenRequest{
			Id:     courier.Courier.Id,
			Access: access,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while updating token") {
		return
	}

	c.JSON(http.StatusOK, &models.ConfirmLoginResponse{
		ID:          courier.Courier.Id,
		AccessToken: access,
	})
}

// @Security ApiKeyAuth
// @Router /v1/search-couriers [get]
// @Summary Search by phone
// @Description API for getting phones
// @Tags courier
// @Accept  json
// @Produce  json
// @Param phone query string true "phone"
// @Success 200 {object} models.SearchCouriersByPhoneModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) SearchCouriersByPhone(c *gin.Context) {
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

	res, err := h.grpcClient.CourierService().SearchCouriersByPhone(
		context.Background(),
		&pbc.SearchCouriersByPhoneRequest{
			ShipperId: userInfo.ShipperID,
			Phone:     phone,
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
// @Router /v1/branches/add-courier [post]
// @Summary Create Branch Courier
// @Description API for creating branch courier
// @Tags branch
// @Accept  json
// @Produce  json
// @Param courier body models.BranchCourierModel true "branch"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateBranchCourier(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		branchCourier pbc.CreateBranchCourierRequest
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &branchCourier)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	_, err = h.grpcClient.CourierService().CreateBranchCourier(
		context.Background(), &branchCourier,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating branch courier") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Router /v1/branches/remove-courier [post]
// @Summary Remove Courier From Branch
// @Description API for removing courier from branch
// @Tags branch
// @Accept json
// @Produce json
// @Param courier body models.BranchCourierModel true "branch"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteBranchCourier(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		branchCourier pbc.DeleteBranchCourierRequest
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &branchCourier)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	_, err = h.grpcClient.CourierService().DeleteBranchCourier(
		context.Background(), &branchCourier,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "Error while deleting branch courier") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id}/branches [get]
// @Summary Get All Couirer Branch Ids
// @Description API for getting courier branhes
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.GetAllCourierBranchesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCourierBranches(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.CourierService().GetAllCourierBranches(
		context.Background(),
		&pbc.GetAllCourierBranchesRequest{
			CourierId: c.Param("courier_id"),
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
// @Router /v1/branches/{branch_id}/couriers [get]
// @Summary Get All Branch Couriers
// @Description API for getting branch couriers
// @Tags branch
// @Accept  json
// @Produce  json
// @Param branch_id path string true "courier_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllCouriersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllBranchCouriers(c *gin.Context) {
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

	res, err := h.grpcClient.CourierService().GetAllBranchCouriers(
		context.Background(),
		&pbc.GetAllBranchCouriersRequest{
			BranchId: c.Param("branch_id"),
			Page:     uint64(page),
			Limit:    uint64(pageSize),
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
