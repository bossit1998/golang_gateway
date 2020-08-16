// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: auth_service.proto

package auth_service

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

type OTPLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone     string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	ClientId  string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ShipperId string `protobuf:"bytes,3,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *OTPLoginRequest) Reset() {
	*x = OTPLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPLoginRequest) ProtoMessage() {}

func (x *OTPLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPLoginRequest.ProtoReflect.Descriptor instead.
func (*OTPLoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *OTPLoginRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OTPLoginRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OTPLoginRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

type OTPLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *OTPLoginResponse) Reset() {
	*x = OTPLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPLoginResponse) ProtoMessage() {}

func (x *OTPLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPLoginResponse.ProtoReflect.Descriptor instead.
func (*OTPLoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *OTPLoginResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type OTPConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone     string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code      string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	ClientId  string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ShipperId string `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *OTPConfirmRequest) Reset() {
	*x = OTPConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OTPConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OTPConfirmRequest) ProtoMessage() {}

func (x *OTPConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OTPConfirmRequest.ProtoReflect.Descriptor instead.
func (*OTPConfirmRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *OTPConfirmRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OTPConfirmRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OTPConfirmRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OTPConfirmRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ShipperId string `protobuf:"bytes,3,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RefreshTokenRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RefreshTokenRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

type RefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *RefreshToken) Reset() {
	*x = RefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshToken) ProtoMessage() {}

func (x *RefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshToken.ProtoReflect.Descriptor instead.
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshToken) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Login     string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ShipperId string `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserTypeId string `protobuf:"bytes,2,opt,name=user_type_id,json=userTypeId,proto3" json:"user_type_id,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *LoginResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *LoginResponse) GetUserTypeId() string {
	if x != nil {
		return x.UserTypeId
	}
	return ""
}

type HasAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Method   string `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *HasAccessRequest) Reset() {
	*x = HasAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasAccessRequest) ProtoMessage() {}

func (x *HasAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasAccessRequest.ProtoReflect.Descriptor instead.
func (*HasAccessRequest) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *HasAccessRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *HasAccessRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HasAccessRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type HasAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasAccess bool `protobuf:"varint,1,opt,name=has_access,json=hasAccess,proto3" json:"has_access,omitempty"`
}

func (x *HasAccessResponse) Reset() {
	*x = HasAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasAccessResponse) ProtoMessage() {}

func (x *HasAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasAccessResponse.ProtoReflect.Descriptor instead.
func (*HasAccessResponse) Descriptor() ([]byte, []int) {
	return file_auth_service_proto_rawDescGZIP(), []int{8}
}

func (x *HasAccessResponse) GetHasAccess() bool {
	if x != nil {
		return x.HasAccess
	}
	return false
}

var File_auth_service_proto protoreflect.FileDescriptor

var file_auth_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0f, 0x4f, 0x54,
	0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x28, 0x0a, 0x10, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x79, 0x0a, 0x11, 0x4f, 0x54, 0x50,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x7c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x58, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x10, 0x48, 0x61,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61,
	0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x68, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe9, 0x02, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x12, 0x19, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x54, 0x50, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x54, 0x50,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4f, 0x54, 0x50, 0x12, 0x1b, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x54, 0x50, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_service_proto_rawDescOnce sync.Once
	file_auth_service_proto_rawDescData = file_auth_service_proto_rawDesc
)

func file_auth_service_proto_rawDescGZIP() []byte {
	file_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_service_proto_rawDescData)
	})
	return file_auth_service_proto_rawDescData
}

var file_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auth_service_proto_goTypes = []interface{}{
	(*OTPLoginRequest)(nil),     // 0: genproto.OTPLoginRequest
	(*OTPLoginResponse)(nil),    // 1: genproto.OTPLoginResponse
	(*OTPConfirmRequest)(nil),   // 2: genproto.OTPConfirmRequest
	(*RefreshTokenRequest)(nil), // 3: genproto.RefreshTokenRequest
	(*RefreshToken)(nil),        // 4: genproto.RefreshToken
	(*LoginRequest)(nil),        // 5: genproto.LoginRequest
	(*LoginResponse)(nil),       // 6: genproto.LoginResponse
	(*HasAccessRequest)(nil),    // 7: genproto.HasAccessRequest
	(*HasAccessResponse)(nil),   // 8: genproto.HasAccessResponse
	(*Token)(nil),               // 9: genproto.Token
}
var file_auth_service_proto_depIdxs = []int32{
	9, // 0: genproto.LoginResponse.token:type_name -> genproto.Token
	3, // 1: genproto.AuthService.RefreshToken:input_type -> genproto.RefreshTokenRequest
	5, // 2: genproto.AuthService.Login:input_type -> genproto.LoginRequest
	7, // 3: genproto.AuthService.HasAccess:input_type -> genproto.HasAccessRequest
	0, // 4: genproto.AuthService.GenerateOTP:input_type -> genproto.OTPLoginRequest
	2, // 5: genproto.AuthService.ConfirmOTP:input_type -> genproto.OTPConfirmRequest
	6, // 6: genproto.AuthService.RefreshToken:output_type -> genproto.LoginResponse
	6, // 7: genproto.AuthService.Login:output_type -> genproto.LoginResponse
	8, // 8: genproto.AuthService.HasAccess:output_type -> genproto.HasAccessResponse
	1, // 9: genproto.AuthService.GenerateOTP:output_type -> genproto.OTPLoginResponse
	6, // 10: genproto.AuthService.ConfirmOTP:output_type -> genproto.LoginResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_service_proto_init() }
func file_auth_service_proto_init() {
	if File_auth_service_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPLoginRequest); i {
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
		file_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPLoginResponse); i {
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
		file_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OTPConfirmRequest); i {
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
		file_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshToken); i {
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
		file_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasAccessRequest); i {
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
		file_auth_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasAccessResponse); i {
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
			RawDescriptor: file_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_service_proto_msgTypes,
	}.Build()
	File_auth_service_proto = out.File
	file_auth_service_proto_rawDesc = nil
	file_auth_service_proto_goTypes = nil
	file_auth_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error)
	GenerateOTP(ctx context.Context, in *OTPLoginRequest, opts ...grpc.CallOption) (*OTPLoginResponse, error)
	ConfirmOTP(ctx context.Context, in *OTPConfirmRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) HasAccess(ctx context.Context, in *HasAccessRequest, opts ...grpc.CallOption) (*HasAccessResponse, error) {
	out := new(HasAccessResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthService/HasAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateOTP(ctx context.Context, in *OTPLoginRequest, opts ...grpc.CallOption) (*OTPLoginResponse, error) {
	out := new(OTPLoginResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthService/GenerateOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ConfirmOTP(ctx context.Context, in *OTPConfirmRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/genproto.AuthService/ConfirmOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	RefreshToken(context.Context, *RefreshTokenRequest) (*LoginResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error)
	GenerateOTP(context.Context, *OTPLoginRequest) (*OTPLoginResponse, error)
	ConfirmOTP(context.Context, *OTPConfirmRequest) (*LoginResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (*UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) HasAccess(context.Context, *HasAccessRequest) (*HasAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasAccess not implemented")
}
func (*UnimplementedAuthServiceServer) GenerateOTP(context.Context, *OTPLoginRequest) (*OTPLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateOTP not implemented")
}
func (*UnimplementedAuthServiceServer) ConfirmOTP(context.Context, *OTPConfirmRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOTP not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_HasAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).HasAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthService/HasAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).HasAccess(ctx, req.(*HasAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthService/GenerateOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateOTP(ctx, req.(*OTPLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ConfirmOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ConfirmOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AuthService/ConfirmOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ConfirmOTP(ctx, req.(*OTPConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "HasAccess",
			Handler:    _AuthService_HasAccess_Handler,
		},
		{
			MethodName: "GenerateOTP",
			Handler:    _AuthService_GenerateOTP_Handler,
		},
		{
			MethodName: "ConfirmOTP",
			Handler:    _AuthService_ConfirmOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
