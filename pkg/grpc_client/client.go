package grpc_client

import (
	"fmt"
	pba "genproto/auth_service"
	pb "genproto/catalog_service"
	pbc "genproto/courier_service"
	pbf "genproto/fare_service"
	pbn "genproto/notification_service"
	pbo "genproto/order_service"
	pbr "genproto/report_service"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"

	"google.golang.org/grpc"

	"bitbucket.org/alien_soft/api_getaway/config"
)

//GrpcClientI ...
type GrpcClientI interface {
	CourierService() pbc.CourierServiceClient
	DistributorService() pbc.DistributorServiceClient
	FareService() pbf.FareServiceClient
	OrderService() pbo.OrderServiceClient
	SmsService() pbs.SmsServiceClient
	CustomerService() pbu.CustomerServiceClient
	BranchService() pbu.BranchServiceClient
	ShipperService() pbu.ShipperServiceClient
	SpecificationService() pb.SpecificationServiceClient
	ProductKindService() pb.ProductKindServiceClient
	MeasureService() pb.MeasureServiceClient
	CategoryService() pb.CategoryServiceClient
	ProductService() pb.ProductServiceClient
	AuthService() pba.AuthServiceClient
	NotificationService() pbn.NotificationServiceClient
	ReportService() pbr.ReportServiceClient
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

	if err != nil {
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

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%d err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}

	connAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("auth service dial host: %s port:%d err: %s",
			cfg.AuthServiceHost, cfg.AuthServicePort, err)
	}

	connNotification, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.NotificationServiceHost, cfg.NotificationServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("notification service dial host: %s port:%d err: %s",
			cfg.NotificationServiceHost, cfg.NotificationServicePort, err)
	}

	connReport, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ReportServiceHost, cfg.ReportServicePort),
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, fmt.Errorf("Report service dial host: %s port:%d err: %s",
			cfg.ReportServiceHost, cfg.ReportServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"courier_service":       pbc.NewCourierServiceClient(connCourier),
			"distributor_service":   pbc.NewDistributorServiceClient(connCourier),
			"fare_service":          pbf.NewFareServiceClient(connFare),
			"order_service":         pbo.NewOrderServiceClient(connOrder),
			"sms_service":           pbs.NewSmsServiceClient(connSms),
			"customer_service":      pbu.NewCustomerServiceClient(connUser),
			"branch_service":        pbu.NewBranchServiceClient(connUser),
			"shipper_service":       pbu.NewShipperServiceClient(connUser),
			"system_user_service":   pbu.NewSystemUserServiceClient(connUser),
			"specification_service": pb.NewSpecificationServiceClient(connCatalog),
			"product_kind_service":  pb.NewProductKindServiceClient(connCatalog),
			"measure_service":       pb.NewMeasureServiceClient(connCatalog),
			"category_service":      pb.NewCategoryServiceClient(connCatalog),
			"product_service":       pb.NewProductServiceClient(connCatalog),
			"auth_service":          pba.NewAuthServiceClient(connAuth),
			"platform_service":      pba.NewPlatformServiceClient(connAuth),
			"notification_service":  pbn.NewNotificationServiceClient(connNotification),
			"report_service":        pbr.NewReportServiceClient(connReport),
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

//SmsService ...
func (g *GrpcClient) SmsService() pbs.SmsServiceClient {
	return g.connections["sms_service"].(pbs.SmsServiceClient)
}

//CustomerService ...
func (g *GrpcClient) CustomerService() pbu.CustomerServiceClient {
	return g.connections["customer_service"].(pbu.CustomerServiceClient)
}

//BranchService ...
func (g *GrpcClient) BranchService() pbu.BranchServiceClient {
	return g.connections["branch_service"].(pbu.BranchServiceClient)
}

//SpecificationService ...
func (g *GrpcClient) SpecificationService() pb.SpecificationServiceClient {
	return g.connections["specification_service"].(pb.SpecificationServiceClient)
}

//ProductKindService ...
func (g *GrpcClient) ProductKindService() pb.ProductKindServiceClient {
	return g.connections["product_kind_service"].(pb.ProductKindServiceClient)
}

func (g *GrpcClient) MeasureService() pb.MeasureServiceClient {
	return g.connections["measure_service"].(pb.MeasureServiceClient)
}

func (g *GrpcClient) CategoryService() pb.CategoryServiceClient {
	return g.connections["category_service"].(pb.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pb.ProductServiceClient {
	return g.connections["product_service"].(pb.ProductServiceClient)
}

func (g *GrpcClient) ShipperService() pbu.ShipperServiceClient {
	return g.connections["shipper_service"].(pbu.ShipperServiceClient)
}

func (g *GrpcClient) SystemUserService() pbu.SystemUserServiceClient {
	return g.connections["system_user_service"].(pbu.SystemUserServiceClient)
}

func (g *GrpcClient) AuthService() pba.AuthServiceClient {
	return g.connections["auth_service"].(pba.AuthServiceClient)
}

func (g *GrpcClient) PlatformService() pba.PlatformServiceClient {
	return g.connections["platform_service"].(pba.PlatformServiceClient)
}

func (g *GrpcClient) NotificationService() pbn.NotificationServiceClient {
	return g.connections["notification_service"].(pbn.NotificationServiceClient)
}

func (g *GrpcClient) ReportService() pbr.ReportServiceClient {
	return g.connections["report_service"].(pbr.ReportServiceClient)
}
