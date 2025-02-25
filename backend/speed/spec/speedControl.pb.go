// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/speedControl.proto

package spec

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

type SendSpeed_Train int32

const (
	SendSpeed_UNKNOWN  SendSpeed_Train = 0
	SendSpeed_TAKAO    SendSpeed_Train = 1
	SendSpeed_CHICHIBU SendSpeed_Train = 2
	SendSpeed_HAKONE   SendSpeed_Train = 3
	SendSpeed_OKUTAMA  SendSpeed_Train = 4
	SendSpeed_NIKKO    SendSpeed_Train = 5
	SendSpeed_ENOSHIMA SendSpeed_Train = 6
	SendSpeed_KAMAKURA SendSpeed_Train = 7
	SendSpeed_YOKOSUKA SendSpeed_Train = 8
)

// Enum value maps for SendSpeed_Train.
var (
	SendSpeed_Train_name = map[int32]string{
		0: "UNKNOWN",
		1: "TAKAO",
		2: "CHICHIBU",
		3: "HAKONE",
		4: "OKUTAMA",
		5: "NIKKO",
		6: "ENOSHIMA",
		7: "KAMAKURA",
		8: "YOKOSUKA",
	}
	SendSpeed_Train_value = map[string]int32{
		"UNKNOWN":  0,
		"TAKAO":    1,
		"CHICHIBU": 2,
		"HAKONE":   3,
		"OKUTAMA":  4,
		"NIKKO":    5,
		"ENOSHIMA": 6,
		"KAMAKURA": 7,
		"YOKOSUKA": 8,
	}
)

func (x SendSpeed_Train) Enum() *SendSpeed_Train {
	p := new(SendSpeed_Train)
	*p = x
	return p
}

func (x SendSpeed_Train) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendSpeed_Train) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_speedControl_proto_enumTypes[0].Descriptor()
}

func (SendSpeed_Train) Type() protoreflect.EnumType {
	return &file_proto_speedControl_proto_enumTypes[0]
}

func (x SendSpeed_Train) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendSpeed_Train.Descriptor instead.
func (SendSpeed_Train) EnumDescriptor() ([]byte, []int) {
	return file_proto_speedControl_proto_rawDescGZIP(), []int{0, 0}
}

type StatusCode_Code int32

const (
	StatusCode_UNKNOWN StatusCode_Code = 0
	StatusCode_SUCCESS StatusCode_Code = 1
	StatusCode_FAILED  StatusCode_Code = 2
)

// Enum value maps for StatusCode_Code.
var (
	StatusCode_Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILED",
	}
	StatusCode_Code_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILED":  2,
	}
)

func (x StatusCode_Code) Enum() *StatusCode_Code {
	p := new(StatusCode_Code)
	*p = x
	return p
}

func (x StatusCode_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_speedControl_proto_enumTypes[1].Descriptor()
}

func (StatusCode_Code) Type() protoreflect.EnumType {
	return &file_proto_speedControl_proto_enumTypes[1]
}

func (x StatusCode_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode_Code.Descriptor instead.
func (StatusCode_Code) EnumDescriptor() ([]byte, []int) {
	return file_proto_speedControl_proto_rawDescGZIP(), []int{1, 0}
}

type SendSpeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed int32           `protobuf:"varint,1,opt,name=Speed,proto3" json:"Speed,omitempty"`
	Train SendSpeed_Train `protobuf:"varint,2,opt,name=train,proto3,enum=SendSpeed_Train" json:"train,omitempty"`
}

func (x *SendSpeed) Reset() {
	*x = SendSpeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_speedControl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSpeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSpeed) ProtoMessage() {}

func (x *SendSpeed) ProtoReflect() protoreflect.Message {
	mi := &file_proto_speedControl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSpeed.ProtoReflect.Descriptor instead.
func (*SendSpeed) Descriptor() ([]byte, []int) {
	return file_proto_speedControl_proto_rawDescGZIP(), []int{0}
}

func (x *SendSpeed) GetSpeed() int32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *SendSpeed) GetTrain() SendSpeed_Train {
	if x != nil {
		return x.Train
	}
	return SendSpeed_UNKNOWN
}

type StatusCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StatusCode_Code `protobuf:"varint,1,opt,name=code,proto3,enum=StatusCode_Code" json:"code,omitempty"`
}

func (x *StatusCode) Reset() {
	*x = StatusCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_speedControl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCode) ProtoMessage() {}

func (x *StatusCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_speedControl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCode.ProtoReflect.Descriptor instead.
func (*StatusCode) Descriptor() ([]byte, []int) {
	return file_proto_speedControl_proto_rawDescGZIP(), []int{1}
}

func (x *StatusCode) GetCode() StatusCode_Code {
	if x != nil {
		return x.Code
	}
	return StatusCode_UNKNOWN
}

var File_proto_speedControl_proto protoreflect.FileDescriptor

var file_proto_speedControl_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x09, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x26,
	0x0a, 0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x7b, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x41, 0x4b, 0x41, 0x4f, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x49, 0x43, 0x48,
	0x49, 0x42, 0x55, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x41, 0x4b, 0x4f, 0x4e, 0x45, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4b, 0x55, 0x54, 0x41, 0x4d, 0x41, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x4e, 0x49, 0x4b, 0x4b, 0x4f, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x4f,
	0x53, 0x48, 0x49, 0x4d, 0x41, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x41, 0x4d, 0x41, 0x4b,
	0x55, 0x52, 0x41, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x59, 0x4f, 0x4b, 0x4f, 0x53, 0x55, 0x4b,
	0x41, 0x10, 0x08, 0x22, 0x60, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0x32, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x29,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x0a,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x1a, 0x0b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x73,
	0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_speedControl_proto_rawDescOnce sync.Once
	file_proto_speedControl_proto_rawDescData = file_proto_speedControl_proto_rawDesc
)

func file_proto_speedControl_proto_rawDescGZIP() []byte {
	file_proto_speedControl_proto_rawDescOnce.Do(func() {
		file_proto_speedControl_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_speedControl_proto_rawDescData)
	})
	return file_proto_speedControl_proto_rawDescData
}

var file_proto_speedControl_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_speedControl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_speedControl_proto_goTypes = []interface{}{
	(SendSpeed_Train)(0), // 0: SendSpeed.Train
	(StatusCode_Code)(0), // 1: StatusCode.Code
	(*SendSpeed)(nil),    // 2: SendSpeed
	(*StatusCode)(nil),   // 3: StatusCode
}
var file_proto_speedControl_proto_depIdxs = []int32{
	0, // 0: SendSpeed.train:type_name -> SendSpeed.Train
	1, // 1: StatusCode.code:type_name -> StatusCode.Code
	2, // 2: Speed.ControlSpeed:input_type -> SendSpeed
	3, // 3: Speed.ControlSpeed:output_type -> StatusCode
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_speedControl_proto_init() }
func file_proto_speedControl_proto_init() {
	if File_proto_speedControl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_speedControl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSpeed); i {
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
		file_proto_speedControl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCode); i {
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
			RawDescriptor: file_proto_speedControl_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_speedControl_proto_goTypes,
		DependencyIndexes: file_proto_speedControl_proto_depIdxs,
		EnumInfos:         file_proto_speedControl_proto_enumTypes,
		MessageInfos:      file_proto_speedControl_proto_msgTypes,
	}.Build()
	File_proto_speedControl_proto = out.File
	file_proto_speedControl_proto_rawDesc = nil
	file_proto_speedControl_proto_goTypes = nil
	file_proto_speedControl_proto_depIdxs = nil
}
