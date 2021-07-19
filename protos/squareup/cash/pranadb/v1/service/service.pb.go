// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: squareup/cash/pranadb/service/v1/service.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ColumnType int32

const (
	ColumnType_COLUMN_TYPE_UNSPECIFIED ColumnType = 0
	ColumnType_COLUMN_TYPE_TINY_INT    ColumnType = 1
	ColumnType_COLUMN_TYPE_INT         ColumnType = 2
	ColumnType_COLUMN_TYPE_BIG_INT     ColumnType = 3
	ColumnType_COLUMN_TYPE_DOUBLE      ColumnType = 4
	ColumnType_COLUMN_TYPE_DECIMAL     ColumnType = 5
	ColumnType_COLUMN_TYPE_VARCHAR     ColumnType = 6
	ColumnType_COLUMN_TYPE_TIMESTAMP   ColumnType = 7
)

// Enum value maps for ColumnType.
var (
	ColumnType_name = map[int32]string{
		0: "COLUMN_TYPE_UNSPECIFIED",
		1: "COLUMN_TYPE_TINY_INT",
		2: "COLUMN_TYPE_INT",
		3: "COLUMN_TYPE_BIG_INT",
		4: "COLUMN_TYPE_DOUBLE",
		5: "COLUMN_TYPE_DECIMAL",
		6: "COLUMN_TYPE_VARCHAR",
		7: "COLUMN_TYPE_TIMESTAMP",
	}
	ColumnType_value = map[string]int32{
		"COLUMN_TYPE_UNSPECIFIED": 0,
		"COLUMN_TYPE_TINY_INT":    1,
		"COLUMN_TYPE_INT":         2,
		"COLUMN_TYPE_BIG_INT":     3,
		"COLUMN_TYPE_DOUBLE":      4,
		"COLUMN_TYPE_DECIMAL":     5,
		"COLUMN_TYPE_VARCHAR":     6,
		"COLUMN_TYPE_TIMESTAMP":   7,
	}
)

func (x ColumnType) Enum() *ColumnType {
	p := new(ColumnType)
	*p = x
	return p
}

func (x ColumnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnType) Descriptor() protoreflect.EnumDescriptor {
	return file_squareup_cash_pranadb_service_v1_service_proto_enumTypes[0].Descriptor()
}

func (ColumnType) Type() protoreflect.EnumType {
	return &file_squareup_cash_pranadb_service_v1_service_proto_enumTypes[0]
}

func (x ColumnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnType.Descriptor instead.
func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{0}
}

type ExecuteSQLStatementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *ExecuteSQLStatementRequest) Reset() {
	*x = ExecuteSQLStatementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteSQLStatementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteSQLStatementRequest) ProtoMessage() {}

func (x *ExecuteSQLStatementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteSQLStatementRequest.ProtoReflect.Descriptor instead.
func (*ExecuteSQLStatementRequest) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteSQLStatementRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{1}
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type             ColumnType `protobuf:"varint,2,opt,name=type,proto3,enum=squareup.cash.pranadb.service.v1.ColumnType" json:"type,omitempty"`
	DecimalPrecision *uint32    `protobuf:"varint,3,opt,name=decimal_precision,json=decimalPrecision,proto3,oneof" json:"decimal_precision,omitempty"`
	DecimalScale     *uint32    `protobuf:"varint,4,opt,name=decimal_scale,json=decimalScale,proto3,oneof" json:"decimal_scale,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetType() ColumnType {
	if x != nil {
		return x.Type
	}
	return ColumnType_COLUMN_TYPE_UNSPECIFIED
}

func (x *Column) GetDecimalPrecision() uint32 {
	if x != nil && x.DecimalPrecision != nil {
		return *x.DecimalPrecision
	}
	return 0
}

func (x *Column) GetDecimalScale() uint32 {
	if x != nil && x.DecimalScale != nil {
		return *x.DecimalScale
	}
	return 0
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{3}
}

type ExecuteSQLStatementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row []*Row `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
}

func (x *ExecuteSQLStatementResponse) Reset() {
	*x = ExecuteSQLStatementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteSQLStatementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteSQLStatementResponse) ProtoMessage() {}

func (x *ExecuteSQLStatementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteSQLStatementResponse.ProtoReflect.Descriptor instead.
func (*ExecuteSQLStatementResponse) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteSQLStatementResponse) GetRow() []*Row {
	if x != nil {
		return x.Row
	}
	return nil
}

var File_squareup_cash_pranadb_service_v1_service_proto protoreflect.FileDescriptor

var file_squareup_cash_pranadb_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e,
	0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x32, 0x0a, 0x1a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x51, 0x4c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0xe2, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x40, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e,
	0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72,
	0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x30, 0x0a, 0x11, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x65,
	0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x10,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0c, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x22, 0x05, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x22, 0x56, 0x0a, 0x1b,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x51, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x72,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x52,
	0x03, 0x72, 0x6f, 0x77, 0x2a, 0xd6, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x49, 0x4e, 0x59, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f,
	0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42,
	0x49, 0x47, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4c, 0x55,
	0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x45, 0x43, 0x49, 0x4d, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4c,
	0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x41, 0x52, 0x43, 0x48, 0x41, 0x52,
	0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x07, 0x32, 0xa7, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x61, 0x6e, 0x61, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x94, 0x01, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x51, 0x4c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x51, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75,
	0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x53, 0x51, 0x4c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x70,
	0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x61, 0x6e,
	0x61, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_squareup_cash_pranadb_service_v1_service_proto_rawDescOnce sync.Once
	file_squareup_cash_pranadb_service_v1_service_proto_rawDescData = file_squareup_cash_pranadb_service_v1_service_proto_rawDesc
)

func file_squareup_cash_pranadb_service_v1_service_proto_rawDescGZIP() []byte {
	file_squareup_cash_pranadb_service_v1_service_proto_rawDescOnce.Do(func() {
		file_squareup_cash_pranadb_service_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_squareup_cash_pranadb_service_v1_service_proto_rawDescData)
	})
	return file_squareup_cash_pranadb_service_v1_service_proto_rawDescData
}

var file_squareup_cash_pranadb_service_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_squareup_cash_pranadb_service_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_squareup_cash_pranadb_service_v1_service_proto_goTypes = []interface{}{
	(ColumnType)(0),                     // 0: squareup.cash.pranadb.service.v1.ColumnType
	(*ExecuteSQLStatementRequest)(nil),  // 1: squareup.cash.pranadb.service.v1.ExecuteSQLStatementRequest
	(*Result)(nil),                      // 2: squareup.cash.pranadb.service.v1.Result
	(*Column)(nil),                      // 3: squareup.cash.pranadb.service.v1.Column
	(*Row)(nil),                         // 4: squareup.cash.pranadb.service.v1.Row
	(*ExecuteSQLStatementResponse)(nil), // 5: squareup.cash.pranadb.service.v1.ExecuteSQLStatementResponse
}
var file_squareup_cash_pranadb_service_v1_service_proto_depIdxs = []int32{
	0, // 0: squareup.cash.pranadb.service.v1.Column.type:type_name -> squareup.cash.pranadb.service.v1.ColumnType
	4, // 1: squareup.cash.pranadb.service.v1.ExecuteSQLStatementResponse.row:type_name -> squareup.cash.pranadb.service.v1.Row
	1, // 2: squareup.cash.pranadb.service.v1.PranaDBService.ExecuteSQLStatement:input_type -> squareup.cash.pranadb.service.v1.ExecuteSQLStatementRequest
	5, // 3: squareup.cash.pranadb.service.v1.PranaDBService.ExecuteSQLStatement:output_type -> squareup.cash.pranadb.service.v1.ExecuteSQLStatementResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_squareup_cash_pranadb_service_v1_service_proto_init() }
func file_squareup_cash_pranadb_service_v1_service_proto_init() {
	if File_squareup_cash_pranadb_service_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteSQLStatementRequest); i {
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
		file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteSQLStatementResponse); i {
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
	file_squareup_cash_pranadb_service_v1_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_squareup_cash_pranadb_service_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_squareup_cash_pranadb_service_v1_service_proto_goTypes,
		DependencyIndexes: file_squareup_cash_pranadb_service_v1_service_proto_depIdxs,
		EnumInfos:         file_squareup_cash_pranadb_service_v1_service_proto_enumTypes,
		MessageInfos:      file_squareup_cash_pranadb_service_v1_service_proto_msgTypes,
	}.Build()
	File_squareup_cash_pranadb_service_v1_service_proto = out.File
	file_squareup_cash_pranadb_service_v1_service_proto_rawDesc = nil
	file_squareup_cash_pranadb_service_v1_service_proto_goTypes = nil
	file_squareup_cash_pranadb_service_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PranaDBServiceClient is the client API for PranaDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PranaDBServiceClient interface {
	ExecuteSQLStatement(ctx context.Context, in *ExecuteSQLStatementRequest, opts ...grpc.CallOption) (PranaDBService_ExecuteSQLStatementClient, error)
}

type pranaDBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPranaDBServiceClient(cc grpc.ClientConnInterface) PranaDBServiceClient {
	return &pranaDBServiceClient{cc}
}

func (c *pranaDBServiceClient) ExecuteSQLStatement(ctx context.Context, in *ExecuteSQLStatementRequest, opts ...grpc.CallOption) (PranaDBService_ExecuteSQLStatementClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PranaDBService_serviceDesc.Streams[0], "/squareup.cash.pranadb.service.v1.PranaDBService/ExecuteSQLStatement", opts...)
	if err != nil {
		return nil, err
	}
	x := &pranaDBServiceExecuteSQLStatementClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PranaDBService_ExecuteSQLStatementClient interface {
	Recv() (*ExecuteSQLStatementResponse, error)
	grpc.ClientStream
}

type pranaDBServiceExecuteSQLStatementClient struct {
	grpc.ClientStream
}

func (x *pranaDBServiceExecuteSQLStatementClient) Recv() (*ExecuteSQLStatementResponse, error) {
	m := new(ExecuteSQLStatementResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PranaDBServiceServer is the server API for PranaDBService service.
type PranaDBServiceServer interface {
	ExecuteSQLStatement(*ExecuteSQLStatementRequest, PranaDBService_ExecuteSQLStatementServer) error
}

// UnimplementedPranaDBServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPranaDBServiceServer struct {
}

func (*UnimplementedPranaDBServiceServer) ExecuteSQLStatement(*ExecuteSQLStatementRequest, PranaDBService_ExecuteSQLStatementServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteSQLStatement not implemented")
}

func RegisterPranaDBServiceServer(s *grpc.Server, srv PranaDBServiceServer) {
	s.RegisterService(&_PranaDBService_serviceDesc, srv)
}

func _PranaDBService_ExecuteSQLStatement_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteSQLStatementRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PranaDBServiceServer).ExecuteSQLStatement(m, &pranaDBServiceExecuteSQLStatementServer{stream})
}

type PranaDBService_ExecuteSQLStatementServer interface {
	Send(*ExecuteSQLStatementResponse) error
	grpc.ServerStream
}

type pranaDBServiceExecuteSQLStatementServer struct {
	grpc.ServerStream
}

func (x *pranaDBServiceExecuteSQLStatementServer) Send(m *ExecuteSQLStatementResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _PranaDBService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squareup.cash.pranadb.service.v1.PranaDBService",
	HandlerType: (*PranaDBServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteSQLStatement",
			Handler:       _PranaDBService_ExecuteSQLStatement_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "squareup/cash/pranadb/service/v1/service.proto",
}
