package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bitbucket.org/alien_soft/api_gateway/api/docs" //for swagger

	v1 "bitbucket.org/alien_soft/api_gateway/api/handlers/v1"
	"bitbucket.org/alien_soft/api_gateway/config"
	"bitbucket.org/alien_soft/api_gateway/pkg/grpc_client"
	"bitbucket.org/alien_soft/api_gateway/pkg/logger"
	"bitbucket.org/alien_soft/api_gateway/storage"
)

//Config ...
type Config struct {
	Storage    storage.StorageI
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage:    cnf.Storage,
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "api gateway"})
	})

	//Courier endpoints
	r.GET("/v1/couriers", handlerV1.GetAllCouriers)
	r.GET("/v1/couriers/:courier_id", handlerV1.GetCourier)
	r.GET("/v1/couriers/:courier_id/courier_details", handlerV1.GetCourierDetails)
	r.GET("/v1/couriers/:courier_id/vehicles", handlerV1.GetAllCourierVehicles)
	r.POST("/v1/couriers", handlerV1.CreateCourier)
	r.POST("/v1/couriers/courier_details", handlerV1.CreateCourierDetails)
	r.PUT("/v1/couriers", handlerV1.UpdateCourier)
	r.PUT("/v1/couriers/courier_details", handlerV1.UpdateCourierDetails)
	r.DELETE("/v1/couriers/:courier_id", handlerV1.DeleteCourier)

	//Vehicle endpoints
	r.GET("/v1/vehicles/:vehicle_id", handlerV1.GetCourierVehicle)
	r.POST("/v1/vehicles", handlerV1.CreateCourierVehicle)
	r.PUT("/v1/vehicles", handlerV1.UpdateCourierVehicle)
	r.DELETE("/v1/vehicles/:vehicle_id", handlerV1.DeleteCourierVehicle)

	//Distributor endpoints
	r.GET("/v1/distributors", handlerV1.GetAllDistributors)
	r.GET("/v1/distributors/:distributor_id", handlerV1.GetDistributor)
	r.GET("/v1/distributors/:distributor_id/couriers", handlerV1.GetAllDistributorCouriers)
	r.POST("/v1/distributors", handlerV1.CreateDistributor)
	r.PUT("/v1/distributors", handlerV1.UpdateDistributor)
	r.DELETE("/v1/distributors/:distributor_id", handlerV1.DeleteDistributor)

	//Geo
	r.GET("/v1/geozones/", handlerV1.GetGeozones)

	//Fare endpoints
	//r.POST("/v1/fares/", handlerV1.CreateFare)
	//r.GET("/v1/fares/:id/", handlerV1.GetFare)
	//r.GET("/v1/fares/", handlerV1.GetAllFares)
	//r.PUT("/v1/fares/:id", handlerV1.UpdateFare)
	//r.DELETE("/v1/fares/:id", handlerV1.DeleteFare)

	//Order endpoints
	r.POST("/v1/order", handlerV1.Create)
	r.GET("/v1/order/:order_id", handlerV1.GetOrder)
	r.GET("/v1/order", handlerV1.GetOrders)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
