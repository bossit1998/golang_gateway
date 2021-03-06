// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: notification_service.proto

package notification_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	ReceiverId     string `protobuf:"bytes,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	StatusId       string `protobuf:"bytes,3,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_notification_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeStatusRequest) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *ChangeStatusRequest) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *ChangeStatusRequest) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

type GetAllUserNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverId string `protobuf:"bytes,1,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
}

func (x *GetAllUserNotificationsRequest) Reset() {
	*x = GetAllUserNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserNotificationsRequest) ProtoMessage() {}

func (x *GetAllUserNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notification_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_notification_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllUserNotificationsRequest) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

type GetAllUserNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *GetAllUserNotificationsResponse) Reset() {
	*x = GetAllUserNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserNotificationsResponse) ProtoMessage() {}

func (x *GetAllUserNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notification_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_notification_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUserNotificationsResponse) GetNotifications() []*Notification {
	if x != nil {
		return x.Notifications
	}
	return nil
}

var File_notification_service_proto protoreflect.FileDescriptor

var file_notification_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x8c, 0x02, 0x0a, 0x13, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_notification_service_proto_rawDescOnce sync.Once
	file_notification_service_proto_rawDescData = file_notification_service_proto_rawDesc
)

func file_notification_service_proto_rawDescGZIP() []byte {
	file_notification_service_proto_rawDescOnce.Do(func() {
		file_notification_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_service_proto_rawDescData)
	})
	return file_notification_service_proto_rawDescData
}

var file_notification_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notification_service_proto_goTypes = []interface{}{
	(*ChangeStatusRequest)(nil),             // 0: genproto.ChangeStatusRequest
	(*GetAllUserNotificationsRequest)(nil),  // 1: genproto.GetAllUserNotificationsRequest
	(*GetAllUserNotificationsResponse)(nil), // 2: genproto.GetAllUserNotificationsResponse
	(*Notification)(nil),                    // 3: genproto.Notification
	(*empty.Empty)(nil),                     // 4: google.protobuf.Empty
}
var file_notification_service_proto_depIdxs = []int32{
	3, // 0: genproto.GetAllUserNotificationsResponse.notifications:type_name -> genproto.Notification
	3, // 1: genproto.NotificationService.Create:input_type -> genproto.Notification
	0, // 2: genproto.NotificationService.ChangeStatus:input_type -> genproto.ChangeStatusRequest
	1, // 3: genproto.NotificationService.GetAllUserNotifications:input_type -> genproto.GetAllUserNotificationsRequest
	4, // 4: genproto.NotificationService.Create:output_type -> google.protobuf.Empty
	4, // 5: genproto.NotificationService.ChangeStatus:output_type -> google.protobuf.Empty
	2, // 6: genproto.NotificationService.GetAllUserNotifications:output_type -> genproto.GetAllUserNotificationsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_notification_service_proto_init() }
func file_notification_service_proto_init() {
	if File_notification_service_proto != nil {
		return
	}
	file_notification_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notification_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusRequest); i {
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
		file_notification_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserNotificationsRequest); i {
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
		file_notification_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserNotificationsResponse); i {
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
			RawDescriptor: file_notification_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_service_proto_goTypes,
		DependencyIndexes: file_notification_service_proto_depIdxs,
		MessageInfos:      file_notification_service_proto_msgTypes,
	}.Build()
	File_notification_service_proto = out.File
	file_notification_service_proto_rawDesc = nil
	file_notification_service_proto_goTypes = nil
	file_notification_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Create(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*empty.Empty, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAllUserNotifications(ctx context.Context, in *GetAllUserNotificationsRequest, opts ...grpc.CallOption) (*GetAllUserNotificationsResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Create(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAllUserNotifications(ctx context.Context, in *GetAllUserNotificationsRequest, opts ...grpc.CallOption) (*GetAllUserNotificationsResponse, error) {
	out := new(GetAllUserNotificationsResponse)
	err := c.cc.Invoke(ctx, "/genproto.NotificationService/GetAllUserNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	Create(context.Context, *Notification) (*empty.Empty, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*empty.Empty, error)
	GetAllUserNotifications(context.Context, *GetAllUserNotificationsRequest) (*GetAllUserNotificationsResponse, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) Create(context.Context, *Notification) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNotificationServiceServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (*UnimplementedNotificationServiceServer) GetAllUserNotifications(context.Context, *GetAllUserNotificationsRequest) (*GetAllUserNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserNotifications not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Create(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAllUserNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllUserNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.NotificationService/GetAllUserNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllUserNotifications(ctx, req.(*GetAllUserNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NotificationService_Create_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _NotificationService_ChangeStatus_Handler,
		},
		{
			MethodName: "GetAllUserNotifications",
			Handler:    _NotificationService_GetAllUserNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification_service.proto",
}
