package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	pbo "genproto/order_service"
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
		layout      = "2006-01-02 15:04:05"
		startDate   time.Time
		endDate     time.Time
		model       models.OperatorsReport
		jspbMarshal jsonpb.Marshaler
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	startDate, err = time.Parse(layout, c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_date is invalid",
		})
		return
	}

	endDate, err = time.Parse(layout, c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "end_time is invalid",
		})
		return
	}

	if !startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
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

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
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
		model       []models.BranchReport
		layout      = "2006-01-02 15:04:05"
		startDate   time.Time
		endDate     time.Time
		jspbMarshal jsonpb.Marshaler
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	startDate, err = time.Parse(layout, c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_date is invalid",
		})
		return
	}

	endDate, err = time.Parse(layout, c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "end_time is invalid",
		})
		return
	}

	if !startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
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

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
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
		model       models.ShipperReport
		layout      = "2006-01-02 15:04:05"
		startDate   time.Time
		endDate     time.Time
		jspbMarshal jsonpb.Marshaler
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	startDate, err = time.Parse(layout, c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_date is invalid",
		})
		return
	}

	endDate, err = time.Parse(layout, c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "end_time is invalid",
		})
		return
	}

	if !startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
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

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}
