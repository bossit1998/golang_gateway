// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: report_service.proto

package report_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetBranchesReportExcelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Date      string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetBranchesReportExcelRequest) Reset() {
	*x = GetBranchesReportExcelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchesReportExcelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesReportExcelRequest) ProtoMessage() {}

func (x *GetBranchesReportExcelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchesReportExcelRequest.ProtoReflect.Descriptor instead.
func (*GetBranchesReportExcelRequest) Descriptor() ([]byte, []int) {
	return file_report_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBranchesReportExcelRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetBranchesReportExcelRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetBranchesReportExcelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetBranchesReportExcelResponse) Reset() {
	*x = GetBranchesReportExcelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBranchesReportExcelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBranchesReportExcelResponse) ProtoMessage() {}

func (x *GetBranchesReportExcelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBranchesReportExcelResponse.ProtoReflect.Descriptor instead.
func (*GetBranchesReportExcelResponse) Descriptor() ([]byte, []int) {
	return file_report_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetBranchesReportExcelResponse) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type GetCouriersReportExcelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Date      string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetCouriersReportExcelRequest) Reset() {
	*x = GetCouriersReportExcelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouriersReportExcelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouriersReportExcelRequest) ProtoMessage() {}

func (x *GetCouriersReportExcelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouriersReportExcelRequest.ProtoReflect.Descriptor instead.
func (*GetCouriersReportExcelRequest) Descriptor() ([]byte, []int) {
	return file_report_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCouriersReportExcelRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetCouriersReportExcelRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type GetCouriersReportExcelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetCouriersReportExcelResponse) Reset() {
	*x = GetCouriersReportExcelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouriersReportExcelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouriersReportExcelResponse) ProtoMessage() {}

func (x *GetCouriersReportExcelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_report_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouriersReportExcelResponse.ProtoReflect.Descriptor instead.
func (*GetCouriersReportExcelResponse) Descriptor() ([]byte, []int) {
	return file_report_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCouriersReportExcelResponse) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

var File_report_service_proto protoreflect.FileDescriptor

var file_report_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x34, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45,
	0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x34,
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x32, 0xed, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c,
	0x12, 0x27, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x12,
	0x27, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_service_proto_rawDescOnce sync.Once
	file_report_service_proto_rawDescData = file_report_service_proto_rawDesc
)

func file_report_service_proto_rawDescGZIP() []byte {
	file_report_service_proto_rawDescOnce.Do(func() {
		file_report_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_service_proto_rawDescData)
	})
	return file_report_service_proto_rawDescData
}

var file_report_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_report_service_proto_goTypes = []interface{}{
	(*GetBranchesReportExcelRequest)(nil),  // 0: genproto.GetBranchesReportExcelRequest
	(*GetBranchesReportExcelResponse)(nil), // 1: genproto.GetBranchesReportExcelResponse
	(*GetCouriersReportExcelRequest)(nil),  // 2: genproto.GetCouriersReportExcelRequest
	(*GetCouriersReportExcelResponse)(nil), // 3: genproto.GetCouriersReportExcelResponse
}
var file_report_service_proto_depIdxs = []int32{
	0, // 0: genproto.ReportService.GetBranchesReportExcel:input_type -> genproto.GetBranchesReportExcelRequest
	2, // 1: genproto.ReportService.GetCouriersReportExcel:input_type -> genproto.GetCouriersReportExcelRequest
	1, // 2: genproto.ReportService.GetBranchesReportExcel:output_type -> genproto.GetBranchesReportExcelResponse
	3, // 3: genproto.ReportService.GetCouriersReportExcel:output_type -> genproto.GetCouriersReportExcelResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_report_service_proto_init() }
func file_report_service_proto_init() {
	if File_report_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBranchesReportExcelRequest); i {
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
		file_report_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBranchesReportExcelResponse); i {
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
		file_report_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouriersReportExcelRequest); i {
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
		file_report_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouriersReportExcelResponse); i {
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
			RawDescriptor: file_report_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_service_proto_goTypes,
		DependencyIndexes: file_report_service_proto_depIdxs,
		MessageInfos:      file_report_service_proto_msgTypes,
	}.Build()
	File_report_service_proto = out.File
	file_report_service_proto_rawDesc = nil
	file_report_service_proto_goTypes = nil
	file_report_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportServiceClient interface {
	GetBranchesReportExcel(ctx context.Context, in *GetBranchesReportExcelRequest, opts ...grpc.CallOption) (*GetBranchesReportExcelResponse, error)
	GetCouriersReportExcel(ctx context.Context, in *GetCouriersReportExcelRequest, opts ...grpc.CallOption) (*GetCouriersReportExcelResponse, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) GetBranchesReportExcel(ctx context.Context, in *GetBranchesReportExcelRequest, opts ...grpc.CallOption) (*GetBranchesReportExcelResponse, error) {
	out := new(GetBranchesReportExcelResponse)
	err := c.cc.Invoke(ctx, "/genproto.ReportService/GetBranchesReportExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportServiceClient) GetCouriersReportExcel(ctx context.Context, in *GetCouriersReportExcelRequest, opts ...grpc.CallOption) (*GetCouriersReportExcelResponse, error) {
	out := new(GetCouriersReportExcelResponse)
	err := c.cc.Invoke(ctx, "/genproto.ReportService/GetCouriersReportExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceServer is the server API for ReportService service.
type ReportServiceServer interface {
	GetBranchesReportExcel(context.Context, *GetBranchesReportExcelRequest) (*GetBranchesReportExcelResponse, error)
	GetCouriersReportExcel(context.Context, *GetCouriersReportExcelRequest) (*GetCouriersReportExcelResponse, error)
}

// UnimplementedReportServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (*UnimplementedReportServiceServer) GetBranchesReportExcel(context.Context, *GetBranchesReportExcelRequest) (*GetBranchesReportExcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchesReportExcel not implemented")
}
func (*UnimplementedReportServiceServer) GetCouriersReportExcel(context.Context, *GetCouriersReportExcelRequest) (*GetCouriersReportExcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCouriersReportExcel not implemented")
}

func RegisterReportServiceServer(s *grpc.Server, srv ReportServiceServer) {
	s.RegisterService(&_ReportService_serviceDesc, srv)
}

func _ReportService_GetBranchesReportExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchesReportExcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GetBranchesReportExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReportService/GetBranchesReportExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GetBranchesReportExcel(ctx, req.(*GetBranchesReportExcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportService_GetCouriersReportExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouriersReportExcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceServer).GetCouriersReportExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReportService/GetCouriersReportExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceServer).GetCouriersReportExcel(ctx, req.(*GetCouriersReportExcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBranchesReportExcel",
			Handler:    _ReportService_GetBranchesReportExcel_Handler,
		},
		{
			MethodName: "GetCouriersReportExcel",
			Handler:    _ReportService_GetCouriersReportExcel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "report_service.proto",
}
