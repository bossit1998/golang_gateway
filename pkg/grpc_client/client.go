package grpc_client

import (
	"fmt"

	"bitbucket.org/alien_soft/api_getaway/config"
	pbc "genproto/courier_service"
	pbf "genproto/fare_service"
	pbo "genproto/order_service"
	pbco "genproto/co_service"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"
	pb "genproto/catalog_service"
	"google.golang.org/grpc"
)

//GrpcClientI ...
type GrpcClientI interface {
	CourierService() pbc.CourierServiceClient
	DistributorService() pbc.DistributorServiceClient
	FareService() pbf.FareServiceClient
	OrderService() pbo.OrderServiceClient
	SmsService() pbs.SmsServiceClient
	UserSerivice() pbu.UserServiceClient
	SpecificationService() pb.SpecificationServiceClient
	ProductKindService() pb.ProductKindServiceClient
	MeasureService() pb.MeasureServiceClient
	CategoryService() pb.CategoryServiceClient
	ProductService() pb.ProductServiceClient
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
			cfg.SmsServiceHost, cfg.SmsServicePort, err)
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithInsecure(),
	)

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"courier_service":     		pbc.NewCourierServiceClient(connCourier),
			"distributor_service": 		pbc.NewDistributorServiceClient(connCourier),
			"fare_service":        	 	pbf.NewFareServiceClient(connFare),
			"order_service":       		pbo.NewOrderServiceClient(connOrder),
			"co_service":          		pbco.NewCOServiceClient(connCO),
			"sms_service":		   		pbs.NewSmsServiceClient(connSms),
			"user_service":		   		pbu.NewUserServiceClient(connUser),
			"specification_service":	pb.NewSpecificationServiceClient(connCatalog),
			"product_kind_service":		pb.NewProductKindServiceClient(connCatalog),
			"measure_service":			pb.NewMeasureServiceClient(connCatalog),
			"category_service":			pb.NewCategoryServiceClient(connCatalog),
			"product_service": 			pb.NewProductServiceClient(connCatalog),
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

//UserService ...
func (g *GrpcClient) UserService() pbu.UserServiceClient {
	return g.connections["user_service"].(pbu.UserServiceClient)
}

//SpecificationService ...
func (g *GrpcClient) SpecificationService() pb.SpecificationServiceClient {
	return g.connections["specification_service"].(pb.SpecificationServiceClient)
}

func (g *GrpcClient) ProductKindService() pb.ProductKindServiceClient {
	return g.connections["product_kind_service"].(pb.ProductKindServiceClient)
}

func (g *GrpcClient) MeasureService() pb.MeasureServiceClient {
	return g.connections["measure_service"].(pb.MeasureServiceClient)
}

func (g *GrpcClient) CategortService() pb.CategoryServiceClient {
	return g.connections["category_service"].(pb.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pb.ProductServiceClient {
	return g.connections["product_service"].(pb.ProductServiceClient)
}