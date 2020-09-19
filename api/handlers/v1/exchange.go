package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// @Router /v1/exchange-rates [get]
// @Summary Get Exchange rates
// @Description API for getting exchange rates
// @Tags exchange
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetExchangeRatesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetExchangeRate(c *gin.Context) {
	resp, err := http.Get("https://nbu.uz/en/exchange-rates/json/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: ErrorCodeInternal,
		})
		h.log.Error("error while making a http request", logger.Error(err))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: ErrorCodeInternal,
		})
		h.log.Error("error while making a http request", logger.Error(err))
		return
	}
  c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(body))
}
