// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mcenter/apps/policy/pb/permission.proto

package policy

import (
	request "github.com/infraboard/mcube/http/request"
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

// CheckPermissionRequest todo
type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 用户名
	// @gotags: json:"username"
	Username string `protobuf:"bytes,7,opt,name=username,proto3" json:"username"`
	// 那个域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// 那个空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 服务Id
	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,5,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
	// 访问路径, 比如HTTP Path
	// @gotags: json:"path"
	Path string `protobuf:"bytes,6,opt,name=path,proto3" json:"path"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mcenter_apps_policy_pb_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcenter_apps_policy_pb_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_mcenter_apps_policy_pb_permission_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPermissionRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *CheckPermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CheckPermissionRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CheckPermissionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CheckPermissionRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CheckPermissionRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_mcenter_apps_policy_pb_permission_proto protoreflect.FileDescriptor

var file_mcenter_apps_policy_pb_permission_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x1a, 0x18, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcenter_apps_policy_pb_permission_proto_rawDescOnce sync.Once
	file_mcenter_apps_policy_pb_permission_proto_rawDescData = file_mcenter_apps_policy_pb_permission_proto_rawDesc
)

func file_mcenter_apps_policy_pb_permission_proto_rawDescGZIP() []byte {
	file_mcenter_apps_policy_pb_permission_proto_rawDescOnce.Do(func() {
		file_mcenter_apps_policy_pb_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcenter_apps_policy_pb_permission_proto_rawDescData)
	})
	return file_mcenter_apps_policy_pb_permission_proto_rawDescData
}

var file_mcenter_apps_policy_pb_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mcenter_apps_policy_pb_permission_proto_goTypes = []interface{}{
	(*CheckPermissionRequest)(nil), // 0: infraboard.mcenter.policy.CheckPermissionRequest
	(*request.PageRequest)(nil),    // 1: infraboard.mcube.page.PageRequest
}
var file_mcenter_apps_policy_pb_permission_proto_depIdxs = []int32{
	1, // 0: infraboard.mcenter.policy.CheckPermissionRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mcenter_apps_policy_pb_permission_proto_init() }
func file_mcenter_apps_policy_pb_permission_proto_init() {
	if File_mcenter_apps_policy_pb_permission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mcenter_apps_policy_pb_permission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
			RawDescriptor: file_mcenter_apps_policy_pb_permission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcenter_apps_policy_pb_permission_proto_goTypes,
		DependencyIndexes: file_mcenter_apps_policy_pb_permission_proto_depIdxs,
		MessageInfos:      file_mcenter_apps_policy_pb_permission_proto_msgTypes,
	}.Build()
	File_mcenter_apps_policy_pb_permission_proto = out.File
	file_mcenter_apps_policy_pb_permission_proto_rawDesc = nil
	file_mcenter_apps_policy_pb_permission_proto_goTypes = nil
	file_mcenter_apps_policy_pb_permission_proto_depIdxs = nil
}
