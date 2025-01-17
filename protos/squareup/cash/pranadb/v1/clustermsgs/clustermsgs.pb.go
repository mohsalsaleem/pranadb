// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: squareup/cash/pranadb/clustermsgs/v1/clustermsgs.proto

package clustermsgs

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

type DDLStatementInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginatingNodeId int64    `protobuf:"varint,1,opt,name=originating_node_id,json=originatingNodeId,proto3" json:"originating_node_id,omitempty"`
	CommandId         int64    `protobuf:"varint,2,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Phase             int32    `protobuf:"varint,3,opt,name=phase,proto3" json:"phase,omitempty"`
	CommandType       int32    `protobuf:"varint,4,opt,name=command_type,json=commandType,proto3" json:"command_type,omitempty"`
	SchemaName        string   `protobuf:"bytes,5,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Sql               string   `protobuf:"bytes,6,opt,name=sql,proto3" json:"sql,omitempty"`
	TableSequences    []uint64 `protobuf:"varint,7,rep,packed,name=table_sequences,json=tableSequences,proto3" json:"table_sequences,omitempty"`
	ExtraData         []byte   `protobuf:"bytes,8,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (x *DDLStatementInfo) Reset() {
	*x = DDLStatementInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDLStatementInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDLStatementInfo) ProtoMessage() {}

func (x *DDLStatementInfo) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDLStatementInfo.ProtoReflect.Descriptor instead.
func (*DDLStatementInfo) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{0}
}

func (x *DDLStatementInfo) GetOriginatingNodeId() int64 {
	if x != nil {
		return x.OriginatingNodeId
	}
	return 0
}

func (x *DDLStatementInfo) GetCommandId() int64 {
	if x != nil {
		return x.CommandId
	}
	return 0
}

func (x *DDLStatementInfo) GetPhase() int32 {
	if x != nil {
		return x.Phase
	}
	return 0
}

func (x *DDLStatementInfo) GetCommandType() int32 {
	if x != nil {
		return x.CommandType
	}
	return 0
}

func (x *DDLStatementInfo) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *DDLStatementInfo) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

func (x *DDLStatementInfo) GetTableSequences() []uint64 {
	if x != nil {
		return x.TableSequences
	}
	return nil
}

func (x *DDLStatementInfo) GetExtraData() []byte {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type DDLCancelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
}

func (x *DDLCancelMessage) Reset() {
	*x = DDLCancelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DDLCancelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DDLCancelMessage) ProtoMessage() {}

func (x *DDLCancelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DDLCancelMessage.ProtoReflect.Descriptor instead.
func (*DDLCancelMessage) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{1}
}

func (x *DDLCancelMessage) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

type ReloadProtobuf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReloadProtobuf) Reset() {
	*x = ReloadProtobuf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadProtobuf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadProtobuf) ProtoMessage() {}

func (x *ReloadProtobuf) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadProtobuf.ProtoReflect.Descriptor instead.
func (*ReloadProtobuf) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{2}
}

type ClusterProposeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId     int64  `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	RequestBody []byte `protobuf:"bytes,2,opt,name=request_body,json=requestBody,proto3" json:"request_body,omitempty"`
}

func (x *ClusterProposeRequest) Reset() {
	*x = ClusterProposeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterProposeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterProposeRequest) ProtoMessage() {}

func (x *ClusterProposeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterProposeRequest.ProtoReflect.Descriptor instead.
func (*ClusterProposeRequest) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterProposeRequest) GetShardId() int64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *ClusterProposeRequest) GetRequestBody() []byte {
	if x != nil {
		return x.RequestBody
	}
	return nil
}

type ClusterProposeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetVal       int64  `protobuf:"varint,1,opt,name=ret_val,json=retVal,proto3" json:"ret_val,omitempty"`
	ResponseBody []byte `protobuf:"bytes,2,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
}

func (x *ClusterProposeResponse) Reset() {
	*x = ClusterProposeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterProposeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterProposeResponse) ProtoMessage() {}

func (x *ClusterProposeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterProposeResponse.ProtoReflect.Descriptor instead.
func (*ClusterProposeResponse) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{4}
}

func (x *ClusterProposeResponse) GetRetVal() int64 {
	if x != nil {
		return x.RetVal
	}
	return 0
}

func (x *ClusterProposeResponse) GetResponseBody() []byte {
	if x != nil {
		return x.ResponseBody
	}
	return nil
}

type ClusterForwardWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId     int64  `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	RequestBody []byte `protobuf:"bytes,2,opt,name=request_body,json=requestBody,proto3" json:"request_body,omitempty"`
}

func (x *ClusterForwardWriteRequest) Reset() {
	*x = ClusterForwardWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterForwardWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterForwardWriteRequest) ProtoMessage() {}

func (x *ClusterForwardWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterForwardWriteRequest.ProtoReflect.Descriptor instead.
func (*ClusterForwardWriteRequest) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterForwardWriteRequest) GetShardId() int64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *ClusterForwardWriteRequest) GetRequestBody() []byte {
	if x != nil {
		return x.RequestBody
	}
	return nil
}

type ClusterForwardWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterForwardWriteResponse) Reset() {
	*x = ClusterForwardWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterForwardWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterForwardWriteResponse) ProtoMessage() {}

func (x *ClusterForwardWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterForwardWriteResponse.ProtoReflect.Descriptor instead.
func (*ClusterForwardWriteResponse) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{6}
}

type ClusterReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId     int64  `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	RequestBody []byte `protobuf:"bytes,2,opt,name=request_body,json=requestBody,proto3" json:"request_body,omitempty"`
}

func (x *ClusterReadRequest) Reset() {
	*x = ClusterReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReadRequest) ProtoMessage() {}

func (x *ClusterReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReadRequest.ProtoReflect.Descriptor instead.
func (*ClusterReadRequest) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{7}
}

func (x *ClusterReadRequest) GetShardId() int64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *ClusterReadRequest) GetRequestBody() []byte {
	if x != nil {
		return x.RequestBody
	}
	return nil
}

type ClusterReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseBody []byte `protobuf:"bytes,1,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
}

func (x *ClusterReadResponse) Reset() {
	*x = ClusterReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReadResponse) ProtoMessage() {}

func (x *ClusterReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReadResponse.ProtoReflect.Descriptor instead.
func (*ClusterReadResponse) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{8}
}

func (x *ClusterReadResponse) GetResponseBody() []byte {
	if x != nil {
		return x.ResponseBody
	}
	return nil
}

type SourceSetMaxIngestRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaName string `protobuf:"bytes,1,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	SourceName string `protobuf:"bytes,2,opt,name=source_name,json=sourceName,proto3" json:"source_name,omitempty"`
	Rate       int64  `protobuf:"varint,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *SourceSetMaxIngestRate) Reset() {
	*x = SourceSetMaxIngestRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceSetMaxIngestRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceSetMaxIngestRate) ProtoMessage() {}

func (x *SourceSetMaxIngestRate) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceSetMaxIngestRate.ProtoReflect.Descriptor instead.
func (*SourceSetMaxIngestRate) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{9}
}

func (x *SourceSetMaxIngestRate) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *SourceSetMaxIngestRate) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *SourceSetMaxIngestRate) GetRate() int64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type LeaderInfosMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderInfos []*LeaderInfo `protobuf:"bytes,1,rep,name=leader_infos,json=leaderInfos,proto3" json:"leader_infos,omitempty"`
}

func (x *LeaderInfosMessage) Reset() {
	*x = LeaderInfosMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderInfosMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderInfosMessage) ProtoMessage() {}

func (x *LeaderInfosMessage) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderInfosMessage.ProtoReflect.Descriptor instead.
func (*LeaderInfosMessage) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{10}
}

func (x *LeaderInfosMessage) GetLeaderInfos() []*LeaderInfo {
	if x != nil {
		return x.LeaderInfos
	}
	return nil
}

type LeaderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardId int64 `protobuf:"varint,1,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	NodeId  int64 `protobuf:"varint,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *LeaderInfo) Reset() {
	*x = LeaderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderInfo) ProtoMessage() {}

func (x *LeaderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderInfo.ProtoReflect.Descriptor instead.
func (*LeaderInfo) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{11}
}

func (x *LeaderInfo) GetShardId() int64 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *LeaderInfo) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

type RemotingTestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SomeField string `protobuf:"bytes,1,opt,name=some_field,json=someField,proto3" json:"some_field,omitempty"`
}

func (x *RemotingTestMessage) Reset() {
	*x = RemotingTestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemotingTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemotingTestMessage) ProtoMessage() {}

func (x *RemotingTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemotingTestMessage.ProtoReflect.Descriptor instead.
func (*RemotingTestMessage) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP(), []int{12}
}

func (x *RemotingTestMessage) GetSomeField() string {
	if x != nil {
		return x.SomeField
	}
	return ""
}

var File_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto protoreflect.FileDescriptor

var file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6d,
	0x73, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x73,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x95,
	0x02, 0x0a, 0x10, 0x44, 0x44, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x71, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12, 0x27,
	0x0a, 0x0f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x10, 0x44, 0x44, 0x4c, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x52,
	0x65, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x55, 0x0a,
	0x15, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x56, 0x0a, 0x16, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x5a, 0x0a, 0x1a,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x3a, 0x0a, 0x13, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x6e, 0x0a, 0x16, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x22, 0x69, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x53, 0x0a,
	0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x40, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x6f, 0x6d, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75,
	0x70, 0x2f, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f, 0x70,
	0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x6d, 0x73, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescOnce sync.Once
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescData = file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDesc
)

func file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescGZIP() []byte {
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescOnce.Do(func() {
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescData)
	})
	return file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDescData
}

var file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_goTypes = []interface{}{
	(*DDLStatementInfo)(nil),            // 0: squareup.cash.pranadb.clustermsgs.v1.DDLStatementInfo
	(*DDLCancelMessage)(nil),            // 1: squareup.cash.pranadb.clustermsgs.v1.DDLCancelMessage
	(*ReloadProtobuf)(nil),              // 2: squareup.cash.pranadb.clustermsgs.v1.ReloadProtobuf
	(*ClusterProposeRequest)(nil),       // 3: squareup.cash.pranadb.clustermsgs.v1.ClusterProposeRequest
	(*ClusterProposeResponse)(nil),      // 4: squareup.cash.pranadb.clustermsgs.v1.ClusterProposeResponse
	(*ClusterForwardWriteRequest)(nil),  // 5: squareup.cash.pranadb.clustermsgs.v1.ClusterForwardWriteRequest
	(*ClusterForwardWriteResponse)(nil), // 6: squareup.cash.pranadb.clustermsgs.v1.ClusterForwardWriteResponse
	(*ClusterReadRequest)(nil),          // 7: squareup.cash.pranadb.clustermsgs.v1.ClusterReadRequest
	(*ClusterReadResponse)(nil),         // 8: squareup.cash.pranadb.clustermsgs.v1.ClusterReadResponse
	(*SourceSetMaxIngestRate)(nil),      // 9: squareup.cash.pranadb.clustermsgs.v1.SourceSetMaxIngestRate
	(*LeaderInfosMessage)(nil),          // 10: squareup.cash.pranadb.clustermsgs.v1.LeaderInfosMessage
	(*LeaderInfo)(nil),                  // 11: squareup.cash.pranadb.clustermsgs.v1.LeaderInfo
	(*RemotingTestMessage)(nil),         // 12: squareup.cash.pranadb.clustermsgs.v1.RemotingTestMessage
}
var file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_depIdxs = []int32{
	11, // 0: squareup.cash.pranadb.clustermsgs.v1.LeaderInfosMessage.leader_infos:type_name -> squareup.cash.pranadb.clustermsgs.v1.LeaderInfo
	1,  // [1:1] is the sub-list for method output_type
	1,  // [1:1] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_init() }
func file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_init() {
	if File_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDLStatementInfo); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DDLCancelMessage); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadProtobuf); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterProposeRequest); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterProposeResponse); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterForwardWriteRequest); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterForwardWriteResponse); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterReadRequest); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterReadResponse); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceSetMaxIngestRate); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderInfosMessage); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderInfo); i {
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
		file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemotingTestMessage); i {
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
			RawDescriptor: file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_goTypes,
		DependencyIndexes: file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_depIdxs,
		MessageInfos:      file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_msgTypes,
	}.Build()
	File_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto = out.File
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_rawDesc = nil
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_goTypes = nil
	file_squareup_cash_pranadb_clustermsgs_v1_clustermsgs_proto_depIdxs = nil
}
