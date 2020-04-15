package v1

import (
	"bitbucket.org/alien_soft/api_gateway/api/models"
	pbg "bitbucket.org/alien_soft/api_gateway/genproto/geo_service"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//@Summary Get Geozones
//@Description Get Geozones
//@Tags geo
//@Accept  json
//@Produce  json
//@Param parent_id query string false "Parent ID"
//@Param page query int false "PAGE"
//@Param limit query int false "PAGE"
//@Success 200 {object} models.GeozoneModel
//@Failure 404 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
//@Router /v1/geozones/ [get]
func (h *handlerV1) GetGeozones(c *gin.Context) {
	var (
		page, limit int
		parentID    string
		err         error
		hasError    bool
	)

	parentID = c.Query("parent_id")

	page, err = ParsePageQueryParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})

		h.log.Error("Error while parsing page", logger.Error(err))
		return
	}

	limit, err = ParseLimitQueryParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})

		h.log.Error("Error while parsing limit", logger.Error(err))
		return
	}

	geozones, err := h.grpcClient.GeoService().GetAllGeozones(
		context.Background(),
		&pbg.GetAllGeozonesRequest{
			ParentId: parentID,
			Page:     uint64(page),
			Limit:    uint64(limit),
		})

	hasError = handleGrpcErrWithMessage(c, h.log, err, "Error while getting geozones")

	if hasError {
		return
	}

	writeMessageAsJSON(c, h.log, geozones)
}
