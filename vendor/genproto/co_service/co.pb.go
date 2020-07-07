// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: co.proto

package co_service

import (
	proto "github.com/golang/protobuf/proto"
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

type CargoOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Logo        string `protobuf:"bytes,3,opt,name=logo,proto3" json:"logo,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsActive    bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Token       string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	Login       string `protobuf:"bytes,8,opt,name=login,proto3" json:"login,omitempty"`
	Password    string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CargoOwner) Reset() {
	*x = CargoOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_co_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CargoOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoOwner) ProtoMessage() {}

func (x *CargoOwner) ProtoReflect() protoreflect.Message {
	mi := &file_co_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoOwner.ProtoReflect.Descriptor instead.
func (*CargoOwner) Descriptor() ([]byte, []int) {
	return file_co_proto_rawDescGZIP(), []int{0}
}

func (x *CargoOwner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CargoOwner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CargoOwner) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CargoOwner) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CargoOwner) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CargoOwner) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *CargoOwner) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CargoOwner) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CargoOwner) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

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
		mi := &file_co_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_co_proto_msgTypes[1]
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
	return file_co_proto_rawDescGZIP(), []int{1}
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

type CargoOwnerBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CoId               string    `protobuf:"bytes,3,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	Location           *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Address            string    `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	DestinationAddress string    `protobuf:"bytes,6,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Description        string    `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	IsActive           bool      `protobuf:"varint,8,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CargoOwnerBranch) Reset() {
	*x = CargoOwnerBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_co_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CargoOwnerBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CargoOwnerBranch) ProtoMessage() {}

func (x *CargoOwnerBranch) ProtoReflect() protoreflect.Message {
	mi := &file_co_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CargoOwnerBranch.ProtoReflect.Descriptor instead.
func (*CargoOwnerBranch) Descriptor() ([]byte, []int) {
	return file_co_proto_rawDescGZIP(), []int{2}
}

func (x *CargoOwnerBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CargoOwnerBranch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CargoOwnerBranch) GetCoId() string {
	if x != nil {
		return x.CoId
	}
	return ""
}

func (x *CargoOwnerBranch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *CargoOwnerBranch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CargoOwnerBranch) GetDestinationAddress() string {
	if x != nil {
		return x.DestinationAddress
	}
	return ""
}

func (x *CargoOwnerBranch) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CargoOwnerBranch) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

var File_co_proto protoreflect.FileDescriptor

var file_co_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x30, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22, 0x85, 0x02, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x67,
	0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x13, 0x0a, 0x05, 0x63, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_co_proto_rawDescOnce sync.Once
	file_co_proto_rawDescData = file_co_proto_rawDesc
)

func file_co_proto_rawDescGZIP() []byte {
	file_co_proto_rawDescOnce.Do(func() {
		file_co_proto_rawDescData = protoimpl.X.CompressGZIP(file_co_proto_rawDescData)
	})
	return file_co_proto_rawDescData
}

var file_co_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_co_proto_goTypes = []interface{}{
	(*CargoOwner)(nil),       // 0: genproto.CargoOwner
	(*Location)(nil),         // 1: genproto.Location
	(*CargoOwnerBranch)(nil), // 2: genproto.CargoOwnerBranch
}
var file_co_proto_depIdxs = []int32{
	1, // 0: genproto.CargoOwnerBranch.location:type_name -> genproto.Location
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_co_proto_init() }
func file_co_proto_init() {
	if File_co_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_co_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CargoOwner); i {
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
		file_co_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_co_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CargoOwnerBranch); i {
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
			RawDescriptor: file_co_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_co_proto_goTypes,
		DependencyIndexes: file_co_proto_depIdxs,
		MessageInfos:      file_co_proto_msgTypes,
	}.Build()
	File_co_proto = out.File
	file_co_proto_rawDesc = nil
	file_co_proto_goTypes = nil
	file_co_proto_depIdxs = nil
}
