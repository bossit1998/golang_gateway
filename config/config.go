package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	//RoleCargoOwner ...
	RoleCargoOwner = "cargo_owner"
	//RoleSuperAdmin ...
	RoleSuperAdmin = "super_admin"
	RoleCourier = "courier"
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

	LogLevel string
	HTTPPort string

	CasbinConfigPath string
	MapboxToken string
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
