//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: yarn_server_nodemanager_service_protos.proto

package nm

import (
	api "github.com/kom0055/go-hadoop/proto/yarn/api"
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

type ResourceStatusTypeProto int32

const (
	ResourceStatusTypeProto_FETCH_PENDING ResourceStatusTypeProto = 1
	ResourceStatusTypeProto_FETCH_SUCCESS ResourceStatusTypeProto = 2
	ResourceStatusTypeProto_FETCH_FAILURE ResourceStatusTypeProto = 3
)

// Enum value maps for ResourceStatusTypeProto.
var (
	ResourceStatusTypeProto_name = map[int32]string{
		1: "FETCH_PENDING",
		2: "FETCH_SUCCESS",
		3: "FETCH_FAILURE",
	}
	ResourceStatusTypeProto_value = map[string]int32{
		"FETCH_PENDING": 1,
		"FETCH_SUCCESS": 2,
		"FETCH_FAILURE": 3,
	}
)

func (x ResourceStatusTypeProto) Enum() *ResourceStatusTypeProto {
	p := new(ResourceStatusTypeProto)
	*p = x
	return p
}

func (x ResourceStatusTypeProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceStatusTypeProto) Descriptor() protoreflect.EnumDescriptor {
	return file_yarn_server_nodemanager_service_protos_proto_enumTypes[0].Descriptor()
}

func (ResourceStatusTypeProto) Type() protoreflect.EnumType {
	return &file_yarn_server_nodemanager_service_protos_proto_enumTypes[0]
}

func (x ResourceStatusTypeProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ResourceStatusTypeProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ResourceStatusTypeProto(num)
	return nil
}

// Deprecated: Use ResourceStatusTypeProto.Descriptor instead.
func (ResourceStatusTypeProto) EnumDescriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{0}
}

type LocalizerActionProto int32

const (
	LocalizerActionProto_LIVE LocalizerActionProto = 1
	LocalizerActionProto_DIE  LocalizerActionProto = 2
)

// Enum value maps for LocalizerActionProto.
var (
	LocalizerActionProto_name = map[int32]string{
		1: "LIVE",
		2: "DIE",
	}
	LocalizerActionProto_value = map[string]int32{
		"LIVE": 1,
		"DIE":  2,
	}
)

func (x LocalizerActionProto) Enum() *LocalizerActionProto {
	p := new(LocalizerActionProto)
	*p = x
	return p
}

func (x LocalizerActionProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocalizerActionProto) Descriptor() protoreflect.EnumDescriptor {
	return file_yarn_server_nodemanager_service_protos_proto_enumTypes[1].Descriptor()
}

func (LocalizerActionProto) Type() protoreflect.EnumType {
	return &file_yarn_server_nodemanager_service_protos_proto_enumTypes[1]
}

func (x LocalizerActionProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LocalizerActionProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LocalizerActionProto(num)
	return nil
}

// Deprecated: Use LocalizerActionProto.Descriptor instead.
func (LocalizerActionProto) EnumDescriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{1}
}

type LocalResourceStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource  *api.LocalResourceProto       `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	Status    *ResourceStatusTypeProto      `protobuf:"varint,2,opt,name=status,enum=hadoop.yarn.ResourceStatusTypeProto" json:"status,omitempty"`
	LocalPath *api.URLProto                 `protobuf:"bytes,3,opt,name=localPath" json:"localPath,omitempty"`
	LocalSize *int64                        `protobuf:"varint,4,opt,name=localSize" json:"localSize,omitempty"`
	Exception *api.SerializedExceptionProto `protobuf:"bytes,5,opt,name=exception" json:"exception,omitempty"`
}

func (x *LocalResourceStatusProto) Reset() {
	*x = LocalResourceStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalResourceStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalResourceStatusProto) ProtoMessage() {}

func (x *LocalResourceStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalResourceStatusProto.ProtoReflect.Descriptor instead.
func (*LocalResourceStatusProto) Descriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{0}
}

func (x *LocalResourceStatusProto) GetResource() *api.LocalResourceProto {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *LocalResourceStatusProto) GetStatus() ResourceStatusTypeProto {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ResourceStatusTypeProto_FETCH_PENDING
}

func (x *LocalResourceStatusProto) GetLocalPath() *api.URLProto {
	if x != nil {
		return x.LocalPath
	}
	return nil
}

func (x *LocalResourceStatusProto) GetLocalSize() int64 {
	if x != nil && x.LocalSize != nil {
		return *x.LocalSize
	}
	return 0
}

func (x *LocalResourceStatusProto) GetException() *api.SerializedExceptionProto {
	if x != nil {
		return x.Exception
	}
	return nil
}

type LocalizerStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalizerId *string                     `protobuf:"bytes,1,opt,name=localizer_id,json=localizerId" json:"localizer_id,omitempty"`
	Resources   []*LocalResourceStatusProto `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
}

func (x *LocalizerStatusProto) Reset() {
	*x = LocalizerStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizerStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizerStatusProto) ProtoMessage() {}

func (x *LocalizerStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizerStatusProto.ProtoReflect.Descriptor instead.
func (*LocalizerStatusProto) Descriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{1}
}

func (x *LocalizerStatusProto) GetLocalizerId() string {
	if x != nil && x.LocalizerId != nil {
		return *x.LocalizerId
	}
	return ""
}

func (x *LocalizerStatusProto) GetResources() []*LocalResourceStatusProto {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ResourceLocalizationSpecProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource             *api.LocalResourceProto `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	DestinationDirectory *api.URLProto           `protobuf:"bytes,2,opt,name=destination_directory,json=destinationDirectory" json:"destination_directory,omitempty"`
}

func (x *ResourceLocalizationSpecProto) Reset() {
	*x = ResourceLocalizationSpecProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLocalizationSpecProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLocalizationSpecProto) ProtoMessage() {}

func (x *ResourceLocalizationSpecProto) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLocalizationSpecProto.ProtoReflect.Descriptor instead.
func (*ResourceLocalizationSpecProto) Descriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{2}
}

func (x *ResourceLocalizationSpecProto) GetResource() *api.LocalResourceProto {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceLocalizationSpecProto) GetDestinationDirectory() *api.URLProto {
	if x != nil {
		return x.DestinationDirectory
	}
	return nil
}

type LocalizerHeartbeatResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    *LocalizerActionProto            `protobuf:"varint,1,opt,name=action,enum=hadoop.yarn.LocalizerActionProto" json:"action,omitempty"`
	Resources []*ResourceLocalizationSpecProto `protobuf:"bytes,2,rep,name=resources" json:"resources,omitempty"`
}

func (x *LocalizerHeartbeatResponseProto) Reset() {
	*x = LocalizerHeartbeatResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalizerHeartbeatResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalizerHeartbeatResponseProto) ProtoMessage() {}

func (x *LocalizerHeartbeatResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_server_nodemanager_service_protos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalizerHeartbeatResponseProto.ProtoReflect.Descriptor instead.
func (*LocalizerHeartbeatResponseProto) Descriptor() ([]byte, []int) {
	return file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP(), []int{3}
}

func (x *LocalizerHeartbeatResponseProto) GetAction() LocalizerActionProto {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return LocalizerActionProto_LIVE
}

func (x *LocalizerHeartbeatResponseProto) GetResources() []*ResourceLocalizationSpecProto {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_yarn_server_nodemanager_service_protos_proto protoreflect.FileDescriptor

var file_yarn_server_nodemanager_service_protos_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x72, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x1a, 0x11, 0x79, 0x61, 0x72,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad,
	0x02, 0x0a, 0x18, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x55, 0x52, 0x4c, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7e,
	0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xa8,
	0x01, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x15, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x55, 0x52, 0x4c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xa6, 0x01, 0x0a, 0x1f, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2a, 0x52, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x0a,
	0x0d, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x45, 0x54, 0x43, 0x48, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x03, 0x2a, 0x29, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x49, 0x45, 0x10,
	0x02, 0x42, 0x7b, 0x0a, 0x1c, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x22, 0x59, 0x61, 0x72, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x72, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x6d, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_yarn_server_nodemanager_service_protos_proto_rawDescOnce sync.Once
	file_yarn_server_nodemanager_service_protos_proto_rawDescData = file_yarn_server_nodemanager_service_protos_proto_rawDesc
)

func file_yarn_server_nodemanager_service_protos_proto_rawDescGZIP() []byte {
	file_yarn_server_nodemanager_service_protos_proto_rawDescOnce.Do(func() {
		file_yarn_server_nodemanager_service_protos_proto_rawDescData = protoimpl.X.CompressGZIP(file_yarn_server_nodemanager_service_protos_proto_rawDescData)
	})
	return file_yarn_server_nodemanager_service_protos_proto_rawDescData
}

var file_yarn_server_nodemanager_service_protos_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yarn_server_nodemanager_service_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yarn_server_nodemanager_service_protos_proto_goTypes = []interface{}{
	(ResourceStatusTypeProto)(0),            // 0: hadoop.yarn.ResourceStatusTypeProto
	(LocalizerActionProto)(0),               // 1: hadoop.yarn.LocalizerActionProto
	(*LocalResourceStatusProto)(nil),        // 2: hadoop.yarn.LocalResourceStatusProto
	(*LocalizerStatusProto)(nil),            // 3: hadoop.yarn.LocalizerStatusProto
	(*ResourceLocalizationSpecProto)(nil),   // 4: hadoop.yarn.ResourceLocalizationSpecProto
	(*LocalizerHeartbeatResponseProto)(nil), // 5: hadoop.yarn.LocalizerHeartbeatResponseProto
	(*api.LocalResourceProto)(nil),          // 6: hadoop.yarn.LocalResourceProto
	(*api.URLProto)(nil),                    // 7: hadoop.yarn.URLProto
	(*api.SerializedExceptionProto)(nil),    // 8: hadoop.yarn.SerializedExceptionProto
}
var file_yarn_server_nodemanager_service_protos_proto_depIdxs = []int32{
	6, // 0: hadoop.yarn.LocalResourceStatusProto.resource:type_name -> hadoop.yarn.LocalResourceProto
	0, // 1: hadoop.yarn.LocalResourceStatusProto.status:type_name -> hadoop.yarn.ResourceStatusTypeProto
	7, // 2: hadoop.yarn.LocalResourceStatusProto.localPath:type_name -> hadoop.yarn.URLProto
	8, // 3: hadoop.yarn.LocalResourceStatusProto.exception:type_name -> hadoop.yarn.SerializedExceptionProto
	2, // 4: hadoop.yarn.LocalizerStatusProto.resources:type_name -> hadoop.yarn.LocalResourceStatusProto
	6, // 5: hadoop.yarn.ResourceLocalizationSpecProto.resource:type_name -> hadoop.yarn.LocalResourceProto
	7, // 6: hadoop.yarn.ResourceLocalizationSpecProto.destination_directory:type_name -> hadoop.yarn.URLProto
	1, // 7: hadoop.yarn.LocalizerHeartbeatResponseProto.action:type_name -> hadoop.yarn.LocalizerActionProto
	4, // 8: hadoop.yarn.LocalizerHeartbeatResponseProto.resources:type_name -> hadoop.yarn.ResourceLocalizationSpecProto
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yarn_server_nodemanager_service_protos_proto_init() }
func file_yarn_server_nodemanager_service_protos_proto_init() {
	if File_yarn_server_nodemanager_service_protos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yarn_server_nodemanager_service_protos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalResourceStatusProto); i {
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
		file_yarn_server_nodemanager_service_protos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizerStatusProto); i {
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
		file_yarn_server_nodemanager_service_protos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLocalizationSpecProto); i {
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
		file_yarn_server_nodemanager_service_protos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalizerHeartbeatResponseProto); i {
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
			RawDescriptor: file_yarn_server_nodemanager_service_protos_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yarn_server_nodemanager_service_protos_proto_goTypes,
		DependencyIndexes: file_yarn_server_nodemanager_service_protos_proto_depIdxs,
		EnumInfos:         file_yarn_server_nodemanager_service_protos_proto_enumTypes,
		MessageInfos:      file_yarn_server_nodemanager_service_protos_proto_msgTypes,
	}.Build()
	File_yarn_server_nodemanager_service_protos_proto = out.File
	file_yarn_server_nodemanager_service_protos_proto_rawDesc = nil
	file_yarn_server_nodemanager_service_protos_proto_goTypes = nil
	file_yarn_server_nodemanager_service_protos_proto_depIdxs = nil
}
