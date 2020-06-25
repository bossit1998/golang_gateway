package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	jwtg "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"bitbucket.org/alien_soft/api_getaway/api/models"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/grpc_client"
	"bitbucket.org/alien_soft/api_getaway/pkg/jwt"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage"
	"bitbucket.org/alien_soft/api_getaway/storage/repo"
)

type handlerV1 struct {
	storage         storage.StorageI
	log             logger.Logger
	inMemoryStorage repo.InMemoryStorageI
	grpcClient      *grpc_client.GrpcClient
	cfg             config.Config
}

//HandlerV1Config ...
type HandlerV1Config struct {
	Storage         storage.StorageI
	Logger          logger.Logger
	InMemoryStorage repo.InMemoryStorageI
	GrpcClient      *grpc_client.GrpcClient
	Cfg             config.Config
}

const (
	//ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	//ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	//ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	//ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	//ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	//ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	//ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	//ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	//ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	//ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	//ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	//ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
)

var (
	signingKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
)

//New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		storage:         c.Storage,
		inMemoryStorage: c.InMemoryStorage,
		log:             c.Logger,
		grpcClient:      c.GrpcClient,
		cfg:             c.Cfg,
	}
}

//ParsePageQueryParam ...
func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 10)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

//ParsePageSizeQueryParam ...
func ParsePageSizeQueryParam(c *gin.Context) (uint64, error) {
	pageSize, err := strconv.ParseUint(c.DefaultQuery("page_size", "10"), 10, 10)
	if err != nil {
		return 0, err
	}
	if pageSize < 0 {
		return 0, errors.New("page_size must be an positive integer")
	}
	return pageSize, nil
}

//ParseLimitQueryParam ...
func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 10)
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("page_size must be an positive integer")
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func handleGRPCErr(c *gin.Context, l logger.Logger, err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	var errI interface{} = models.InternalServerError{
		Code:    ErrorCodeInternal,
		Message: "Internal Server Error",
	}
	httpCode := http.StatusInternalServerError
	if ok && st.Code() == codes.InvalidArgument {
		httpCode = http.StatusBadRequest
		errI = ErrorBadRequest
	}
	c.JSON(httpCode, models.ResponseError{
		Error: errI,
	})
	if ok {
		l.Error(fmt.Sprintf("code=%d message=%s", st.Code(), st.Message()), logger.Error(err))
	}
	return true
}

func writeMessageAsJSON(c *gin.Context, l logger.Logger, msg proto.Message) {
	if msg == nil {
		c.String(http.StatusOK, "")
		return
	}
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	js, err := jspbMarshal.MarshalToString(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error("Error while marshaling", logger.Error(err))
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

func handleGrpcErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: st.Message(),
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: st.Message(),
			},
		})
		l.Error(message+", not found", logger.Error(err))
		return true
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message+", service unavailable", logger.Error(err))
		return true
	} else if st.Code() == codes.AlreadyExists {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: st.Message(),
			},
		})
		l.Error(message+", already exists", logger.Error(err))
		return true
	} else if st.Code() == codes.InvalidArgument {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Code() == codes.Code(20) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Err() != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	}
	return false
}

func handleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func handleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidJSON,
				Message: "Invalid Json",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}

func handleStorageErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: "Not found",
			},
		})
		l.Error(message+", not found", logger.Error(err))
		return true
	} else if err == repo.ErrAlreadyExists {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Already Exists",
			},
		})
		l.Error(message+", already exists", logger.Error(err))
		return true
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func getDistance(fromLocation models.Location, toLocation models.Location, cfg config.Config) float64 {
	coordinates := fmt.Sprintf("%f,%f;%f,%f", fromLocation.Long, fromLocation.Lat, toLocation.Long, toLocation.Lat)
	url := "https://api.mapbox.com/directions/v5/mapbox/driving/" + coordinates + "/?approaches=unrestricted;curb&access_token=" + cfg.MapboxToken + ""

	resp, err := http.Get(url)
	if err != nil {
		return 0
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	geodriving := models.GeoDrivingAPIResponse{}
	json.Unmarshal(body, &geodriving)

	dist := geodriving.RoutesList[0].LegsList[0].Distance

	return dist
}

func getOptimizedTrip(tripData models.TripsDataModel, cfg config.Config) models.OptimizedTrips {

	var tripCoordinates string
	tripCoordinates = fmt.Sprintf("%f,%f;", tripData.CurrentLocation.Long, tripData.CurrentLocation.Lat)

	if len(tripData.Origins) < 10 {
		for j := 0; j < len(tripData.Origins); j++ {
			tripCoordinates += fmt.Sprintf("%f,%f;", tripData.Origins[j].Long, tripData.Origins[j].Lat)
		}
	} else {
		fmt.Println("Too many origins")
	}

	tripCoordinates += fmt.Sprintf("%f,%f", tripData.Destination.Long, tripData.Destination.Lat)

	fmt.Println(tripCoordinates)

	url := "https://api.mapbox.com/optimized-trips/v1/mapbox/driving/" + tripCoordinates + "?source=first&destination=last&roundtrip=true&access_token=" + cfg.MapboxToken + ""

	resp, err := http.Get(url)

	fmt.Println(url)
	if err != nil {
		fmt.Println("to many origins | InvalidInput")
		//return 0
		// handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	optimizedTrip := models.OptimizedTrips{}
	json.Unmarshal(body, &optimizedTrip)

	return optimizedTrip
}

func userInfo(h *handlerV1, c *gin.Context) (models.UserInfo, error) {
	claims, err := GetClaims(h, c)

	if err != nil {
		return models.UserInfo{}, err
	}

	h.log.Info("claims", logger.String("", claims))
	userID := claims["sub"].(string)
	userRole := claims["role"].(string)

	return models.UserInfo{
		ID:   userID,
		Role: userRole,
	}, nil
}

func GetClaims(h *handlerV1, c *gin.Context) (jwtg.MapClaims, error) {
	var (
		ErrUnauthorized = errors.New("unauthorized")
		authorization   models.AuthorizationModel
		claims          jwtg.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return nil, ErrUnauthorized
	}

	claims, err = jwt.ExtractClaims(authorization.Token, signingKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: ", logger.Error(err))
		return nil, ErrUnauthorized
	}

	return claims, nil
}
