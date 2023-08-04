// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/notify/pb/mail.proto

package notify

import (
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

type MailConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mail服务器地址
	// @gotags: bson:"host" json:"host" env:"MAIL_HOST"
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host" bson:"host" env:"MAIL_HOST"`
	// mail服务器端口
	// @gotags: bson:"port" json:"port" env:"MAIL_HORT"
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port" bson:"port" env:"MAIL_HORT"`
	// 用户名称
	// @gotags: bson:"username" json:"username" env:"MAIL_USERNAME"
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username" bson:"username" env:"MAIL_USERNAME"`
	// 用户密码
	// @gotags: bson:"password" json:"password" env:"MAIL_PASSWORD"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password" bson:"password" env:"MAIL_PASSWORD"`
}

func (x *MailConfig) Reset() {
	*x = MailConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_notify_pb_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailConfig) ProtoMessage() {}

func (x *MailConfig) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_notify_pb_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailConfig.ProtoReflect.Descriptor instead.
func (*MailConfig) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_notify_pb_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *MailConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *MailConfig) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MailConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_mcenter_apps_notify_pb_mail_proto protoreflect.FileDescriptor

var file_mcenter_apps_notify_pb_mail_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x6c,
	0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mcenter_apps_notify_pb_mail_proto_rawDescOnce sync.Once
	file_mcenter_apps_notify_pb_mail_proto_rawDescData = file_mcenter_apps_notify_pb_mail_proto_rawDesc
)

func file_mcenter_apps_notify_pb_mail_proto_rawDescGZIP() []byte {
	file_mcenter_apps_notify_pb_mail_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_notify_pb_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_notify_pb_mail_proto_rawDescData)
	})
	return file_mcenter_apps_notify_pb_mail_proto_rawDescData
}

var file_mcenter_apps_notify_pb_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcenter_apps_notify_pb_mail_proto_goTypes = []interface{}{
	(*MailConfig)(nil), // 0: infraboard.mcenter.notify.MailConfig
}
var file_mcenter_apps_notify_pb_mail_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcenter_apps_notify_pb_mail_proto_init() }
func file_mcenter_apps_notify_pb_mail_proto_init() {
	if File_mcenter_apps_notify_pb_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_notify_pb_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailConfig); i {
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
			RawDescriptor: file_mcenter_apps_notify_pb_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_notify_pb_mail_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_notify_pb_mail_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_notify_pb_mail_proto_msgTypes,
	}.Build()
	File_mcenter_apps_notify_pb_mail_proto = out.File
	file_mcenter_apps_notify_pb_mail_proto_rawDesc = nil
	file_mcenter_apps_notify_pb_mail_proto_goTypes = nil
	file_mcenter_apps_notify_pb_mail_proto_depIdxs = nil
}