package v1

import (
	"context"
	pba "genproto/auth_service"
	"net/http"
	"strings"

	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Router /v1/auth/login [POST]
// @Summary User Login
// @Description API that returns token based on user credential
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.LoginRequest true "login"
// @Param client header string true "client"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Login(c *gin.Context) {
	var (
		login models.LoginRequest
	)

	clientID := c.GetHeader("client")
	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client_id not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

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
	c.JSON(http.StatusOK, &models.LoginResponse{
		ID:           res.Token.Id,
		UserID:       res.Token.UserId,
		ClientID:     res.Token.ClientId,
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
		UserRoleID:   res.Token.UserRoleId,
		UserType:     res.UserType,
	})
}

// @Router /v1/auth/generate-otp [POST]
// @Summary Generate otp for a user
// @Description API that checks whether user exists then generates random otp
// @Tags auth
// @Accept json
// @Param login body models.OTPLoginRequest true "login"
// @Param client header string true "client"
// @Param shipper header string true "shipper"
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GenerateOTP(c *gin.Context) {
	var (
		login models.OTPLoginRequest
	)

	shipperID := c.GetHeader("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	clientID := c.GetHeader("client")
	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client_id not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.log.Error("error while binding login parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "phone field is required",
			"code":    ErrorBadRequest,
		})
		return
	}

	login.Phone = strings.TrimSpace(login.Phone)

	_, err = h.grpcClient.AuthService().GenerateOTP(context.Background(), &pba.OTPLoginRequest{
		Phone:     login.Phone,
		ShipperId: shipperID,
		ClientId:  clientID,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "phone is incorrect",
			"code":    ErrorCodeNotFound,
		})
		return
	}

	c.Status(http.StatusOK)
}

// @Router /v1/auth/confirm-otp [POST]
// @Summary Confirm otp for a user
// @Description API that confirms user otp
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.OTPConfirmRequest true "login"
// @Param client header string true "client"
// @Param shipper header string true "shipper"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmOTP(c *gin.Context) {
	var (
		login models.OTPConfirmRequest
	)

	shipperID := c.GetHeader("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	clientID := c.GetHeader("client")
	if clientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client_id not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.log.Error("error while binding login parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "phone and code fields are required",
			"code":    ErrorBadRequest,
		})
		return
	}

	login.Phone = strings.TrimSpace(login.Phone)
	login.Code = strings.TrimSpace(login.Code)

	res, err := h.grpcClient.AuthService().ConfirmOTP(context.Background(), &pba.OTPConfirmRequest{
		Phone:     login.Phone,
		Code:      login.Code,
		ShipperId: shipperID,
		ClientId:  clientID,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "phone or code is incorrect",
			"code":    ErrorCodeNotFound,
		})
		return
	}

	c.JSON(http.StatusOK, &models.LoginResponse{
		ID:           res.Token.Id,
		UserID:       res.Token.UserId,
		ClientID:     res.Token.ClientId,
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
		UserRoleID:   res.Token.UserRoleId,
		UserType:     res.UserType,
	})
}
