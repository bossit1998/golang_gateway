// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fare.proto

package fare_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Fare struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DeliveryTime         int64    `protobuf:"varint,3,opt,name=delivery_time,json=deliveryTime,proto3" json:"delivery_time,omitempty"`
	PricePerKm           int64    `protobuf:"varint,4,opt,name=price_per_km,json=pricePerKm,proto3" json:"price_per_km,omitempty"`
	MinPrice             int64    `protobuf:"varint,5,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MinDistance          int64    `protobuf:"varint,6,opt,name=min_distance,json=minDistance,proto3" json:"min_distance,omitempty"`
	IsActive             bool     `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fare) Reset()         { *m = Fare{} }
func (m *Fare) String() string { return proto.CompactTextString(m) }
func (*Fare) ProtoMessage()    {}
func (*Fare) Descriptor() ([]byte, []int) {
	return fileDescriptor_af2909dcbc3131d2, []int{0}
}

func (m *Fare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fare.Unmarshal(m, b)
}
func (m *Fare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fare.Marshal(b, m, deterministic)
}
func (m *Fare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fare.Merge(m, src)
}
func (m *Fare) XXX_Size() int {
	return xxx_messageInfo_Fare.Size(m)
}
func (m *Fare) XXX_DiscardUnknown() {
	xxx_messageInfo_Fare.DiscardUnknown(m)
}

var xxx_messageInfo_Fare proto.InternalMessageInfo

func (m *Fare) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Fare) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fare) GetDeliveryTime() int64 {
	if m != nil {
		return m.DeliveryTime
	}
	return 0
}

func (m *Fare) GetPricePerKm() int64 {
	if m != nil {
		return m.PricePerKm
	}
	return 0
}

func (m *Fare) GetMinPrice() int64 {
	if m != nil {
		return m.MinPrice
	}
	return 0
}

func (m *Fare) GetMinDistance() int64 {
	if m != nil {
		return m.MinDistance
	}
	return 0
}

func (m *Fare) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Fare) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Fare) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Fare) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*Fare)(nil), "genproto.Fare")
}

func init() { proto.RegisterFile("fare.proto", fileDescriptor_af2909dcbc3131d2) }

var fileDescriptor_af2909dcbc3131d2 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x80, 0x61, 0xda, 0xad, 0x6b, 0x3b, 0x56, 0x0f, 0x01, 0x31, 0x20, 0x42, 0xd5, 0x4b, 0x4f,
	0x7a, 0xf0, 0x09, 0x2a, 0x45, 0x58, 0xbc, 0x2c, 0xc5, 0x93, 0x97, 0x10, 0xdb, 0x51, 0x06, 0x37,
	0xd9, 0x92, 0xc6, 0x82, 0xcf, 0xe4, 0x4b, 0x4a, 0xa6, 0xe9, 0x6d, 0xe6, 0xff, 0x5a, 0x08, 0x03,
	0xf0, 0xa9, 0x1d, 0x3e, 0x8c, 0xee, 0xe8, 0x8f, 0x22, 0xff, 0x42, 0xcb, 0xd3, 0xdd, 0x5f, 0x0a,
	0xd9, 0x8b, 0x76, 0x28, 0x2e, 0x20, 0xdd, 0xb5, 0x32, 0xa9, 0x92, 0xba, 0xe8, 0xd2, 0x5d, 0x2b,
	0x04, 0x64, 0x56, 0x1b, 0x94, 0x29, 0x17, 0x9e, 0xc5, 0x3d, 0x9c, 0x0f, 0x78, 0xa0, 0x19, 0xdd,
	0xaf, 0xf2, 0x64, 0x50, 0x6e, 0xaa, 0xa4, 0xde, 0x74, 0xe5, 0x1a, 0xdf, 0xc8, 0xa0, 0xa8, 0xa0,
	0x1c, 0x1d, 0xf5, 0xa8, 0x46, 0x74, 0xea, 0xdb, 0xc8, 0x8c, 0xbf, 0x01, 0x6e, 0x7b, 0x74, 0xaf,
	0x46, 0x5c, 0x43, 0x61, 0xc8, 0x2a, 0x2e, 0xf2, 0x84, 0x39, 0x37, 0x64, 0xf7, 0x61, 0x17, 0xb7,
	0x50, 0x06, 0x1c, 0x68, 0xf2, 0xda, 0xf6, 0x28, 0xb7, 0xec, 0x67, 0x86, 0x6c, 0x1b, 0x53, 0xf8,
	0x9f, 0x26, 0xa5, 0x7b, 0x4f, 0x33, 0xca, 0xd3, 0x2a, 0xa9, 0xf3, 0x2e, 0xa7, 0xa9, 0xe1, 0x5d,
	0xdc, 0x00, 0xf4, 0x0e, 0xb5, 0xc7, 0x41, 0x69, 0x2f, 0x73, 0x7e, 0x7d, 0x11, 0x4b, 0xe3, 0x03,
	0xff, 0x8c, 0xc3, 0xca, 0xc5, 0xc2, 0xb1, 0x2c, 0x3c, 0xe0, 0x01, 0x23, 0xc3, 0xc2, 0xb1, 0x34,
	0xfe, 0xf9, 0xea, 0xfd, 0x72, 0xbd, 0xdc, 0x63, 0x38, 0xa7, 0x9a, 0xd0, 0xcd, 0xd4, 0xe3, 0xc7,
	0x96, 0xdb, 0xd3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xba, 0xb7, 0xb3, 0x65, 0x01, 0x00,
	0x00,
}
