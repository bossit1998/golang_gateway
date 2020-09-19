package middleware

import (
	"net/http"

	v1 "bitbucket.org/alien_soft/api_getaway/api/handlers/v1"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Authorizer(key []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientID := c.GetHeader("client")
		token := c.GetHeader("Authorization")

		if clientID == "" && token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "login credentials is incorrect",
			})
			c.Abort()
			return
		}

		if clientID != "" {
			_, err := uuid.Parse(clientID)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "client_id format is incorrect",
					"code":    v1.ErrorBadRequest,
				})
				c.Abort()
				return
			}
		}
	}
}
