// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order_service

import (
	fmt "fmt"
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
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId             string                `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	FromLocation         *Location             `protobuf:"bytes,3,opt,name=from_location,json=fromLocation,proto3" json:"from_location,omitempty"`
	FromAddress          string                `protobuf:"bytes,4,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToLocation           *Location             `protobuf:"bytes,5,opt,name=to_location,json=toLocation,proto3" json:"to_location,omitempty"`
	ToAddress            string                `protobuf:"bytes,6,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	PhoneNumber          string                `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CourierId            *wrappers.StringValue `protobuf:"bytes,8,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	FareId               string                `protobuf:"bytes,9,opt,name=fare_id,json=fareId,proto3" json:"fare_id,omitempty"`
	StatusId             string                `protobuf:"bytes,10,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	CreatedAt            string                `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Description          string                `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	CoId                 string                `protobuf:"bytes,13,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	CreatorTypeId        string                `protobuf:"bytes,14,opt,name=creator_type_id,json=creatorTypeId,proto3" json:"creator_type_id,omitempty"`
	UserId               string                `protobuf:"bytes,15,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Products             []*Product            `protobuf:"bytes,16,rep,name=products,proto3" json:"products,omitempty"`
	Fare                 *fare_service.Fare    `protobuf:"bytes,17,opt,name=fare,proto3" json:"fare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
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

func (m *Order) GetFromLocation() *Location {
	if m != nil {
		return m.FromLocation
	}
	return nil
}

func (m *Order) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
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

func (m *Order) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
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

func (m *Order) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Order) GetFare() *fare_service.Fare {
	if m != nil {
		return m.Fare
	}
	return nil
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
	return fileDescriptor_cd01338c35d87077, []int{2}
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
	proto.RegisterType((*Product)(nil), "genproto.Product")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077) }

var fileDescriptor_cd01338c35d87077 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x3f, 0x8f, 0xd3, 0x4e,
	0x10, 0x55, 0x7c, 0x4e, 0x62, 0x8f, 0x93, 0xdc, 0xdd, 0xfe, 0x7e, 0xe2, 0xac, 0xf0, 0x47, 0xb9,
	0x14, 0x28, 0x0d, 0x0e, 0xba, 0x2b, 0x28, 0xa8, 0x42, 0x81, 0x64, 0x09, 0x01, 0x32, 0x88, 0x82,
	0xc6, 0xda, 0x78, 0x37, 0xb9, 0x95, 0x1c, 0x8f, 0x59, 0xaf, 0x41, 0x69, 0xf9, 0xac, 0x7c, 0x10,
	0xb4, 0xb3, 0x76, 0x82, 0x84, 0xe8, 0xe6, 0xbd, 0x19, 0xbf, 0x7d, 0xf3, 0xc7, 0x10, 0xa1, 0x16,
	0x52, 0x27, 0xb5, 0x46, 0x83, 0x2c, 0xd8, 0xcb, 0x8a, 0xa2, 0xf9, 0xb3, 0x3d, 0xe2, 0xbe, 0x94,
	0x6b, 0x42, 0xdb, 0x76, 0xb7, 0xfe, 0xa1, 0x79, 0x5d, 0x4b, 0xdd, 0xb8, 0xca, 0xf9, 0xcd, 0x8e,
	0x6b, 0x99, 0x37, 0x52, 0x7f, 0x57, 0x85, 0x5c, 0x5b, 0xe0, 0x12, 0xcb, 0x97, 0x10, 0xbc, 0xc3,
	0x82, 0x1b, 0x85, 0x15, 0x63, 0xe0, 0x97, 0x58, 0xed, 0xe3, 0xc1, 0x62, 0xb0, 0xf2, 0x32, 0x8a,
	0xd9, 0x15, 0x5c, 0x94, 0xdc, 0xc4, 0x1e, 0x51, 0x36, 0x5c, 0xfe, 0xf2, 0x61, 0xf8, 0xc1, 0x9a,
	0x60, 0x33, 0xf0, 0x94, 0xa0, 0xea, 0x30, 0xf3, 0x94, 0x60, 0x8f, 0x21, 0xdc, 0x6a, 0x5e, 0x15,
	0x0f, 0xb9, 0x12, 0xf4, 0x45, 0x98, 0x05, 0x8e, 0x48, 0x05, 0x7b, 0x05, 0xd3, 0x9d, 0xc6, 0x43,
	0x5e, 0x76, 0xaf, 0xc5, 0x17, 0x8b, 0xc1, 0x2a, 0xba, 0x63, 0x49, 0xdf, 0x43, 0xd2, 0xfb, 0xc8,
	0x26, 0xb6, 0xf0, 0xe4, 0xea, 0x16, 0x08, 0xe7, 0x5c, 0x08, 0x2d, 0x9b, 0x26, 0xf6, 0x49, 0x38,
	0xb2, 0xdc, 0xc6, 0x51, 0xec, 0x1e, 0x22, 0x83, 0x67, 0xe5, 0xe1, 0x3f, 0x95, 0xc1, 0xe0, 0x49,
	0xf7, 0x29, 0x80, 0xc1, 0x93, 0xea, 0x88, 0x54, 0x43, 0x83, 0xbd, 0xe6, 0x2d, 0x4c, 0xea, 0x07,
	0xac, 0x64, 0x5e, 0xb5, 0x87, 0xad, 0xd4, 0xf1, 0xd8, 0x3d, 0x4b, 0xdc, 0x7b, 0xa2, 0xd8, 0x6b,
	0x80, 0x02, 0x5b, 0xad, 0xa4, 0xb6, 0x0d, 0x07, 0xf4, 0xea, 0x93, 0xc4, 0x6d, 0x22, 0xe9, 0x37,
	0x91, 0x7c, 0x32, 0x5a, 0x55, 0xfb, 0x2f, 0xbc, 0x6c, 0x65, 0x16, 0x76, 0xf5, 0xa9, 0x60, 0x37,
	0x30, 0xa6, 0x9d, 0x28, 0x11, 0x87, 0x24, 0x3d, 0xb2, 0x30, 0xa5, 0x29, 0x36, 0x86, 0x9b, 0xb6,
	0xb1, 0x29, 0x70, 0x53, 0x74, 0x44, 0x2a, 0xac, 0xe9, 0x42, 0x4b, 0x6e, 0xa4, 0xc8, 0xb9, 0x89,
	0x23, 0x67, 0xba, 0x63, 0x36, 0x86, 0x2d, 0x20, 0x12, 0xb2, 0x29, 0xb4, 0xaa, 0x69, 0x10, 0x13,
	0xe7, 0xf9, 0x0f, 0x8a, 0xfd, 0x07, 0xc3, 0x02, 0xad, 0xf2, 0x94, 0x72, 0x7e, 0x81, 0xa9, 0x60,
	0xcf, 0xe1, 0x92, 0x34, 0x50, 0xe7, 0xe6, 0x58, 0x93, 0xa7, 0x19, 0xa5, 0xa7, 0x1d, 0xfd, 0xf9,
	0x58, 0x4b, 0xe7, 0xb9, 0x6d, 0x5c, 0xb7, 0x97, 0xce, 0xb3, 0x85, 0xa9, 0x60, 0x2f, 0x20, 0xa8,
	0x35, 0x8a, 0xb6, 0x30, 0x4d, 0x7c, 0xb5, 0xb8, 0x58, 0x45, 0x77, 0xd7, 0xe7, 0xe9, 0x7f, 0x74,
	0x99, 0xec, 0x54, 0xc2, 0x96, 0xe0, 0xdb, 0x66, 0xe3, 0x6b, 0x1a, 0xd9, 0xec, 0x5c, 0xfa, 0x96,
	0x6b, 0x99, 0x51, 0x6e, 0xf9, 0x73, 0x00, 0xe3, 0xee, 0xcb, 0xbf, 0x0e, 0x8d, 0x81, 0x5f, 0xf1,
	0x83, 0xec, 0x6e, 0x8c, 0x62, 0x36, 0x87, 0xe0, 0x5b, 0xcb, 0x2b, 0xa3, 0xcc, 0x91, 0x4e, 0xcb,
	0xcb, 0x4e, 0x98, 0xfd, 0x0f, 0xc3, 0x5a, 0xab, 0x42, 0xd2, 0xed, 0x78, 0x99, 0x03, 0x76, 0xc3,
	0x06, 0x0d, 0x2f, 0x73, 0x7e, 0xc0, 0xb6, 0x32, 0x74, 0x36, 0x5e, 0x16, 0x11, 0xb7, 0x21, 0xea,
	0x4d, 0xfc, 0xf5, 0x51, 0xef, 0x6d, 0x4d, 0x3f, 0x5e, 0xff, 0x0b, 0x6d, 0x47, 0x44, 0xde, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x12, 0xb9, 0x79, 0x86, 0x90, 0x03, 0x00, 0x00,
}
