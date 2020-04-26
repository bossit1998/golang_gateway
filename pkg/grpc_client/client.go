package grpc_client

import (
	"fmt"

	"bitbucket.org/alien_soft/api_getaway/config"
	pbc "genproto/courier_service"
	pbf "genproto/fare_service"
	pbo "genproto/order_service"
	pbco "genproto/co_service"
	pbs "genproto/sms_service"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	CourierService() pbc.CourierServiceClient
	DistributorService() pbc.DistributorServiceClient
	FareService() pbf.FareServiceClient
	OrderService() pbo.OrderServiceClient
	SmsService() pbs.SmsServiceClient
}

//GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connCourier, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CourierServiceHost, cfg.CourierServicePort),
		grpc.WithInsecure())
	fmt.Println(connCourier)

	if err != nil {
		fmt.Println("error yoq")
		return nil, fmt.Errorf("courier service dial host: %s port:%d err: %s",
			cfg.CourierServiceHost, cfg.CourierServicePort, err)
	}

	connFare, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.FareServiceHost, cfg.FareServicePort),
		grpc.WithInsecure())

	if err != nil {
		return nil, fmt.Errorf("fare service dial host: %s port:%d err: %s",
			cfg.FareServiceHost, cfg.FareServicePort, err)
	}

	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("order service dial host: %s port:%d err: %s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	connCO, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.COServiceHost, cfg.COServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("cargo_owner service dial host: %s port:%d err: %s",
			cfg.COServiceHost, cfg.COServicePort, err)
	}

	connSms, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.SmsServiceHost, cfg.SmsServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("sms service dial host: %s port:%d err: %s",
			cfg.COServiceHost, cfg.COServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"courier_service":     pbc.NewCourierServiceClient(connCourier),
			"distributor_service": pbc.NewDistributorServiceClient(connCourier),
			"fare_service":        pbf.NewFareServiceClient(connFare),
			"order_service":       pbo.NewOrderServiceClient(connOrder),
			"co_service":          pbco.NewCOServiceClient(connCO),
			"sms_service":		   pbs.NewSmsServiceClient(connSms),
		},
	}, nil
}

//CourierService ...
func (g *GrpcClient) CourierService() pbc.CourierServiceClient {
	return g.connections["courier_service"].(pbc.CourierServiceClient)
}

//DistributorService ...
func (g *GrpcClient) DistributorService() pbc.DistributorServiceClient {
	return g.connections["distributor_service"].(pbc.DistributorServiceClient)
}

//FareService ...
func (g *GrpcClient) FareService() pbf.FareServiceClient {
	return g.connections["fare_service"].(pbf.FareServiceClient)
}

//OrderService ...
func (g *GrpcClient) OrderService() pbo.OrderServiceClient {
	return g.connections["order_service"].(pbo.OrderServiceClient)
}

//COService ...
func (g *GrpcClient) COService() pbco.COServiceClient {
	return g.connections["co_service"].(pbco.COServiceClient)
}

//SmsService ...
func (g *GrpcClient) SmsService() pbs.SmsServiceClient {
	return g.connections["sms_service"].(pbs.SmsServiceClient)
}
