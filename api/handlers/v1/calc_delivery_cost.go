package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"context"
	"fmt"
	pbf "genproto/fare_service"
	"github.com/gin-gonic/gin"
)

// @Router /v1/calc-delivery-cost [post]
// @Summary Calculate Delivery Price For Clients
// @Description API for getting total delivery cost
// @Tags geo
// @Accept  json
// @Produce  json
// @Param calc body models.CalcDeliveryCostRequest true "calc-delivery-cost"
// @Success 200 {object} models.CalcDeliveryCostResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CalcDeliveryCost(c *gin.Context) {
	var (
		calcCostModel models.CalcDeliveryCostRequest
	)
	err := c.ShouldBindJSON(&calcCostModel)

	if handleGrpcErrWithMessage(c, h.log, err, "error while binding") {
		return
	}

	distance := getDistance(calcCostModel.FromLocation, calcCostModel.ToLocation, h.cfg)

	totalDeliveryCost := calcDeliveryCost(calcCostModel.MinDistance, calcCostModel.MinPrice, calcCostModel.PerKmPrice, distance)

	c.JSON(200, models.CalcDeliveryCostResponse{
		Distance: distance,
		Price: totalDeliveryCost,
	})
}

func calcDeliveryCost(limitDistance float64, initialPrice float64, unitPrice float64, distance float64) float64{
	totalDeliveryCost := 0.0
	fmt.Println(distance, limitDistance)

	if distance < limitDistance {
		totalDeliveryCost = initialPrice
	} else {
		price := int((distance - limitDistance) * unitPrice / 1000)
		price = (price/100) * 100
		totalDeliveryCost = initialPrice + float64(price)
	}
	return totalDeliveryCost
}

func calcDeliveryPriceWithFare(c *gin.Context, h *handlerV1, fareID string, fromLocation, toLocation models.Location) (float64, error) {
	fare, err := h.grpcClient.FareService().GetFare(context.Background(),
		&pbf.GetFareRequest{
			Id:fareID,
		})

	if err != nil {
		return 0, err
	}

	distance := getDistance(fromLocation, toLocation, h.cfg)

	price := calcDeliveryCost(float64(fare.Fare.MinDistance), float64(fare.Fare.MinPrice), float64(fare.Fare.PricePerKm), distance)

	return price, nil
}

