// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order_service

import (
	fmt "fmt"
	courier_service "genproto/courier_service"
	fare_service "genproto/fare_service"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Location struct {
	Long                 float32  `protobuf:"fixed32,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat                  float32  `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLong() float32 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

type Order struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerName         string                   `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	ToLocation           *Location                `protobuf:"bytes,3,opt,name=to_location,json=toLocation,proto3" json:"to_location,omitempty"`
	ToAddress            string                   `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	CourierId            *wrappers.StringValue    `protobuf:"bytes,5,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	FareId               string                   `protobuf:"bytes,6,opt,name=fare_id,json=fareId,proto3" json:"fare_id,omitempty"`
	StatusId             string                   `protobuf:"bytes,7,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	CreatedAt            string                   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description          string                   `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	CoId                 string                   `protobuf:"bytes,10,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	CreatorTypeId        string                   `protobuf:"bytes,11,opt,name=creator_type_id,json=creatorTypeId,proto3" json:"creator_type_id,omitempty"`
	UserId               string                   `protobuf:"bytes,12,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Steps                []*Step                  `protobuf:"bytes,13,rep,name=steps,proto3" json:"steps,omitempty"`
	Fare                 *fare_service.Fare       `protobuf:"bytes,14,opt,name=fare,proto3" json:"fare,omitempty"`
	CoDeliveryPrice      float32                  `protobuf:"fixed32,15,opt,name=co_delivery_price,json=coDeliveryPrice,proto3" json:"co_delivery_price,omitempty"`
	DeliveryPrice        float32                  `protobuf:"fixed32,16,opt,name=delivery_price,json=deliveryPrice,proto3" json:"delivery_price,omitempty"`
	Courier              *courier_service.Courier `protobuf:"bytes,17,opt,name=courier,proto3" json:"courier,omitempty"`
	CustomerPhoneNumber  string                   `protobuf:"bytes,18,opt,name=customer_phone_number,json=customerPhoneNumber,proto3" json:"customer_phone_number,omitempty"`
	FinishedAt           string                   `protobuf:"bytes,19,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	OrderAmount          float32                  `protobuf:"fixed32,20,opt,name=order_amount,json=orderAmount,proto3" json:"order_amount,omitempty"`
	ExternalOrderId      uint64                   `protobuf:"varint,21,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	CustomerId           *wrappers.StringValue    `protobuf:"bytes,22,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	VendorId             *wrappers.StringValue    `protobuf:"bytes,23,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *Order) GetToLocation() *Location {
	if m != nil {
		return m.ToLocation
	}
	return nil
}

func (m *Order) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *Order) GetCourierId() *wrappers.StringValue {
	if m != nil {
		return m.CourierId
	}
	return nil
}

func (m *Order) GetFareId() string {
	if m != nil {
		return m.FareId
	}
	return ""
}

func (m *Order) GetStatusId() string {
	if m != nil {
		return m.StatusId
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Order) GetCoId() string {
	if m != nil {
		return m.CoId
	}
	return ""
}

func (m *Order) GetCreatorTypeId() string {
	if m != nil {
		return m.CreatorTypeId
	}
	return ""
}

func (m *Order) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Order) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Order) GetFare() *fare_service.Fare {
	if m != nil {
		return m.Fare
	}
	return nil
}

func (m *Order) GetCoDeliveryPrice() float32 {
	if m != nil {
		return m.CoDeliveryPrice
	}
	return 0
}

func (m *Order) GetDeliveryPrice() float32 {
	if m != nil {
		return m.DeliveryPrice
	}
	return 0
}

func (m *Order) GetCourier() *courier_service.Courier {
	if m != nil {
		return m.Courier
	}
	return nil
}

func (m *Order) GetCustomerPhoneNumber() string {
	if m != nil {
		return m.CustomerPhoneNumber
	}
	return ""
}

func (m *Order) GetFinishedAt() string {
	if m != nil {
		return m.FinishedAt
	}
	return ""
}

func (m *Order) GetOrderAmount() float32 {
	if m != nil {
		return m.OrderAmount
	}
	return 0
}

func (m *Order) GetExternalOrderId() uint64 {
	if m != nil {
		return m.ExternalOrderId
	}
	return 0
}

func (m *Order) GetCustomerId() *wrappers.StringValue {
	if m != nil {
		return m.CustomerId
	}
	return nil
}

func (m *Order) GetVendorId() *wrappers.StringValue {
	if m != nil {
		return m.VendorId
	}
	return nil
}

type Step struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchName           string     `protobuf:"bytes,2,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Location             *Location  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Address              string     `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	DestinationAddress   string     `protobuf:"bytes,5,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	PhoneNumber          string     `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Products             []*Product `protobuf:"bytes,7,rep,name=products,proto3" json:"products,omitempty"`
	Description          string     `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	OrderNo              uint64     `protobuf:"varint,9,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	StepAmount           float32    `protobuf:"fixed32,11,opt,name=step_amount,json=stepAmount,proto3" json:"step_amount,omitempty"`
	ExternalOrderId      uint64     `protobuf:"varint,12,opt,name=external_order_id,json=externalOrderId,proto3" json:"external_order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Step) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Step) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Step) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Step) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *Step) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Step) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Step) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Step) GetOrderNo() uint64 {
	if m != nil {
		return m.OrderNo
	}
	return 0
}

func (m *Step) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Step) GetStepAmount() float32 {
	if m != nil {
		return m.StepAmount
	}
	return 0
}

func (m *Step) GetExternalOrderId() uint64 {
	if m != nil {
		return m.ExternalOrderId
	}
	return 0
}

type Product struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity             float32               `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                float32               `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalAmount          float32               `protobuf:"fixed32,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	ProductId            *wrappers.StringValue `protobuf:"bytes,6,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Product) GetQuantity() float32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetTotalAmount() float32 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *Product) GetProductId() *wrappers.StringValue {
	if m != nil {
		return m.ProductId
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "genproto.Location")
	proto.RegisterType((*Order)(nil), "genproto.Order")
	proto.RegisterType((*Step)(nil), "genproto.Step")
	proto.RegisterType((*Product)(nil), "genproto.Product")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x85, 0x1d, 0x39, 0x92, 0xaf, 0x9c, 0xb8, 0xa1, 0xdb, 0x84, 0xcb, 0xd6, 0xd5, 0xf5, 0x3e,
	0x60, 0x6c, 0x98, 0x5d, 0xa4, 0x4f, 0x43, 0xb1, 0x07, 0x6f, 0xc3, 0x00, 0x01, 0x43, 0x16, 0xa8,
	0xc3, 0x1e, 0xf6, 0x22, 0xd0, 0x22, 0xe3, 0x10, 0x90, 0x45, 0x8d, 0xa2, 0xb2, 0xf9, 0x75, 0x3f,
	0x66, 0xff, 0x6a, 0xff, 0x65, 0xe0, 0xa5, 0xa8, 0x3a, 0xcd, 0xd0, 0xe4, 0x4d, 0xf7, 0x9c, 0x43,
	0xf2, 0x7e, 0x1d, 0x1b, 0x62, 0xa5, 0xb9, 0xd0, 0x8b, 0x4a, 0x2b, 0xa3, 0x48, 0xb4, 0x11, 0x25,
	0x7e, 0x9d, 0x7f, 0xba, 0x51, 0x6a, 0x53, 0x88, 0x25, 0x46, 0xeb, 0xe6, 0x7a, 0xf9, 0xa7, 0x66,
	0x55, 0x25, 0x74, 0xed, 0x94, 0xe7, 0x67, 0xd7, 0x4c, 0x8b, 0xac, 0x16, 0xfa, 0x56, 0xe6, 0x62,
	0x69, 0x83, 0x96, 0x78, 0x9e, 0xab, 0x46, 0x4b, 0xa1, 0x3b, 0xae, 0x8d, 0x1d, 0x3d, 0x7b, 0x05,
	0xd1, 0xcf, 0x2a, 0x67, 0x46, 0xaa, 0x92, 0x10, 0x08, 0x0a, 0x55, 0x6e, 0x68, 0x6f, 0xda, 0x9b,
	0xf7, 0x53, 0xfc, 0x26, 0x4f, 0xe0, 0xa0, 0x60, 0x86, 0xf6, 0x11, 0xb2, 0x9f, 0xb3, 0xbf, 0x43,
	0x18, 0xfc, 0x62, 0x73, 0x24, 0xc7, 0xd0, 0x97, 0x1c, 0xd5, 0xc3, 0xb4, 0x2f, 0x39, 0xf9, 0x0c,
	0x8e, 0xf2, 0xa6, 0x36, 0x6a, 0x2b, 0x74, 0x56, 0xb2, 0xad, 0xc0, 0x53, 0xc3, 0x74, 0xe4, 0xc1,
	0x4b, 0xb6, 0x15, 0xe4, 0x35, 0xc4, 0x46, 0x65, 0x45, 0xfb, 0x26, 0x3d, 0x98, 0xf6, 0xe6, 0xf1,
	0x05, 0x59, 0xf8, 0x42, 0x17, 0x3e, 0x9b, 0x14, 0x8c, 0xea, 0x32, 0x7b, 0x0e, 0x60, 0x54, 0xc6,
	0x38, 0xd7, 0xa2, 0xae, 0x69, 0x80, 0xd7, 0x0e, 0x8d, 0x5a, 0x39, 0x80, 0xbc, 0x01, 0xf0, 0x55,
	0x4a, 0x4e, 0x07, 0x78, 0xe5, 0x27, 0x0b, 0xd7, 0xb1, 0x85, 0xef, 0xd8, 0xe2, 0xad, 0xd1, 0xb2,
	0xdc, 0xfc, 0xc6, 0x8a, 0x46, 0xa4, 0xc3, 0x56, 0x9f, 0x70, 0x72, 0x06, 0x21, 0xf6, 0x4e, 0x72,
	0x7a, 0x88, 0x17, 0x1f, 0xda, 0x30, 0xe1, 0xe4, 0x63, 0x18, 0xd6, 0x86, 0x99, 0xa6, 0xb6, 0x54,
	0x88, 0x54, 0xe4, 0x80, 0x84, 0xdb, 0x8c, 0x72, 0x2d, 0x98, 0x11, 0x3c, 0x63, 0x86, 0x46, 0x2e,
	0xa3, 0x16, 0x59, 0x19, 0x32, 0x85, 0x98, 0x8b, 0x3a, 0xd7, 0xb2, 0xc2, 0x2a, 0x87, 0xc8, 0xef,
	0x43, 0x64, 0x02, 0x83, 0x5c, 0xd9, 0x9b, 0x01, 0xb9, 0x20, 0x57, 0x09, 0x27, 0x5f, 0xc2, 0x18,
	0xef, 0x50, 0x3a, 0x33, 0xbb, 0x0a, 0x73, 0x8a, 0x91, 0x3e, 0x6a, 0xe1, 0x5f, 0x77, 0x95, 0x70,
	0x39, 0x37, 0xb5, 0xab, 0x76, 0xe4, 0x72, 0xb6, 0x61, 0xc2, 0xc9, 0xe7, 0x30, 0xa8, 0x8d, 0xa8,
	0x6a, 0x7a, 0x34, 0x3d, 0x98, 0xc7, 0x17, 0xc7, 0xef, 0xfa, 0xfa, 0xd6, 0x88, 0x2a, 0x75, 0x24,
	0x99, 0x41, 0x60, 0x6b, 0xa4, 0xc7, 0xd8, 0xa9, 0x3d, 0xd1, 0x4f, 0x4c, 0x8b, 0x14, 0x39, 0xf2,
	0x15, 0x9c, 0xe4, 0x2a, 0xe3, 0xa2, 0x90, 0xb7, 0x42, 0xef, 0xb2, 0x4a, 0xcb, 0x5c, 0xd0, 0x31,
	0xae, 0xc1, 0x38, 0x57, 0x3f, 0xb6, 0xf8, 0x95, 0x85, 0xc9, 0x17, 0x70, 0xfc, 0x9e, 0xf0, 0x09,
	0x0a, 0x8f, 0xf8, 0x1d, 0xd9, 0xd7, 0x10, 0xb6, 0x6d, 0xa7, 0x27, 0xf8, 0xf2, 0xc9, 0xbb, 0x97,
	0x7f, 0x70, 0x44, 0xea, 0x15, 0xe4, 0x02, 0x9e, 0x75, 0xcb, 0x54, 0xdd, 0xa8, 0x52, 0x64, 0x65,
	0xb3, 0x5d, 0x0b, 0x4d, 0x09, 0x16, 0x3c, 0xf1, 0xe4, 0x95, 0xe5, 0x2e, 0x91, 0x22, 0x2f, 0x20,
	0xbe, 0x96, 0xa5, 0xac, 0x6f, 0xdc, 0x54, 0x26, 0xa8, 0x04, 0x0f, 0xad, 0x0c, 0x79, 0x09, 0x23,
	0xb4, 0x57, 0xc6, 0xb6, 0xaa, 0x29, 0x0d, 0x7d, 0x8a, 0x69, 0x3a, 0xcb, 0xad, 0x10, 0xb2, 0x75,
	0x8b, 0xbf, 0x8c, 0xd0, 0x25, 0x2b, 0x32, 0xa7, 0x95, 0x9c, 0x3e, 0x9b, 0xf6, 0xe6, 0x41, 0x3a,
	0xf6, 0x04, 0xae, 0x7f, 0xc2, 0xc9, 0x77, 0x10, 0x77, 0x39, 0x4a, 0x4e, 0x4f, 0x1f, 0xb1, 0x78,
	0xe0, 0x0f, 0x24, 0x9c, 0x7c, 0x0b, 0xc3, 0x5b, 0x51, 0x72, 0x85, 0x87, 0xcf, 0x1e, 0x71, 0x38,
	0x72, 0xf2, 0x84, 0xcf, 0xfe, 0x39, 0x80, 0xc0, 0x4e, 0xf4, 0x9e, 0x07, 0x5f, 0x40, 0xbc, 0xd6,
	0xac, 0xcc, 0x6f, 0xf6, 0x1d, 0x08, 0x0e, 0x42, 0xff, 0x2d, 0x20, 0x7a, 0x84, 0xf9, 0x3a, 0x0d,
	0xa1, 0x10, 0xde, 0xf5, 0x9d, 0x0f, 0xc9, 0x12, 0x26, 0x5c, 0xd4, 0x46, 0x96, 0x28, 0xec, 0xdc,
	0x39, 0x40, 0x15, 0xd9, 0xa3, 0xbc, 0x4d, 0x5f, 0xc2, 0xe8, 0xce, 0x24, 0x9d, 0xdd, 0xe2, 0x6a,
	0x6f, 0x82, 0xdf, 0x40, 0x54, 0x69, 0xc5, 0x9b, 0xdc, 0xd4, 0x34, 0xc4, 0x15, 0xde, 0xdb, 0x91,
	0x2b, 0xc7, 0xa4, 0x9d, 0xe4, 0x7d, 0x9b, 0x45, 0xf7, 0x6d, 0xf6, 0x11, 0x44, 0x6e, 0x8a, 0xa5,
	0x42, 0x17, 0x06, 0x69, 0x88, 0xf1, 0xa5, 0x22, 0xa7, 0x70, 0xe8, 0xec, 0xdc, 0x5a, 0xb0, 0x8d,
	0x6c, 0x0b, 0xad, 0x4d, 0xfc, 0x8e, 0xc4, 0xb8, 0x23, 0x60, 0xa1, 0x0f, 0xad, 0xc8, 0xe8, 0x7f,
	0x57, 0x64, 0xf6, 0x6f, 0x0f, 0xc2, 0x36, 0xef, 0x7b, 0xb3, 0x7a, 0x05, 0x41, 0x37, 0xa4, 0x87,
	0x46, 0x8f, 0x4a, 0x72, 0x0e, 0xd1, 0x1f, 0x0d, 0x2b, 0x8d, 0x34, 0x3b, 0x1c, 0x5e, 0x3f, 0xed,
	0x62, 0xf2, 0x14, 0x06, 0xce, 0x7b, 0x01, 0x12, 0x2e, 0xb0, 0x3d, 0x37, 0xca, 0xb0, 0xc2, 0x57,
	0x33, 0x70, 0x1b, 0x8f, 0x58, 0x5b, 0xce, 0x1b, 0x80, 0xb6, 0xa1, 0xfe, 0x37, 0xf0, 0xc1, 0x5f,
	0xcf, 0x56, 0x9f, 0xf0, 0xef, 0xe9, 0xef, 0xa7, 0x7e, 0x3e, 0x4b, 0xd7, 0x8b, 0xf6, 0x7f, 0x66,
	0x7d, 0x88, 0xe0, 0xeb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xa6, 0x61, 0x3d, 0xd1, 0x06,
	0x00, 0x00,
}
