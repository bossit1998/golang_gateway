package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
)


func (h *handlerV1) GetTotalDeliveryCost(c *gin.Context) {

	var(
	LimitDistance float64
	InitialPrice float64
	UnitPrice float64
	Distance float64
	)
	params := c.Params

	if s, err := strconv.ParseFloat(params.ByName("limit_distance"), 64); err == nil {
		LimitDistance = s
	}
	if s, err := strconv.ParseFloat(params.ByName("initial_price"), 64); err == nil {
		InitialPrice = s
	}
	if s, err := strconv.ParseFloat(params.ByName("unit_price"), 64); err == nil {
		UnitPrice = s
	}
	if s, err := strconv.ParseFloat(params.ByName("distance"), 64); err == nil {
		Distance = s
	}

	//distance := getDistance(location, token)
	//totalDeliveryCost := calcDeliveryCost(3000, 5000, 500, 9000)
	totalDeliveryCost := calcDeliveryCost(LimitDistance, InitialPrice, UnitPrice, Distance)

	c.JSON(200, gin.H{"total_delivery_cost": totalDeliveryCost})
}



func calcDeliveryCost(limit_distance float64, inital_price float64, unit_price float64, distance float64) float64{

	total_delivery_cost := 0.0

	if distance < limit_distance {
		total_delivery_cost = inital_price
	} else {
		total_delivery_cost = inital_price + distance*unit_price/1000
	}
	return total_delivery_cost
}


