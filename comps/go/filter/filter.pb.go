// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: filter.proto

package filter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_ASC  Direction = 0
	Direction_DESC Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Direction_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_filter_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_filter_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{0}
}

type Operator int32

const (
	Operator_Equals   Operator = 0
	Operator_Contains Operator = 1
	Operator_Starts   Operator = 2
	Operator_In       Operator = 3
	Operator_Gt       Operator = 4
	Operator_Gte      Operator = 5
	Operator_Lt       Operator = 6
	Operator_Lte      Operator = 7
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "Equals",
		1: "Contains",
		2: "Starts",
		3: "In",
		4: "Gt",
		5: "Gte",
		6: "Lt",
		7: "Lte",
	}
	Operator_value = map[string]int32{
		"Equals":   0,
		"Contains": 1,
		"Starts":   2,
		"In":       3,
		"Gt":       4,
		"Gte":      5,
		"Lt":       6,
		"Lte":      7,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_filter_proto_enumTypes[1].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_filter_proto_enumTypes[1]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{1}
}

// generic filter
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainFilter string       `protobuf:"bytes,1,opt,name=main_filter,json=mainFilter,proto3" json:"main_filter,omitempty"`
	Ids        []string     `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Conditions []*Condition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetMainFilter() string {
	if x != nil {
		return x.MainFilter
	}
	return ""
}

func (x *Filter) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *Filter) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName string     `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Operator  Operator   `protobuf:"varint,2,opt,name=operator,proto3,enum=Operator" json:"operator,omitempty"`
	Value     *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Not       bool       `protobuf:"varint,4,opt,name=not,proto3" json:"not,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{1}
}

func (x *Condition) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Condition) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_Equals
}

func (x *Condition) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Condition) GetNot() bool {
	if x != nil {
		return x.Not
	}
	return false
}

type Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName string    `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=Direction" json:"direction,omitempty"`
}

func (x *Sorting) Reset() {
	*x = Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorting) ProtoMessage() {}

func (x *Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorting.ProtoReflect.Descriptor instead.
func (*Sorting) Descriptor() ([]byte, []int) {
	return file_filter_proto_rawDescGZIP(), []int{2}
}

func (x *Sorting) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Sorting) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_ASC
}

var File_filter_proto protoreflect.FileDescriptor

var file_filter_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64,
	0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8f, 0x01,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x22,
	0x52, 0x0a, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x1e, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53,
	0x43, 0x10, 0x01, 0x2a, 0x5a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x0a, 0x0a, 0x06, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x73, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x73, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x6e, 0x10, 0x03, 0x12, 0x06, 0x0a,
	0x02, 0x47, 0x74, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x74, 0x65, 0x10, 0x05, 0x12, 0x06,
	0x0a, 0x02, 0x4c, 0x74, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x74, 0x65, 0x10, 0x07, 0x42,
	0x0f, 0x5a, 0x0d, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filter_proto_rawDescOnce sync.Once
	file_filter_proto_rawDescData = file_filter_proto_rawDesc
)

func file_filter_proto_rawDescGZIP() []byte {
	file_filter_proto_rawDescOnce.Do(func() {
		file_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_filter_proto_rawDescData)
	})
	return file_filter_proto_rawDescData
}

var file_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_filter_proto_goTypes = []interface{}{
	(Direction)(0),    // 0: direction
	(Operator)(0),     // 1: operator
	(*Filter)(nil),    // 2: Filter
	(*Condition)(nil), // 3: condition
	(*Sorting)(nil),   // 4: Sorting
	(*anypb.Any)(nil), // 5: google.protobuf.Any
}
var file_filter_proto_depIdxs = []int32{
	3, // 0: Filter.conditions:type_name -> condition
	1, // 1: condition.operator:type_name -> operator
	5, // 2: condition.value:type_name -> google.protobuf.Any
	0, // 3: Sorting.direction:type_name -> direction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_filter_proto_init() }
func file_filter_proto_init() {
	if File_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sorting); i {
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
			RawDescriptor: file_filter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filter_proto_goTypes,
		DependencyIndexes: file_filter_proto_depIdxs,
		EnumInfos:         file_filter_proto_enumTypes,
		MessageInfos:      file_filter_proto_msgTypes,
	}.Build()
	File_filter_proto = out.File
	file_filter_proto_rawDesc = nil
	file_filter_proto_goTypes = nil
	file_filter_proto_depIdxs = nil
}
