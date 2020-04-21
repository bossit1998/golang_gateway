package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/helpers"
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/etc"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"context"
	pbco "genproto/co_service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// @Router /v1/cargo-owner [post]
// @Summary Create Cargo Owner
// @Description API for creating cargo owner
// @Tags cargo-owner
// @Accept  json
// @Produce  json
// @Param cargo_owner body models.CreateCargoOwner true "cargo_owner"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateCO(c *gin.Context) {
	var (
		jspbUnmarshal jsonpb.Unmarshaler
		cargoOwner pbco.CargoOwner
		cargoOwnerID uuid.UUID
		err error
	)

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &cargoOwner)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while unmarshal") {
		return
	}
	cargoOwner.Login = strings.TrimSpace(cargoOwner.Login)

	err = helpers.ValidateLogin(cargoOwner.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = helpers.ValidatePassword(cargoOwner.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	passwordHash, err := etc.GeneratePasswordHash(cargoOwner.Password)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while hashing password") {
		return
	}

	for {
		cargoOwnerID, err = uuid.NewRandom()

		if err == nil {
			break
		}
	}

	cargoOwner.Id = cargoOwnerID.String()

	token, err := jwt.GenerateJWT(cargoOwner.Id, config.RoleCargoOwner, newSigningKey)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while generating token") {
		return
	}

	cargoOwner.Password = string(passwordHash)
	cargoOwner.Token = token

	_, err = h.grpcClient.COService().Create(
		context.Background(),
		&cargoOwner,
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating cargo owner") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "cargo owner created successfully",
	})
}

// @Router /v1/cargo-owner/check-name [post]
// @Summary Check Cargo Owner Name
// @Description API for checking cargo owner name exists or not in the table
// @Tags cargo-owner
// @Accept  json
// @Produce  json
// @Param cargo_owner_name body models.CheckNameRequest true "cargo_owner_name"
// @Success 200 {object} models.CheckExistsResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckCOName(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		checkNameModel models.CheckNameRequest
	)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err := c.ShouldBindJSON(&checkNameModel)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	resp, err := h.grpcClient.COService().CheckExists(
		context.Background(),
		&pbco.CheckExistsRequest{
			Column:"name",
			Value:strings.ToLower(strings.ReplaceAll(checkNameModel.Name, " ", "")),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while checking name") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)

	if handleGrpcErrWithMessage(c, h.log, err, "marshaling error while checking name") {
		return
	}

	c.Header("content-type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/cargo-owner/check-login [post]
// @Summary Check Cargo Owner Login
// @Description API for checking cargo owner login exists or not in the table
// @Tags cargo-owner
// @Accept  json
// @Produce  json
// @Param cargo_owner_login body models.CheckLoginRequest true "cargo_owner_login"
// @Success 200 {object} models.CheckExistsResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckLogin(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		checkLoginModel models.CheckLoginRequest
	)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err := c.ShouldBindJSON(&checkLoginModel)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}
	checkLoginModel.Login = strings.TrimSpace(checkLoginModel.Login)


	err = helpers.ValidateLogin(checkLoginModel.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	resp, err := h.grpcClient.COService().CheckExists(
		context.Background(),
		&pbco.CheckExistsRequest{
			Column:"login",
			Value:strings.ToLower(strings.ReplaceAll(checkLoginModel.Login, " ", "")),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while checking name") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)

	if handleGrpcErrWithMessage(c, h.log, err, "marshaling error while checking name") {
		return
	}

	c.Header("content-type", "application/json")
	c.String(http.StatusOK, js)
}