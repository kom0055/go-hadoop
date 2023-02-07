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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/server/common/yarn_server_common_protos.proto

package common

import (
	api "github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/api"
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

type NodeActionProto int32

const (
	NodeActionProto_NORMAL   NodeActionProto = 0
	NodeActionProto_RESYNC   NodeActionProto = 1
	NodeActionProto_SHUTDOWN NodeActionProto = 2
)

// Enum value maps for NodeActionProto.
var (
	NodeActionProto_name = map[int32]string{
		0: "NORMAL",
		1: "RESYNC",
		2: "SHUTDOWN",
	}
	NodeActionProto_value = map[string]int32{
		"NORMAL":   0,
		"RESYNC":   1,
		"SHUTDOWN": 2,
	}
)

func (x NodeActionProto) Enum() *NodeActionProto {
	p := new(NodeActionProto)
	*p = x
	return p
}

func (x NodeActionProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeActionProto) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_enumTypes[0].Descriptor()
}

func (NodeActionProto) Type() protoreflect.EnumType {
	return &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_enumTypes[0]
}

func (x NodeActionProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NodeActionProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NodeActionProto(num)
	return nil
}

// Deprecated: Use NodeActionProto.Descriptor instead.
func (NodeActionProto) EnumDescriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{0}
}

type NodeStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId                        *api.NodeIdProto                    `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	ResponseId                    *int32                              `protobuf:"varint,2,opt,name=response_id,json=responseId" json:"response_id,omitempty"`
	ContainersStatuses            []*api.ContainerStatusProto         `protobuf:"bytes,3,rep,name=containersStatuses" json:"containersStatuses,omitempty"`
	NodeHealthStatus              *NodeHealthStatusProto              `protobuf:"bytes,4,opt,name=nodeHealthStatus" json:"nodeHealthStatus,omitempty"`
	KeepAliveApplications         []*api.ApplicationIdProto           `protobuf:"bytes,5,rep,name=keep_alive_applications,json=keepAliveApplications" json:"keep_alive_applications,omitempty"`
	ContainersUtilization         *api.ResourceUtilizationProto       `protobuf:"bytes,6,opt,name=containers_utilization,json=containersUtilization" json:"containers_utilization,omitempty"`
	NodeUtilization               *api.ResourceUtilizationProto       `protobuf:"bytes,7,opt,name=node_utilization,json=nodeUtilization" json:"node_utilization,omitempty"`
	IncreasedContainers           []*api.ContainerProto               `protobuf:"bytes,8,rep,name=increased_containers,json=increasedContainers" json:"increased_containers,omitempty"`
	OpportunisticContainersStatus *OpportunisticContainersStatusProto `protobuf:"bytes,9,opt,name=opportunistic_containers_status,json=opportunisticContainersStatus" json:"opportunistic_containers_status,omitempty"`
}

func (x *NodeStatusProto) Reset() {
	*x = NodeStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusProto) ProtoMessage() {}

func (x *NodeStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusProto.ProtoReflect.Descriptor instead.
func (*NodeStatusProto) Descriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatusProto) GetNodeId() *api.NodeIdProto {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *NodeStatusProto) GetResponseId() int32 {
	if x != nil && x.ResponseId != nil {
		return *x.ResponseId
	}
	return 0
}

func (x *NodeStatusProto) GetContainersStatuses() []*api.ContainerStatusProto {
	if x != nil {
		return x.ContainersStatuses
	}
	return nil
}

func (x *NodeStatusProto) GetNodeHealthStatus() *NodeHealthStatusProto {
	if x != nil {
		return x.NodeHealthStatus
	}
	return nil
}

func (x *NodeStatusProto) GetKeepAliveApplications() []*api.ApplicationIdProto {
	if x != nil {
		return x.KeepAliveApplications
	}
	return nil
}

func (x *NodeStatusProto) GetContainersUtilization() *api.ResourceUtilizationProto {
	if x != nil {
		return x.ContainersUtilization
	}
	return nil
}

func (x *NodeStatusProto) GetNodeUtilization() *api.ResourceUtilizationProto {
	if x != nil {
		return x.NodeUtilization
	}
	return nil
}

func (x *NodeStatusProto) GetIncreasedContainers() []*api.ContainerProto {
	if x != nil {
		return x.IncreasedContainers
	}
	return nil
}

func (x *NodeStatusProto) GetOpportunisticContainersStatus() *OpportunisticContainersStatusProto {
	if x != nil {
		return x.OpportunisticContainersStatus
	}
	return nil
}

type OpportunisticContainersStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunningOpportContainers *int32 `protobuf:"varint,1,opt,name=running_opport_containers,json=runningOpportContainers" json:"running_opport_containers,omitempty"`
	OpportMemoryUsed        *int64 `protobuf:"varint,2,opt,name=opport_memory_used,json=opportMemoryUsed" json:"opport_memory_used,omitempty"`
	OpportCoresUsed         *int32 `protobuf:"varint,3,opt,name=opport_cores_used,json=opportCoresUsed" json:"opport_cores_used,omitempty"`
	QueuedOpportContainers  *int32 `protobuf:"varint,4,opt,name=queued_opport_containers,json=queuedOpportContainers" json:"queued_opport_containers,omitempty"`
	WaitQueueLength         *int32 `protobuf:"varint,5,opt,name=wait_queue_length,json=waitQueueLength" json:"wait_queue_length,omitempty"`
	EstimatedQueueWaitTime  *int32 `protobuf:"varint,6,opt,name=estimated_queue_wait_time,json=estimatedQueueWaitTime" json:"estimated_queue_wait_time,omitempty"`
	OpportQueueCapacity     *int32 `protobuf:"varint,7,opt,name=opport_queue_capacity,json=opportQueueCapacity" json:"opport_queue_capacity,omitempty"`
}

func (x *OpportunisticContainersStatusProto) Reset() {
	*x = OpportunisticContainersStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpportunisticContainersStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpportunisticContainersStatusProto) ProtoMessage() {}

func (x *OpportunisticContainersStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpportunisticContainersStatusProto.ProtoReflect.Descriptor instead.
func (*OpportunisticContainersStatusProto) Descriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{1}
}

func (x *OpportunisticContainersStatusProto) GetRunningOpportContainers() int32 {
	if x != nil && x.RunningOpportContainers != nil {
		return *x.RunningOpportContainers
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetOpportMemoryUsed() int64 {
	if x != nil && x.OpportMemoryUsed != nil {
		return *x.OpportMemoryUsed
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetOpportCoresUsed() int32 {
	if x != nil && x.OpportCoresUsed != nil {
		return *x.OpportCoresUsed
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetQueuedOpportContainers() int32 {
	if x != nil && x.QueuedOpportContainers != nil {
		return *x.QueuedOpportContainers
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetWaitQueueLength() int32 {
	if x != nil && x.WaitQueueLength != nil {
		return *x.WaitQueueLength
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetEstimatedQueueWaitTime() int32 {
	if x != nil && x.EstimatedQueueWaitTime != nil {
		return *x.EstimatedQueueWaitTime
	}
	return 0
}

func (x *OpportunisticContainersStatusProto) GetOpportQueueCapacity() int32 {
	if x != nil && x.OpportQueueCapacity != nil {
		return *x.OpportQueueCapacity
	}
	return 0
}

type MasterKeyProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyId *int32 `protobuf:"varint,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Bytes []byte `protobuf:"bytes,2,opt,name=bytes" json:"bytes,omitempty"`
}

func (x *MasterKeyProto) Reset() {
	*x = MasterKeyProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MasterKeyProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MasterKeyProto) ProtoMessage() {}

func (x *MasterKeyProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MasterKeyProto.ProtoReflect.Descriptor instead.
func (*MasterKeyProto) Descriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{2}
}

func (x *MasterKeyProto) GetKeyId() int32 {
	if x != nil && x.KeyId != nil {
		return *x.KeyId
	}
	return 0
}

func (x *MasterKeyProto) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type NodeHealthStatusProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsNodeHealthy        *bool   `protobuf:"varint,1,opt,name=is_node_healthy,json=isNodeHealthy" json:"is_node_healthy,omitempty"`
	HealthReport         *string `protobuf:"bytes,2,opt,name=health_report,json=healthReport" json:"health_report,omitempty"`
	LastHealthReportTime *int64  `protobuf:"varint,3,opt,name=last_health_report_time,json=lastHealthReportTime" json:"last_health_report_time,omitempty"`
}

func (x *NodeHealthStatusProto) Reset() {
	*x = NodeHealthStatusProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeHealthStatusProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeHealthStatusProto) ProtoMessage() {}

func (x *NodeHealthStatusProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeHealthStatusProto.ProtoReflect.Descriptor instead.
func (*NodeHealthStatusProto) Descriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{3}
}

func (x *NodeHealthStatusProto) GetIsNodeHealthy() bool {
	if x != nil && x.IsNodeHealthy != nil {
		return *x.IsNodeHealthy
	}
	return false
}

func (x *NodeHealthStatusProto) GetHealthReport() string {
	if x != nil && x.HealthReport != nil {
		return *x.HealthReport
	}
	return ""
}

func (x *NodeHealthStatusProto) GetLastHealthReportTime() int64 {
	if x != nil && x.LastHealthReportTime != nil {
		return *x.LastHealthReportTime
	}
	return 0
}

type VersionProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MajorVersion *int32 `protobuf:"varint,1,opt,name=major_version,json=majorVersion" json:"major_version,omitempty"`
	MinorVersion *int32 `protobuf:"varint,2,opt,name=minor_version,json=minorVersion" json:"minor_version,omitempty"`
}

func (x *VersionProto) Reset() {
	*x = VersionProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionProto) ProtoMessage() {}

func (x *VersionProto) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionProto.ProtoReflect.Descriptor instead.
func (*VersionProto) Descriptor() ([]byte, []int) {
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP(), []int{4}
}

func (x *VersionProto) GetMajorVersion() int32 {
	if x != nil && x.MajorVersion != nil {
		return *x.MajorVersion
	}
	return 0
}

func (x *VersionProto) GetMinorVersion() int32 {
	if x != nil && x.MinorVersion != nil {
		return *x.MinorVersion
	}
	return 0
}

var File_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto protoreflect.FileDescriptor

var file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d,
	0x30, 0x30, 0x35, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x79, 0x61,
	0x72, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x79, 0x61, 0x72, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d,
	0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2f, 0x67,
	0x6f, 0x2d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x79, 0x61, 0x72, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x79, 0x61, 0x72, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x08, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x54, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x64, 0x12, 0x74, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2e, 0x67,
	0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x7b, 0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d, 0x30, 0x30,
	0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x10,
	0x6e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x7a, 0x0a, 0x17, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x42, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x15, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7f, 0x0a, 0x16,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35,
	0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x73, 0x0a,
	0x10, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d, 0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f,
	0x6d, 0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x13, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x1f, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x75, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x5c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x6f, 0x6d,
	0x30, 0x30, 0x35, 0x35, 0x2e, 0x67, 0x6f, 0x5f, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x79, 0x61, 0x72, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x1d, 0x6f,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8f, 0x03, 0x0a,
	0x22, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x75, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x0a, 0x19, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6f,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6f, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x11, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x72, 0x65, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x64, 0x4f, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x77, 0x61, 0x69, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x39, 0x0a, 0x19, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x16, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x57, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6f, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x3d,
	0x0a, 0x0e, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x9b, 0x01,
	0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x69, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x37, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d,
	0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x55, 0x54, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x42, 0x40,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6d,
	0x30, 0x30, 0x35, 0x35, 0x2f, 0x67, 0x6f, 0x2d, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x79, 0x61,
	0x72, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
}

var (
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescOnce sync.Once
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescData = file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDesc
)

func file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescGZIP() []byte {
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescOnce.Do(func() {
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescData)
	})
	return file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDescData
}

var file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_goTypes = []interface{}{
	(NodeActionProto)(0),                       // 0: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeActionProto
	(*NodeStatusProto)(nil),                    // 1: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto
	(*OpportunisticContainersStatusProto)(nil), // 2: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.OpportunisticContainersStatusProto
	(*MasterKeyProto)(nil),                     // 3: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.MasterKeyProto
	(*NodeHealthStatusProto)(nil),              // 4: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeHealthStatusProto
	(*VersionProto)(nil),                       // 5: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.VersionProto
	(*api.NodeIdProto)(nil),                    // 6: github.com.kom0055.go_hadoop.v1alpha1.yarn.api.NodeIdProto
	(*api.ContainerStatusProto)(nil),           // 7: github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ContainerStatusProto
	(*api.ApplicationIdProto)(nil),             // 8: github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ApplicationIdProto
	(*api.ResourceUtilizationProto)(nil),       // 9: github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ResourceUtilizationProto
	(*api.ContainerProto)(nil),                 // 10: github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ContainerProto
}
var file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_depIdxs = []int32{
	6,  // 0: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.node_id:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.NodeIdProto
	7,  // 1: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.containersStatuses:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ContainerStatusProto
	4,  // 2: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.nodeHealthStatus:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeHealthStatusProto
	8,  // 3: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.keep_alive_applications:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ApplicationIdProto
	9,  // 4: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.containers_utilization:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ResourceUtilizationProto
	9,  // 5: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.node_utilization:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ResourceUtilizationProto
	10, // 6: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.increased_containers:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.api.ContainerProto
	2,  // 7: github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.NodeStatusProto.opportunistic_containers_status:type_name -> github.com.kom0055.go_hadoop.v1alpha1.yarn.server.common.OpportunisticContainersStatusProto
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() {
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_init()
}
func file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_init() {
	if File_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatusProto); i {
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
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpportunisticContainersStatusProto); i {
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
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MasterKeyProto); i {
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
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeHealthStatusProto); i {
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
		file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionProto); i {
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
			RawDescriptor: file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_goTypes,
		DependencyIndexes: file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_depIdxs,
		EnumInfos:         file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_enumTypes,
		MessageInfos:      file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_msgTypes,
	}.Build()
	File_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto = out.File
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_rawDesc = nil
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_goTypes = nil
	file_github_com_kom0055_go_hadoop_proto_v1alpha1_yarn_server_common_yarn_server_common_protos_proto_depIdxs = nil
}