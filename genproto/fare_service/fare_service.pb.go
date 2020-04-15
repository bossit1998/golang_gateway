// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fare_service.proto

package fare_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateFareResponse struct {
	Fare                 *Fare    `protobuf:"bytes,1,opt,name=fare,proto3" json:"fare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFareResponse) Reset()         { *m = CreateFareResponse{} }
func (m *CreateFareResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFareResponse) ProtoMessage()    {}
func (*CreateFareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{0}
}

func (m *CreateFareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFareResponse.Unmarshal(m, b)
}
func (m *CreateFareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFareResponse.Marshal(b, m, deterministic)
}
func (m *CreateFareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFareResponse.Merge(m, src)
}
func (m *CreateFareResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFareResponse.Size(m)
}
func (m *CreateFareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFareResponse proto.InternalMessageInfo

func (m *CreateFareResponse) GetFare() *Fare {
	if m != nil {
		return m.Fare
	}
	return nil
}

type UpdateFareResponse struct {
	Fare                 *Fare    `protobuf:"bytes,1,opt,name=fare,proto3" json:"fare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFareResponse) Reset()         { *m = UpdateFareResponse{} }
func (m *UpdateFareResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateFareResponse) ProtoMessage()    {}
func (*UpdateFareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{1}
}

func (m *UpdateFareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFareResponse.Unmarshal(m, b)
}
func (m *UpdateFareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFareResponse.Marshal(b, m, deterministic)
}
func (m *UpdateFareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFareResponse.Merge(m, src)
}
func (m *UpdateFareResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateFareResponse.Size(m)
}
func (m *UpdateFareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFareResponse proto.InternalMessageInfo

func (m *UpdateFareResponse) GetFare() *Fare {
	if m != nil {
		return m.Fare
	}
	return nil
}

type GetFareRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFareRequest) Reset()         { *m = GetFareRequest{} }
func (m *GetFareRequest) String() string { return proto.CompactTextString(m) }
func (*GetFareRequest) ProtoMessage()    {}
func (*GetFareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{2}
}

func (m *GetFareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFareRequest.Unmarshal(m, b)
}
func (m *GetFareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFareRequest.Marshal(b, m, deterministic)
}
func (m *GetFareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFareRequest.Merge(m, src)
}
func (m *GetFareRequest) XXX_Size() int {
	return xxx_messageInfo_GetFareRequest.Size(m)
}
func (m *GetFareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFareRequest proto.InternalMessageInfo

func (m *GetFareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetFareResponse struct {
	Fare                 *Fare    `protobuf:"bytes,1,opt,name=fare,proto3" json:"fare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFareResponse) Reset()         { *m = GetFareResponse{} }
func (m *GetFareResponse) String() string { return proto.CompactTextString(m) }
func (*GetFareResponse) ProtoMessage()    {}
func (*GetFareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{3}
}

func (m *GetFareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFareResponse.Unmarshal(m, b)
}
func (m *GetFareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFareResponse.Marshal(b, m, deterministic)
}
func (m *GetFareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFareResponse.Merge(m, src)
}
func (m *GetFareResponse) XXX_Size() int {
	return xxx_messageInfo_GetFareResponse.Size(m)
}
func (m *GetFareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFareResponse proto.InternalMessageInfo

func (m *GetFareResponse) GetFare() *Fare {
	if m != nil {
		return m.Fare
	}
	return nil
}

type GetAllFaresRequest struct {
	Limit                uint64   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 uint64   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllFaresRequest) Reset()         { *m = GetAllFaresRequest{} }
func (m *GetAllFaresRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllFaresRequest) ProtoMessage()    {}
func (*GetAllFaresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{4}
}

func (m *GetAllFaresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFaresRequest.Unmarshal(m, b)
}
func (m *GetAllFaresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFaresRequest.Marshal(b, m, deterministic)
}
func (m *GetAllFaresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFaresRequest.Merge(m, src)
}
func (m *GetAllFaresRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllFaresRequest.Size(m)
}
func (m *GetAllFaresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFaresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFaresRequest proto.InternalMessageInfo

func (m *GetAllFaresRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetAllFaresRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

type GetAllFaresResponse struct {
	Fares                []*Fare  `protobuf:"bytes,1,rep,name=fares,proto3" json:"fares,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllFaresResponse) Reset()         { *m = GetAllFaresResponse{} }
func (m *GetAllFaresResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllFaresResponse) ProtoMessage()    {}
func (*GetAllFaresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{5}
}

func (m *GetAllFaresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllFaresResponse.Unmarshal(m, b)
}
func (m *GetAllFaresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllFaresResponse.Marshal(b, m, deterministic)
}
func (m *GetAllFaresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllFaresResponse.Merge(m, src)
}
func (m *GetAllFaresResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllFaresResponse.Size(m)
}
func (m *GetAllFaresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllFaresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllFaresResponse proto.InternalMessageInfo

func (m *GetAllFaresResponse) GetFares() []*Fare {
	if m != nil {
		return m.Fares
	}
	return nil
}

func (m *GetAllFaresResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DeleteFareRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFareRequest) Reset()         { *m = DeleteFareRequest{} }
func (m *DeleteFareRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteFareRequest) ProtoMessage()    {}
func (*DeleteFareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_581d3f182827d35a, []int{6}
}

func (m *DeleteFareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFareRequest.Unmarshal(m, b)
}
func (m *DeleteFareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFareRequest.Marshal(b, m, deterministic)
}
func (m *DeleteFareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFareRequest.Merge(m, src)
}
func (m *DeleteFareRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteFareRequest.Size(m)
}
func (m *DeleteFareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFareRequest proto.InternalMessageInfo

func (m *DeleteFareRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateFareResponse)(nil), "genproto.CreateFareResponse")
	proto.RegisterType((*UpdateFareResponse)(nil), "genproto.UpdateFareResponse")
	proto.RegisterType((*GetFareRequest)(nil), "genproto.GetFareRequest")
	proto.RegisterType((*GetFareResponse)(nil), "genproto.GetFareResponse")
	proto.RegisterType((*GetAllFaresRequest)(nil), "genproto.GetAllFaresRequest")
	proto.RegisterType((*GetAllFaresResponse)(nil), "genproto.GetAllFaresResponse")
	proto.RegisterType((*DeleteFareRequest)(nil), "genproto.DeleteFareRequest")
}

func init() {
	proto.RegisterFile("fare_service.proto", fileDescriptor_581d3f182827d35a)
}

var fileDescriptor_581d3f182827d35a = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xa5, 0xa5, 0xa0, 0x0e, 0x09, 0xc6, 0xf1, 0xab, 0x16, 0x4c, 0x9a, 0xd5, 0x03, 0xa7, 0x92,
	0x60, 0x4c, 0x38, 0xf9, 0xad, 0x5c, 0xbc, 0x88, 0xf1, 0xe2, 0xc5, 0x14, 0x18, 0x9a, 0x26, 0x85,
	0xd6, 0xee, 0x62, 0xe2, 0xdf, 0xf4, 0x17, 0x99, 0xdd, 0xa5, 0x7c, 0x56, 0x13, 0xbd, 0xed, 0xcc,
	0x9b, 0xf7, 0x26, 0x7d, 0x6f, 0x0a, 0x38, 0xf4, 0x53, 0x7a, 0xe3, 0x94, 0x7e, 0x84, 0x7d, 0xf2,
	0x92, 0x34, 0x16, 0x31, 0x6e, 0x06, 0x34, 0x56, 0x2f, 0x07, 0x24, 0xaa, 0xbb, 0x4e, 0x2d, 0x88,
	0xe3, 0x20, 0xa2, 0xa6, 0xaa, 0x7a, 0x93, 0x61, 0x93, 0x46, 0x89, 0xf8, 0xd4, 0x20, 0x6b, 0x03,
	0xde, 0xa6, 0xe4, 0x0b, 0x7a, 0xf0, 0x53, 0xea, 0x12, 0x4f, 0xe2, 0x31, 0x27, 0x64, 0x60, 0x49,
	0x01, 0xdb, 0x70, 0x8d, 0x46, 0xa5, 0x55, 0xf5, 0x32, 0x5d, 0x4f, 0x4d, 0x29, 0x4c, 0x32, 0x5f,
	0x92, 0xc1, 0x7f, 0x98, 0x2e, 0x54, 0x3b, 0x24, 0x34, 0xed, 0x7d, 0x42, 0x5c, 0x60, 0x15, 0xcc,
	0x70, 0xa0, 0x38, 0x5b, 0x5d, 0x33, 0x1c, 0xb0, 0x73, 0xd8, 0x9e, 0x4d, 0xfc, 0x41, 0xf8, 0x02,
	0xb0, 0x43, 0xe2, 0x3a, 0x8a, 0x64, 0x8f, 0x67, 0xe2, 0x7b, 0x50, 0x8a, 0xc2, 0x51, 0x28, 0xec,
	0xa2, 0x6b, 0x34, 0xac, 0xae, 0x2e, 0x10, 0xc1, 0x4a, 0xfc, 0x80, 0x6c, 0x4b, 0x35, 0xd5, 0x9b,
	0x3d, 0xc1, 0xee, 0x12, 0x7f, 0xba, 0xfa, 0x14, 0x4a, 0x52, 0x9e, 0xdb, 0x86, 0x5b, 0xcc, 0xd9,
	0xad, 0x41, 0xb9, 0xa6, 0x1f, 0x4f, 0xc6, 0xc2, 0x36, 0xf5, 0x1a, 0x55, 0xb0, 0x13, 0xd8, 0xb9,
	0xa3, 0x88, 0x32, 0x97, 0x72, 0x3f, 0xb7, 0xf5, 0x65, 0x42, 0x45, 0xe2, 0xcf, 0x3a, 0x4d, 0x6c,
	0x43, 0x59, 0x87, 0x82, 0x2b, 0xbb, 0x9c, 0xfa, 0xbc, 0x5e, 0x8f, 0x8d, 0x15, 0x24, 0x53, 0x87,
	0xf2, 0x1b, 0x73, 0x3d, 0x36, 0x56, 0xc0, 0x2b, 0xd8, 0x98, 0x5a, 0x8e, 0xf6, 0x7c, 0x74, 0x39,
	0x27, 0xe7, 0x28, 0x07, 0x99, 0x29, 0x3c, 0x42, 0x65, 0xc1, 0x3d, 0xac, 0x2f, 0xcd, 0xae, 0x84,
	0xe2, 0x1c, 0xff, 0x80, 0xce, 0xd4, 0x2e, 0xa1, 0xac, 0x8d, 0xc3, 0xda, 0x7c, 0x74, 0xcd, 0x4a,
	0xe7, 0xc0, 0xd3, 0xd7, 0xed, 0x65, 0xd7, 0xed, 0xdd, 0xcb, 0xeb, 0x66, 0x85, 0x9b, 0xc3, 0xd7,
	0xfd, 0x8c, 0xd7, 0x5c, 0xfc, 0x57, 0x7a, 0x65, 0xd5, 0x3b, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0xa2, 0x18, 0x89, 0x42, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FareServiceClient is the client API for FareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FareServiceClient interface {
	Create(ctx context.Context, in *Fare, opts ...grpc.CallOption) (*CreateFareResponse, error)
	Update(ctx context.Context, in *Fare, opts ...grpc.CallOption) (*UpdateFareResponse, error)
	GetFare(ctx context.Context, in *GetFareRequest, opts ...grpc.CallOption) (*GetFareResponse, error)
	GetAllFares(ctx context.Context, in *GetAllFaresRequest, opts ...grpc.CallOption) (*GetAllFaresResponse, error)
	Delete(ctx context.Context, in *DeleteFareRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type fareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFareServiceClient(cc grpc.ClientConnInterface) FareServiceClient {
	return &fareServiceClient{cc}
}

func (c *fareServiceClient) Create(ctx context.Context, in *Fare, opts ...grpc.CallOption) (*CreateFareResponse, error) {
	out := new(CreateFareResponse)
	err := c.cc.Invoke(ctx, "/genproto.FareService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) Update(ctx context.Context, in *Fare, opts ...grpc.CallOption) (*UpdateFareResponse, error) {
	out := new(UpdateFareResponse)
	err := c.cc.Invoke(ctx, "/genproto.FareService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) GetFare(ctx context.Context, in *GetFareRequest, opts ...grpc.CallOption) (*GetFareResponse, error) {
	out := new(GetFareResponse)
	err := c.cc.Invoke(ctx, "/genproto.FareService/GetFare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) GetAllFares(ctx context.Context, in *GetAllFaresRequest, opts ...grpc.CallOption) (*GetAllFaresResponse, error) {
	out := new(GetAllFaresResponse)
	err := c.cc.Invoke(ctx, "/genproto.FareService/GetAllFares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) Delete(ctx context.Context, in *DeleteFareRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.FareService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FareServiceServer is the server API for FareService service.
type FareServiceServer interface {
	Create(context.Context, *Fare) (*CreateFareResponse, error)
	Update(context.Context, *Fare) (*UpdateFareResponse, error)
	GetFare(context.Context, *GetFareRequest) (*GetFareResponse, error)
	GetAllFares(context.Context, *GetAllFaresRequest) (*GetAllFaresResponse, error)
	Delete(context.Context, *DeleteFareRequest) (*empty.Empty, error)
}

// UnimplementedFareServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFareServiceServer struct {
}

func (*UnimplementedFareServiceServer) Create(ctx context.Context, req *Fare) (*CreateFareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedFareServiceServer) Update(ctx context.Context, req *Fare) (*UpdateFareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedFareServiceServer) GetFare(ctx context.Context, req *GetFareRequest) (*GetFareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFare not implemented")
}
func (*UnimplementedFareServiceServer) GetAllFares(ctx context.Context, req *GetAllFaresRequest) (*GetAllFaresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFares not implemented")
}
func (*UnimplementedFareServiceServer) Delete(ctx context.Context, req *DeleteFareRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterFareServiceServer(s *grpc.Server, srv FareServiceServer) {
	s.RegisterService(&_FareService_serviceDesc, srv)
}

func _FareService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.FareService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).Create(ctx, req.(*Fare))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Fare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.FareService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).Update(ctx, req.(*Fare))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_GetFare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).GetFare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.FareService/GetFare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).GetFare(ctx, req.(*GetFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_GetAllFares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFaresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).GetAllFares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.FareService/GetAllFares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).GetAllFares(ctx, req.(*GetAllFaresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.FareService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).Delete(ctx, req.(*DeleteFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.FareService",
	HandlerType: (*FareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FareService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FareService_Update_Handler,
		},
		{
			MethodName: "GetFare",
			Handler:    _FareService_GetFare_Handler,
		},
		{
			MethodName: "GetAllFares",
			Handler:    _FareService_GetAllFares_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FareService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fare_service.proto",
}
