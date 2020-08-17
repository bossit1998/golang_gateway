package v1

import (
	"context"
	"net/http"

	pbo "genproto/order_service"
	pbr "genproto/report_service"
	pbu "genproto/user_service"

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

func (h *handlerV1) GetCustomersReport(c *gin.Context) {
}

func (h *handlerV1) GetOperatorsReport(c *gin.Context) {
	var (
		err      error
		userInfo models.UserInfo
		// reports models.GetAllOperatorsReportModel
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	users, err := h.grpcClient.SystemUserService().GetAllSystemUsers(
		context.Background(),
		&pbu.GetAllSystemUsersRequest{
			// ShipperId: userInfo.ShipperID,
			Page:  1,
			Limit: 1000,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all operators") {
		return
	}

	// getting all orders
	orders, err := h.grpcClient.OrderService().GetAll(context.Background(), &pbo.OrdersRequest{
		ShipperId: userInfo.ShipperID,
		StartTime: c.Query("start_time"),
		EndTime:   c.Query("end_time"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all order") {
		return
	}

	for _, user := range users.SystemUsers {
		for _, order := range orders.Orders {
			if order.CreatorId == user.Id {

			}
		}
	}

	c.JSON(http.StatusOK, "ok")
}
