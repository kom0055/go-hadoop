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
// source: ProtobufRpcEngine.proto

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
// This message is the header for the Protobuf Rpc Engine
// when sending a RPC request from  RPC client to the RPC server.
// The actual request (serialized as protobuf) follows this request.
//
// No special header is needed for the Rpc Response for Protobuf Rpc Engine.
// The normal RPC response header (see RpcHeader.proto) are sufficient.
type RequestHeaderProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* Name of the RPC method
	MethodName *string `protobuf:"bytes,1,req,name=methodName" json:"methodName,omitempty"`
	//*
	// RPCs for a particular interface (ie protocol) are done using a
	// IPC connection that is setup using rpcProxy.
	// The rpcProxy's has a declared protocol name that is
	// sent form client to server at connection time.
	//
	// Each Rpc call also sends a protocol name
	// (called declaringClassprotocolName). This name is usually the same
	// as the connection protocol name except in some cases.
	// For example metaProtocols such ProtocolInfoProto which get metainfo
	// about the protocol reuse the connection but need to indicate that
	// the actual protocol is different (i.e. the protocol is
	// ProtocolInfoProto) since they reuse the connection; in this case
	// the declaringClassProtocolName field is set to the ProtocolInfoProto
	DeclaringClassProtocolName *string `protobuf:"bytes,2,req,name=declaringClassProtocolName" json:"declaringClassProtocolName,omitempty"`
	//* protocol version of class declaring the called method
	ClientProtocolVersion *uint64 `protobuf:"varint,3,req,name=clientProtocolVersion" json:"clientProtocolVersion,omitempty"`
}

func (x *RequestHeaderProto) Reset() {
	*x = RequestHeaderProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtobufRpcEngine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeaderProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeaderProto) ProtoMessage() {}

func (x *RequestHeaderProto) ProtoReflect() protoreflect.Message {
	mi := &file_ProtobufRpcEngine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeaderProto.ProtoReflect.Descriptor instead.
func (*RequestHeaderProto) Descriptor() ([]byte, []int) {
	return file_ProtobufRpcEngine_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeaderProto) GetMethodName() string {
	if x != nil && x.MethodName != nil {
		return *x.MethodName
	}
	return ""
}

func (x *RequestHeaderProto) GetDeclaringClassProtocolName() string {
	if x != nil && x.DeclaringClassProtocolName != nil {
		return *x.DeclaringClassProtocolName
	}
	return ""
}

func (x *RequestHeaderProto) GetClientProtocolVersion() uint64 {
	if x != nil && x.ClientProtocolVersion != nil {
		return *x.ClientProtocolVersion
	}
	return 0
}

var File_ProtobufRpcEngine_proto protoreflect.FileDescriptor

var file_ProtobufRpcEngine_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x52, 0x70, 0x63, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xaa, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x1a, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x1a, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x15,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x67, 0x0a, 0x1e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61,
	0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x52, 0x70, 0x63, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d,
	0x30, 0x30, 0x35, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01,
}

var (
	file_ProtobufRpcEngine_proto_rawDescOnce sync.Once
	file_ProtobufRpcEngine_proto_rawDescData = file_ProtobufRpcEngine_proto_rawDesc
)

func file_ProtobufRpcEngine_proto_rawDescGZIP() []byte {
	file_ProtobufRpcEngine_proto_rawDescOnce.Do(func() {
		file_ProtobufRpcEngine_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtobufRpcEngine_proto_rawDescData)
	})
	return file_ProtobufRpcEngine_proto_rawDescData
}

var file_ProtobufRpcEngine_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ProtobufRpcEngine_proto_goTypes = []interface{}{
	(*RequestHeaderProto)(nil), // 0: hadoop.common.RequestHeaderProto
}
var file_ProtobufRpcEngine_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProtobufRpcEngine_proto_init() }
func file_ProtobufRpcEngine_proto_init() {
	if File_ProtobufRpcEngine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtobufRpcEngine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeaderProto); i {
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
			RawDescriptor: file_ProtobufRpcEngine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProtobufRpcEngine_proto_goTypes,
		DependencyIndexes: file_ProtobufRpcEngine_proto_depIdxs,
		MessageInfos:      file_ProtobufRpcEngine_proto_msgTypes,
	}.Build()
	File_ProtobufRpcEngine_proto = out.File
	file_ProtobufRpcEngine_proto_rawDesc = nil
	file_ProtobufRpcEngine_proto_goTypes = nil
	file_ProtobufRpcEngine_proto_depIdxs = nil
}
