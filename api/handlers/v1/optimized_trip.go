package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"github.com/gin-gonic/gin"
)

// @Router /v1/optimized-trip [post]
// @Summary Get Optimized Trip location with index
// @Description API for getting optimized trip
// @Tags geo
// @Accept  json
// @Produce  json
// @Param tripdata body models.TripsDataModel true "Current location"
// @Param tripdata body models.TripsDataModel true "Current location"
// @Success 200 {object} models.OptimizedTrips
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) OptimizedTrip(c *gin.Context) {
	var (
		tripDataModel = models.TripsDataModel{}
	)

	err := c.ShouldBindJSON(&tripDataModel)

	if handleGrpcErrWithMessage(c, h.log, err, "error while binding") {
		return
	}

	optimizedTrip := getOptimizedTrip(tripDataModel, h.cfg)

	c.JSON(200, optimizedTrip.Waypoints)
}
