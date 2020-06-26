package v1

import (
	"context"
	"fmt"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"
	"net/http"
	"strconv"
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

// @Security ApiKeyAuth
// @Router /v1/branches [post]
// @Summary Create Branch
// @Description API for creating branch
// @Tags branch
// @Accept  json
// @Produce  json
// @Param branch body models.CreateBranchModel true "branch"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateBranch(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		branch        pbu.Branch
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &branch)

	if handleInternalWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	id, err := uuid.NewRandom()
	if handleInternalWithMessage(c, h.log, err, "Error while generating UUID") {
		return
	}

	accessToken, err := jwt.GenerateJWT(id.String(), "branch", signingKey)
	if handleInternalWithMessage(c, h.log, err, "Error while generating access token") {
		return
	}

	branch.Id = id.String()
	branch.ShipperId = userInfo.ShipperID
	branch.AccessToken = accessToken

	res, err := h.grpcClient.BranchService().CreateBranch(
		context.Background(), &pbu.CreateBranchRequest{
			Branch: &branch,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating branch") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res.Branch)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/branches [put]
// @Summary Update Branch
// @Description API for updating branch
// @Tags branch
// @Accept  json
// @Produce  json
// @Param branch body models.UpdateBranchModel true "branch"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateBranch(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		branch        pbu.Branch
	)

	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &branch)
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

	res, err := h.grpcClient.BranchService().UpdateBranch(
		context.Background(),
		&pbu.UpdateBranchRequest{
			Branch: &branch,
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
		h.log.Error("Error while updating branch", logger.Error(err))
		return
	}
	if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while updating branch, service unavailable", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetBranch())
	if err != nil {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags branch
// @Router /v1/branches/{branch_id} [delete]
// @Summary Delete Branch
// @Description API for deleting branch
// @Accept  json
// @Produce  json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteBranch(c *gin.Context) {

	_, err := h.grpcClient.BranchService().DeleteBranch(
		context.Background(),
		&pbu.DeleteBranchRequest{
			Id: c.Param("branch_id"),
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
		h.log.Error("Error while deleting branch", logger.Error(err))
		return
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		h.log.Error("Error while deleting branch, not found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while deleting branch, service unavailable", logger.Error(err))
		return
	}
	c.Status(http.StatusOK)
}

// @Tags branch
// @Router /v1/branches/{branch_id} [get]
// @Summary Get Branch
// @Description API for getting branch info
// @Accept  json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBranch(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	res, err := h.grpcClient.BranchService().GetBranch(
		context.Background(), &pbu.GetBranchRequest{
			Id: c.Param("branch_id"),
		},
	)
	st, ok := status.FromError(err)
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Branch Not Found",
			},
		})
		h.log.Error("Error while getting branch, Branch Not Found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Server unavailable",
			},
		})
		h.log.Error("Error while getting branch, service unavailable", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting branch", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res.GetBranch())

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/branches [get]
// @Summary Get All Branches
// @Description API for getting branches
// @Tags branch
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllBranchesModel
// @Failure 404 {object} models.ResponseError

// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllBranches(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

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

	res, err := h.grpcClient.BranchService().GetAllBranches(
		context.Background(),
		&pbu.GetAllBranchesRequest{
			ShipperId: userInfo.ShipperID,
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

// @Router /v1/branches/check-login/ [POST]
// @Summary Check Branch Login
// @Description API that checks whether branch exists
// @Description and if exists sends sms to their number
// @Tags branch
// @Accept  json
// @Produce  json
// @Param check_login body models.CheckBranchLoginRequest true "check login"
// @Success 200 {object} models.CheckBranchLoginResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckBranchLogin(c *gin.Context) {
	var (
		checkBranchLoginModel models.CheckBranchLoginRequest
		code                  string
	)

	err := c.ShouldBindJSON(&checkBranchLoginModel)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	checkBranchLoginModel.Phone = strings.TrimSpace(checkBranchLoginModel.Phone)

	resp, err := h.grpcClient.BranchService().ExistsBranch(
		context.Background(), &pbu.ExistsBranchRequest{
			Phone: checkBranchLoginModel.Phone,
		},
	)
	if handleStorageErrWithMessage(c, h.log, err, "Error while checking branch") {
		return
	}

	if !resp.Exists {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Branch not found",
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
				Recipients: []string{checkBranchLoginModel.Phone},
			},
		)
		if handleGrpcErrWithMessage(c, h.log, err, "Error while sending sms") {
			return
		}
	}

	err = h.inMemoryStorage.SetWithTTl(checkBranchLoginModel.Phone, code, 1800)
	if handleInternalWithMessage(c, h.log, err, "Error while setting map for code") {
		return
	}

	c.Status(http.StatusOK)
}

// @Router /v1/branches/confirm-login/ [POST]
// @Summary Confirm Branch Login
// @Description API that checks whether - branch entered
// @Description valid token
// @Tags branch
// @Accept  json
// @Produce  json
// @Param confirm_phone body models.ConfirmBranchLoginRequest true "confirm login"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmBranchLogin(c *gin.Context) {
	var (
		cb models.ConfirmBranchLoginRequest
	)

	err := c.ShouldBindJSON(&cb)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	cb.Code = strings.TrimSpace(cb.Code)

	//Getting code from redis
	key := cb.Phone
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
	if cb.Code != s {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidCode,
				Message: "Code is invalid",
			},
		})
		h.log.Error("Code is invalid", logger.Error(err))
		return
	}

	_, err = h.grpcClient.BranchService().GetBranch(
		context.Background(), &pbu.GetBranchRequest{
			Id: cb.Phone,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting branch") {
		return
	}

	c.Status(http.StatusOK)
}

// @Tags branch
// @Router /v1/nearest-branch [get]
// @Summary Get Nearest Branch
// @Description API for getting branch info
// @Accept  json
// @Produce json
// @Param long query string false "long"
// @Param lat query string false "lat"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetNearestBranch(c *gin.Context) {

	var jspbMarshal jsonpb.Marshaler
	var location pbu.Location
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	longString,_ := c.GetQuery("long")
	fmt.Println()
	long, _ := strconv.ParseFloat(longString,64)
	latString,_ := c.GetQuery("lat")
	lat, _ := strconv.ParseFloat(latString,64)
	location.Long = long
	location.Lat = lat
	res, err := h.grpcClient.BranchService().GetNearestBranch(
		context.Background(),
		&pbu.GetNearestBranchRequest{
			Location : &location,
			},
		)
	if handleGRPCErr(c, h.log, err) {
		return 
	}

	st, ok := status.FromError(err)
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Branch Not Found",
			},
		})
		h.log.Error("Error while getting branch, Branch Not Found", logger.Error(err))
		return
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Server unavailable",
			},
		})
		h.log.Error("Error while getting branch, service unavailable", logger.Error(err))
		return
	} else if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while getting branch", logger.Error(err))
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
