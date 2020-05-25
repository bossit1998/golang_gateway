// Code generated by protoc-gen-go. DO NOT EDIT.
// source: measure_service.proto

package catalog_service

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

type GetAllMeasureResponse struct {
	Measures             []*Measure `protobuf:"bytes,1,rep,name=measures,proto3" json:"measures,omitempty"`
	Count                int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllMeasureResponse) Reset()         { *m = GetAllMeasureResponse{} }
func (m *GetAllMeasureResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllMeasureResponse) ProtoMessage()    {}
func (*GetAllMeasureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e7a8d389942f64, []int{0}
}

func (m *GetAllMeasureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllMeasureResponse.Unmarshal(m, b)
}
func (m *GetAllMeasureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllMeasureResponse.Marshal(b, m, deterministic)
}
func (m *GetAllMeasureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllMeasureResponse.Merge(m, src)
}
func (m *GetAllMeasureResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllMeasureResponse.Size(m)
}
func (m *GetAllMeasureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllMeasureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllMeasureResponse proto.InternalMessageInfo

func (m *GetAllMeasureResponse) GetMeasures() []*Measure {
	if m != nil {
		return m.Measures
	}
	return nil
}

func (m *GetAllMeasureResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetAllMeasureResponse)(nil), "genproto.GetAllMeasureResponse")
}

func init() { proto.RegisterFile("measure_service.proto", fileDescriptor_43e7a8d389942f64) }

var fileDescriptor_43e7a8d389942f64 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x93, 0x37, 0xbc, 0xa1, 0x8c, 0x28, 0x38, 0xb4, 0x1a, 0xd6, 0x83, 0x65, 0x4f, 0xbd,
	0xb8, 0x91, 0x8a, 0x78, 0xae, 0x7f, 0xc8, 0xc9, 0x4b, 0xc4, 0x8b, 0x08, 0x92, 0xc6, 0x31, 0x08,
	0xdb, 0x6c, 0xcc, 0x4e, 0x04, 0xbf, 0xb9, 0x47, 0x31, 0x9b, 0xb5, 0xd8, 0xe2, 0x6d, 0x67, 0xe6,
	0x37, 0x33, 0xcf, 0x3e, 0x03, 0x93, 0x15, 0x15, 0xb6, 0x6b, 0xe9, 0xc9, 0x52, 0xfb, 0xfe, 0x5a,
	0x92, 0x6a, 0x5a, 0xc3, 0x06, 0x47, 0x15, 0xd5, 0xfd, 0x4b, 0xec, 0x96, 0x05, 0x17, 0xda, 0x54,
	0xae, 0x20, 0x8e, 0x2a, 0x63, 0x2a, 0x4d, 0x69, 0x1f, 0x2d, 0xbb, 0x97, 0x94, 0x56, 0x0d, 0x7f,
	0xb8, 0xa2, 0x7c, 0x84, 0x49, 0x46, 0xbc, 0xd0, 0xfa, 0xd6, 0x0d, 0xcd, 0xc9, 0x36, 0xa6, 0xb6,
	0x84, 0x27, 0x30, 0x1a, 0xf6, 0xd8, 0x24, 0x9c, 0x46, 0xb3, 0x9d, 0xf9, 0xbe, 0xf2, 0x1b, 0x94,
	0x87, 0x7f, 0x10, 0x1c, 0xc3, 0xff, 0xd2, 0x74, 0x35, 0x27, 0xff, 0xa6, 0xe1, 0x2c, 0xca, 0x5d,
	0x30, 0xff, 0x0c, 0x61, 0x6f, 0x60, 0xef, 0x9c, 0x58, 0xbc, 0x80, 0xf8, 0xaa, 0xa5, 0x82, 0x09,
	0xb7, 0xe7, 0x89, 0x64, 0x9d, 0x72, 0x90, 0x97, 0x23, 0x03, 0x3c, 0x85, 0x28, 0x23, 0xc6, 0xf1,
	0x1a, 0xc9, 0x88, 0x73, 0x7a, 0xeb, 0xc8, 0xb2, 0xd8, 0x9e, 0x25, 0x03, 0xbc, 0x86, 0xd8, 0xfd,
	0x0d, 0x0f, 0x7f, 0x35, 0x2d, 0xb4, 0xf6, 0x7d, 0xc7, 0x9b, 0x85, 0x0d, 0x1b, 0x64, 0x80, 0xe7,
	0x10, 0xdf, 0x37, 0xcf, 0x7f, 0x08, 0x3e, 0x50, 0xce, 0x5c, 0xe5, 0xcd, 0x55, 0x37, 0xdf, 0xe6,
	0xca, 0xe0, 0x52, 0x3c, 0x24, 0x9e, 0x4e, 0x87, 0x7b, 0xf8, 0x83, 0x2d, 0xe3, 0x3e, 0x7d, 0xf6,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x13, 0xb3, 0xe4, 0xfc, 0xca, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeasureServiceClient is the client API for MeasureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeasureServiceClient interface {
	Create(ctx context.Context, in *Measure, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Measure, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllMeasureResponse, error)
	Update(ctx context.Context, in *Measure, opts ...grpc.CallOption) (*empty.Empty, error)
}

type measureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeasureServiceClient(cc grpc.ClientConnInterface) MeasureServiceClient {
	return &measureServiceClient{cc}
}

func (c *measureServiceClient) Create(ctx context.Context, in *Measure, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/genproto.MeasureService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Measure, error) {
	out := new(Measure)
	err := c.cc.Invoke(ctx, "/genproto.MeasureService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllMeasureResponse, error) {
	out := new(GetAllMeasureResponse)
	err := c.cc.Invoke(ctx, "/genproto.MeasureService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *measureServiceClient) Update(ctx context.Context, in *Measure, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.MeasureService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeasureServiceServer is the server API for MeasureService service.
type MeasureServiceServer interface {
	Create(context.Context, *Measure) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*Measure, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllMeasureResponse, error)
	Update(context.Context, *Measure) (*empty.Empty, error)
}

// UnimplementedMeasureServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMeasureServiceServer struct {
}

func (*UnimplementedMeasureServiceServer) Create(ctx context.Context, req *Measure) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedMeasureServiceServer) Get(ctx context.Context, req *GetRequest) (*Measure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedMeasureServiceServer) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllMeasureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedMeasureServiceServer) Update(ctx context.Context, req *Measure) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterMeasureServiceServer(s *grpc.Server, srv MeasureServiceServer) {
	s.RegisterService(&_MeasureService_serviceDesc, srv)
}

func _MeasureService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Measure)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MeasureService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Create(ctx, req.(*Measure))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MeasureService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MeasureService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeasureService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Measure)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeasureServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.MeasureService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeasureServiceServer).Update(ctx, req.(*Measure))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeasureService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.MeasureService",
	HandlerType: (*MeasureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MeasureService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MeasureService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _MeasureService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MeasureService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "measure_service.proto",
}