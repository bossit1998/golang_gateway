package v1

import (
	pbo "bitbucket.org/alien_soft/api_gateway/genproto/order_service"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
)

func (h *handlerV1) Create(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order pbo.Order
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	_, err = h.grpcClient.OrderService().Create(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	c.JSON(200, "")
}
