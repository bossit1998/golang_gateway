package v1

import (
	"context"
	"net/http"

	pbo "genproto/order_service"
	pbr "genproto/report_service"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
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
func (h *handlerV1) GetBranchesReportExcel(c *gin.Context) {
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
func (h *handlerV1) GetCouriersReportExcel(c *gin.Context) {
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

// @Security ApiKeyAuth
// @Router /v1/reports/operators [get]
// @Summary Get Operators Report
// @Description API for getting operators report
// @Tags report
// @Accept  json
// @Produce  json
// @Param start_date query string true "start_date"
// @Param end_date query string true "end_date"
// @Success 200 {object} models.OperatorsReport
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOperatorsReport(c *gin.Context) {
	var (
		userInfo    models.UserInfo
		jspbMarshal jsonpb.Marshaler
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = helpers.ValidateDates(c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	reports, err := h.grpcClient.OrderService().GetOperatorsReport(context.Background(), &pbo.GetReportsRequest{
		ShipperId: userInfo.ShipperID,
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting operators report") {
		return
	}

	js, err := jspbMarshal.MarshalToString(reports)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/reports/branches [get]
// @Summary Get Branches Report
// @Description API for getting branches report
// @Tags report
// @Accept  json
// @Produce  json
// @Param start_date query string true "start_date"
// @Param end_date query string true "end_date"
// @Success 200 {object} models.BranchesReport
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBranchesReport(c *gin.Context) {
	var (
		userInfo    models.UserInfo
		jspbMarshal jsonpb.Marshaler
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = helpers.ValidateDates(c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	// getting branches report
	reports, err := h.grpcClient.OrderService().GetBranchesReport(
		context.Background(),
		&pbo.GetReportsRequest{
			ShipperId: userInfo.ShipperID,
			StartDate: c.Query("start_date"),
			EndDate:   c.Query("end_date"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(reports)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/reports/shipper [get]
// @Summary Get Shipper Report
// @Description API for getting shipper report
// @Tags report
// @Accept  json
// @Produce  json
// @Param start_date query string true "start_date"
// @Param end_date query string true "end_date"
// @Success 200 {object} models.ShipperReport
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetShipperReport(c *gin.Context) {
	var (
		userInfo    models.UserInfo
		jspbMarshal jsonpb.Marshaler
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = helpers.ValidateDates(c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err,
		})
		return
	}

	// getting all orders
	report, err := h.grpcClient.OrderService().GetShippersReport(
		context.Background(),
		&pbo.GetReportsRequest{
			ShipperId: userInfo.ShipperID,
			StartDate: c.Query("start_date"),
			EndDate:   c.Query("end_date"),
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while getting shipper report") {
		return
	}

	js, err := jspbMarshal.MarshalToString(report)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
