package grpc_client

import (
	"fmt"

	"bitbucket.org/alien_soft/api_gateway/config"
	pbc "bitbucket.org/alien_soft/api_gateway/genproto/courier_service"
	pbf "bitbucket.org/alien_soft/api_gateway/genproto/fare_service"
	pbg "bitbucket.org/alien_soft/api_gateway/genproto/geo_service"
	"google.golang.org/grpc"
)

type GrpcClientI interface {
	GeoService() pbg.GeoServiceClient
	CourierService() pbc.CourierServiceClient
	DistributorService() pbc.DistributorServiceClient
	FareService() pbf.FareServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {

	connGeo, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GeoServiceHost, cfg.GeoServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("geo service dial host: %s port:%d err: %s",
			cfg.GeoServiceHost, cfg.GeoServicePort, err)
	}

	connCourier, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CourierServiceHost, cfg.CourierServicePort),
		grpc.WithInsecure())
	fmt.Println(connCourier)

	if err != nil {
		fmt.Println("error yoq")
		return nil, fmt.Errorf("courier service dial host: %s port:%d err: %s",
			cfg.CourierServiceHost, cfg.CourierServicePort, err)
	}

	connDistributor, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.DistributorServiceHost, cfg.DistributorServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("Distributor service dial host: %s port:%d err: %s",
			cfg.DistributorServiceHost, cfg.DistributorServicePort, err)
	}

	connFare, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.FareServiceHost, cfg.FareServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("fare service dial host: %s port:%d err: %s",
			cfg.FareServiceHost, cfg.FareServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"geo_service":         pbg.NewGeoServiceClient(connGeo),
			"courier_service":     pbc.NewCourierServiceClient(connCourier),
			"distributor_service": pbc.NewDistributorServiceClient(connDistributor),
			"fare_service":        pbf.NewFareServiceClient(connFare),
		},
	}, nil
}

func (g *GrpcClient) GeoService() pbg.GeoServiceClient {
	return g.connections["geo_service"].(pbg.GeoServiceClient)
}

func (g *GrpcClient) CourierService() pbc.CourierServiceClient {
	return g.connections["courier_service"].(pbc.CourierServiceClient)
}

func (g *GrpcClient) DistributorService() pbc.DistributorServiceClient {
	return g.connections["distributor_service"].(pbc.DistributorServiceClient)
}

func (g *GrpcClient) FareService() pbf.FareServiceClient {
	return g.connections["fare_service"].(pbf.FareServiceClient)
}
