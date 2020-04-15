package middleware

import (
	"net/http"
	s "strings"

	"bitbucket.org/alien_soft/api_gateway/api/models"

	"github.com/casbin/casbin/v2"
	jwtg "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"bitbucket.org/alien_soft/api_gateway/pkg/jwt"
)

var (
	mySigningKey = []byte("secretphrase")
	newSigningKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
)

//NewAuthorizer is a middleware for gin to get role and
//allow or deny access to endpointns
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &JWTRoleAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		allow, err := a.CheckPermission(c.Request)
		if err != nil {
			v, _ := err.(*jwtg.ValidationError)
			if v.Errors == jwtg.ValidationErrorExpired {
				a.RequireRefresh(c)
			} else {
				a.RequirePermission(c)
			}
		} else if !allow {
			a.RequirePermission(c)
		}
	}
}

//JWTRoleAuthorizer ...
type JWTRoleAuthorizer struct {
	enforcer *casbin.Enforcer
}

//GetRole ...
func (a *JWTRoleAuthorizer) GetRole(r *http.Request) (string, error) {
	var (
		role string
		claims jwtg.MapClaims
		err error
	)

	token := r.Header.Get("Authorization")
	if token == "" {
		return "unauthorized", nil
	} else if s.Contains(token, "Basic") {
		return "unauthorized", nil
	}

	claims, err = jwt.ExtractClaims(token, mySigningKey)
	if err != nil {
		claims, err = jwt.ExtractClaims(token, newSigningKey)
		if err != nil {
			return "", err
		}
	}


	if claims["role"].(string) == "cluber" {
		role = "cluber"
	} else if claims["role"].(string) == "club" {
		role = "club"
	} else if claims["role"].(string) == "promoter" {
		role = "promoter"
	} else if claims["role"].(string) == "admin" {
		role = "admin"
	} else if claims["role"].(string) == "authorized" {
		role = "authorized"
	} else {
		role = "unknown"
	}

	return role, nil
}

//CheckPermission ...
func (a *JWTRoleAuthorizer) CheckPermission(r *http.Request) (bool, error) {
	user, err := a.GetRole(r)
	if err != nil {
		return false, err
	}
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		panic(err)
	}

	return allowed, nil
}

//RequirePermission ...
func (a *JWTRoleAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}

//RequireRefresh ...
func (a *JWTRoleAuthorizer) RequireRefresh(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, models.ResponseError{
		Error: models.InternalServerError{
			Code:    "UNAUTHORIZED",
			Message: "Token is expired",
		},
	})
	c.AbortWithStatus(401)
}
