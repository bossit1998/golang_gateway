package v1

import (
	"context"
	pbu "genproto/user_service"
	"net/http"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

// @Router /v1/system-users [post]
// @Summary Create SystemUser
// @Description API for creating systemUser
// @Tags systemUser
// @Accept  json
// @Produce  json
// @Param systemUser body models.CreateSystemUserModel true "systemUser"
// @Success 200 {object} models.GetSystemUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateSystemUser(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		systemUser    pbu.SystemUser
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &systemUser)
	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	err = helpers.ValidateLogin(systemUser.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = helpers.ValidatePassword(systemUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	passwordHash, err := etc.GeneratePasswordHash(systemUser.Password)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while hashing password") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	// accessToken, err := jwt.GenerateJWT(id.String(), "systemUser", signingKey)
	// if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
	// 	return
	// }

	systemUser.Id = id.String()
	systemUser.Password = string(passwordHash)
	// systemUser.AccessToken = accessToken
	// systemUser.UserRoleId = id.String()

	_, err = h.grpcClient.SystemUserService().CreateSystemUser(
		context.Background(), &pbu.CreateSystemUserRequest{
			SystemUser: &systemUser,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating SystemUser") {
		return
	}

	c.JSON(http.StatusOK, models.Response{
		ID: id.String(),
	})
}

// @Router /v1/system-users [put]
// @Summary Update SystemUser
// @Description API for updating systemUser
// @Tags systemUser
// @Accept  json
// @Produce  json
// @Param systemUser body models.UpdateSystemUserModel true "systemUser"
// @Success 200 {object} models.GetSystemUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateSystemUser(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		systemUser    pbu.SystemUser
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &systemUser)
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

	res, err := h.grpcClient.SystemUserService().UpdateSystemUser(
		context.Background(),
		&pbu.UpdateSystemUserRequest{
			SystemUser: &systemUser,
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
		h.log.Error("Error while updating SystemUser", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating SystemUser, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetSystemUser())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags systemUser
// @Router /v1/system-users/{systemUser_id} [delete]
// @Summary Delete SystemUser
// @Description API for deleting systemUser
// @Accept  json
// @Produce  json
// @Param systemUser_id path string true "systemUser_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteSystemUser(c *gin.Context) {

	_, err := h.grpcClient.SystemUserService().DeleteSystemUser(
		context.Background(),
		&pbu.DeleteSystemUserRequest{
			Id: c.Param("systemUser_id"),
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
		h.log.Error("Error while deleting SystemUser", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting systemUser, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting systemUser, service unavailable", logger.Error(err))
		return
	}
	c.Status(http.StatusOK)
}

// @Tags systemUser
// @Router /v1/system-users/{systemUser_id} [get]
// @Summary Get SystemUser
// @Description API for getting systemUser info
// @Accept  json
// @Produce json
// @Param systemUser_id path string true "systemUser_id"
// @Success 200 {object} models.GetSystemUserModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetSystemUser(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.SystemUserService().GetSystemUser(
		context.Background(), &pbu.GetSystemUserRequest{
			Id: c.Param("systemUser_id"),
		},
	)
	st, ok := status.FromError(err)
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "SystemUser Not Found",
			},
		})
		h.log.Error("Error while getting SystemUser, SystemUser Not Found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Server unavailable",
			},
		})
		h.log.Error("Error while getting SystemUser, service unavailable", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting SystemUser", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetSystemUser())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/system-users [get]
// @Summary Get All systemUsers
// @Description API for getting systemUsers
// @Tags systemUser
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllSystemUsersModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllSystemUsers(c *gin.Context) {
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

	res, err := h.grpcClient.SystemUserService().GetAllSystemUsers(
		context.Background(),
		&pbu.GetAllSystemUsersRequest{
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

// // @Router /v1/system-users/login [POST]
// // @Summary Check SystemUser Login
// // @Description API that checks whether systemUser exists
// // @Tags systemUser
// // @Accept  json
// // @Produce  json
// // @Param login body models.SystemUserLogin true "login"
// // @Success 200 {object} models.GetSystemUserModel
// // @Failure 404 {object} models.ResponseError
// // @Failure 500 {object} models.ResponseError
// func (h *handlerV1) SystemUserLogin(c *gin.Context) {
// 	var (
// 		model   models.SystemUserLogin
// 		isMatch = true
// 	)
// 	err := c.ShouldBindJSON(&model)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: models.InternalServerError{
// 				Code:    ErrorBadRequest,
// 				Message: err.Error(),
// 			},
// 		})
// 		return
// 	}

// 	systemUser, err := h.grpcClient.SystemUserService().GetByLogin(
// 		context.Background(),
// 		&pbu.SystemUser{
// 			Username: model.Login,
// 		})

// 	if err != nil {
// 		isMatch = false
// 	} else {
// 		err = bcrypt.CompareHashAndPassword([]byte(systemUser.Password), []byte(model.Password))
// 		if err != nil {
// 			isMatch = false
// 		}
// 	}

// 	if !isMatch {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: models.InternalServerError{
// 				Code:    ErrorBadRequest,
// 				Message: "login or password is incorrect",
// 			},
// 		})
// 		return
// 	}

// 	m := map[interface{}]interface{}{
// 		"user_type":     "systemUser",
// 		"systemUser_id": systemUser.Id,
// 		"sub":           systemUser.Id,
// 	}
// 	accessToken, _, err := jwt.GenJWT(m, signingKey)

// 	systemUser.AccessToken = accessToken

// 	c.JSON(http.StatusOK, systemUser)
// }

// @Security ApiKeyAuth
// @Router /v1/system-users/change-password [POST]
// @Summary Change systemUser password
// @Description API that change systemUser password
// @Tags systemUser
// @Accept  json
// @Produce  json
// @Param change_password body models.SystemUserChangePassword true "change_password"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ChangeSystemUserPassword(c *gin.Context) {
	var (
		model    models.SystemUserChangePassword
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

	_, err = h.grpcClient.SystemUserService().ChangePassword(
		context.Background(),
		&pbu.SystemUser{
			Id:       userInfo.ID,
			Password: string(passwordHash),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while changing systemUser password") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{Message: "successfully password changed"})
}
