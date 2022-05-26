// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: apps/instance/pb/instance.proto

package instance

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

// 服务实例信息, 比如 阿里云_杭州/生产环境/app01/default/instance01
type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 实例所属环境, 默认default
	// @gotags: bson:"region" json:"region"
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region" bson:"region"`
	// 实例所属环境, 默认default
	// @gotags: bson:"environment" json:"environment"
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment" bson:"environment"`
	// 实例所属分组,默认default
	// @gotags: bson:"group" json:"group"
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group" bson:"group"`
	// 实例所属应用Id
	// @gotags: bson:"application_id" json:"application_id"
	ApplicationId string `protobuf:"bytes,4,opt,name=application_id,json=applicationId,proto3" json:"application_id" bson:"application_id"`
	// 实例名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name" bson:"name"`
	// 实例地址
	// @gotags: bson:"address" json:"address"
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address" bson:"address"`
	// 实例标签, 可以根据标签快速过滤实例
	// @gotags: bson:"lables" json:"lables"
	Lables map[string]string `protobuf:"bytes,7,rep,name=lables,proto3" json:"lables" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"lables"`
	// 实例构建信息
	// @gotags: bson:"build" json:"build"
	Build *Build `protobuf:"bytes,8,opt,name=build,proto3" json:"build" bson:"build"`
	// 实例状态
	// @gotags: bson:"status" json:"status"
	Status *Status `protobuf:"bytes,9,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_instance_pb_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_apps_instance_pb_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_apps_instance_pb_instance_proto_rawDescGZIP(), []int{0}
}

func (x *Instance) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Instance) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *Instance) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Instance) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Instance) GetLables() map[string]string {
	if x != nil {
		return x.Lables
	}
	return nil
}

func (x *Instance) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *Instance) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 实例构建相关信息
type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 实例版本
	// @gotags: bson:"version" json:"version"
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version" bson:"version"`
	// 实例代码构建对应分支
	// @gotags: bson:"git_branch" json:"git_branch"
	GitBranch string `protobuf:"bytes,2,opt,name=git_branch,json=gitBranch,proto3" json:"git_branch" bson:"git_branch"`
	// 实例代码对应commit号
	// @gotags: bson:"git_commit" json:"git_commit"
	GitCommit string `protobuf:"bytes,3,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit" bson:"git_commit"`
	// 实例构建时间
	// @gotags: bson:"build_at" json:"build_at"
	BuildAt int64 `protobuf:"varint,4,opt,name=build_at,json=buildAt,proto3" json:"build_at" bson:"build_at"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_instance_pb_instance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_apps_instance_pb_instance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_apps_instance_pb_instance_proto_rawDescGZIP(), []int{1}
}

func (x *Build) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Build) GetGitBranch() string {
	if x != nil {
		return x.GitBranch
	}
	return ""
}

func (x *Build) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *Build) GetBuildAt() int64 {
	if x != nil {
		return x.BuildAt
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 上线时间
	// @gotags: bson:"online" json:"online"
	Online int64 `protobuf:"varint,5,opt,name=online,proto3" json:"online" bson:"online"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_instance_pb_instance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_instance_pb_instance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_instance_pb_instance_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetOnline() int64 {
	if x != nil {
		return x.Online
	}
	return 0
}

var File_apps_instance_pb_instance_proto protoreflect.FileDescriptor

var file_apps_instance_pb_instance_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xac,
	0x03, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7a, 0x0a,
	0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_instance_pb_instance_proto_rawDescOnce sync.Once
	file_apps_instance_pb_instance_proto_rawDescData = file_apps_instance_pb_instance_proto_rawDesc
)

func file_apps_instance_pb_instance_proto_rawDescGZIP() []byte {
	file_apps_instance_pb_instance_proto_rawDescOnce.Do(func() {
		file_apps_instance_pb_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_instance_pb_instance_proto_rawDescData)
	})
	return file_apps_instance_pb_instance_proto_rawDescData
}

var file_apps_instance_pb_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_instance_pb_instance_proto_goTypes = []interface{}{
	(*Instance)(nil), // 0: infraboard.mcenter.instance.Instance
	(*Build)(nil),    // 1: infraboard.mcenter.instance.Build
	(*Status)(nil),   // 2: infraboard.mcenter.instance.Status
	nil,              // 3: infraboard.mcenter.instance.Instance.LablesEntry
}
var file_apps_instance_pb_instance_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.instance.Instance.lables:type_name -> infraboard.mcenter.instance.Instance.LablesEntry
	1, // 1: infraboard.mcenter.instance.Instance.build:type_name -> infraboard.mcenter.instance.Build
	2, // 2: infraboard.mcenter.instance.Instance.status:type_name -> infraboard.mcenter.instance.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_instance_pb_instance_proto_init() }
func file_apps_instance_pb_instance_proto_init() {
	if File_apps_instance_pb_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_instance_pb_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
		file_apps_instance_pb_instance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
		file_apps_instance_pb_instance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_apps_instance_pb_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_instance_pb_instance_proto_goTypes,
		DependencyIndexes: file_apps_instance_pb_instance_proto_depIdxs,
		MessageInfos:      file_apps_instance_pb_instance_proto_msgTypes,
	}.Build()
	File_apps_instance_pb_instance_proto = out.File
	file_apps_instance_pb_instance_proto_rawDesc = nil
	file_apps_instance_pb_instance_proto_goTypes = nil
	file_apps_instance_pb_instance_proto_depIdxs = nil
}
