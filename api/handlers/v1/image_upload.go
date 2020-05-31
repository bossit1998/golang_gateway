package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v6"
	url2 "net/url"
	"os"
	"time"

	"github.com/google/uuid"
	//"io/ioutil"
	"mime/multipart"
	"net/http"
	//"github.com/minio/minio-go/v6"
)

var allowedExtensions = []string{"image/png", "image/jpeg"}

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type Path struct {
	Filename string `json:"filename"`
} 

func validation(ext string) bool {
	for _, val := range allowedExtensions {
		if val == ext {
			return true
		}
	}
	return false
}

// @Router /v1/upload [post]
// @Tags image
// @Param file formData file true "file"
// @Success 200 {object} Path
func (h *handlerV1) ImageUpload(c *gin.Context) {
	var (
		file File
	)
	err := c.ShouldBind(&file)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while binding file",
			},
		})
		h.log.Error("error while binding file", logger.Error(err))
		return
	}

	if !validation(file.File.Header["Content-Type"][0]) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message: "Content-Type not allowed",
			},
		})
		h.log.Error("content-type not allowed", logger.String("content-type", c.GetHeader("Content-Type")))
		return
	}

	fName, _ := uuid.NewRandom()
	file.File.Filename = fName.String()
	dst, _ := os.Getwd()

	minioClient, err := minio.New(h.cfg.MinioEndpoint, h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccesKey, false)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:models.InternalServerError{
				Code:    ErrorBadRequest,
			},
		})
		h.log.Error("error while connecting minio", logger.Error(err))
		return
	}
	err = c.SaveUploadedFile(file.File, dst+"/"+fName.String())

	_, err = minioClient.FPutObject("delever", fName.String(), dst+"/"+fName.String(), minio.PutObjectOptions{ContentType:"image/jpeg"})
	
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:ErrorBadRequest,
				Message:"error while saving file",
			},
		})
		h.log.Error("error while saving minio", logger.Error(err))
		os.Remove(dst+"/"+fName.String())
		
		return
	}
	os.Remove(dst+"/"+fName.String())

	url, err := minioClient.PresignedGetObject("delever", fName.String(), time.Hour, url2.Values{})
	fmt.Println(url)
	
	c.JSON(http.StatusOK, Path{
		Filename: fName.String(),
	})
}
