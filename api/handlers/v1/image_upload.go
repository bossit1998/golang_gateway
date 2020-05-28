package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

var allowedExtensions = []string{"image/png", "image/jpeg"}

type Path struct {
	Path string `json:"path"`
}

func validation(ext string) bool {
	for _, val := range allowedExtensions {
		if val == ext {
			return true
		}
	}
	return false
}

func (h *handlerV1) ImageUpload(c *gin.Context) {
	if !validation(c.GetHeader("Content-Type")) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message: "Content-Type not allowed",
			},
		})
		h.log.Error("content-type not allowed", logger.String("content-type", c.GetHeader("Content-Type")))
		return
	}

	file, err := c.GetRawData()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
			},
		})
		h.log.Error("error while reading binary data", logger.Error(err))
		return
	}

	id, _ := uuid.NewRandom()

	err = ioutil.WriteFile(fmt.Sprintf("static/images/%s", id.String()), file, 0644)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while saving file",
			},
		})
		h.log.Error("error while saving file", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, Path{
		Path: fmt.Sprintf("images/%s", id.String()),
	})
}
