// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: order.proto

package order_service

import (
	courier_service "genproto/courier_service"
	fare_service "genproto/fare_service"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Long float32 `protobuf:"fixed32,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat  float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Location) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientName        string                   `protobuf:"bytes,2,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	ToLocation        *Location                `protobuf:"bytes,3,opt,name=to_location,json=toLocation,proto3" json:"to_location,omitempty"`
	ToAddress         string                   `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	CourierId         *wrappers.StringValue    `protobuf:"bytes,5,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	FareId            string                   `protobuf:"bytes,6,opt,name=fare_id,json=fareId,proto3" json:"fare_id,omitempty"`
	StatusId          string                   `protobuf:"bytes,7,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	CreatedAt         string                   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description       string                   `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	ShipperId         string                   `protobuf:"bytes,10,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	CreatorTypeId     string                   `protobuf:"bytes,11,opt,name=creator_type_id,json=creatorTypeId,proto3" json:"creator_type_id,omitempty"`
	CreatorId         string                   `protobuf:"bytes,12,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Steps             []*Step                  `protobuf:"bytes,13,rep,name=steps,proto3" json:"steps,omitempty"`
	Fare              *fare_service.Fare       `protobuf:"bytes,14,opt,name=fare,proto3" json:"fare,omitempty"`
	CoDeliveryPrice   float32                  `protobuf:"fixed32,15,opt,name=co_delivery_price,json=coDeliveryPrice,proto3" json:"co_delivery_price,omitempty"`
	DeliveryPrice     float32                  `protobuf:"fixed32,16,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	Courier           *courier_service.Courier `protobuf:"bytes,17,opt,name=courier,proto3" json:"courier,omitempty"`
	ClientPhoneNumber string                   `protobuf:"bytes,18,opt,name=client_phone_number,json=clientPhoneNumber,proto3" json:"client_phone_number,omitempty"`
	FinishedAt        string                   `protobuf:"bytes,19,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	OrderAmount       float32                  `protobuf:"fixed32,20,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	ExternalOrderId   *wrappers.Int64Value     `protobuf:"bytes,21,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	ClientId          string                   `protobuf:"bytes,22,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	PaymentType       string                   `protobuf:"bytes,23,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Order) GetToLocation() *Location {
	if x != nil {
		return x.ToLocation
	}
	return nil
}

func (x *Order) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Order) GetCourierId() *wrappers.StringValue {
	if x != nil {
		return x.CourierId
	}
	return nil
}

func (x *Order) GetFareId() string {
	if x != nil {
		return x.FareId
	}
	return ""
}

func (x *Order) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *Order) GetCreatorTypeId() string {
	if x != nil {
		return x.CreatorTypeId
	}
	return ""
}

func (x *Order) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *Order) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Order) GetFare() *fare_service.Fare {
	if x != nil {
		return x.Fare
	}
	return nil
}

func (x *Order) GetCoDeliveryPrice() float32 {
	if x != nil {
		return x.CoDeliveryPrice
	}
	return 0
}

func (x *Order) GetDeliveryPrice() float32 {
	if x != nil {
		return x.DeliveryPrice
	}
	return 0
}

func (x *Order) GetCourier() *courier_service.Courier {
	if x != nil {
		return x.Courier
	}
	return nil
}

func (x *Order) GetClientPhoneNumber() string {
	if x != nil {
		return x.ClientPhoneNumber
	}
	return ""
}

func (x *Order) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *Order) GetOrderAmount() float32 {
	if x != nil {
		return x.OrderAmount
	}
	return 0
}

func (x *Order) GetExternalOrderId() *wrappers.Int64Value {
	if x != nil {
		return x.ExternalOrderId
	}
	return nil
}

func (x *Order) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Order) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchName         *wrappers.StringValue `protobuf:"bytes,2,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Location           *Location             `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Address            *wrappers.StringValue `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	DestinationAddress *wrappers.StringValue `protobuf:"bytes,5,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	PhoneNumber        *wrappers.StringValue `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Products           []*Product            `protobuf:"bytes,7,rep,name=products,proto3" json:"products,omitempty"`
	Description        string                `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	OrderNo            uint64                `protobuf:"varint,9,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	Status             string                `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	StepAmount         float32               `protobuf:"fixed32,11,opt,name=step_amount,json=stepAmount,proto3" json:"step_amount,omitempty"`
	ExternalStepId     *wrappers.Int64Value  `protobuf:"bytes,12,opt,name=external_step_id,json=externalStepId,proto3" json:"external_step_id,omitempty"`
	BranchId           *wrappers.StringValue `protobuf:"bytes,13,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *Step) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Step) GetBranchName() *wrappers.StringValue {
	if x != nil {
		return x.BranchName
	}
	return nil
}

func (x *Step) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Step) GetAddress() *wrappers.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Step) GetDestinationAddress() *wrappers.StringValue {
	if x != nil {
		return x.DestinationAddress
	}
	return nil
}

func (x *Step) GetPhoneNumber() *wrappers.StringValue {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *Step) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Step) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Step) GetOrderNo() uint64 {
	if x != nil {
		return x.OrderNo
	}
	return 0
}

func (x *Step) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Step) GetStepAmount() float32 {
	if x != nil {
		return x.StepAmount
	}
	return 0
}

func (x *Step) GetExternalStepId() *wrappers.Int64Value {
	if x != nil {
		return x.ExternalStepId
	}
	return nil
}

func (x *Step) GetBranchId() *wrappers.StringValue {
	if x != nil {
		return x.BranchId
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity          float32               `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price             float32               `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalAmount       float32               `protobuf:"fixed32,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	ProductId         *wrappers.StringValue `protobuf:"bytes,6,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ExternalProductId *wrappers.Int64Value  `protobuf:"bytes,7,opt,name=external_product_id,json=externalProductId,proto3" json:"external_product_id,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Product) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetTotalAmount() float32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Product) GetProductId() *wrappers.StringValue {
	if x != nil {
		return x.ProductId
	}
	return nil
}

func (x *Product) GetExternalProductId() *wrappers.Int64Value {
	if x != nil {
		return x.ExternalProductId
	}
	return nil
}

type CustomerAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Address  string    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CustomerAddress) Reset() {
	*x = CustomerAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerAddress) ProtoMessage() {}

func (x *CustomerAddress) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerAddress.ProtoReflect.Descriptor instead.
func (*CustomerAddress) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerAddress) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *CustomerAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x66, 0x61, 0x72, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x22, 0xed, 0x06, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0b,
	0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74,
	0x65, 0x70, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73,
	0x12, 0x22, 0x0a, 0x04, 0x66, 0x61, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x61, 0x72, 0x65, 0x52, 0x04,
	0x66, 0x61, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0f, 0x63, 0x6f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x69, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xf4, 0x04, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x4d, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x12, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x74, 0x65, 0x70,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x22, 0xaa, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_proto_goTypes = []interface{}{
	(*Location)(nil),                // 0: genproto.Location
	(*Order)(nil),                   // 1: genproto.Order
	(*Step)(nil),                    // 2: genproto.Step
	(*Product)(nil),                 // 3: genproto.Product
	(*CustomerAddress)(nil),         // 4: genproto.CustomerAddress
	(*wrappers.StringValue)(nil),    // 5: google.protobuf.StringValue
	(*fare_service.Fare)(nil),       // 6: genproto.Fare
	(*courier_service.Courier)(nil), // 7: genproto.Courier
	(*wrappers.Int64Value)(nil),     // 8: google.protobuf.Int64Value
}
var file_order_proto_depIdxs = []int32{
	0,  // 0: genproto.Order.to_location:type_name -> genproto.Location
	5,  // 1: genproto.Order.courier_id:type_name -> google.protobuf.StringValue
	2,  // 2: genproto.Order.steps:type_name -> genproto.Step
	6,  // 3: genproto.Order.fare:type_name -> genproto.Fare
	7,  // 4: genproto.Order.courier:type_name -> genproto.Courier
	8,  // 5: genproto.Order.external_order_id:type_name -> google.protobuf.Int64Value
	5,  // 6: genproto.Step.branch_name:type_name -> google.protobuf.StringValue
	0,  // 7: genproto.Step.location:type_name -> genproto.Location
	5,  // 8: genproto.Step.address:type_name -> google.protobuf.StringValue
	5,  // 9: genproto.Step.destination_address:type_name -> google.protobuf.StringValue
	5,  // 10: genproto.Step.phone_number:type_name -> google.protobuf.StringValue
	3,  // 11: genproto.Step.products:type_name -> genproto.Product
	8,  // 12: genproto.Step.external_step_id:type_name -> google.protobuf.Int64Value
	5,  // 13: genproto.Step.branch_id:type_name -> google.protobuf.StringValue
	5,  // 14: genproto.Product.name:type_name -> google.protobuf.StringValue
	5,  // 15: genproto.Product.product_id:type_name -> google.protobuf.StringValue
	8,  // 16: genproto.Product.external_product_id:type_name -> google.protobuf.Int64Value
	0,  // 17: genproto.CustomerAddress.location:type_name -> genproto.Location
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
