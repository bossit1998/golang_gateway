// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: shipper.proto

package user_service

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

type Shipper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Username    string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Logo        string   `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Password    string   `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	Phone       []string `protobuf:"bytes,8,rep,name=phone,proto3" json:"phone,omitempty"`
	IsActive    bool     `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt   string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string   `protobuf:"bytes,12,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Shipper) Reset() {
	*x = Shipper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipper) ProtoMessage() {}

func (x *Shipper) ProtoReflect() protoreflect.Message {
	mi := &file_shipper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipper.ProtoReflect.Descriptor instead.
func (*Shipper) Descriptor() ([]byte, []int) {
	return file_shipper_proto_rawDescGZIP(), []int{0}
}

func (x *Shipper) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shipper) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Shipper) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shipper) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Shipper) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Shipper) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Shipper) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Shipper) GetPhone() []string {
	if x != nil {
		return x.Phone
	}
	return nil
}

func (x *Shipper) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Shipper) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Shipper) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Shipper) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

var File_shipper_proto protoreflect.FileDescriptor

var file_shipper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x07, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipper_proto_rawDescOnce sync.Once
	file_shipper_proto_rawDescData = file_shipper_proto_rawDesc
)

func file_shipper_proto_rawDescGZIP() []byte {
	file_shipper_proto_rawDescOnce.Do(func() {
		file_shipper_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipper_proto_rawDescData)
	})
	return file_shipper_proto_rawDescData
}

var file_shipper_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shipper_proto_goTypes = []interface{}{
	(*Shipper)(nil), // 0: genproto.Shipper
}
var file_shipper_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipper_proto_init() }
func file_shipper_proto_init() {
	if File_shipper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipper); i {
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
			RawDescriptor: file_shipper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shipper_proto_goTypes,
		DependencyIndexes: file_shipper_proto_depIdxs,
		MessageInfos:      file_shipper_proto_msgTypes,
	}.Build()
	File_shipper_proto = out.File
	file_shipper_proto_rawDesc = nil
	file_shipper_proto_goTypes = nil
	file_shipper_proto_depIdxs = nil
}
