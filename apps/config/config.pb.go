// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/config/pb/config.proto

package config

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

type TYPE int32

const (
	// KV格式, 环境变量格式
	TYPE_KV TYPE = 0
	// 文件
	TYPE_FILE TYPE = 1
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "KV",
		1: "FILE",
	}
	TYPE_value = map[string]int32{
		"KV":   0,
		"FILE": 1,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_config_pb_config_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_config_pb_config_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{0}
}

type FORMAT int32

const (
	// 纯文本
	FORMAT_TEXT FORMAT = 0
	// toml
	FORMAT_TOML FORMAT = 1
	// json
	FORMAT_JSON FORMAT = 2
	// yaml
	FORMAT_YAML FORMAT = 3
	// xml
	FORMAT_XML FORMAT = 4
)

// Enum value maps for FORMAT.
var (
	FORMAT_name = map[int32]string{
		0: "TEXT",
		1: "TOML",
		2: "JSON",
		3: "YAML",
		4: "XML",
	}
	FORMAT_value = map[string]int32{
		"TEXT": 0,
		"TOML": 1,
		"JSON": 2,
		"YAML": 3,
		"XML":  4,
	}
)

func (x FORMAT) Enum() *FORMAT {
	p := new(FORMAT)
	*p = x
	return p
}

func (x FORMAT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FORMAT) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_config_pb_config_proto_enumTypes[1].Descriptor()
}

func (FORMAT) Type() protoreflect.EnumType {
	return &file_apps_config_pb_config_proto_enumTypes[1]
}

func (x FORMAT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FORMAT.Descriptor instead.
func (FORMAT) EnumDescriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{1}
}

type STATUS int32

const (
	// 草稿
	STATUS_DRAFT STATUS = 0
	// 已发布
	STATUS_PUBLISHED STATUS = 1
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0: "DRAFT",
		1: "PUBLISHED",
	}
	STATUS_value = map[string]int32{
		"DRAFT":     0,
		"PUBLISHED": 1,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_config_pb_config_proto_enumTypes[2].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_apps_config_pb_config_proto_enumTypes[2]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{2}
}

type ConfigSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 列表
	// @gotags: json:"items"
	Items []*Config `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *ConfigSet) Reset() {
	*x = ConfigSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_config_pb_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSet) ProtoMessage() {}

func (x *ConfigSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_config_pb_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSet.ProtoReflect.Descriptor instead.
func (*ConfigSet) Descriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ConfigSet) GetItems() []*Config {
	if x != nil {
		return x.Items
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属域
	// @gotags: bson:"domain" json:"domain"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain" bson:"domain"`
	// 范围
	// @gotags: bson:"namespace" json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 创建的时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: bson:"update_by" json:"update_by"
	UpdateBy string `protobuf:"bytes,5,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 微服务名称
	// @gotags: bson:"srevice_name" json:"srevice_name"
	SreviceName string `protobuf:"bytes,6,opt,name=srevice_name,json=sreviceName,proto3" json:"srevice_name" bson:"srevice_name"`
	// 实例所属环境, 默认default
	// @gotags: bson:"environment" json:"environment" validate:"required"
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment" bson:"environment" validate:"required"`
	// 实例所属分组,默认default
	// @gotags: bson:"group" json:"group" validate:"required"
	Group string `protobuf:"bytes,8,opt,name=group,proto3" json:"group" bson:"group" validate:"required"`
	// 配置名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name" bson:"name"`
	// 配置描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,10,opt,name=description,proto3" json:"description" bson:"description"`
	// 配置类型
	// @gotags: bson:"type" json:"type"
	Type TYPE `protobuf:"varint,11,opt,name=type,proto3,enum=infraboard.mcenter.config.TYPE" json:"type" bson:"type"`
	// 标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,12,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
	// 状态
	// @gotags: bson:"status" json:"status"
	Status STATUS `protobuf:"varint,13,opt,name=status,proto3,enum=infraboard.mcenter.config.STATUS" json:"status" bson:"status"`
	// 发布时间
	// @gotags: bson:"publish_at" json:"publish_at"
	PublishAt int64 `protobuf:"varint,14,opt,name=publish_at,json=publishAt,proto3" json:"publish_at" bson:"publish_at"`
	// kv配置
	// @gotags: bson:"kv_config" json:"kv_config"
	KvConfig *KVConfig `protobuf:"bytes,15,opt,name=kv_config,json=kvConfig,proto3" json:"kv_config" bson:"kv_config"`
	// 文件配置
	// @gotags: bson:"file_config" json:"file_config"
	FileConfig *FileConfig `protobuf:"bytes,16,opt,name=file_config,json=fileConfig,proto3" json:"file_config" bson:"file_config"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_config_pb_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_apps_config_pb_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Config) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Config) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Config) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Config) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *Config) GetSreviceName() string {
	if x != nil {
		return x.SreviceName
	}
	return ""
}

func (x *Config) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *Config) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Config) GetType() TYPE {
	if x != nil {
		return x.Type
	}
	return TYPE_KV
}

func (x *Config) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Config) GetStatus() STATUS {
	if x != nil {
		return x.Status
	}
	return STATUS_DRAFT
}

func (x *Config) GetPublishAt() int64 {
	if x != nil {
		return x.PublishAt
	}
	return 0
}

func (x *Config) GetKvConfig() *KVConfig {
	if x != nil {
		return x.KvConfig
	}
	return nil
}

func (x *Config) GetFileConfig() *FileConfig {
	if x != nil {
		return x.FileConfig
	}
	return nil
}

// 文件格式配置
type FileConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件内容的格式
	// @gotags: bson:"format" json:"format"
	Format FORMAT `protobuf:"varint,1,opt,name=format,proto3,enum=infraboard.mcenter.config.FORMAT" json:"format" bson:"format"`
	// 文件内容
	// @gotags: bson:"content" json:"content"
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content" bson:"content"`
	// 文件挂载时路径名称, 可以下载到程序本地保留
	// @gotags: bson:"path" json:"path"
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path" bson:"path"`
	// 文件挂载时的权限, 比如066
	// @gotags: bson:"mode" json:"mode"
	Mode string `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode" bson:"mode"`
	// 配置文件是否加密
	// @gotags: bson:"is_encrypt" json:"is_encrypt"
	IsEncrypt bool `protobuf:"varint,5,opt,name=is_encrypt,json=isEncrypt,proto3" json:"is_encrypt" bson:"is_encrypt"`
	// 文件内容的hash值
	// @gotags: bson:"md5_hash" json:"md5_hash"
	Md5Hash string `protobuf:"bytes,6,opt,name=md5_hash,json=md5Hash,proto3" json:"md5_hash" bson:"md5_hash"`
}

func (x *FileConfig) Reset() {
	*x = FileConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_config_pb_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileConfig) ProtoMessage() {}

func (x *FileConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_config_pb_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileConfig.ProtoReflect.Descriptor instead.
func (*FileConfig) Descriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{2}
}

func (x *FileConfig) GetFormat() FORMAT {
	if x != nil {
		return x.Format
	}
	return FORMAT_TEXT
}

func (x *FileConfig) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FileConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileConfig) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *FileConfig) GetIsEncrypt() bool {
	if x != nil {
		return x.IsEncrypt
	}
	return false
}

func (x *FileConfig) GetMd5Hash() string {
	if x != nil {
		return x.Md5Hash
	}
	return ""
}

// kv格式配置
type KVConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 具体的配置项
	// @gotags: bson:"items" json:"items"
	Items []*KVItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *KVConfig) Reset() {
	*x = KVConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_config_pb_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVConfig) ProtoMessage() {}

func (x *KVConfig) ProtoReflect() protoreflect.Message {
	mi := &file_apps_config_pb_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVConfig.ProtoReflect.Descriptor instead.
func (*KVConfig) Descriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{3}
}

func (x *KVConfig) GetItems() []*KVItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 配置项
type KVItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建的时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,1,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,2,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 更新人
	// @gotags: bson:"update_by" json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by" bson:"update_by"`
	// 配置项名称
	// @gotags: bson:"name" json:"name"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name" bson:"name"`
	// 配置项值
	// @gotags: bson:"value" json:"value"
	Value string `protobuf:"bytes,6,opt,name=value,proto3" json:"value" bson:"value"`
	// 配置项值是否加密
	// @gotags: bson:"is_encrypt" json:"is_encrypt"
	IsEncrypt bool `protobuf:"varint,7,opt,name=is_encrypt,json=isEncrypt,proto3" json:"is_encrypt" bson:"is_encrypt"`
	// 配置项描述
	// @gotags: bson:"description" json:"description"
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description" bson:"description"`
}

func (x *KVItem) Reset() {
	*x = KVItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_config_pb_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVItem) ProtoMessage() {}

func (x *KVItem) ProtoReflect() protoreflect.Message {
	mi := &file_apps_config_pb_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVItem.ProtoReflect.Descriptor instead.
func (*KVItem) Descriptor() ([]byte, []int) {
	return file_apps_config_pb_config_proto_rawDescGZIP(), []int{4}
}

func (x *KVItem) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *KVItem) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *KVItem) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *KVItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KVItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *KVItem) GetIsEncrypt() bool {
	if x != nil {
		return x.IsEncrypt
	}
	return false
}

func (x *KVItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_apps_config_pb_config_proto protoreflect.FileDescriptor

var file_apps_config_pb_config_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5a, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xc1, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x72, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x72, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x6b, 0x76, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4b, 0x56, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x6b, 0x76, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc3, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x64, 0x35, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x64, 0x35, 0x48, 0x61, 0x73, 0x68, 0x22, 0x43,
	0x0a, 0x08, 0x4b, 0x56, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x06, 0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2a, 0x18, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x56, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x2a, 0x39, 0x0a, 0x06, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x4f, 0x4d, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x41, 0x4d, 0x4c, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x58, 0x4d, 0x4c, 0x10, 0x04, 0x2a, 0x22, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55,
	0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_config_pb_config_proto_rawDescOnce sync.Once
	file_apps_config_pb_config_proto_rawDescData = file_apps_config_pb_config_proto_rawDesc
)

func file_apps_config_pb_config_proto_rawDescGZIP() []byte {
	file_apps_config_pb_config_proto_rawDescOnce.Do(func() {
		file_apps_config_pb_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_config_pb_config_proto_rawDescData)
	})
	return file_apps_config_pb_config_proto_rawDescData
}

var file_apps_config_pb_config_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_apps_config_pb_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_config_pb_config_proto_goTypes = []interface{}{
	(TYPE)(0),          // 0: infraboard.mcenter.config.TYPE
	(FORMAT)(0),        // 1: infraboard.mcenter.config.FORMAT
	(STATUS)(0),        // 2: infraboard.mcenter.config.STATUS
	(*ConfigSet)(nil),  // 3: infraboard.mcenter.config.ConfigSet
	(*Config)(nil),     // 4: infraboard.mcenter.config.Config
	(*FileConfig)(nil), // 5: infraboard.mcenter.config.FileConfig
	(*KVConfig)(nil),   // 6: infraboard.mcenter.config.KVConfig
	(*KVItem)(nil),     // 7: infraboard.mcenter.config.KVItem
	nil,                // 8: infraboard.mcenter.config.Config.LabelsEntry
}
var file_apps_config_pb_config_proto_depIdxs = []int32{
	4, // 0: infraboard.mcenter.config.ConfigSet.items:type_name -> infraboard.mcenter.config.Config
	0, // 1: infraboard.mcenter.config.Config.type:type_name -> infraboard.mcenter.config.TYPE
	8, // 2: infraboard.mcenter.config.Config.labels:type_name -> infraboard.mcenter.config.Config.LabelsEntry
	2, // 3: infraboard.mcenter.config.Config.status:type_name -> infraboard.mcenter.config.STATUS
	6, // 4: infraboard.mcenter.config.Config.kv_config:type_name -> infraboard.mcenter.config.KVConfig
	5, // 5: infraboard.mcenter.config.Config.file_config:type_name -> infraboard.mcenter.config.FileConfig
	1, // 6: infraboard.mcenter.config.FileConfig.format:type_name -> infraboard.mcenter.config.FORMAT
	7, // 7: infraboard.mcenter.config.KVConfig.items:type_name -> infraboard.mcenter.config.KVItem
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_apps_config_pb_config_proto_init() }
func file_apps_config_pb_config_proto_init() {
	if File_apps_config_pb_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_config_pb_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSet); i {
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
		file_apps_config_pb_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_apps_config_pb_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileConfig); i {
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
		file_apps_config_pb_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVConfig); i {
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
		file_apps_config_pb_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVItem); i {
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
			RawDescriptor: file_apps_config_pb_config_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_config_pb_config_proto_goTypes,
		DependencyIndexes: file_apps_config_pb_config_proto_depIdxs,
		EnumInfos:         file_apps_config_pb_config_proto_enumTypes,
		MessageInfos:      file_apps_config_pb_config_proto_msgTypes,
	}.Build()
	File_apps_config_pb_config_proto = out.File
	file_apps_config_pb_config_proto_rawDesc = nil
	file_apps_config_pb_config_proto_goTypes = nil
	file_apps_config_pb_config_proto_depIdxs = nil
}
