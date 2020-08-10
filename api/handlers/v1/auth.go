package v1

import (
	"context"
	pba "genproto/auth_service"
	"net/http"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Router /v1/auth/login [POST]
// @Summary Check User Login
// @Description API that checks whether user exists
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.Login true "login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Login(c *gin.Context) {
	var (
		login models.Login
	)
	clientID := c.GetHeader("client_id")

	err := c.ShouldBindJSON(&login)

	if err != nil {
		h.log.Error("error while binding login parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login and password required fields",
			"code":    ErrorBadRequest,
		})
		return
	}

	err = helpers.ValidateLogin(login.Login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    ErrorBadRequest,
		})
		return
	}

	err = helpers.ValidatePassword(login.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code":    ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.AuthService().Login(context.Background(), &pba.LoginRequest{
		Login:    login.Login,
		Password: login.Password,
		ClientId: clientID,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "login or password is incorrect",
			"code":    ErrorCodeNotFound,
		})
		return
	}

	// switch res.UserType {
	// case "shipper":
	// 	c.JSON(httpswitch res.UserType {
	// case "shipper":
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"shipper_id":    res.Token.UserId,
	// 		"access_token":  res.Token.AccessToken,
	// 		"refresh_token": res.Token.RefreshToken,
	// 		"user_role_id":  res.Token.UserRoleId,
	// 	})
	// default:
	// 	c.JSON(http.StatusOK, res.Token)
	// }.StatusOK, gin.H{
	// 		"shipper_id":    res.Token.UserId,
	// 		"access_token":  res.Token.AccessToken,
	// 		"refresh_token": res.Token.RefreshToken,
	// 		"user_role_id":  res.Token.UserRoleId,
	// 	})
	// default:
	// 	c.JSON(http.StatusOK, res.Token)
	// }
	c.JSON(http.StatusOK, res.Token)
}

// @Security ApiKeyAuth
// @Router /v1/auth/platforms [GET]
// @Summary Get All Platforms
// @Description API for getting platforms
// @Tags auth
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllPlatformsModel
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// func (h *handlerV1) GetAllPlatforms(c *gin.Context) {
// 	var (
// 		jspbMarshal jsonpb.Marshaler
// 	)

// 	jspbMarshal.OrigName = true
// 	jspbMarshal.EmitDefaults = true

// 	page, err := ParsePageQueryParam(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: ErrorBadRequest,
// 		})
// 		return
// 	}

// 	limit, err := ParseLimitQueryParam(c)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, models.ResponseError{
// 			Error: ErrorBadRequest,
// 		})
// 		return
// 	}

// 	res, err := h.grpcClient.AuthService().GetAllPlatforms(
// 		context.Background(),
// 		&pba.GetAllRequest{
// 			Page:  page,
// 			Limit: limit,
// 		},
// 	)
// 	if handleGRPCErr(c, h.log, err) {
// 		return
// 	}
// 	js, err := jspbMarshal.MarshalToString(res)

// 	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
// 		return
// 	}

// 	c.Header("Content-Type", "application/json")
// 	c.String(http.StatusOK, js)
// }
