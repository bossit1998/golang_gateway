package v1

import (
	"context"
	"fmt"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage/redis"
)

// @Router /v1/users [post]
// @Summary Create User
// @Description API for creating user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.CreateUserModel true "user"
// @Success 200 {object} models.GetUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateClient(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		client        pbu.Client
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &client)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
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

	client.Id = id.String()
	client.AccessToken = accessToken

	res, err := h.grpcClient.UserService().CreateClient(
		context.Background(), &client,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating user") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Client)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/users [put]
// @Summary Update User
// @Description API for updating user
// @Tags user
// @Accept  json
// @Produce  json
// @Param courier body models.UpdateUserModel true "user"
// @Success 200 {object} models.GetUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateClient(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		client        pbu.Client
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &client)
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

	res, err := h.grpcClient.UserService().UpdateClient(
		context.Background(),
		&client,
	)
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating user", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating user, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetClient())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags user
// @Router /v1/users/{user_id} [delete]
// @Summary Delete User
// @Description API for deleting user
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteClient(c *gin.Context) {

	_, err := h.grpcClient.UserService().DeleteClient(
		context.Background(),
		&pbu.DeleteClientRequest{
			Id: c.Param("user_id"),
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
		h.log.Error("Error while deleting user", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting user, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting user, service unavailable", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Tags user
// @Router /v1/users/{user_id} [get]
// @Summary Get User
// @Description API for getting user info
// @Accept  json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} models.GetUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetClient(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	fmt.Println()
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.UserService().GetClient(
		context.Background(), &pbu.GetClientRequest{
			Id: c.Param("user_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "User Not Found",
			},
		})
		return
	}
	js, err := jspbMarshal.MarshalToString(res.GetClient())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/users [get]
// @Summary Get All Users
// @Description API for getting users
// @Tags user
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllUsersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllClients(c *gin.Context) {
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

	pageSize, err := ParsePageSizeQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.UserService().GetAllClients(
		context.Background(),
		&pbu.GetAllClientsRequest{
			Page:  uint64(page),
			Limit: uint64(pageSize),
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

// @Router /v1/users/check-login/ [POST]
// @Summary Check User Login
// @Description API that checks whether user exists
// @Description and if exists sends sms to their number
// @Tags user
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckUserLoginRequest true "check login"
// @Success 200 {object} models.CheckUserLoginResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckUserLogin(c *gin.Context) {
	var (
		checkUserLoginModel models.CheckUserLoginRequest
		code            string
	)
	
	err := c.ShouldBindJSON(&checkUserLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	checkUserLoginModel.Phone = strings.TrimSpace(checkUserLoginModel.Phone)

	resp, err := h.grpcClient.UserService().ExistsClient(
		context.Background(), &pbu.ExistsClientRequest{
			Phone: checkUserLoginModel.Phone,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking user") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "User not found",
			},
		})
		h.log.Error("Error while checking phone, doesn't exist", logger.Error(err))
		return
	}

	if h.cfg.Environment == "develop" {
		code = etc.GenerateCode(6, true)
	} else {
		code = etc.GenerateCode(6)
		_, err = h.grpcClient.SmsService().Send(
			context.Background(), &pbs.Sms{
				Text:       code,
				Recipients: []string{checkUserLoginModel.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(checkUserLoginModel.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.JSON(http.StatusOK, models.CheckUserLoginResponse{
		Code:  code,
		Phone: checkUserLoginModel.Phone,
	})
}

// @Router /v1/users/confirm-login/ [POST]
// @Summary Confirm User Login
// @Description API that checks whether user entered
// @Description valid token
// @Tags user
// @Accept  json
// @Produce  json
// @Param confirm_phone body models.ConfirmUserLoginRequest true "confirm login"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmUserLogin(c *gin.Context) {
	var (
		cm models.ConfirmUserLoginRequest
	)

	err := c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	cm.Code = strings.TrimSpace(cm.Code)

	//Getting code from redis
	key := cm.Phone
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
	if cm.Code != s {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "Code is invalid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	user, err := h.grpcClient.UserService().GetClient(
		context.Background(), &pbu.GetClientRequest{
			Id: cm.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting client") {
		return
	}

	c.JSON(http.StatusOK, &models.ConfirmUserLoginResponse{
		ID:          user.Client.Id,
		AccessToken: user.Client.AccessToken,
	})
}
