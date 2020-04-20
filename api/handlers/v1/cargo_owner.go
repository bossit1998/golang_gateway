package v1

import (
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
)

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
	cargoOwner.Token = token
	passwordHash, err := etc.GeneratePasswordHash(cargoOwner.Password)

	if handleBadRequestErrWithMessage(c, h.log, err, "error while hashing password") {
		return
	}
	cargoOwner.Password = string(passwordHash)

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
