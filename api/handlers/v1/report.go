package v1

import (
	"context"
	"net/http"

	pbr "genproto/report_service"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Security ApiKeyAuth
// @Router /v1/branches-report-excel/ [get]
// @Summary Get Branches Report Excel
// @Description API for getting branches report excel
// @Tags report
// @Accept  json
// @Produce  json
// @Param date query string false "date"
// @Success 200 {object} models.GetReportModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBranchesReport(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.ReportService().GetBranchesReportExcel(
		context.Background(),
		&pbr.GetBranchesReportExcelRequest{
			ShipperId: userInfo.ShipperID,
			Date:      c.Query("date"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting branches report") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/couriers-report-excel/ [get]
// @Summary Get Couriers Report Excel
// @Description API for getting couriers report excel
// @Tags report
// @Accept  json
// @Produce  json
// @Param date query string false "date"
// @Success 200 {object} models.GetReportModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCouriersReport(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.ReportService().GetCouriersReportExcel(
		context.Background(),
		&pbr.GetCouriersReportExcelRequest{
			ShipperId: userInfo.ShipperID,
			Date:      c.Query("date"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting couriers report") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
