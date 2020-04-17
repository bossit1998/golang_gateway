package config

import (
	"os"

	"github.com/spf13/cast"
)

var (
	FollowTypeID = "086032bc-fae4-4a85-aedd-83e37b52f309"
	LikeTypeId   = "23f0ed87-6874-40be-993a-140fa31d7050"
	SystemTypeID = "70e21be4-d511-4284-b8e2-f1d94205ef1e"
	NewStatusId  = "2cb46139-7c5f-4e0d-88f6-50fe79986fa6"
	SentStatusId = "f9c31d41-7e91-42a1-859b-e0991357566c"
	ReadStatusId = "303d5102-7595-4ba4-9c67-0c5d969d6a3e"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	GeoServiceHost string
	GeoServicePort int

	CourierServiceHost string
	CourierServicePort int

	FareServiceHost string
	FareServicePort int

	OrderServiceHost string
	OrderServicePort int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "127.0.0.1"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "courier_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "12345"))

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

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
