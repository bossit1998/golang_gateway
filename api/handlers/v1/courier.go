package v1

import (
	"context"
	"net/http"
	"strings"

	pbc "genproto/courier_service"
	pbs "genproto/sms_service"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

// @Router /v1/couriers/{courier_id}/courier_details [get]
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
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &courier)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	accessToken, err := jwt.GenerateJWT(id.String(), "courier", newSigningKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	courier.Id = id.String()
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

// @Router /v1/couriers/courier_details [post]
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

// @Router /v1/couriers/courier_details [put]
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

// @Router /v1/couriers/{courier_id}/block [post]
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

// @Router /v1/couriers/{courier_id}/unblock [post]
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

// @Router /v1/vehicle [post]
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

// @Router /v1/vehicle [put]
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

// @Router /v1/courier/check-login [POST]
// @Summary Check Courier Login
// @Description API that checks whether courier exists
// @Description and if exists sends sms to their number
// @Tags courier
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckLoginRequest true "check login"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckCourierLogin(c *gin.Context) {
	var (
		checkLoginModel models.CheckLoginRequest
		code string
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
				Text: code,
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
		Code: code,
		Phone: checkLoginModel.Login,
	})
}
/*
func (h *handlerV1) ConfirmCourierLogin(c *gin.Context) {

}*/