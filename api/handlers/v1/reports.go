package v1

import (
	"context"
	pbu "genproto/user_service"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Security ApiKeyAuth
// @Router /v1/couriers/{courier_id} [get]
// @Summary Get Courier
// @Description API for getting courier
// @Tags courier
// @Accept  json
// @Produce  json
// @Param courier_id path string true "courier_id"
// @Success 200 {object} models.GetCourierModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomersReport(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	_, err := h.grpcClient.CustomerService().GetAllCustomers(
		context.Background(), &pbu.GetAllCustomersRequest{},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}
}
