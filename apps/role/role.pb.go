// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/role/pb/role.proto

package role

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

// RoleType 角色
type RoleType int32

const (
	// 内建角色, 系统初始时创建
	RoleType_BUILDIN RoleType = 0
	// 管理员创建的一些角色, 全局可用
	RoleType_GLOBAL RoleType = 1
	// 用户自定义的角色, 仅域内可见
	RoleType_CUSTOM RoleType = 2
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "BUILDIN",
		1: "GLOBAL",
		2: "CUSTOM",
	}
	RoleType_value = map[string]int32{
		"BUILDIN": 0,
		"GLOBAL":  1,
		"CUSTOM":  2,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_role_pb_role_proto_enumTypes[0].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_apps_role_pb_role_proto_enumTypes[0]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{0}
}

// Role is rbac's role
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 范围, 角色范围限制, 由策略引擎动态补充
	// @gotags: bson:"-" json:"scope"
	Scope string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope" bson:"-"`
	// 角色描述信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateRoleRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Role) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Role) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Role) GetSpec() *CreateRoleRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type RoleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// @gotags: bson:"items" json:"items"
	Items []*Role `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RoleSet) Reset() {
	*x = RoleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSet) ProtoMessage() {}

func (x *RoleSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSet.ProtoReflect.Descriptor instead.
func (*RoleSet) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RoleSet) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

// CreateRoleRequest 创建角色请求
type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 创建者ID
	// @gotags: bson:"create_by" json:"create_by"
	CreateBy string `protobuf:"bytes,2,opt,name=create_by,json=createBy,proto3" json:"create_by" bson:"create_by"`
	// 角色类型
	// @gotags: bson:"type" json:"type"
	Type RoleType `protobuf:"varint,3,opt,name=type,proto3,enum=infraboard.mcenter.role.RoleType" json:"type" bson:"type"`
	// 角色名称
	// @gotags: bson:"name" json:"name" validate:"required,lte=30"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=30"`
	// 角色描述
	// @gotags: bson:"description" json:"description" validate:"lte=400"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description" bson:"description" validate:"lte=400"`
	// 角色关联的其他信息，比如展示的视图
	// @gotags: bson:"meta" json:"meta" validate:"lte=400"
	Meta map[string]string `protobuf:"bytes,6,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta" validate:"lte=400"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateRoleRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_BUILDIN
}

func (x *CreateRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRoleRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRoleRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

// QueryRoleRequest 列表查询
type QueryRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// @gotags: json:"type"
	Type *RoleType `protobuf:"varint,2,opt,name=type,proto3,enum=infraboard.mcenter.role.RoleType,oneof" json:"type"`
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRoleRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRoleRequest) GetType() RoleType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return RoleType_BUILDIN
}

func (x *QueryRoleRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// DescribeRoleRequest role详情
type DescribeRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// @gotags: json:"name,omitempty" validate:"required,lte=64"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required,lte=64"`
	// @gotags: bson:"type" json:"type"
	Type RoleType `protobuf:"varint,4,opt,name=type,proto3,enum=infraboard.mcenter.role.RoleType" json:"type" bson:"type"`
}

func (x *DescribeRoleRequest) Reset() {
	*x = DescribeRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeRoleRequest) ProtoMessage() {}

func (x *DescribeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeRoleRequest.ProtoReflect.Descriptor instead.
func (*DescribeRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DescribeRoleRequest) GetType() RoleType {
	if x != nil {
		return x.Type
	}
	return RoleType_BUILDIN
}

// DeleteRoleRequest role删除
type DeleteRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required,lte=64"`
	// @gotags: json:"delete_policy"
	DeletePolicy bool `protobuf:"varint,2,opt,name=delete_policy,json=deletePolicy,proto3" json:"delete_policy"`
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_role_pb_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_role_pb_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_apps_role_pb_role_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRoleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRoleRequest) GetDeletePolicy() bool {
	if x != nil {
		return x.DeletePolicy
	}
	return false
}

var File_apps_role_pb_role_proto protoreflect.FileDescriptor

var file_apps_role_pb_role_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x54, 0x0a, 0x07, 0x52,
	0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xb8, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x35, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa7, 0x01, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x70, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2a, 0x2f, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47,
	0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x10, 0x02, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_role_pb_role_proto_rawDescOnce sync.Once
	file_apps_role_pb_role_proto_rawDescData = file_apps_role_pb_role_proto_rawDesc
)

func file_apps_role_pb_role_proto_rawDescGZIP() []byte {
	file_apps_role_pb_role_proto_rawDescOnce.Do(func() {
		file_apps_role_pb_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_role_pb_role_proto_rawDescData)
	})
	return file_apps_role_pb_role_proto_rawDescData
}

var file_apps_role_pb_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_role_pb_role_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_apps_role_pb_role_proto_goTypes = []interface{}{
	(RoleType)(0),               // 0: infraboard.mcenter.role.RoleType
	(*Role)(nil),                // 1: infraboard.mcenter.role.Role
	(*RoleSet)(nil),             // 2: infraboard.mcenter.role.RoleSet
	(*CreateRoleRequest)(nil),   // 3: infraboard.mcenter.role.CreateRoleRequest
	(*QueryRoleRequest)(nil),    // 4: infraboard.mcenter.role.QueryRoleRequest
	(*DescribeRoleRequest)(nil), // 5: infraboard.mcenter.role.DescribeRoleRequest
	(*DeleteRoleRequest)(nil),   // 6: infraboard.mcenter.role.DeleteRoleRequest
	nil,                         // 7: infraboard.mcenter.role.CreateRoleRequest.MetaEntry
	(*request.PageRequest)(nil), // 8: infraboard.mcube.page.PageRequest
}
var file_apps_role_pb_role_proto_depIdxs = []int32{
	3, // 0: infraboard.mcenter.role.Role.spec:type_name -> infraboard.mcenter.role.CreateRoleRequest
	1, // 1: infraboard.mcenter.role.RoleSet.items:type_name -> infraboard.mcenter.role.Role
	0, // 2: infraboard.mcenter.role.CreateRoleRequest.type:type_name -> infraboard.mcenter.role.RoleType
	7, // 3: infraboard.mcenter.role.CreateRoleRequest.meta:type_name -> infraboard.mcenter.role.CreateRoleRequest.MetaEntry
	8, // 4: infraboard.mcenter.role.QueryRoleRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 5: infraboard.mcenter.role.QueryRoleRequest.type:type_name -> infraboard.mcenter.role.RoleType
	0, // 6: infraboard.mcenter.role.DescribeRoleRequest.type:type_name -> infraboard.mcenter.role.RoleType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_role_pb_role_proto_init() }
func file_apps_role_pb_role_proto_init() {
	if File_apps_role_pb_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_role_pb_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_apps_role_pb_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSet); i {
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
		file_apps_role_pb_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_apps_role_pb_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
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
		file_apps_role_pb_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeRoleRequest); i {
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
		file_apps_role_pb_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleRequest); i {
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
	file_apps_role_pb_role_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_role_pb_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_role_pb_role_proto_goTypes,
		DependencyIndexes: file_apps_role_pb_role_proto_depIdxs,
		EnumInfos:         file_apps_role_pb_role_proto_enumTypes,
		MessageInfos:      file_apps_role_pb_role_proto_msgTypes,
	}.Build()
	File_apps_role_pb_role_proto = out.File
	file_apps_role_pb_role_proto_rawDesc = nil
	file_apps_role_pb_role_proto_goTypes = nil
	file_apps_role_pb_role_proto_depIdxs = nil
}
