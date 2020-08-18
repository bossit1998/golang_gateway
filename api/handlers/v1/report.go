package v1

import (
	"context"
	"net/http"
	"time"

	pbo "genproto/order_service"
	pbr "genproto/report_service"
	pbu "genproto/user_service"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
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
		userInfo  models.UserInfo
		reports   []models.OperatorReport
		layout    = "2006-01-02 15:04:05"
		startDate time.Time
		endDate   time.Time
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

	if startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
		})
		return
	}

	users, err := h.grpcClient.SystemUserService().GetAllSystemUsers(
		context.Background(),
		&pbu.GetAllSystemUsersRequest{
			ShipperId:  userInfo.ShipperID,
			UserRoleId: config.OperatorRoleId,
			Page:       1,
			Limit:      1000,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all operators") {
		return
	}

	// getting all orders
	orders, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		ShipperId: userInfo.ShipperID,
		StartTime: c.Query("start_date"),
		EndTime:   c.Query("end_date"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	// calculate hours between startTime and endTime
	diff := endDate.Sub(startDate)

	for _, user := range users.SystemUsers {
		var report models.OperatorReport
		for _, order := range orders.Orders {
			if order.CreatorId == user.Id {
				report.TotalOrdersCount++

				if order.Source == "bot" {
					report.BotOrdersCount++
				} else if order.Source == "admin_panel" {
					report.AdminPanelOrdersCount++
				} else if order.Source == "ios" || order.Source == "android" {
					report.AppOrdersCount++
				} else if order.Source == "website" {
					report.WebsiteOrdersCount++
				}
			}
		}
		report.AvgPerHour = float32(float64(report.TotalOrdersCount) / diff.Hours())
		reports = append(reports, report)
	}

	c.JSON(http.StatusOK, &models.OperatorsReport{
		Reports: reports,
	})
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
		userInfo  models.UserInfo
		reports   []models.BranchReport
		layout    = "2006-01-02 15:04:05"
		startDate time.Time
		endDate   time.Time
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

	if startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
		})
		return
	}

	respBranch, err := h.grpcClient.BranchService().GetAllBranches(
		context.Background(),
		&pbu.GetAllBranchesRequest{
			ShipperId: userInfo.ShipperID,
			Page:      1,
			Limit:     1000,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all operators") {
		return
	}

	// getting all orders
	respOrder, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		ShipperId: userInfo.ShipperID,
		StartTime: c.Query("start_date"),
		EndTime:   c.Query("end_date"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	for _, branch := range respBranch.Branches {
		var report models.BranchReport
		for _, order := range respOrder.Orders {
			if order.Steps[0].BranchId.Value == branch.Id {
				report.TotalCount++

				if order.Source == "bot" {
					report.BotOrdersCount++
				} else if order.Source == "admin_panel" {
					report.AdminPanelOrdersCount++
				} else if order.Source == "ios" || order.Source == "android" {
					report.AppOrdersCount++
				} else if order.Source == "website" {
					report.WebsiteOrdersCount++
				}

				report.TotalSum += order.OrderAmount

				if order.PaymentType == "cash" {
					report.TotalSumCash += order.OrderAmount
				} else if order.PaymentType == "payme" {
					report.TotalSumPayme += order.OrderAmount
				} else if order.PaymentType == "click" {
					report.TotalSumClick += order.OrderAmount
				}
				report.TotalSum += order.OrderAmount
			}
		}

		reports = append(reports, report)
	}

	c.JSON(http.StatusOK, &models.BranchesReport{
		Reports: reports,
	})
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
		userInfo  models.UserInfo
		report    models.ShipperReport
		layout    = "2006-01-02 15:04:05"
		startDate time.Time
		endDate   time.Time
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

	if startDate.Before(endDate) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: "start_time can not be greater than end_time",
		})
		return
	}

	// getting all orders
	respOrder, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		ShipperId: userInfo.ShipperID,
		StartTime: c.Query("start_date"),
		EndTime:   c.Query("end_date"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	report.TotalOrdersCount = respOrder.Count
	for _, order := range respOrder.Orders {
		if order.Source == "bot" {
			report.BotOrdersCount++
		} else if order.Source == "admin_panel" {
			report.AdminPanelOrdersCount++
		} else if order.Source == "ios" || order.Source == "android" {
			report.AppOrdersCount++
		} else if order.Source == "website" {
			report.WebsiteOrdersCount++
		}

		if order.PaymentType == "cash" {
			report.TotalSumCash += order.OrderAmount
		} else if order.PaymentType == "payme" {
			report.TotalSumPayme += order.OrderAmount
		} else if order.PaymentType == "click" {
			report.TotalSumClick += order.OrderAmount
		}
		report.TotalSum += order.OrderAmount
	}

	c.JSON(http.StatusOK, report)
}
