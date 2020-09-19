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

// @Router /v1/auth/refresh-token [POST]
// @Summary User Refresh Token
// @Description API that returns token based on user credential
// @Tags auth
// @Accept  json
// @Produce  json
// @Param refresh_token body models.RefreshTokenRequest true "refresh-token"
// @Param client header string true "client"
// @Success 200 {object} models.LoginResponse
// @Failure 403 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) RefreshToken(c *gin.Context) {
	var (
		refreshTokenRequest models.RefreshTokenRequest
	)

	err := c.ShouldBindJSON(&refreshTokenRequest)

	if err != nil {
		h.log.Error("error while binding refreshTokenRequest parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh token field is required ",
			"code":    ErrorBadRequest,
		})
		return
	}
	userInfo, err := ReturnUserInfo(refreshTokenRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusForbidden, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	clientID := c.GetHeader("client")
	if clientID == "" || clientID != userInfo.ClientID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client_id not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.AuthService().RefreshToken(context.Background(), &pba.RefreshTokenRequest{
		ClientId:  userInfo.ClientID,
		UserId:    userInfo.UserID,
		ShipperId: userInfo.ShipperID,
	})

	if err != nil {
		c.JSON(http.StatusForbidden, models.ResponseError{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, &models.LoginResponse{
		ID:           res.Token.Id,
		UserID:       res.Token.UserId,
		ClientID:     res.Token.ClientId,
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
		UserRoleID:   res.Token.UserRoleId,
		UserTypeID:   res.UserTypeId,
	})
}

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

	c.JSON(http.StatusOK, &models.LoginResponse{
		ID:           res.Token.Id,
		UserID:       res.Token.UserId,
		ClientID:     res.Token.ClientId,
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
		UserRoleID:   res.Token.UserRoleId,
		UserTypeID:   res.UserTypeId,
	})
}

// @Router /v1/auth/generate-otp [POST]
// @Summary Generate otp for a user
// @Description API that checks whether user exists then generates random otp
// @Tags auth
// @Accept json
// @Param login body models.OTPLoginRequest true "login"
// @Param client header string true "client"
// @Param shipper header string false "shipper"
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GenerateOTP(c *gin.Context) {
	var (
		login models.OTPLoginRequest
	)

	shipperID := c.GetHeader("shipper")
	// if shipperID == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "shipper not found in header",
	// 		"code":    ErrorBadRequest,
	// 	})
	// 	return
	// }

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
// @Param shipper header string false "shipper"
// @Param fcm header string false "fcm"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseError
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmOTP(c *gin.Context) {
	var (
		login models.OTPConfirmRequest
	)

	fcmToken := c.GetHeader("fcm")
	shipperID := c.GetHeader("shipper")
	// if shipperID == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "shipper not found in header",
	// 		"code":    ErrorBadRequest,
	// 	})
	// 	return
	// }

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
		ClientId:  clientID,
		ShipperId: shipperID,
		FcmToken:  fcmToken,
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
		UserTypeID:   res.UserTypeId,
	})
}
