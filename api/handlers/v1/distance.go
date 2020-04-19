

package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

// @Router /v1/getdistance/from/{from_lat}/{from_long}/to/{to_lat}/{to_long} [get]
// @Summary Get Distance
// @Description API for getting distance
// @Tags geo
// @Accept  json
// @Produce  json
// @Param from_lat path string true "from_lat"
// @Param from_long path string true "from_long"
// @Param to_lat path string true "to_lat"
// @Param to_long path string true "to_long"
// @Success 200 {object} models.GetDistanceResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetDistance(c *gin.Context) {
	var location models.Coordinate
	token:= "pk.eyJ1IjoidGRvc3RvbiIsImEiOiJjazh0cmRrMnowMWszM29sc2Y5c3A5NTZ4In0.mtrOXD4cD4QKZ-dnZ_vKdA"
	params := c.Params

	if s, err := strconv.ParseFloat(params.ByName("from_lat"), 64); err == nil {
		location.FromLat = s
	}
	if s, err := strconv.ParseFloat(params.ByName("from_long"), 64); err == nil {
		location.FromLong = s
	}
	if s, err := strconv.ParseFloat(params.ByName("to_lat"), 64); err == nil {
		location.ToLat = s
	}
	if s, err := strconv.ParseFloat(params.ByName("to_long"), 64); err == nil {
		location.ToLong = s
	}

	distance := getDistance(location, token)

	c.JSON(200, gin.H{"distance": distance})
}




func getDistance(location models.Coordinate , token string) float64{
	coordinates := fmt.Sprintf("%f,%f;%f,%f", location.FromLat, location.FromLong, location.ToLat, location.ToLong)
	url := "https://api.mapbox.com/directions/v5/mapbox/driving/"+ coordinates +"/?approaches=unrestricted;curb&access_token="+ token + ""
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	geodriving := models.GeoDrivingAPIResponse{}
	json.Unmarshal([]byte(body), &geodriving)

	var dist float64
	dist = geodriving.RoutesList[0].LegsList[0].Distance
	fmt.Print(dist)
	return dist
}