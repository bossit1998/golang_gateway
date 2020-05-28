package v1

import (
	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v6"
	"os"

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

func validation(ext string) bool {
	for _, val := range allowedExtensions {
		if val == ext {
			return true
		}
	}
	return false
}

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
	fName, _ := uuid.NewRandom()
	file.File.Filename = fName.String()
	dst, _ := os.Getwd()


	minioClient, err := minio.New(h.cfg.MinioEndpoint, h.cfg.MinioAccessKeyID, h.cfg.MinioSecretAccesKey, true)

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

	fmt.Println(dst+"/"+fName.String())

	n, err := minioClient.FPutObject("delever", fName.String(), fName.String(), minio.PutObjectOptions{ContentType:"image/jpeg"})
	fmt.Println(err)
	c.JSON(http.StatusOK, n)

	//file.File.
	//if !validation(c.GetHeader("Content-Type")) {
	//	c.JSON(http.StatusBadRequest, models.ResponseError{
	//		Error: models.InternalServerError{
	//			Code:ErrorBadRequest,
	//			Message: "Content-Type not allowed",
	//		},
	//	})
	//	h.log.Error("content-type not allowed", logger.String("content-type", c.GetHeader("Content-Type")))
	//	return
	//}
	//
	//file, err := c.GetRawData()
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, models.ResponseError{
	//		Error: models.InternalServerError{
	//			Code:ErrorBadRequest,
	//		},
	//	})
	//	h.log.Error("error while reading binary data", logger.Error(err))
	//	return
	//}
	//
	//id, _ := uuid.NewRandom()
	//
	//err = ioutil.WriteFile(fmt.Sprintf("static/images/%s", id.String()), file, 0644)
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, models.ResponseError{
	//		Error: models.InternalServerError{
	//			Code:    ErrorBadRequest,
	//			Message: "error while saving file",
	//		},
	//	})
	//	h.log.Error("error while saving file", logger.Error(err))
	//	return
	//}
	//
	//c.JSON(http.StatusOK, Path{
	//	Path: fmt.Sprintf("images/%s", id.String()),
	//})
}
