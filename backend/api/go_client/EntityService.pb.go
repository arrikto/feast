// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: EntityService.proto

package go_client

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *CreateEntityRequest) Reset() {
	*x = CreateEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEntityRequest) ProtoMessage() {}

func (x *CreateEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEntityRequest.ProtoReflect.Descriptor instead.
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEntityRequest) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type GetEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetEntityRequest) Reset() {
	*x = GetEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityRequest) ProtoMessage() {}

func (x *GetEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityRequest.ProtoReflect.Descriptor instead.
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{1}
}

func (x *GetEntityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEntityRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type UpdateEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *UpdateEntityRequest) Reset() {
	*x = UpdateEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEntityRequest) ProtoMessage() {}

func (x *UpdateEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEntityRequest.ProtoReflect.Descriptor instead.
func (*UpdateEntityRequest) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateEntityRequest) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type DeleteEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *DeleteEntityRequest) Reset() {
	*x = DeleteEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntityRequest) ProtoMessage() {}

func (x *DeleteEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntityRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntityRequest) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEntityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteEntityRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type ListEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ListEntitiesRequest) Reset() {
	*x = ListEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntitiesRequest) ProtoMessage() {}

func (x *ListEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntitiesRequest.ProtoReflect.Descriptor instead.
func (*ListEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{4}
}

func (x *ListEntitiesRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type ListEntitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *ListEntitiesResponse) Reset() {
	*x = ListEntitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEntitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEntitiesResponse) ProtoMessage() {}

func (x *ListEntitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEntitiesResponse.ProtoReflect.Descriptor instead.
func (*ListEntitiesResponse) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{5}
}

func (x *ListEntitiesResponse) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the entity.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Name of Feast project that this entity belongs to.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Type of the entity.
	ValueType ValueType_Enum `protobuf:"varint,3,opt,name=value_type,json=valueType,proto3,enum=api.ValueType_Enum" json:"value_type,omitempty"`
	// Description of the entity.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Join key for the entity (i.e. name of the column the entity maps to).
	JoinKey string `protobuf:"bytes,5,opt,name=join_key,json=joinKey,proto3" json:"join_key,omitempty"`
	// User defined metadata.
	Tags map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Owner of the entity.
	Owner string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	// Creation time of the entity.
	CreatedTimestamp *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_timestamp,json=createdTimestamp,proto3" json:"created_timestamp,omitempty"`
	// Last update time of the entity.
	LastUpdatedTimestamp *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_updated_timestamp,json=lastUpdatedTimestamp,proto3" json:"last_updated_timestamp,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_EntityService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_EntityService_proto_rawDescGZIP(), []int{6}
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Entity) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Entity) GetValueType() ValueType_Enum {
	if x != nil {
		return x.ValueType
	}
	return ValueType_INVALID
}

func (x *Entity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Entity) GetJoinKey() string {
	if x != nil {
		return x.JoinKey
	}
	return ""
}

func (x *Entity) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Entity) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Entity) GetCreatedTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTimestamp
	}
	return nil
}

func (x *Entity) GetLastUpdatedTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdatedTimestamp
	}
	return nil
}

var File_EntityService_proto protoreflect.FileDescriptor

var file_EntityService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0f, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x43, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xbc, 0x03, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x32, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x50, 0x0a, 0x16,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x37,
	0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xb5, 0x03, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x0d, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x43, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x0d, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x3a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42,
	0x82, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x65, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x92, 0x41, 0x4d, 0x52, 0x1c, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x12, 0x11, 0x12, 0x0f, 0x0a, 0x0d, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5a, 0x1f, 0x0a, 0x1d, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityService_proto_rawDescOnce sync.Once
	file_EntityService_proto_rawDescData = file_EntityService_proto_rawDesc
)

func file_EntityService_proto_rawDescGZIP() []byte {
	file_EntityService_proto_rawDescOnce.Do(func() {
		file_EntityService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityService_proto_rawDescData)
	})
	return file_EntityService_proto_rawDescData
}

var file_EntityService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_EntityService_proto_goTypes = []interface{}{
	(*CreateEntityRequest)(nil),  // 0: api.CreateEntityRequest
	(*GetEntityRequest)(nil),     // 1: api.GetEntityRequest
	(*UpdateEntityRequest)(nil),  // 2: api.UpdateEntityRequest
	(*DeleteEntityRequest)(nil),  // 3: api.DeleteEntityRequest
	(*ListEntitiesRequest)(nil),  // 4: api.ListEntitiesRequest
	(*ListEntitiesResponse)(nil), // 5: api.ListEntitiesResponse
	(*Entity)(nil),               // 6: api.Entity
	nil,                          // 7: api.Entity.TagsEntry
	(ValueType_Enum)(0),          // 8: api.ValueType.Enum
	(*timestamp.Timestamp)(nil),  // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_EntityService_proto_depIdxs = []int32{
	6,  // 0: api.CreateEntityRequest.entity:type_name -> api.Entity
	6,  // 1: api.UpdateEntityRequest.entity:type_name -> api.Entity
	6,  // 2: api.ListEntitiesResponse.entities:type_name -> api.Entity
	8,  // 3: api.Entity.value_type:type_name -> api.ValueType.Enum
	7,  // 4: api.Entity.tags:type_name -> api.Entity.TagsEntry
	9,  // 5: api.Entity.created_timestamp:type_name -> google.protobuf.Timestamp
	9,  // 6: api.Entity.last_updated_timestamp:type_name -> google.protobuf.Timestamp
	0,  // 7: api.EntityService.CreateEntity:input_type -> api.CreateEntityRequest
	1,  // 8: api.EntityService.GetEntity:input_type -> api.GetEntityRequest
	2,  // 9: api.EntityService.UpdateEntity:input_type -> api.UpdateEntityRequest
	3,  // 10: api.EntityService.DeleteEntity:input_type -> api.DeleteEntityRequest
	4,  // 11: api.EntityService.ListEntities:input_type -> api.ListEntitiesRequest
	6,  // 12: api.EntityService.CreateEntity:output_type -> api.Entity
	6,  // 13: api.EntityService.GetEntity:output_type -> api.Entity
	6,  // 14: api.EntityService.UpdateEntity:output_type -> api.Entity
	10, // 15: api.EntityService.DeleteEntity:output_type -> google.protobuf.Empty
	5,  // 16: api.EntityService.ListEntities:output_type -> api.ListEntitiesResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_EntityService_proto_init() }
func file_EntityService_proto_init() {
	if File_EntityService_proto != nil {
		return
	}
	file_ValueType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEntityRequest); i {
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
		file_EntityService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityRequest); i {
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
		file_EntityService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEntityRequest); i {
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
		file_EntityService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntityRequest); i {
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
		file_EntityService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntitiesRequest); i {
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
		file_EntityService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEntitiesResponse); i {
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
		file_EntityService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
			RawDescriptor: file_EntityService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_EntityService_proto_goTypes,
		DependencyIndexes: file_EntityService_proto_depIdxs,
		MessageInfos:      file_EntityService_proto_msgTypes,
	}.Build()
	File_EntityService_proto = out.File
	file_EntityService_proto_rawDesc = nil
	file_EntityService_proto_goTypes = nil
	file_EntityService_proto_depIdxs = nil
}
