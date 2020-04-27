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
	BranchId             string                   `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
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
	CustomerName         string                   `protobuf:"bytes,18,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	CustomerPhoneNumber  string                   `protobuf:"bytes,19,opt,name=customer_phone_number,json=customerPhoneNumber,proto3" json:"customer_phone_number,omitempty"`
	FinishedAt           string                   `protobuf:"bytes,20,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
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

func (m *Order) GetBranchId() string {
	if m != nil {
		return m.BranchId
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

func (m *Order) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
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

type Product struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity             float32  `protobuf:"fixed32,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                float32  `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalAmount          float32  `protobuf:"fixed32,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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

func init() {
	proto.RegisterType((*Location)(nil), "genproto.Location")
	proto.RegisterType((*Order)(nil), "genproto.Order")
	proto.RegisterType((*Step)(nil), "genproto.Step")
	proto.RegisterType((*Product)(nil), "genproto.Product")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xbb, 0x74, 0x4d, 0x6f, 0xd6, 0x6e, 0x73, 0x07, 0x33, 0x83, 0xb1, 0x52, 0x3e, 0x34,
	0x81, 0x48, 0xd1, 0xf6, 0xc8, 0x53, 0x01, 0x21, 0x55, 0x42, 0x63, 0xca, 0x10, 0x0f, 0xbc, 0x44,
	0x6e, 0xec, 0x75, 0x96, 0xd2, 0x38, 0xd8, 0xce, 0x50, 0x5f, 0xf9, 0x67, 0xfc, 0x09, 0x7e, 0x0f,
	0xca, 0x75, 0x92, 0x75, 0xdb, 0x0b, 0x6f, 0xbe, 0xe7, 0x9c, 0x3a, 0xf7, 0x9c, 0x9c, 0x06, 0x02,
	0xa5, 0xb9, 0xd0, 0x61, 0xae, 0x95, 0x55, 0xc4, 0x5f, 0x88, 0x0c, 0x4f, 0x07, 0x4f, 0x17, 0x4a,
	0x2d, 0x52, 0x31, 0xc1, 0x69, 0x5e, 0x5c, 0x4e, 0x7e, 0x69, 0x96, 0xe7, 0x42, 0x1b, 0xa7, 0x3c,
	0xd8, 0xbf, 0x64, 0x5a, 0xc4, 0x46, 0xe8, 0x6b, 0x99, 0x88, 0x49, 0x39, 0x54, 0xc4, 0x61, 0xa2,
	0x0a, 0x2d, 0x85, 0x6e, 0xb8, 0x6a, 0x76, 0xf4, 0xf8, 0x1d, 0xf8, 0x5f, 0x54, 0xc2, 0xac, 0x54,
	0x19, 0x21, 0xe0, 0xa5, 0x2a, 0x5b, 0xd0, 0xd6, 0xa8, 0x75, 0xdc, 0x8e, 0xf0, 0x4c, 0x76, 0x60,
	0x23, 0x65, 0x96, 0xb6, 0x11, 0x2a, 0x8f, 0xe3, 0xbf, 0x1d, 0xe8, 0x7c, 0x2d, 0x77, 0x24, 0x03,
	0x68, 0x4b, 0x8e, 0xea, 0x5e, 0xd4, 0x96, 0x9c, 0x3c, 0x86, 0xde, 0x5c, 0xb3, 0x2c, 0xb9, 0x8a,
	0x25, 0xc7, 0x5f, 0xf4, 0x22, 0xdf, 0x01, 0x33, 0x4e, 0x4e, 0x21, 0xb0, 0x2a, 0x4e, 0xab, 0x67,
	0xd1, 0x8d, 0x51, 0xeb, 0x38, 0x38, 0x21, 0x61, 0x6d, 0x30, 0xac, 0xb7, 0x88, 0xc0, 0xaa, 0x66,
	0xa3, 0x43, 0x00, 0xab, 0x62, 0xc6, 0xb9, 0x16, 0xc6, 0x50, 0x0f, 0xaf, 0xec, 0x59, 0x35, 0x75,
	0x00, 0x79, 0x0f, 0x50, 0xbb, 0x93, 0x9c, 0x76, 0xf0, 0xca, 0x27, 0xa1, 0x4b, 0x2a, 0xac, 0x93,
	0x0a, 0x2f, 0xac, 0x96, 0xd9, 0xe2, 0x3b, 0x4b, 0x0b, 0x11, 0xf5, 0x2a, 0xfd, 0x8c, 0x93, 0x7d,
	0xe8, 0x62, 0x66, 0x92, 0xd3, 0x4d, 0xbc, 0x78, 0xb3, 0x1c, 0x67, 0x68, 0xc3, 0x58, 0x66, 0x0b,
	0x53, 0x52, 0x5d, 0x67, 0xc3, 0x01, 0x33, 0x5e, 0x6e, 0x94, 0x68, 0xc1, 0xac, 0xe0, 0x31, 0xb3,
	0xd4, 0x77, 0x1b, 0x55, 0xc8, 0xd4, 0x92, 0x11, 0x04, 0x5c, 0x98, 0x44, 0xcb, 0x1c, 0x5d, 0xf6,
	0x90, 0x5f, 0x87, 0xc8, 0x10, 0x3a, 0x89, 0x2a, 0x6f, 0x06, 0xe4, 0xbc, 0x44, 0xcd, 0x38, 0x79,
	0x05, 0xdb, 0x78, 0x87, 0xd2, 0xb1, 0x5d, 0xe5, 0xb8, 0x53, 0x80, 0x74, 0xbf, 0x82, 0xbf, 0xad,
	0x72, 0xe1, 0x76, 0x2e, 0x8c, 0x73, 0xbb, 0xe5, 0x76, 0x2e, 0xc7, 0x19, 0x27, 0x2f, 0xa0, 0x63,
	0xac, 0xc8, 0x0d, 0xed, 0x8f, 0x36, 0x8e, 0x83, 0x93, 0xc1, 0x4d, 0xae, 0x17, 0x56, 0xe4, 0x91,
	0x23, 0xc9, 0x18, 0xbc, 0xd2, 0x23, 0x1d, 0x60, 0x52, 0x6b, 0xa2, 0xcf, 0x4c, 0x8b, 0x08, 0x39,
	0xf2, 0x1a, 0x76, 0x13, 0x15, 0x73, 0x91, 0xca, 0x6b, 0xa1, 0x57, 0x71, 0xae, 0x65, 0x22, 0xe8,
	0x36, 0xbe, 0xfe, 0xed, 0x44, 0x7d, 0xaa, 0xf0, 0xf3, 0x12, 0x26, 0x2f, 0x61, 0x70, 0x47, 0xb8,
	0x83, 0xc2, 0x3e, 0xbf, 0x25, 0x7b, 0x03, 0xdd, 0x2a, 0x76, 0xba, 0x8b, 0x4f, 0xde, 0xbd, 0x79,
	0xf2, 0x47, 0x47, 0x44, 0xb5, 0x82, 0x3c, 0x87, 0x7e, 0x52, 0x18, 0xab, 0x96, 0x42, 0xc7, 0x19,
	0x5b, 0x0a, 0x4a, 0xd0, 0xe8, 0x56, 0x0d, 0x9e, 0xb1, 0xa5, 0x20, 0x27, 0xf0, 0xa0, 0x11, 0xe5,
	0x57, 0x2a, 0x13, 0x71, 0x56, 0x2c, 0xe7, 0x42, 0xd3, 0x21, 0x8a, 0x87, 0x35, 0x79, 0x5e, 0x72,
	0x67, 0x48, 0x91, 0x23, 0x08, 0x2e, 0x65, 0x26, 0xcd, 0x95, 0x7b, 0x75, 0x7b, 0xa8, 0x84, 0x1a,
	0x9a, 0xda, 0xf1, 0x9f, 0x36, 0x78, 0x65, 0x5a, 0xf7, 0x7a, 0x7d, 0x04, 0x41, 0xd5, 0x6b, 0x5c,
	0xc8, 0x35, 0x1b, 0x1c, 0x84, 0xeb, 0x84, 0xe0, 0xff, 0x47, 0xb1, 0x1b, 0x0d, 0xa1, 0xd0, 0xbd,
	0xdd, 0xe9, 0x7a, 0x24, 0x13, 0x18, 0x72, 0x61, 0xac, 0xcc, 0x50, 0xd8, 0x34, 0xbf, 0x83, 0x2a,
	0xb2, 0x46, 0xd5, 0x7f, 0x81, 0x67, 0xb0, 0x75, 0x2b, 0x00, 0x57, 0xe5, 0x20, 0x5f, 0x33, 0xfe,
	0x16, 0xfc, 0x5c, 0x2b, 0x5e, 0x24, 0xd6, 0xd0, 0x2e, 0xd6, 0x63, 0x2d, 0xff, 0x73, 0xc7, 0x44,
	0x8d, 0xe4, 0x6e, 0x85, 0xfd, 0xfb, 0x15, 0x7e, 0x04, 0x3e, 0x7e, 0xa4, 0xe2, 0x4c, 0x61, 0xc3,
	0xbd, 0xa8, 0x8b, 0xf3, 0x99, 0x1a, 0xff, 0x6e, 0x41, 0xb7, 0xba, 0xf2, 0x5e, 0x8c, 0x04, 0xbc,
	0xb5, 0xfc, 0xf0, 0x4c, 0x0e, 0xc0, 0xff, 0x59, 0xb0, 0xcc, 0x4a, 0xbb, 0xc2, 0xe4, 0xda, 0x51,
	0x33, 0x93, 0x3d, 0xe8, 0xb8, 0x52, 0x79, 0x48, 0xb8, 0xa1, 0x34, 0x6c, 0x95, 0x65, 0x69, 0xcc,
	0x96, 0xaa, 0xc8, 0x2c, 0x46, 0xd3, 0x8e, 0x02, 0xc4, 0xa6, 0x08, 0x7d, 0xa0, 0x3f, 0x1e, 0xd6,
	0xfe, 0x26, 0x6e, 0xd1, 0xea, 0xdb, 0x37, 0xdf, 0x44, 0xf0, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb7, 0x48, 0x6f, 0xff, 0x65, 0x05, 0x00, 0x00,
}
