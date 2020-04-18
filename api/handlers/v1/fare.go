package v1

import (
	"context"
	"net/http"

	"bitbucket.org/alien_soft/api_gateway/api/models"
	pb "bitbucket.org/alien_soft/api_gateway/genproto/fare_service"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handlerV1) GetFare(c *gin.Context) {
	fareResponse, err := h.grpcClient.FareService().GetFare(
		context.Background(), &pb.GetFareRequest{
			Id: c.Param("fare_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if fareResponse == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Fare Not Found",
			},
		})
		return
	}
	fare := fareResponse.Fare
	c.JSON(http.StatusOK, models.GetFareResponseModel{
		ID:           fare.GetId(),
		Name:         fare.GetName(),
		DeliveryTime: fare.GetDeliveryTime(),
		PricePerKm:   fare.GetPricePerKm(),
		MinPrice:     fare.GetMinPrice(),
		CreatedAt:    fare.GetCreatedAt(),
		UpdatedAt:    fare.GetUpdatedAt(),
	})
}

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

	createdFare, err := h.grpcClient.FareService().Create(
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

	js, err := jspbMarshal.MarshalToString(createdFare.Fare)
	if err != nil {
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

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

	updatedFare, err := h.grpcClient.FareService().Update(
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

	js, err := jspbMarshal.MarshalToString(updatedFare.Fare)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

func (h *handlerV1) GetAllFares(c *gin.Context) {
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

	resp, err := h.grpcClient.FareService().GetAllFares(
		context.Background(),
		&pb.GetAllFaresRequest{
			Page:  uint64(page),
			Limit: uint64(pageSize),
		},
	)
	if handleGRPCErr(c, h.log, err) {
		return
	}

	generalResp := models.GetAllFaresResponseModel{Count: int(resp.GetCount())}

	for _, e := range resp.GetFares() {
		generalResp.Fares = append(generalResp.Fares, models.GetFareResponseModel{
			ID:        e.Id,
			CreatedAt: e.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, generalResp)
}

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
