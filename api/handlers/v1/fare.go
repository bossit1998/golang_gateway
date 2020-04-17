package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/alien_soft/api_gateway/api/models"
	pb "bitbucket.org/alien_soft/api_gateway/genproto/fare_service"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*// @Summary Get Fare
// @Description Get Fare API returns Fare
// @Tags Fare
// @Accept  json
// @Produce  json
// @Param id path string true "Fare Id"
// @Success 200 {object} models.GetFareResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/fares/{id}/ [get]*/
func (h *handlerV1) GetFare(c *gin.Context) {

	fareResponse, err := h.grpcClient.FareService().GetFare(
		context.Background(), &pb.GetFareRequest{
			Id: c.Param("id"),
		},
	)
	fmt.Println(fareResponse)
	if handleGRPCErr(c, h.log, err) {
		return
	}

	if fareResponse == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Event Not Found",
			},
		})
		return
	}
	fare := fareResponse.Fare
	c.JSON(http.StatusOK, models.GetFareResponseModel{
		ID:   fare.ID,
		Name: fare.Name,
		// DeliveryTime: fare.DeliveryTime,
		// PricePerKm:   fare.PricePerKm,
		// MinPrice:     fare.MinPrice,
		CreatedAt: fare.CreatedAt,
	})

}

/*// @Summary Create fare
// @Description Get Profile API creates fare
// @Tags fare
// @Accept  json
// @Produce  json
// @Param fare body models.CreateFareRequestModel true "createFare"
// @Failure 500 {object} models.ResponseError
// @Router /v1/fares/ [POST]*/
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

/*// @Summary Update fare
// @Description  updates fares
// @Tags fare
// @Accept  json
// @Produce  json
// @Param  models.CreateFareRequestModel true "updateFare"
// @Failure 500 {object} models.ResponseError
// @Router /v1/fares/ [PUT]*/
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

	_, err = h.grpcClient.FareService().Update(
		context.Background(),
		&fare,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while updating fare") {
		return
	}
	c.JSON(http.StatusOK, "")
}

/*// @Summary Get All Fares
// @Tags fare
// @Produce  json
// @Param
// @Param page query string false "page"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetAllFareResponseModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/fares/ [GET]*/
func (h *handlerV1) GetAllFares(c *gin.Context) {
	var (
		jspbMarshal           jsonpb.Marshaler
		pageValue, limitValue string
		page, limit           uint64
		err                   error
	)

	pageValue = c.Query("page")
	limitValue = c.Query("limit")
	if pageValue == "" {
		page = 1
	} else {
		page, err = strconv.ParseUint(pageValue, 10, 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: models.InternalServerError{
					Code:    ErrorCodeInvalidURL,
					Message: "Invalid query",
				},
			})
			h.log.Error("Error while parsing page", logger.Error(err))
			return
		}
	}

	if page == 0 {
		page = 1
	}

	if limitValue == "" {
		limit = 10
	} else {
		limit, err = strconv.ParseUint(limitValue, 10, 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, models.ResponseError{
				Error: models.InternalServerError{
					Code:    ErrorCodeInvalidURL,
					Message: "Invalid query",
				},
			})
			h.log.Error("Error while parsing limit", logger.Error(err))
			return
		}
	}

	if limit == 0 {
		limit = 10
	}

	jspbMarshal.OrigName = true

	fares, err := h.grpcClient.FareService().GetAllFares(
		context.Background(),
		&pb.GetAllFaresRequest{
			Limit: limit,
			Page:  page,
		})

	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting all fare", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while getting all fares, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting  all fares, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(fares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while marshalling", logger.Error(err))
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)

}

/*// @Summary DeleteFare
// @Description DeleteFare API is for deleting fare
// @Tags fare
// @Accept json
// @Produce json
// @Param  body models.DeleteFareModel true "delete fare"
// @Success 200
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/fare/delete_fare/ [delete]*/
func (h *handlerV1) DeleteFare(c *gin.Context) {
	var (
		deleteFare models.DeleteFareModel
	)
	err := c.ShouldBindJSON(&deleteFare)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidJSON,
				Message: "Invalid Json",
			},
		})
		h.log.Error("Error binding json", logger.Error(err))
		return
	}
	_, err = h.grpcClient.FareService().Delete(
		context.Background(),
		&pb.DeleteFareRequest{
			Id: deleteFare.ID,
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
