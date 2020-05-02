package api

import (
	"net/http"

	//"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bitbucket.org/alien_soft/api_getaway/api/docs" //for swagger
	"bitbucket.org/alien_soft/api_getaway/storage/repo"

	v1 "bitbucket.org/alien_soft/api_getaway/api/handlers/v1"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/grpc_client"

	//"bitbucket.org/alien_soft/api_getaway/pkg/http/middleware"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
	"bitbucket.org/alien_soft/api_getaway/storage"
)

//Config ...
type Config struct {
	Storage         storage.StorageI
	Logger          logger.Logger
	InMemoryStorage repo.InMemoryStorageI
	GrpcClient      *grpc_client.GrpcClient
	Cfg             config.Config
	//CasbinEnforcer  *casbin.Enforcer
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//r.Use(middleware.NewAuthorizer(cnf.CasbinEnforcer))

	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type")
		context.Header("Access-Control-Allow-Headers","Authorization")
	})

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage:         cnf.Storage,
		InMemoryStorage: cnf.InMemoryStorage,
		Logger:          cnf.Logger,
		GrpcClient:      cnf.GrpcClient,
		Cfg:             cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "bratan api getaway bu"})
	})

	//Courier endpoints
	r.GET("/v1/couriers", handlerV1.GetAllCouriers)
	r.GET("/v1/couriers/:courier_id", handlerV1.GetCourier)
	r.GET("/v1/couriers/:courier_id/courier-details", handlerV1.GetCourierDetails)
	r.GET("/v1/couriers/:courier_id/vehicles", handlerV1.GetAllCourierVehicles)
	r.POST("/v1/couriers", handlerV1.CreateCourier)
	r.POST("/v1/couriers/courier-details", handlerV1.CreateCourierDetails)
	r.PATCH("/v1/couriers/:courier_id/block", handlerV1.BlockCourier)
	r.PATCH("/v1/couriers/:courier_id/unblock", handlerV1.UnblockCourier)
	r.PUT("/v1/couriers", handlerV1.UpdateCourier)
	r.PUT("/v1/couriers/courier-details", handlerV1.UpdateCourierDetails)
	r.DELETE("/v1/couriers/:courier_id", handlerV1.DeleteCourier)
	r.POST("/v1/couriers/check-login/", handlerV1.CheckCourierLogin)
	r.POST("/v1/couriers/confirm-login/", handlerV1.ConfirmCourierLogin)

	//Vehicle endpoints
	r.GET("/v1/vehicles/:vehicle_id", handlerV1.GetCourierVehicle)
	r.POST("/v1/vehicles", handlerV1.CreateCourierVehicle)
	r.PUT("/v1/vehicles", handlerV1.UpdateCourierVehicle)
	r.DELETE("/v1/vehicles/:vehicle_id", handlerV1.DeleteCourierVehicle)

	//Distributor endpoints
	r.GET("/v1/distributors/", handlerV1.GetAllDistributors)
	r.GET("/v1/distributors/:distributor_id/", handlerV1.GetDistributor)
	r.GET("/v1/distributors/:distributor_id/couriers/", handlerV1.GetAllDistributorCouriers)
	r.GET("/v1/distributors/:distributor_id/parks/", handlerV1.GetAllDistributorParks)
	r.POST("/v1/distributors/", handlerV1.CreateDistributor)
	r.PUT("/v1/distributors/", handlerV1.UpdateDistributor)
	r.DELETE("/v1/distributors/:distributor_id/", handlerV1.DeleteDistributor)

	//Park endpoints
	r.GET("/v1/parks/:park_id", handlerV1.GetPark)
	r.POST("/v1/parks", handlerV1.CreatePark)

	//GetDistanse
	r.POST("/v1/calc-delivery-cost", handlerV1.CalcDeliveryCost)
	//GetOptimizedTrips
	r.POST("/v1/optimized-trip", handlerV1.OptimizedTrip)

	//Fare endpoints
	r.GET("/v1/fares/:fare_id", handlerV1.GetFare)
	r.GET("/v1/fares", handlerV1.GetAllFares)
	r.POST("/v1/fares", handlerV1.CreateFare)
	r.PUT("/v1/fares", handlerV1.UpdateFare)
	r.DELETE("/v1/fares/:fare_id", handlerV1.DeleteFare)

	//Order endpoints
	r.POST("/v1/order", handlerV1.CreateOrder)
	r.PUT("/v1/order/:order_id", handlerV1.UpdateOrder)
	r.GET("/v1/order/:order_id", handlerV1.GetOrder)
	r.GET("/v1/order", handlerV1.GetOrders)
	r.GET("/v1/new-order", handlerV1.NewOrders)
	r.PATCH("v1/order/:order_id/change-status", handlerV1.ChangeOrderStatus)
	r.GET("/v1/order-statuses", handlerV1.GetStatuses)
	r.PATCH("v1/order/:order_id/add-courier", handlerV1.AddCourier)
	r.PATCH("v1/order/:order_id/remove-courier", handlerV1.RemoveCourier)
	r.GET("/v1/courier/order", handlerV1.GetCourierOrders)
	r.PATCH("/v1/order-step/:step_id/take", handlerV1.TakeOrderStep)

	//Cargo owner
	r.POST("/v1/cargo-owner", handlerV1.CreateCO)
	r.GET("/v1/cargo-owner/", handlerV1.GetCO)
	r.POST("/v1/cargo-owner/check-name", handlerV1.CheckCOName)
	r.POST("/v1/cargo-owner/check-login", handlerV1.CheckLogin)
	r.POST("/v1/cargo-owner/refresh-token", handlerV1.RefreshToken)
	r.POST("/v1/cargo-owner/change-credentials", handlerV1.ChangeLoginPassword)

	//Login endpoints
	r.POST("/v1/check_code/")

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
