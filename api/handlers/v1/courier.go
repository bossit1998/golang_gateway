package v1

import (
	"context"
	"net/http"

	"bitbucket.org/alien_soft/api_gateway/api/models"
	pbc "bitbucket.org/alien_soft/api_gateway/genproto/courier_service"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
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

	courierResp, err := h.grpcClient.CourierService().GetCourier(
		context.Background(), &pbc.GetCourierRequest{
			Id: c.Param("courier_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if courierResp == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Courier Not Found",
			},
		})
		return
	}
	courier := courierResp.Courier

	c.JSON(http.StatusOK, models.GetCourierModel{
		ID:        courier.Id,
		Phone:     courier.Phone,
		FirstName: courier.FirstName,
		LastName:  courier.LastName,
		CreatedAt: courier.CreatedAt,
	})
}

func (h *handlerV1) GetCourierDetails(c *gin.Context) {

	resp, err := h.grpcClient.CourierService().GetCourierDetails(
		context.Background(), &pbc.GetCourierDetailsRequest{
			CourierId: c.Param("courier_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if resp == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Details Not Found",
			},
		})
		return
	}
	cd := resp.CourierDetails

	c.JSON(http.StatusOK, models.CourierDetailsModel{
		PassportNumber:    cd.GetPassportNumber(),
		Gender:            cd.GetGender().GetValue(),
		BirthDate:         cd.GetBirthDate(),
		Address:           cd.GetAddress().GetValue(),
		Img:               cd.GetImg().GetValue(),
		LisenseNumber:     cd.GetLisenseNumber(),
		LisenseGivenDate:  cd.GetLisenseGivenDate(),
		LisenseExpiryDate: cd.GetLisenseExpiryDate(),
	})
}

func (h *handlerV1) GetAllCouriers(c *gin.Context) {

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

	resp, err := h.grpcClient.CourierService().GetAllCouriers(
		context.Background(),
		&pbc.GetAllCouriersRequest{
			Page:  uint64(page),
			Limit: uint64(pageSize),
		},
	)
	if handleGRPCErr(c, h.log, err) {
		return
	}

	generalResp := models.GetAllCouriersModel{Count: int(resp.GetCount())}

	for _, e := range resp.GetCouriers() {
		generalResp.Couriers = append(generalResp.Couriers, models.GetCourierModel{
			ID:        e.Id,
			Phone:     e.Phone,
			FirstName: e.FirstName,
			LastName:  e.LastName,
			CreatedAt: e.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, generalResp)
}

func (h *handlerV1) CreateCourier(c *gin.Context) {
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

	createdCourier, err := h.grpcClient.CourierService().Create(
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
		h.log.Error("Error while creating event", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating event, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(createdCourier.Courier)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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
		h.log.Error("Error while creating event", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating event, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(cd.CourierDetails)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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

	updatedCourier, err := h.grpcClient.CourierService().Update(
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

	js, err := jspbMarshal.MarshalToString(updatedCourier.Courier)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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

	cd, err := h.grpcClient.CourierService().UpdateCourierDetails(
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

	js, err := jspbMarshal.MarshalToString(cd.CourierDetails)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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

func (h *handlerV1) GetCourierVehicle(c *gin.Context) {

	courierResp, err := h.grpcClient.CourierService().GetCourierVehicle(
		context.Background(), &pbc.GetCourierVehicleRequest{
			Id: c.Param("vehicle_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if courierResp == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Event Not Found",
			},
		})
		return
	}
	vehicle := courierResp.CourierVehicle

	c.JSON(http.StatusOK, models.GetCourierModel{
		ID:        vehicle.Id,
		Phone:     vehicle.Model,
		FirstName: vehicle.VehicleNumber,
		CreatedAt: vehicle.CreatedAt,
	})
}

func (h *handlerV1) GetAllCourierVehicles(c *gin.Context) {

	resp, err := h.grpcClient.CourierService().GetAllCourierVehicles(
		context.Background(),
		&pbc.GetAllCourierVehiclesRequest{},
	)
	if handleGRPCErr(c, h.log, err) {
		return
	}

	generalResp := models.GetAllCouriersModel{}

	for _, e := range resp.GetCourierVehicles() {
		generalResp.Couriers = append(generalResp.Couriers, models.GetCourierModel{
			ID:        e.Id,
			Phone:     e.Model,
			FirstName: e.VehicleNumber,
			CreatedAt: e.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, generalResp)
}

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

	v, err := h.grpcClient.CourierService().CreateCourierVehicle(
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

	js, err := jspbMarshal.MarshalToString(v.CourierVehicle)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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

	v, err := h.grpcClient.CourierService().UpdateCourierVehicle(
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

		h.log.Error("Error while updating vehiclee, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(v.CourierVehicle)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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
