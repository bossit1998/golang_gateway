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

// @Security ApiKeyAuth
// @Summary Get Courier
// @Description Get Courier API returns event
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "Courier Id"
// @Success 200 {object} models.GetCourierResp
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/couirer/{courier_id}/ [get]
func (h *handlerV1) GetCourier(c *gin.Context) {

	courierResp, err := h.grpcClient.CourierService().GetCourier(
		context.Background(), &pbc.GetCourierRequest{
			Id: c.Param("id"),
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
	courier := courierResp.Courier

	c.JSON(http.StatusOK, models.GetCourierResponseModel{
		ID:        courier.Id,
		Phone:     courier.Phone,
		FirstName: courier.FirstName,
		LastName:  courier.LastName,
		CreatedAt: courier.CreatedAt,
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

	generalResp := models.GetAllCouriersResponseModel{Count: int(resp.GetCount())}

	for _, e := range resp.GetCouriers() {
		generalResp.Couriers = append(generalResp.Couriers, models.GetCourierResponseModel{
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
		// claims        jwtg.MapClaims
	)

	jspbMarshal.OrigName = true
	// claims = GetClaims(h, c)

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
	courier.Id = c.Param("id")
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

	js, err := jspbMarshal.MarshalToString(updatedCourier.Courier)
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
			Id: c.Param("id"),
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
