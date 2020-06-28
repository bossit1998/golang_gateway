package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"context"
	pba "genproto/auth_service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *handlerV1) Login(c *gin.Context) {
	var (
		login models.Login
		isMatch = true
	)
	clientID := c.GetHeader("client_id")

	err := c.ShouldBindJSON(&login)

	if err != nil {
		h.log.Error("error while binding login parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login and password required fields",
			"code": ErrorBadRequest,
		})
		return
	}

	err = helpers.ValidateLogin(login.Login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code": ErrorBadRequest,
		})
		return
	}

	err = helpers.ValidatePassword(login.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code": ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.AuthService().Login(context.Background(), &pba.LoginRequest{
		Login: login.Login,
		ClientId:clientID,
	})

	if err != nil {
		isMatch = false
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(login.Password))

		if err != nil {
			isMatch = false
		}
	}

	if !isMatch {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "login or password is incorrect",
			"code": ErrorCodeNotFound,
		})
		return
	}

	m := map[interface{}]interface{}{
		"user_type": res.Role,
		"shipper_id": res.Id,
		"sub": res.Id,
	}

	accessToken, refreshToken, err := jwt.GenJWT(m, signingKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"code": ErrorBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shipper_id": res.Id,
		"access_token": accessToken,
		"refresh_token": refreshToken,
	})
}
