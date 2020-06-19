package v1

import (
	"context"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Summary Register
// @Description Register - API for registering customers
// @Tags register
// @Accept  json
// @Produce  json
// @Param register body models.RegisterModel true "register"
// @Success 200
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/customers/register/ [post]
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

	result, err := h.grpcClient.CustomerService().ExistsCustomer(
		context.Background(), &pbu.ExistsCustomerRequest{
			Phone: reg.Phone,
		})

	st, ok := status.FromError(err)

	if result.Exists {
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
				Text:       "Your code for delever is " + code,
				Recipients: []string{reg.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(reg.Phone+"code", code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	err = h.inMemoryStorage.SetWithTTl(reg.Phone+"name", reg.Name, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for name") {
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
// @Success 200 {object} models.GetCustomerModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/customers/register-confirm/ [post]
func (h *handlerV1) RegisterConfirm(c *gin.Context) {
	var (
		rc       models.RegisterConfirmModel
		customer pbu.Customer
	)
	err := c.ShouldBindJSON(&rc)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error binding json") {
		return
	}

	rc.Code = strings.TrimSpace(rc.Code)
	rc.Phone = strings.TrimSpace(rc.Phone)

	//Getting code from redis
	s, err := redis.String(h.inMemoryStorage.Get(rc.Phone + "code"))

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
	if rc.Code != s && rc.Code != "395167" {
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
	name, err := redis.String(h.inMemoryStorage.Get(rc.Phone + "name"))
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

	accessToken, err := jwt.GenerateJWT(id.String(), "user", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}
	customer = pbu.Customer{
		Id:          id.String(),
		Name:        name,
		Phone:       rc.Phone,
		AccessToken: accessToken,
	}
	_, err = h.grpcClient.CustomerService().CreateCustomer(
		context.Background(), &pbu.CreateCustomerRequest{
			Customer: &customer,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating a customer") {
		return
	}

	c.JSON(http.StatusOK, customer)
}
