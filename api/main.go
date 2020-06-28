package api

import (
	"net/http"

	//"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "bitbucket.org/alien_soft/api_getaway/api/docs" //for swagger
	v1 "bitbucket.org/alien_soft/api_getaway/api/handlers/v1"
	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/grpc_client"
	"bitbucket.org/alien_soft/api_getaway/storage/repo"

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

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	//r.Use(middleware.Authorizer(v1.SigningKey))

	//r.Use(middleware.NewAuthorizer(cnf.CasbinEnforcer))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowHeaders = append(config.AllowHeaders, "image/jpeg")
	config.AllowHeaders = append(config.AllowHeaders, "image/png")
	config.AllowHeaders = append(config.AllowHeaders, "shipper_id")

	r.Use(cors.New(config))

	//r.Use(func(context *gin.Context) {
	//	context.Header("Access-Control-Allow-Origin", "*")
	//	context.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH")
	//	context.Header("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	//	context.Header("Access-Control-Allow-Credentials", "true")
	//})

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Storage:         cnf.Storage,
		InMemoryStorage: cnf.InMemoryStorage,
		Logger:          cnf.Logger,
		GrpcClient:      cnf.GrpcClient,
		Cfg:             cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})
	// Register endpoints
	r.POST("/v1/customers/register", handlerV1.Register)
	r.POST("/v1/customers/register-confirm", handlerV1.RegisterConfirm)

	//Customer endpoints
	r.POST("/v1/customers", handlerV1.CreateCustomer)
	r.PUT("/v1/customers", handlerV1.UpdateCustomer)
	r.DELETE("/v1/customers/:customer_id", handlerV1.DeleteCustomer)
	r.GET("/v1/search-customers", handlerV1.SearchByPhone)
	r.GET("/v1/customers/:customer_id", handlerV1.GetCustomer)
	r.GET("/v1/customers", handlerV1.GetAllCustomers)
	r.POST("/v1/customers/login", handlerV1.CheckCustomerLogin)
	r.POST("/v1/customers/confirm-login", handlerV1.ConfirmCustomerLogin)

	//Branch endpoints
	r.POST("/v1/branches", handlerV1.CreateBranch)
	r.PUT("/v1/branches", handlerV1.UpdateBranch)
	r.DELETE("/v1/branches/:branch_id", handlerV1.DeleteBranch)
	r.GET("/v1/branches/:branch_id", handlerV1.GetBranch)
	r.GET("/v1/branches", handlerV1.GetAllBranches)
	r.POST("/v1/branches/check-login", handlerV1.CheckBranchLogin)
	r.POST("/v1/branches/confirm-login", handlerV1.ConfirmBranchLogin)
	r.GET("/v1/nearest-branch", handlerV1.GetNearestBranch)
	r.POST("/v1/branches/add-courier", handlerV1.CreateBranchCourier)
	r.POST("/v1/branches/remove-courier", handlerV1.DeleteBranchCourier)
	r.GET("/v1/branches/:branch_id/couriers", handlerV1.GetAllBranchCouriers)

	//Shipper endpoints
	r.POST("/v1/shippers", handlerV1.CreateShipper)
	r.PUT("/v1/shippers", handlerV1.UpdateShipper)
	r.DELETE("/v1/shippers/:shipper_id", handlerV1.DeleteShipper)
	r.GET("/v1/shippers/:shipper_id", handlerV1.GetShipper)
	r.GET("/v1/shippers", handlerV1.GetAllShippers)
	r.POST("/v1/shippers/check-login", handlerV1.CheckShipperLogin)
	r.POST("/v1/shippers/confirm-login", handlerV1.ConfirmShipperLogin)
	r.PATCH("/v1/shippers/change-password", handlerV1.ChangePassword)
	r.POST("/v1/shippers/login", handlerV1.ShipperLogin)

	//Courier endpoints
	r.GET("/v1/couriers", handlerV1.GetAllCouriers)
	r.GET("/v1/couriers/:courier_id", handlerV1.GetCourier)
	r.GET("/v1/search-couriers", handlerV1.SearchCouriersByPhone)
	r.GET("/v1/couriers/:courier_id/courier-details", handlerV1.GetCourierDetails)
	r.GET("/v1/couriers/:courier_id/vehicles", handlerV1.GetAllCourierVehicles)
	r.GET("/v1/couriers/:courier_id/branches", handlerV1.GetAllCourierBranches)
	r.POST("/v1/couriers", handlerV1.CreateCourier)
	r.POST("/v1/couriers/courier-details", handlerV1.CreateCourierDetails)
	r.PATCH("/v1/couriers/:courier_id/block", handlerV1.BlockCourier)
	r.PATCH("/v1/couriers/:courier_id/unblock", handlerV1.UnblockCourier)
	r.PUT("/v1/couriers", handlerV1.UpdateCourier)
	r.PUT("/v1/couriers/courier-details", handlerV1.UpdateCourierDetails)
	r.DELETE("/v1/couriers/:courier_id", handlerV1.DeleteCourier)
	r.POST("/v1/couriers/check-login", handlerV1.CheckCourierLogin)
	r.POST("/v1/couriers/confirm-login", handlerV1.ConfirmCourierLogin)

	//Vehicle endpoints
	r.GET("/v1/vehicles/:vehicle_id", handlerV1.GetCourierVehicle)
	r.POST("/v1/vehicles", handlerV1.CreateCourierVehicle)
	r.PUT("/v1/vehicles", handlerV1.UpdateCourierVehicle)
	r.DELETE("/v1/vehicles/:vehicle_id", handlerV1.DeleteCourierVehicle)

	//Distributor endpoints
	r.GET("/v1/distributors", handlerV1.GetAllDistributors)
	r.GET("/v1/distributors/:distributor_id", handlerV1.GetDistributor)
	r.GET("/v1/distributors/:distributor_id/couriers", handlerV1.GetAllDistributorCouriers)
	r.GET("/v1/distributors/:distributor_id/parks", handlerV1.GetAllDistributorParks)
	r.POST("/v1/distributors", handlerV1.CreateDistributor)
	r.PUT("/v1/distributors", handlerV1.UpdateDistributor)
	r.DELETE("/v1/distributors/:distributor_id", handlerV1.DeleteDistributor)

	//Park endpoints
	r.GET("/v1/parks/:park_id", handlerV1.GetPark)
	r.POST("/v1/parks", handlerV1.CreatePark)
	r.PUT("/v1/parks", handlerV1.UpdatePark)
	r.DELETE("/v1/parks/:park_id", handlerV1.DeletePark)

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
	r.POST("/v1/demand-order", handlerV1.CreateDemandOrder)
	r.POST("/v1/ondemand-order", handlerV1.CreateOnDemandOrder)
	r.PUT("/v1/order/:order_id", handlerV1.UpdateOrder)
	r.GET("/v1/order/:order_id", handlerV1.GetOrder)
	r.GET("/v1/order", handlerV1.GetOrders)
	r.GET("/v1/new-order", handlerV1.CourierNewOrders)
	r.PATCH("/v1/order/:order_id/change-status", handlerV1.ChangeOrderStatus)
	r.GET("/v1/order-statuses", handlerV1.GetStatuses)
	r.PATCH("/v1/order/:order_id/add-courier", handlerV1.AddCourier)
	r.PATCH("/v1/order/:order_id/remove-courier", handlerV1.RemoveCourier)
	r.GET("/v1/courier/order", handlerV1.GetCourierOrders)
	r.GET("/v1/customers/:customer_id/orders", handlerV1.GetCustomerOrders)
	r.GET("/v1/branches/:branch_id/orders", handlerV1.GetBranchOrders)
	r.PATCH("/v1/order-step/:step_id/take", handlerV1.TakeOrderStep)
	r.GET("/v1/customer-addresses/:phone", handlerV1.GetCustomerAddresses)
	r.PATCH("/v1/order/:order_id/add-branch", handlerV1.AddBranchID)

	//Cargo owner
	//r.POST("/v1/cargo-owner", handlerV1.CreateCO)
	//r.GET("/v1/cargo-owner/", handlerV1.GetCO)
	//r.POST("/v1/cargo-owner/check-name", handlerV1.CheckCOName)
	//r.POST("/v1/cargo-owner/check-login", handlerV1.CheckLogin)
	//r.POST("/v1/cargo-owner/refresh-token", handlerV1.RefreshToken)
	//r.POST("/v1/cargo-owner/change-credentials", handlerV1.ChangeLoginPassword)

	//Login endpoints
	r.POST("/v1/check_code/")

	//Specification endpoints
	r.POST("/v1/specification", handlerV1.CreateSpecification)
	r.GET("/v1/specification", handlerV1.GetAllSpecification)

	//Product kind endpoints
	r.POST("/v1/product-kind", handlerV1.CreateProductKind)
	r.GET("/v1/product-kind", handlerV1.GetAllProductKind)

	//Measure endpoints
	r.POST("/v1/measure", handlerV1.CreateMeasure)
	r.GET("/v1/measure", handlerV1.GetAllMeasure)

	//Category endpoints
	r.POST("/v1/category", handlerV1.CreateCategory)
	r.GET("/v1/category", handlerV1.GetAllCategory)
	r.PUT("/v1/category/:category_id", handlerV1.UpdateCategory)
	r.DELETE("/v1/category/:category_id", handlerV1.DeleteCategory)

	//Product Service
	r.POST("/v1/product", handlerV1.CreateProduct)
	r.GET("/v1/product", handlerV1.GetAllProducts)
	r.GET("/v1/product/:product_id", handlerV1.GetProduct)
	r.PUT("/v1/product/:product_id", handlerV1.UpdateProduct)
	r.DELETE("/v1/product/:product_id", handlerV1.DeleteProduct)

	//Upload File
	r.POST("/v1/upload", handlerV1.ImageUpload)

	//Auth
	r.POST("/v1/login", handlerV1.Login)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
