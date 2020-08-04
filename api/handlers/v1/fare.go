package v1

import (
	"context"
	"net/http"

	pb "genproto/fare_service"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Security ApiKeyAuth
// @Router /v1/fares/{fare_id} [get]
// @Summary Get Fare
// @Description API for getting fare
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare_id path string true "fare_id"
// @Success 200 {object} models.GetFareModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetFare(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.FareService().GetFare(
		context.Background(), &pb.GetFareRequest{
			Id: c.Param("fare_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Fare Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetFare())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/fares [post]
// @Summary Create Fare
// @Description API for creating fare
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare body models.CreateFareModel true "fare"
// @Success 200 {object} models.GetFareModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateFare(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		fare          pb.Fare
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &fare)
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

	res, err := h.grpcClient.FareService().Create(
		context.Background(),
		&fare,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating fare", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while creating fare, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetFare())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/fares [put]
// @Summary Update Fare
// @Description API for updating fare
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare body models.UpdateFareModel true "fare"
// @Success 200 {object} models.GetFareModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateFare(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		fare          pb.Fare
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &fare)
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

	res, err := h.grpcClient.FareService().Update(
		context.Background(),
		&fare,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while updating fare") {
		return
	}

	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating fare", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating fare, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetFare())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/fares [get]
// @Summary Get Fares
// @Description API for getting fares
// @Tags fare
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllFaresModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllFares(c *gin.Context) {
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

	res, err := h.grpcClient.FareService().GetAllFares(
		context.Background(),
		&pb.GetAllFaresRequest{
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

// @Security ApiKeyAuth
// @Router /v1/fares/{fare_id} [delete]
// @Summary Delete Fare
// @Description API for deleting fare
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare_id path string true "fare_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteFare(c *gin.Context) {
	_, err := h.grpcClient.FareService().Delete(
		context.Background(),
		&pb.DeleteFareRequest{
			Id: c.Param("fare_id"),
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
		h.log.Error("Error while deleting fare", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting fare, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting fare, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/delivery-price [get]
// @Summary Get Delivery Price
// @Description API for getting delivery price for shipper
// @Tags fare
// @Accept  json
// @Produce  json
// @Success 200 {object} models.DeliveryPriceModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetDeliveryPrice(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	res, err := h.grpcClient.FareService().GetDeliveryPrice(
		context.Background(), &pb.GetDeliveryPriceRequest{
			ShipperId: userInfo.ShipperID,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while delivery price") {
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
// @Router /v1/delivery-price [post]
// @Summary Create Delivery Price for Shipper
// @Description API for creating delivery price for shipper
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare body models.DeliveryPriceModel true "delivery_price"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateDeliveryPrice(c *gin.Context) {
	var (
		jspbUnmarshal jsonpb.Unmarshaler
		dp            pb.DeliveryPrice
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &dp)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	dp.ShipperId = userInfo.ShipperID

	_, err = h.grpcClient.FareService().CreateDeliveryPrice(
		context.Background(),
		&dp,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "Successfully created",
	})
}

// @Security ApiKeyAuth
// @Router /v1/delivery-price [put]
// @Summary Update Delivery Price
// @Description API for updating delivery price
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare body models.DeliveryPriceModel true "delivery_price"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateDeliveryPrice(c *gin.Context) {
	var (
		jspbUnmarshal jsonpb.Unmarshaler
		dp            pb.DeliveryPrice
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &dp)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	dp.ShipperId = userInfo.ShipperID

	_, err = h.grpcClient.FareService().UpdateDeliveryPrice(
		context.Background(),
		&dp,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while updating order") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "Successfully updated",
	})
}
