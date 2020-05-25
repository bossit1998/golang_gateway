package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	RoleCargoOwnerAdmin  = "cargo_owner_admin"
	RoleAdmin            = "admin"
	RoleDistributorAdmin = "distributor_admin"
	RoleCargoAPI         = "cargo_api"
	RoleCourier          = "courier"
	RoleUnknown          = "unknown"
	NEW_STATUS_ID        = "986a0d09-7b4d-4ca9-8567-aa1c6d770505"
	CANCELLED_STATUS_ID  = "6ba783a3-1c2e-479c-9626-25526b3d9d36"
	ACCEPTED_STATUS_ID   = "8781af8e-f74d-4fb6-ae23-fd997f4a2ee0"
	PICKED_UP_STATUS_ID  = "84be5a2f-3a92-4469-8283-220ca34a0de4"
	DELIVERED_STATUS_ID  = "79413606-a56f-45ed-97c3-f3f18e645972"
	FINISHED_STATUS_ID   = "e665273d-5415-4243-a329-aee410e39465"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	RedisHost string
	RedisPort int

	GeoServiceHost string
	GeoServicePort int

	UserServiceHost string
	UserServicePort int

	CourierServiceHost string
	CourierServicePort int

	FareServiceHost string
	FareServicePort int

	OrderServiceHost string
	OrderServicePort int

	COServiceHost string
	COServicePort int

	SmsServiceHost string
	SmsServicePort int

	UserServiceHost string
	UserServicePort int

	CatalogServiceHost string
	CatalogServicePort int

	LogLevel string
	HTTPPort string

	CasbinConfigPath string
	MapboxToken      string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "127.0.0.1"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "deleverdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "delever"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "delever"))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "redis"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1235"))

	c.GeoServiceHost = cast.ToString(getOrReturnDefault("GEO_SERVICE_HOST", "geo_service"))
	c.GeoServicePort = cast.ToInt(getOrReturnDefault("GEO_SERVICE_PORT", 80))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user_service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 80))

	c.CourierServiceHost = cast.ToString(getOrReturnDefault("COURIER_SERVICE_HOST", "courier_service"))
	c.CourierServicePort = cast.ToInt(getOrReturnDefault("COURIER_SERVICE_PORT", 80))

	c.FareServiceHost = cast.ToString(getOrReturnDefault("FARE_SERVICE_HOST", "fare_service"))
	c.FareServicePort = cast.ToInt(getOrReturnDefault("FARE_SERVICE_PORT", 80))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "order_service"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 80))

	c.COServiceHost = cast.ToString(getOrReturnDefault("CO_SERVICE_HOST", "co_service"))
	c.COServicePort = cast.ToInt(getOrReturnDefault("CO_SERVICE_PORT", 80))

	c.SmsServiceHost = cast.ToString(getOrReturnDefault("SMS_SERVICE_HOST", "sms_service"))
	c.SmsServicePort = cast.ToInt(getOrReturnDefault("SMS_SERVICE_PORT", 80))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user_service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 80))

	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_PORT", "catalog_service"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefault("CATALOG_SERVICE_PORT", 80))

	c.CasbinConfigPath = cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf"))

	c.MapboxToken = cast.ToString(getOrReturnDefault("MAPBOX_TOKEN", "pk.eyJ1IjoidGRvc3RvbiIsImEiOiJjazh0cmRrMnowMWszM29sc2Y5c3A5NTZ4In0.mtrOXD4cD4QKZ-dnZ_vKdA"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
