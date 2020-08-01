package v1

import (
	"context"
	pbu "genproto/user_service"
	"net/http"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Router /v1/shippers [post]
// @Summary Create Shipper
// @Description API for creating shipper
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param shipper body models.CreateShipperModel true "shipper"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateShipper(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		shipper       pbu.Shipper
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &shipper)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	err = helpers.ValidateLogin(shipper.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = helpers.ValidatePassword(shipper.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	passwordHash, err := etc.GeneratePasswordHash(shipper.Password)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while hashing password") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	accessToken, err := jwt.GenerateJWT(id.String(), "shipper", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	shipper.Id = id.String()
	shipper.Password = string(passwordHash)
	shipper.AccessToken = accessToken

	_, err = h.grpcClient.ShipperService().CreateShipper(
		context.Background(), &pbu.CreateShipperRequest{
			Shipper: &shipper,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating Shipper") {
		return
	}

	c.JSON(http.StatusOK, models.Response{
		ID: id.String(),
	})
}

// @Router /v1/shippers [put]
// @Summary Update Shipper
// @Description API for updating shipper
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param shipper body models.UpdateShipperModel true "shipper"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateShipper(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		shipper       pbu.Shipper
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &shipper)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while unmarshalling data", logger.Error(err))
		return
	}

	res, err := h.grpcClient.ShipperService().UpdateShipper(
		context.Background(),
		&pbu.UpdateShipperRequest{
			Shipper: &shipper,
		},
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating Shipper", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating Shipper, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetShipper())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags shipper
// @Router /v1/shippers/{shipper_id} [delete]
// @Summary Delete Shipper
// @Description API for deleting shipper
// @Accept  json
// @Produce  json
// @Param shipper_id path string true "shipper_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteShipper(c *gin.Context) {

	_, err := h.grpcClient.ShipperService().DeleteShipper(
		context.Background(),
		&pbu.DeleteShipperRequest{
			Id: c.Param("shipper_id"),
		},
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting Shipper", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting shipper, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting shipper, service unavailable", logger.Error(err))
		return
	}
	c.Status(http.StatusOK)
}

// @Tags shipper
// @Router /v1/shippers/{shipper_id} [get]
// @Summary Get Shipper
// @Description API for getting shipper info
// @Accept  json
// @Produce json
// @Param shipper_id path string true "shipper_id"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetShipper(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.ShipperService().GetShipper(
		context.Background(), &pbu.GetShipperRequest{
			Id: c.Param("shipper_id"),
		},
	)
	st, ok := status.FromError(err)
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Shipper Not Found",
			},
		})
		h.log.Error("Error while getting Shipper, Shipper Not Found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Server unavailable",
			},
		})
		h.log.Error("Error while getting Shipper, service unavailable", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting Shipper", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetShipper())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/shippers [get]
// @Summary Get All shippers
// @Description API for getting shippers
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllShippersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllShippers(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.ShipperService().GetAllShippers(
		context.Background(),
		&pbu.GetAllShippersRequest{
			Page:  uint64(page),
			Limit: uint64(limit),
		},
	)
	if handleGRPCErr(c, h.log, err) {
		return
	}
	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/shippers/login [POST]
// @Summary Check Shipper Login
// @Description API that checks whether shipper exists
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param login body models.ShipperLogin true "login"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ShipperLogin(c *gin.Context) {
	var (
		model   models.ShipperLogin
		isMatch = true
	)
	err := c.ShouldBindJSON(&model)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	shipper, err := h.grpcClient.ShipperService().GetByLogin(
		context.Background(),
		&pbu.Shipper{
			Username: model.Login,
		})

	if err != nil {
		isMatch = false
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(shipper.Password), []byte(model.Password))
		if err != nil {
			isMatch = false
		}
	}

	if !isMatch {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "login or password is incorrect",
			},
		})
		return
	}

	m := map[interface{}]interface{}{
		"user_type":  "shipper",
		"shipper_id": shipper.Id,
		"sub":        shipper.Id,
	}
	accessToken, _, err := jwt.GenJWT(m, signingKey)

	shipper.AccessToken = accessToken

	c.JSON(http.StatusOK, shipper)
}

// @Security ApiKeyAuth
// @Router /v1/shippers/change-password [POST]
// @Summary Change shipper password
// @Description API that change shipper password
// @Tags shipper
// @Accept  json
// @Produce  json
// @Param change_password body models.ShipperChangePassword true "change_password"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ChangePassword(c *gin.Context) {
	var (
		model    models.ShipperChangePassword
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	err = c.ShouldBindJSON(&model)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: err.Error(),
			},
		})
		return
	}

	passwordHash, err := etc.GeneratePasswordHash(model.Password)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while hashing password") {
		return
	}

	_, err = h.grpcClient.ShipperService().ChangePassword(
		context.Background(),
		&pbu.Shipper{
			Id:       userInfo.ID,
			Password: string(passwordHash),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while changing shipper password") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{Message: "successfully password changed"})
}
