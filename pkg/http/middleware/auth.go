package middleware

import (
	"bitbucket.org/alien_soft/api_getaway/config"
	"net/http"
	s "strings"

	"bitbucket.org/alien_soft/api_getaway/api/models"

	"github.com/casbin/casbin/v2"
	jwtg "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
)

var (
	signingKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
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

//GetRole gets role from jwt token if it is sent
//or sets role to unauthorized if token is not sent
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

	claims, err = jwt.ExtractClaims(token, signingKey)
	if err != nil {
		return "", err
	}


	if claims["role"].(string) == config.RoleCargoOwnerAdmin {
		role = "cargo_owner_admin"
	} else if claims["role"].(string) == config.RoleCargoAPI {
		role = "cargo_api"
	} else if claims["role"].(string) == config.RoleDistributorAdmin {
		role = "distributor_admin"
	} else if claims["role"].(string) == config.RoleCourier {
		role = "courier"
	} else if claims["role"].(string) == config.RoleAdmin {
		role = "admin"
	} else {
		role = config.RoleUnknown
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
