package v1

import (
	"github.com/google/uuid"
	"github.com/gomodule/redigo/redis"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"net/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	"strings"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"github.com/gin-gonic/gin"
	pbu "genproto/user_service"
	pbs "genproto/sms_service"
)

// @Summary Register
// @Description Register - API for registering users
// @Tags register
// @Accept  json
// @Produce  json
// @Param register body models.RegisterModel true "register"
// @Success 200
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/register/ [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		reg  models.RegisterModel
		code string
	)

	err := c.ShouldBindJSON(&reg)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error binding json") {
		return
	}

	reg.Phone = strings.TrimSpace(reg.Phone)
	reg.Name = strings.TrimSpace(reg.Name)

	_, err = h.grpcClient.UserService().GetClient(
		context.Background(), &pbu.GetClientRequest{
			Id: reg.Phone,
		},
	)
	st, ok := status.FromError(err)
	if st.Code() != codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Phone already exists",
			},
		})
		h.log.Error("Error while checking phone, Already exists", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Service Error",
			},
		})
		h.log.Error("Error while checking phone", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		h.log.Error("Error while checking phone, unavailable", logger.Error(err))
		return
	}


	if h.cfg.Environment == "develop" {
		code = etc.GenerateCode(6, true)
	} else {
		code = etc.GenerateCode(6)
		_, err := h.grpcClient.SmsService().Send(
			context.Background(), &pbs.Sms{
				Text: "Your code for delever is " + code,
				Recipients: []string{reg.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(reg.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	key := reg.Phone + "name"
	err = h.inMemoryStorage.SetWithTTl(key, reg.Name, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.Status(http.StatusOK)
}

// @Security ApiKeyAuth
// @Summary Register confirm
// @Description Register - API for confirming and inserting user to DB
// @Tags register
// @Accept  json
// @Produce  json
// @Param register_confirm body models.RegisterConfirmModel true "register_confirm"
// @Success 200 {object} models.RegisterConfirmModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/register_confirm/ [post]
func (h *handlerV1) RegisterConfirm(c *gin.Context) {
	var rc models.RegisterConfirmModel

	err := c.ShouldBindJSON(&rc)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error binding json") {
		return
	}

	rc.ActivationCode = strings.TrimSpace(rc.ActivationCode)
	rc.Phone = strings.TrimSpace(rc.Phone)

	//Getting code from redis
	key := rc.Phone
	s, err := redis.String(h.inMemoryStorage.Get(key))
	if err != nil || s == "" {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Key does not exist", logger.Error(err))
		return
	}

	//Checking whether received code is valid
	if rc.ActivationCode != s {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "code is not valid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	//Getting name from redis
	key = rc.Phone + "name"
	name, err := redis.String(h.inMemoryStorage.Get(key))
	if err != nil || s == "" {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Key does not exist", logger.Error(err))
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	_, err = h.grpcClient.UserService().CreateClient(
		context.Background(), &pbu.Client{
			Id: id.String(),
			FirstName: name,
			Phone: rc.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating a client") {
		return
	}

	
}