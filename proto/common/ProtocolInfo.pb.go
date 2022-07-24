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

//*
// These .proto interfaces are private and stable.
// Please see http://wiki.apache.org/hadoop/Compatibility
// for what changes are allowed for a *stable* .proto interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: ProtocolInfo.proto

package common

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

//*
// Request to get protocol versions for all supported rpc kinds.
type GetProtocolVersionsRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol *string `protobuf:"bytes,1,req,name=protocol" json:"protocol,omitempty"` // Protocol name
}

func (x *GetProtocolVersionsRequestProto) Reset() {
	*x = GetProtocolVersionsRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtocolVersionsRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtocolVersionsRequestProto) ProtoMessage() {}

func (x *GetProtocolVersionsRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtocolVersionsRequestProto.ProtoReflect.Descriptor instead.
func (*GetProtocolVersionsRequestProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GetProtocolVersionsRequestProto) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

//*
// Protocol version with corresponding rpc kind.
type ProtocolVersionProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpcKind  *string  `protobuf:"bytes,1,req,name=rpcKind" json:"rpcKind,omitempty"`    //RPC kind
	Versions []uint64 `protobuf:"varint,2,rep,name=versions" json:"versions,omitempty"` //Protocol version corresponding to the rpc kind.
}

func (x *ProtocolVersionProto) Reset() {
	*x = ProtocolVersionProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolVersionProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolVersionProto) ProtoMessage() {}

func (x *ProtocolVersionProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolVersionProto.ProtoReflect.Descriptor instead.
func (*ProtocolVersionProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{1}
}

func (x *ProtocolVersionProto) GetRpcKind() string {
	if x != nil && x.RpcKind != nil {
		return *x.RpcKind
	}
	return ""
}

func (x *ProtocolVersionProto) GetVersions() []uint64 {
	if x != nil {
		return x.Versions
	}
	return nil
}

//*
// Get protocol version response.
type GetProtocolVersionsResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolVersions []*ProtocolVersionProto `protobuf:"bytes,1,rep,name=protocolVersions" json:"protocolVersions,omitempty"`
}

func (x *GetProtocolVersionsResponseProto) Reset() {
	*x = GetProtocolVersionsResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtocolVersionsResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtocolVersionsResponseProto) ProtoMessage() {}

func (x *GetProtocolVersionsResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtocolVersionsResponseProto.ProtoReflect.Descriptor instead.
func (*GetProtocolVersionsResponseProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{2}
}

func (x *GetProtocolVersionsResponseProto) GetProtocolVersions() []*ProtocolVersionProto {
	if x != nil {
		return x.ProtocolVersions
	}
	return nil
}

//*
// Get protocol signature request.
type GetProtocolSignatureRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol *string `protobuf:"bytes,1,req,name=protocol" json:"protocol,omitempty"` // Protocol name
	RpcKind  *string `protobuf:"bytes,2,req,name=rpcKind" json:"rpcKind,omitempty"`   // RPC kind
}

func (x *GetProtocolSignatureRequestProto) Reset() {
	*x = GetProtocolSignatureRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtocolSignatureRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtocolSignatureRequestProto) ProtoMessage() {}

func (x *GetProtocolSignatureRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtocolSignatureRequestProto.ProtoReflect.Descriptor instead.
func (*GetProtocolSignatureRequestProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{3}
}

func (x *GetProtocolSignatureRequestProto) GetProtocol() string {
	if x != nil && x.Protocol != nil {
		return *x.Protocol
	}
	return ""
}

func (x *GetProtocolSignatureRequestProto) GetRpcKind() string {
	if x != nil && x.RpcKind != nil {
		return *x.RpcKind
	}
	return ""
}

//*
// Get protocol signature response.
type GetProtocolSignatureResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtocolSignature []*ProtocolSignatureProto `protobuf:"bytes,1,rep,name=protocolSignature" json:"protocolSignature,omitempty"`
}

func (x *GetProtocolSignatureResponseProto) Reset() {
	*x = GetProtocolSignatureResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProtocolSignatureResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProtocolSignatureResponseProto) ProtoMessage() {}

func (x *GetProtocolSignatureResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProtocolSignatureResponseProto.ProtoReflect.Descriptor instead.
func (*GetProtocolSignatureResponseProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{4}
}

func (x *GetProtocolSignatureResponseProto) GetProtocolSignature() []*ProtocolSignatureProto {
	if x != nil {
		return x.ProtocolSignature
	}
	return nil
}

type ProtocolSignatureProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *uint64  `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Methods []uint32 `protobuf:"varint,2,rep,name=methods" json:"methods,omitempty"`
}

func (x *ProtocolSignatureProto) Reset() {
	*x = ProtocolSignatureProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtocolInfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolSignatureProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolSignatureProto) ProtoMessage() {}

func (x *ProtocolSignatureProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtocolInfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolSignatureProto.ProtoReflect.Descriptor instead.
func (*ProtocolSignatureProto) Descriptor() ([]byte, []int) {
	return file_ProtocolInfo_proto_rawDescGZIP(), []int{5}
}

func (x *ProtocolSignatureProto) GetVersion() uint64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *ProtocolSignatureProto) GetMethods() []uint32 {
	if x != nil {
		return x.Methods
	}
	return nil
}

var File_ProtocolInfo_proto protoreflect.FileDescriptor

var file_ProtocolInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x70,
	0x63, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x72, 0x70, 0x63,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x73, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4f, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x70, 0x63, 0x4b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x72, 0x70, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x22,
	0x78, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x53, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4c, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x32, 0x88, 0x02, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x76, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x79, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x2f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x30, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x65, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65,
	0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2f, 0x67, 0x6f, 0x2d,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_ProtocolInfo_proto_rawDescOnce sync.Once
	file_ProtocolInfo_proto_rawDescData = file_ProtocolInfo_proto_rawDesc
)

func file_ProtocolInfo_proto_rawDescGZIP() []byte {
	file_ProtocolInfo_proto_rawDescOnce.Do(func() {
		file_ProtocolInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtocolInfo_proto_rawDescData)
	})
	return file_ProtocolInfo_proto_rawDescData
}

var file_ProtocolInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ProtocolInfo_proto_goTypes = []interface{}{
	(*GetProtocolVersionsRequestProto)(nil),   // 0: hadoop.common.GetProtocolVersionsRequestProto
	(*ProtocolVersionProto)(nil),              // 1: hadoop.common.ProtocolVersionProto
	(*GetProtocolVersionsResponseProto)(nil),  // 2: hadoop.common.GetProtocolVersionsResponseProto
	(*GetProtocolSignatureRequestProto)(nil),  // 3: hadoop.common.GetProtocolSignatureRequestProto
	(*GetProtocolSignatureResponseProto)(nil), // 4: hadoop.common.GetProtocolSignatureResponseProto
	(*ProtocolSignatureProto)(nil),            // 5: hadoop.common.ProtocolSignatureProto
}
var file_ProtocolInfo_proto_depIdxs = []int32{
	1, // 0: hadoop.common.GetProtocolVersionsResponseProto.protocolVersions:type_name -> hadoop.common.ProtocolVersionProto
	5, // 1: hadoop.common.GetProtocolSignatureResponseProto.protocolSignature:type_name -> hadoop.common.ProtocolSignatureProto
	0, // 2: hadoop.common.ProtocolInfoService.getProtocolVersions:input_type -> hadoop.common.GetProtocolVersionsRequestProto
	3, // 3: hadoop.common.ProtocolInfoService.getProtocolSignature:input_type -> hadoop.common.GetProtocolSignatureRequestProto
	2, // 4: hadoop.common.ProtocolInfoService.getProtocolVersions:output_type -> hadoop.common.GetProtocolVersionsResponseProto
	4, // 5: hadoop.common.ProtocolInfoService.getProtocolSignature:output_type -> hadoop.common.GetProtocolSignatureResponseProto
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ProtocolInfo_proto_init() }
func file_ProtocolInfo_proto_init() {
	if File_ProtocolInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtocolInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtocolVersionsRequestProto); i {
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
		file_ProtocolInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolVersionProto); i {
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
		file_ProtocolInfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtocolVersionsResponseProto); i {
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
		file_ProtocolInfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtocolSignatureRequestProto); i {
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
		file_ProtocolInfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProtocolSignatureResponseProto); i {
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
		file_ProtocolInfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolSignatureProto); i {
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
			RawDescriptor: file_ProtocolInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProtocolInfo_proto_goTypes,
		DependencyIndexes: file_ProtocolInfo_proto_depIdxs,
		MessageInfos:      file_ProtocolInfo_proto_msgTypes,
	}.Build()
	File_ProtocolInfo_proto = out.File
	file_ProtocolInfo_proto_rawDesc = nil
	file_ProtocolInfo_proto_goTypes = nil
	file_ProtocolInfo_proto_depIdxs = nil
}