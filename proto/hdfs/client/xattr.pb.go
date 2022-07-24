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
// source: xattr.proto

package client

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

type XAttrSetFlagProto int32

const (
	XAttrSetFlagProto_XATTR_CREATE  XAttrSetFlagProto = 1
	XAttrSetFlagProto_XATTR_REPLACE XAttrSetFlagProto = 2
)

// Enum value maps for XAttrSetFlagProto.
var (
	XAttrSetFlagProto_name = map[int32]string{
		1: "XATTR_CREATE",
		2: "XATTR_REPLACE",
	}
	XAttrSetFlagProto_value = map[string]int32{
		"XATTR_CREATE":  1,
		"XATTR_REPLACE": 2,
	}
)

func (x XAttrSetFlagProto) Enum() *XAttrSetFlagProto {
	p := new(XAttrSetFlagProto)
	*p = x
	return p
}

func (x XAttrSetFlagProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XAttrSetFlagProto) Descriptor() protoreflect.EnumDescriptor {
	return file_xattr_proto_enumTypes[0].Descriptor()
}

func (XAttrSetFlagProto) Type() protoreflect.EnumType {
	return &file_xattr_proto_enumTypes[0]
}

func (x XAttrSetFlagProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *XAttrSetFlagProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = XAttrSetFlagProto(num)
	return nil
}

// Deprecated: Use XAttrSetFlagProto.Descriptor instead.
func (XAttrSetFlagProto) EnumDescriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{0}
}

type XAttrProto_XAttrNamespaceProto int32

const (
	XAttrProto_USER     XAttrProto_XAttrNamespaceProto = 0
	XAttrProto_TRUSTED  XAttrProto_XAttrNamespaceProto = 1
	XAttrProto_SECURITY XAttrProto_XAttrNamespaceProto = 2
	XAttrProto_SYSTEM   XAttrProto_XAttrNamespaceProto = 3
	XAttrProto_RAW      XAttrProto_XAttrNamespaceProto = 4
)

// Enum value maps for XAttrProto_XAttrNamespaceProto.
var (
	XAttrProto_XAttrNamespaceProto_name = map[int32]string{
		0: "USER",
		1: "TRUSTED",
		2: "SECURITY",
		3: "SYSTEM",
		4: "RAW",
	}
	XAttrProto_XAttrNamespaceProto_value = map[string]int32{
		"USER":     0,
		"TRUSTED":  1,
		"SECURITY": 2,
		"SYSTEM":   3,
		"RAW":      4,
	}
)

func (x XAttrProto_XAttrNamespaceProto) Enum() *XAttrProto_XAttrNamespaceProto {
	p := new(XAttrProto_XAttrNamespaceProto)
	*p = x
	return p
}

func (x XAttrProto_XAttrNamespaceProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XAttrProto_XAttrNamespaceProto) Descriptor() protoreflect.EnumDescriptor {
	return file_xattr_proto_enumTypes[1].Descriptor()
}

func (XAttrProto_XAttrNamespaceProto) Type() protoreflect.EnumType {
	return &file_xattr_proto_enumTypes[1]
}

func (x XAttrProto_XAttrNamespaceProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *XAttrProto_XAttrNamespaceProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = XAttrProto_XAttrNamespaceProto(num)
	return nil
}

// Deprecated: Use XAttrProto_XAttrNamespaceProto.Descriptor instead.
func (XAttrProto_XAttrNamespaceProto) EnumDescriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{0, 0}
}

type XAttrProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *XAttrProto_XAttrNamespaceProto `protobuf:"varint,1,req,name=namespace,enum=hadoop.hdfs.XAttrProto_XAttrNamespaceProto" json:"namespace,omitempty"`
	Name      *string                         `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Value     []byte                          `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (x *XAttrProto) Reset() {
	*x = XAttrProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XAttrProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XAttrProto) ProtoMessage() {}

func (x *XAttrProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XAttrProto.ProtoReflect.Descriptor instead.
func (*XAttrProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{0}
}

func (x *XAttrProto) GetNamespace() XAttrProto_XAttrNamespaceProto {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return XAttrProto_USER
}

func (x *XAttrProto) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *XAttrProto) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetXAttrRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src   *string     `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttr *XAttrProto `protobuf:"bytes,2,opt,name=xAttr" json:"xAttr,omitempty"`
	Flag  *uint32     `protobuf:"varint,3,opt,name=flag" json:"flag,omitempty"` //bits set using XAttrSetFlagProto
}

func (x *SetXAttrRequestProto) Reset() {
	*x = SetXAttrRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetXAttrRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetXAttrRequestProto) ProtoMessage() {}

func (x *SetXAttrRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetXAttrRequestProto.ProtoReflect.Descriptor instead.
func (*SetXAttrRequestProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{1}
}

func (x *SetXAttrRequestProto) GetSrc() string {
	if x != nil && x.Src != nil {
		return *x.Src
	}
	return ""
}

func (x *SetXAttrRequestProto) GetXAttr() *XAttrProto {
	if x != nil {
		return x.XAttr
	}
	return nil
}

func (x *SetXAttrRequestProto) GetFlag() uint32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

type SetXAttrResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetXAttrResponseProto) Reset() {
	*x = SetXAttrResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetXAttrResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetXAttrResponseProto) ProtoMessage() {}

func (x *SetXAttrResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetXAttrResponseProto.ProtoReflect.Descriptor instead.
func (*SetXAttrResponseProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{2}
}

type GetXAttrsRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src    *string       `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttrs []*XAttrProto `protobuf:"bytes,2,rep,name=xAttrs" json:"xAttrs,omitempty"`
}

func (x *GetXAttrsRequestProto) Reset() {
	*x = GetXAttrsRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetXAttrsRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetXAttrsRequestProto) ProtoMessage() {}

func (x *GetXAttrsRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetXAttrsRequestProto.ProtoReflect.Descriptor instead.
func (*GetXAttrsRequestProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{3}
}

func (x *GetXAttrsRequestProto) GetSrc() string {
	if x != nil && x.Src != nil {
		return *x.Src
	}
	return ""
}

func (x *GetXAttrsRequestProto) GetXAttrs() []*XAttrProto {
	if x != nil {
		return x.XAttrs
	}
	return nil
}

type GetXAttrsResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAttrs []*XAttrProto `protobuf:"bytes,1,rep,name=xAttrs" json:"xAttrs,omitempty"`
}

func (x *GetXAttrsResponseProto) Reset() {
	*x = GetXAttrsResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetXAttrsResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetXAttrsResponseProto) ProtoMessage() {}

func (x *GetXAttrsResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetXAttrsResponseProto.ProtoReflect.Descriptor instead.
func (*GetXAttrsResponseProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{4}
}

func (x *GetXAttrsResponseProto) GetXAttrs() []*XAttrProto {
	if x != nil {
		return x.XAttrs
	}
	return nil
}

type ListXAttrsRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
}

func (x *ListXAttrsRequestProto) Reset() {
	*x = ListXAttrsRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListXAttrsRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListXAttrsRequestProto) ProtoMessage() {}

func (x *ListXAttrsRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListXAttrsRequestProto.ProtoReflect.Descriptor instead.
func (*ListXAttrsRequestProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{5}
}

func (x *ListXAttrsRequestProto) GetSrc() string {
	if x != nil && x.Src != nil {
		return *x.Src
	}
	return ""
}

type ListXAttrsResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XAttrs []*XAttrProto `protobuf:"bytes,1,rep,name=xAttrs" json:"xAttrs,omitempty"`
}

func (x *ListXAttrsResponseProto) Reset() {
	*x = ListXAttrsResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListXAttrsResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListXAttrsResponseProto) ProtoMessage() {}

func (x *ListXAttrsResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListXAttrsResponseProto.ProtoReflect.Descriptor instead.
func (*ListXAttrsResponseProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{6}
}

func (x *ListXAttrsResponseProto) GetXAttrs() []*XAttrProto {
	if x != nil {
		return x.XAttrs
	}
	return nil
}

type RemoveXAttrRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src   *string     `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XAttr *XAttrProto `protobuf:"bytes,2,opt,name=xAttr" json:"xAttr,omitempty"`
}

func (x *RemoveXAttrRequestProto) Reset() {
	*x = RemoveXAttrRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveXAttrRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveXAttrRequestProto) ProtoMessage() {}

func (x *RemoveXAttrRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveXAttrRequestProto.ProtoReflect.Descriptor instead.
func (*RemoveXAttrRequestProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveXAttrRequestProto) GetSrc() string {
	if x != nil && x.Src != nil {
		return *x.Src
	}
	return ""
}

func (x *RemoveXAttrRequestProto) GetXAttr() *XAttrProto {
	if x != nil {
		return x.XAttr
	}
	return nil
}

type RemoveXAttrResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveXAttrResponseProto) Reset() {
	*x = RemoveXAttrResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xattr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveXAttrResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveXAttrResponseProto) ProtoMessage() {}

func (x *RemoveXAttrResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_xattr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveXAttrResponseProto.ProtoReflect.Descriptor instead.
func (*RemoveXAttrResponseProto) Descriptor() ([]byte, []int) {
	return file_xattr_proto_rawDescGZIP(), []int{8}
}

var File_xattr_proto protoreflect.FileDescriptor

var file_xattr_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x78, 0x61, 0x74, 0x74, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x0a, 0x58,
	0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x49, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4f,
	0x0a, 0x13, 0x58, 0x41, 0x74, 0x74, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x52, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59,
	0x53, 0x54, 0x45, 0x4d, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x04, 0x22,
	0x6b, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x58, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x2d, 0x0a, 0x05, 0x78, 0x41, 0x74,
	0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x05, 0x78, 0x41, 0x74, 0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x22, 0x17, 0x0a, 0x15,
	0x53, 0x65, 0x74, 0x58, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x58, 0x41, 0x74, 0x74,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63,
	0x12, 0x2f, 0x0a, 0x06, 0x78, 0x41, 0x74, 0x74, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x58,
	0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x78, 0x41, 0x74, 0x74, 0x72,
	0x73, 0x22, 0x49, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x58, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x06, 0x78,
	0x41, 0x74, 0x74, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x78, 0x41, 0x74, 0x74, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x58, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x4a, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x58, 0x41, 0x74, 0x74, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x06, 0x78, 0x41, 0x74, 0x74, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66,
	0x73, 0x2e, 0x58, 0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x78, 0x41,
	0x74, 0x74, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x58, 0x41,
	0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72,
	0x63, 0x12, 0x2d, 0x0a, 0x05, 0x78, 0x41, 0x74, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x58,
	0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x78, 0x41, 0x74, 0x74, 0x72,
	0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x58, 0x41, 0x74, 0x74, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x38, 0x0a, 0x11,
	0x58, 0x41, 0x74, 0x74, 0x72, 0x53, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x0a, 0x0c, 0x58, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x58, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x52, 0x45, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x10, 0x02, 0x42, 0x67, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x66, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0x0b, 0x58, 0x41, 0x74, 0x74, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35,
	0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x68, 0x64, 0x66, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xa0, 0x01, 0x01,
}

var (
	file_xattr_proto_rawDescOnce sync.Once
	file_xattr_proto_rawDescData = file_xattr_proto_rawDesc
)

func file_xattr_proto_rawDescGZIP() []byte {
	file_xattr_proto_rawDescOnce.Do(func() {
		file_xattr_proto_rawDescData = protoimpl.X.CompressGZIP(file_xattr_proto_rawDescData)
	})
	return file_xattr_proto_rawDescData
}

var file_xattr_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_xattr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_xattr_proto_goTypes = []interface{}{
	(XAttrSetFlagProto)(0),              // 0: hadoop.hdfs.XAttrSetFlagProto
	(XAttrProto_XAttrNamespaceProto)(0), // 1: hadoop.hdfs.XAttrProto.XAttrNamespaceProto
	(*XAttrProto)(nil),                  // 2: hadoop.hdfs.XAttrProto
	(*SetXAttrRequestProto)(nil),        // 3: hadoop.hdfs.SetXAttrRequestProto
	(*SetXAttrResponseProto)(nil),       // 4: hadoop.hdfs.SetXAttrResponseProto
	(*GetXAttrsRequestProto)(nil),       // 5: hadoop.hdfs.GetXAttrsRequestProto
	(*GetXAttrsResponseProto)(nil),      // 6: hadoop.hdfs.GetXAttrsResponseProto
	(*ListXAttrsRequestProto)(nil),      // 7: hadoop.hdfs.ListXAttrsRequestProto
	(*ListXAttrsResponseProto)(nil),     // 8: hadoop.hdfs.ListXAttrsResponseProto
	(*RemoveXAttrRequestProto)(nil),     // 9: hadoop.hdfs.RemoveXAttrRequestProto
	(*RemoveXAttrResponseProto)(nil),    // 10: hadoop.hdfs.RemoveXAttrResponseProto
}
var file_xattr_proto_depIdxs = []int32{
	1, // 0: hadoop.hdfs.XAttrProto.namespace:type_name -> hadoop.hdfs.XAttrProto.XAttrNamespaceProto
	2, // 1: hadoop.hdfs.SetXAttrRequestProto.xAttr:type_name -> hadoop.hdfs.XAttrProto
	2, // 2: hadoop.hdfs.GetXAttrsRequestProto.xAttrs:type_name -> hadoop.hdfs.XAttrProto
	2, // 3: hadoop.hdfs.GetXAttrsResponseProto.xAttrs:type_name -> hadoop.hdfs.XAttrProto
	2, // 4: hadoop.hdfs.ListXAttrsResponseProto.xAttrs:type_name -> hadoop.hdfs.XAttrProto
	2, // 5: hadoop.hdfs.RemoveXAttrRequestProto.xAttr:type_name -> hadoop.hdfs.XAttrProto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_xattr_proto_init() }
func file_xattr_proto_init() {
	if File_xattr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xattr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XAttrProto); i {
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
		file_xattr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetXAttrRequestProto); i {
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
		file_xattr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetXAttrResponseProto); i {
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
		file_xattr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetXAttrsRequestProto); i {
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
		file_xattr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetXAttrsResponseProto); i {
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
		file_xattr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListXAttrsRequestProto); i {
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
		file_xattr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListXAttrsResponseProto); i {
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
		file_xattr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveXAttrRequestProto); i {
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
		file_xattr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveXAttrResponseProto); i {
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
			RawDescriptor: file_xattr_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xattr_proto_goTypes,
		DependencyIndexes: file_xattr_proto_depIdxs,
		EnumInfos:         file_xattr_proto_enumTypes,
		MessageInfos:      file_xattr_proto_msgTypes,
	}.Build()
	File_xattr_proto = out.File
	file_xattr_proto_rawDesc = nil
	file_xattr_proto_goTypes = nil
	file_xattr_proto_depIdxs = nil
}
