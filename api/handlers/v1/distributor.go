package v1

import (
	"context"
	"fmt"
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
// @Summary Get Distributor
// @Description Get Distributor API returns event
// @Tags distributor
// @Accept  json
// @Produce  json
// @Param distributor_id path string true "distributor Id"
// @Success 200 {object} models.GetDistributorResp
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/couirer/{distributor_id}/ [get]
func (h *handlerV1) GetDistributor(c *gin.Context) {

	distributorResp, err := h.grpcClient.DistributorService().GetDistributor(
		context.Background(), &pbc.GetDistributorRequest{
			Id: c.Param("id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if distributorResp == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not Found",
			},
		})
		return
	}
	distributor := distributorResp.Distributor

	c.JSON(http.StatusOK, models.GetDistributorResponseModel{
		ID:        distributor.Id,
		Phone:     distributor.Phone,
		Name:      distributor.Name,
		CreatedAt: distributor.CreatedAt,
	})
}

func (h *handlerV1) GetAllDistributors(c *gin.Context) {

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
	fmt.Print("keldi")
	resp, err := h.grpcClient.DistributorService().GetAllDistributors(
		context.Background(),
		&pbc.GetAllDistributorsRequest{
			Page:  uint64(page),
			Limit: uint64(pageSize),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	generalResp := models.GetAllDistributorsResponseModel{Count: int(resp.GetCount())}

	for _, e := range resp.GetDistributors() {
		generalResp.Distributors = append(generalResp.Distributors, models.GetDistributorResponseModel{
			ID:        e.Id,
			Phone:     e.Phone,
			Name:      e.Name,
			CreatedAt: e.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, generalResp)
}

func (h *handlerV1) CreateDistributor(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		distributor   pbc.Distributor
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &distributor)
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

	createdDistributor, err := h.grpcClient.DistributorService().Create(
		context.Background(),
		&distributor,
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

	js, err := jspbMarshal.MarshalToString(createdDistributor.Distributor)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

func (h *handlerV1) UpdateDistributor(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		distributor   pbc.Distributor
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &distributor)
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
	distributor.Id = c.Param("id")
	d, err := h.grpcClient.DistributorService().Update(
		context.Background(),
		&distributor,
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

	js, err := jspbMarshal.MarshalToString(d.Distributor)
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

func (h *handlerV1) DeleteDistributor(c *gin.Context) {

	_, err := h.grpcClient.DistributorService().Delete(
		context.Background(),
		&pbc.DeleteDistributorRequest{
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
		h.log.Error("Error while deleting event", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting event, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting distributor, service unavailable", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}
