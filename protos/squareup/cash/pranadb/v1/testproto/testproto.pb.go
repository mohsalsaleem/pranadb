// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: squareup/cash/pranadb/testproto/v1/testproto.proto

package testproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoubleField         float64           `protobuf:"fixed64,1,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	FloatField          float32           `protobuf:"fixed32,2,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	Int32Field          int32             `protobuf:"varint,3,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field          int64             `protobuf:"varint,4,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	Uint32Field         uint32            `protobuf:"varint,5,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field         uint64            `protobuf:"varint,6,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	BoolField           bool              `protobuf:"varint,7,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	StringField         string            `protobuf:"bytes,8,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	BytesField          []byte            `protobuf:"bytes,9,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	NestedField         *TestTypes_Nested `protobuf:"bytes,10,opt,name=nested_field,json=nestedField,proto3" json:"nested_field,omitempty"`
	RepeatedStringField []string          `protobuf:"bytes,11,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
	RecursiveField      *Recursive        `protobuf:"bytes,12,opt,name=recursive_field,json=recursiveField,proto3" json:"recursive_field,omitempty"`
	// Types that are assignable to OneofField:
	//	*TestTypes_OneString
	//	*TestTypes_OneInt64
	OneofField      isTestTypes_OneofField  `protobuf_oneof:"oneof_field"`
	StringMapField  map[string]string       `protobuf:"bytes,15,rep,name=string_map_field,json=stringMapField,proto3" json:"string_map_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IntMapField     map[int32]string        `protobuf:"bytes,16,rep,name=int_map_field,json=intMapField,proto3" json:"int_map_field,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MapMessageField map[string]*SimpleValue `protobuf:"bytes,17,rep,name=map_message_field,json=mapMessageField,proto3" json:"map_message_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestTypes) Reset() {
	*x = TestTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTypes) ProtoMessage() {}

func (x *TestTypes) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTypes.ProtoReflect.Descriptor instead.
func (*TestTypes) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescGZIP(), []int{0}
}

func (x *TestTypes) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *TestTypes) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *TestTypes) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *TestTypes) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *TestTypes) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *TestTypes) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *TestTypes) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *TestTypes) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *TestTypes) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *TestTypes) GetNestedField() *TestTypes_Nested {
	if x != nil {
		return x.NestedField
	}
	return nil
}

func (x *TestTypes) GetRepeatedStringField() []string {
	if x != nil {
		return x.RepeatedStringField
	}
	return nil
}

func (x *TestTypes) GetRecursiveField() *Recursive {
	if x != nil {
		return x.RecursiveField
	}
	return nil
}

func (m *TestTypes) GetOneofField() isTestTypes_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (x *TestTypes) GetOneString() string {
	if x, ok := x.GetOneofField().(*TestTypes_OneString); ok {
		return x.OneString
	}
	return ""
}

func (x *TestTypes) GetOneInt64() int64 {
	if x, ok := x.GetOneofField().(*TestTypes_OneInt64); ok {
		return x.OneInt64
	}
	return 0
}

func (x *TestTypes) GetStringMapField() map[string]string {
	if x != nil {
		return x.StringMapField
	}
	return nil
}

func (x *TestTypes) GetIntMapField() map[int32]string {
	if x != nil {
		return x.IntMapField
	}
	return nil
}

func (x *TestTypes) GetMapMessageField() map[string]*SimpleValue {
	if x != nil {
		return x.MapMessageField
	}
	return nil
}

type isTestTypes_OneofField interface {
	isTestTypes_OneofField()
}

type TestTypes_OneString struct {
	OneString string `protobuf:"bytes,13,opt,name=one_string,json=oneString,proto3,oneof"`
}

type TestTypes_OneInt64 struct {
	OneInt64 int64 `protobuf:"varint,14,opt,name=one_int64,json=oneInt64,proto3,oneof"`
}

func (*TestTypes_OneString) isTestTypes_OneofField() {}

func (*TestTypes_OneInt64) isTestTypes_OneofField() {}

type Recursive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField    string     `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	RecursiveField *Recursive `protobuf:"bytes,2,opt,name=recursive_field,json=recursiveField,proto3" json:"recursive_field,omitempty"`
}

func (x *Recursive) Reset() {
	*x = Recursive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recursive) ProtoMessage() {}

func (x *Recursive) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recursive.ProtoReflect.Descriptor instead.
func (*Recursive) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescGZIP(), []int{1}
}

func (x *Recursive) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *Recursive) GetRecursiveField() *Recursive {
	if x != nil {
		return x.RecursiveField
	}
	return nil
}

type SimpleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleValue) Reset() {
	*x = SimpleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleValue) ProtoMessage() {}

func (x *SimpleValue) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleValue.ProtoReflect.Descriptor instead.
func (*SimpleValue) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TestTypes_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedString         string            `protobuf:"bytes,1,opt,name=nested_string,json=nestedString,proto3" json:"nested_string,omitempty"`
	NestedRepeatedString []string          `protobuf:"bytes,2,rep,name=nested_repeated_string,json=nestedRepeatedString,proto3" json:"nested_repeated_string,omitempty"`
	NestedMap            map[string]string `protobuf:"bytes,3,rep,name=nested_map,json=nestedMap,proto3" json:"nested_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestTypes_Nested) Reset() {
	*x = TestTypes_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestTypes_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestTypes_Nested) ProtoMessage() {}

func (x *TestTypes_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestTypes_Nested.ProtoReflect.Descriptor instead.
func (*TestTypes_Nested) Descriptor() ([]byte, []int) {
	return file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TestTypes_Nested) GetNestedString() string {
	if x != nil {
		return x.NestedString
	}
	return ""
}

func (x *TestTypes_Nested) GetNestedRepeatedString() []string {
	if x != nil {
		return x.NestedRepeatedString
	}
	return nil
}

func (x *TestTypes_Nested) GetNestedMap() map[string]string {
	if x != nil {
		return x.NestedMap
	}
	return nil
}

var File_squareup_cash_pranadb_testproto_v1_testproto_proto protoreflect.FileDescriptor

var file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f,
	0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x0b, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x57, 0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x0b,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x56, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6f,
	0x6e, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x09, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x6b, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x41, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x62, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e,
	0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x69, 0x6e, 0x74,
	0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x6e, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x85, 0x02, 0x0a, 0x06, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x62,
	0x0a, 0x0a, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d,
	0x61, 0x70, 0x1a, 0x3c, 0x0a, 0x0e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x41, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x73, 0x0a, 0x14, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x61,
	0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x73, 0x69, 0x76, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x56, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x75,
	0x72, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2e, 0x63, 0x61, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64, 0x62, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x52, 0x0e, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0x23, 0x0a, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x75, 0x70, 0x2f, 0x70, 0x72, 0x61,
	0x6e, 0x61, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x75, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x61, 0x6e, 0x61, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescOnce sync.Once
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescData = file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDesc
)

func file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescGZIP() []byte {
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescOnce.Do(func() {
		file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescData)
	})
	return file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDescData
}

var file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_squareup_cash_pranadb_testproto_v1_testproto_proto_goTypes = []interface{}{
	(*TestTypes)(nil),        // 0: squareup.cash.pranadb.testproto.v1.TestTypes
	(*Recursive)(nil),        // 1: squareup.cash.pranadb.testproto.v1.Recursive
	(*SimpleValue)(nil),      // 2: squareup.cash.pranadb.testproto.v1.SimpleValue
	(*TestTypes_Nested)(nil), // 3: squareup.cash.pranadb.testproto.v1.TestTypes.Nested
	nil,                      // 4: squareup.cash.pranadb.testproto.v1.TestTypes.StringMapFieldEntry
	nil,                      // 5: squareup.cash.pranadb.testproto.v1.TestTypes.IntMapFieldEntry
	nil,                      // 6: squareup.cash.pranadb.testproto.v1.TestTypes.MapMessageFieldEntry
	nil,                      // 7: squareup.cash.pranadb.testproto.v1.TestTypes.Nested.NestedMapEntry
}
var file_squareup_cash_pranadb_testproto_v1_testproto_proto_depIdxs = []int32{
	3, // 0: squareup.cash.pranadb.testproto.v1.TestTypes.nested_field:type_name -> squareup.cash.pranadb.testproto.v1.TestTypes.Nested
	1, // 1: squareup.cash.pranadb.testproto.v1.TestTypes.recursive_field:type_name -> squareup.cash.pranadb.testproto.v1.Recursive
	4, // 2: squareup.cash.pranadb.testproto.v1.TestTypes.string_map_field:type_name -> squareup.cash.pranadb.testproto.v1.TestTypes.StringMapFieldEntry
	5, // 3: squareup.cash.pranadb.testproto.v1.TestTypes.int_map_field:type_name -> squareup.cash.pranadb.testproto.v1.TestTypes.IntMapFieldEntry
	6, // 4: squareup.cash.pranadb.testproto.v1.TestTypes.map_message_field:type_name -> squareup.cash.pranadb.testproto.v1.TestTypes.MapMessageFieldEntry
	1, // 5: squareup.cash.pranadb.testproto.v1.Recursive.recursive_field:type_name -> squareup.cash.pranadb.testproto.v1.Recursive
	7, // 6: squareup.cash.pranadb.testproto.v1.TestTypes.Nested.nested_map:type_name -> squareup.cash.pranadb.testproto.v1.TestTypes.Nested.NestedMapEntry
	2, // 7: squareup.cash.pranadb.testproto.v1.TestTypes.MapMessageFieldEntry.value:type_name -> squareup.cash.pranadb.testproto.v1.SimpleValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_squareup_cash_pranadb_testproto_v1_testproto_proto_init() }
func file_squareup_cash_pranadb_testproto_v1_testproto_proto_init() {
	if File_squareup_cash_pranadb_testproto_v1_testproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestTypes); i {
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
		file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recursive); i {
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
		file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleValue); i {
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
		file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestTypes_Nested); i {
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
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TestTypes_OneString)(nil),
		(*TestTypes_OneInt64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_squareup_cash_pranadb_testproto_v1_testproto_proto_goTypes,
		DependencyIndexes: file_squareup_cash_pranadb_testproto_v1_testproto_proto_depIdxs,
		MessageInfos:      file_squareup_cash_pranadb_testproto_v1_testproto_proto_msgTypes,
	}.Build()
	File_squareup_cash_pranadb_testproto_v1_testproto_proto = out.File
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_rawDesc = nil
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_goTypes = nil
	file_squareup_cash_pranadb_testproto_v1_testproto_proto_depIdxs = nil
}
