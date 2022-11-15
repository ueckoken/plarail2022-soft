// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/block.proto

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

type NotifyStateRequest_State int32

const (
	NotifyStateRequest_UNKNOWN NotifyStateRequest_State = 0
	NotifyStateRequest_OPEN    NotifyStateRequest_State = 1
	NotifyStateRequest_CLOSE   NotifyStateRequest_State = 2
)

// Enum value maps for NotifyStateRequest_State.
var (
	NotifyStateRequest_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "OPEN",
		2: "CLOSE",
	}
	NotifyStateRequest_State_value = map[string]int32{
		"UNKNOWN": 0,
		"OPEN":    1,
		"CLOSE":   2,
	}
)

func (x NotifyStateRequest_State) Enum() *NotifyStateRequest_State {
	p := new(NotifyStateRequest_State)
	*p = x
	return p
}

func (x NotifyStateRequest_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyStateRequest_State) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_block_proto_enumTypes[0].Descriptor()
}

func (NotifyStateRequest_State) Type() protoreflect.EnumType {
	return &file_proto_block_proto_enumTypes[0]
}

func (x NotifyStateRequest_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyStateRequest_State.Descriptor instead.
func (NotifyStateRequest_State) EnumDescriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{0, 0}
}

type NotifyStateResponse_Response int32

const (
	NotifyStateResponse_UNKNOWN NotifyStateResponse_Response = 0
	NotifyStateResponse_SUCCESS NotifyStateResponse_Response = 1
	NotifyStateResponse_FAILED  NotifyStateResponse_Response = 2
)

// Enum value maps for NotifyStateResponse_Response.
var (
	NotifyStateResponse_Response_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "FAILED",
	}
	NotifyStateResponse_Response_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"FAILED":  2,
	}
)

func (x NotifyStateResponse_Response) Enum() *NotifyStateResponse_Response {
	p := new(NotifyStateResponse_Response)
	*p = x
	return p
}

func (x NotifyStateResponse_Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyStateResponse_Response) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_block_proto_enumTypes[1].Descriptor()
}

func (NotifyStateResponse_Response) Type() protoreflect.EnumType {
	return &file_proto_block_proto_enumTypes[1]
}

func (x NotifyStateResponse_Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyStateResponse_Response.Descriptor instead.
func (NotifyStateResponse_Response) EnumDescriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{1, 0}
}

type Blocks_BlockId int32

const (
	Blocks_unknown        Blocks_BlockId = 0
	Blocks_shinjuku_b1    Blocks_BlockId = 1
	Blocks_shinjuku_b2    Blocks_BlockId = 2
	Blocks_sakurajosui_b1 Blocks_BlockId = 11
	Blocks_sakurajosui_b2 Blocks_BlockId = 12
	Blocks_sakurajosui_b3 Blocks_BlockId = 13
	Blocks_sakurajosui_b4 Blocks_BlockId = 14
	Blocks_sakurajosui_b5 Blocks_BlockId = 15
	Blocks_sakurajosui_b6 Blocks_BlockId = 16
	Blocks_chofu_b1       Blocks_BlockId = 21
	Blocks_chofu_b2       Blocks_BlockId = 22
	Blocks_chofu_b3       Blocks_BlockId = 23
	Blocks_chofu_b4       Blocks_BlockId = 24
	Blocks_chofu_b5       Blocks_BlockId = 25
	Blocks_hashimoto_b1   Blocks_BlockId = 31
	Blocks_hashimoto_b2   Blocks_BlockId = 32
	Blocks_hachioji_b1    Blocks_BlockId = 41
	Blocks_hashioji_b2    Blocks_BlockId = 42
)

// Enum value maps for Blocks_BlockId.
var (
	Blocks_BlockId_name = map[int32]string{
		0:  "unknown",
		1:  "shinjuku_b1",
		2:  "shinjuku_b2",
		11: "sakurajosui_b1",
		12: "sakurajosui_b2",
		13: "sakurajosui_b3",
		14: "sakurajosui_b4",
		15: "sakurajosui_b5",
		16: "sakurajosui_b6",
		21: "chofu_b1",
		22: "chofu_b2",
		23: "chofu_b3",
		24: "chofu_b4",
		25: "chofu_b5",
		31: "hashimoto_b1",
		32: "hashimoto_b2",
		41: "hachioji_b1",
		42: "hashioji_b2",
	}
	Blocks_BlockId_value = map[string]int32{
		"unknown":        0,
		"shinjuku_b1":    1,
		"shinjuku_b2":    2,
		"sakurajosui_b1": 11,
		"sakurajosui_b2": 12,
		"sakurajosui_b3": 13,
		"sakurajosui_b4": 14,
		"sakurajosui_b5": 15,
		"sakurajosui_b6": 16,
		"chofu_b1":       21,
		"chofu_b2":       22,
		"chofu_b3":       23,
		"chofu_b4":       24,
		"chofu_b5":       25,
		"hashimoto_b1":   31,
		"hashimoto_b2":   32,
		"hachioji_b1":    41,
		"hashioji_b2":    42,
	}
)

func (x Blocks_BlockId) Enum() *Blocks_BlockId {
	p := new(Blocks_BlockId)
	*p = x
	return p
}

func (x Blocks_BlockId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Blocks_BlockId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_block_proto_enumTypes[2].Descriptor()
}

func (Blocks_BlockId) Type() protoreflect.EnumType {
	return &file_proto_block_proto_enumTypes[2]
}

func (x Blocks_BlockId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Blocks_BlockId.Descriptor instead.
func (Blocks_BlockId) EnumDescriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{2, 0}
}

type NotifyStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State NotifyStateRequest_State `protobuf:"varint,2,opt,name=state,proto3,enum=NotifyStateRequest_State" json:"state,omitempty"`
	Block *Blocks                  `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *NotifyStateRequest) Reset() {
	*x = NotifyStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyStateRequest) ProtoMessage() {}

func (x *NotifyStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyStateRequest.ProtoReflect.Descriptor instead.
func (*NotifyStateRequest) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyStateRequest) GetState() NotifyStateRequest_State {
	if x != nil {
		return x.State
	}
	return NotifyStateRequest_UNKNOWN
}

func (x *NotifyStateRequest) GetBlock() *Blocks {
	if x != nil {
		return x.Block
	}
	return nil
}

type NotifyStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response NotifyStateResponse_Response `protobuf:"varint,1,opt,name=response,proto3,enum=NotifyStateResponse_Response" json:"response,omitempty"`
}

func (x *NotifyStateResponse) Reset() {
	*x = NotifyStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyStateResponse) ProtoMessage() {}

func (x *NotifyStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyStateResponse.ProtoReflect.Descriptor instead.
func (*NotifyStateResponse) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyStateResponse) GetResponse() NotifyStateResponse_Response {
	if x != nil {
		return x.Response
	}
	return NotifyStateResponse_UNKNOWN
}

type Blocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId Blocks_BlockId `protobuf:"varint,3,opt,name=blockId,proto3,enum=Blocks_BlockId" json:"blockId,omitempty"`
}

func (x *Blocks) Reset() {
	*x = Blocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blocks) ProtoMessage() {}

func (x *Blocks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blocks.ProtoReflect.Descriptor instead.
func (*Blocks) Descriptor() ([]byte, []int) {
	return file_proto_block_proto_rawDescGZIP(), []int{2}
}

func (x *Blocks) GetBlockId() Blocks_BlockId {
	if x != nil {
		return x.BlockId
	}
	return Blocks_unknown
}

var File_proto_block_proto protoreflect.FileDescriptor

var file_proto_block_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x29, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x10, 0x02, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0xf2, 0x02, 0x0a, 0x06, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x22, 0xbc, 0x02, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07,
	0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68, 0x69,
	0x6e, 0x6a, 0x75, 0x6b, 0x75, 0x5f, 0x62, 0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x68,
	0x69, 0x6e, 0x6a, 0x75, 0x6b, 0x75, 0x5f, 0x62, 0x32, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x73,
	0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x31, 0x10, 0x0b, 0x12,
	0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62,
	0x32, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73,
	0x75, 0x69, 0x5f, 0x62, 0x33, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72,
	0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x34, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x73,
	0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62, 0x35, 0x10, 0x0f, 0x12,
	0x12, 0x0a, 0x0e, 0x73, 0x61, 0x6b, 0x75, 0x72, 0x61, 0x6a, 0x6f, 0x73, 0x75, 0x69, 0x5f, 0x62,
	0x36, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x31, 0x10,
	0x15, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x32, 0x10, 0x16, 0x12,
	0x0c, 0x0a, 0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x33, 0x10, 0x17, 0x12, 0x0c, 0x0a,
	0x08, 0x63, 0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x34, 0x10, 0x18, 0x12, 0x0c, 0x0a, 0x08, 0x63,
	0x68, 0x6f, 0x66, 0x75, 0x5f, 0x62, 0x35, 0x10, 0x19, 0x12, 0x10, 0x0a, 0x0c, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x6d, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x31, 0x10, 0x1f, 0x12, 0x10, 0x0a, 0x0c, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x6d, 0x6f, 0x74, 0x6f, 0x5f, 0x62, 0x32, 0x10, 0x20, 0x12, 0x0f, 0x0a,
	0x0b, 0x68, 0x61, 0x63, 0x68, 0x69, 0x6f, 0x6a, 0x69, 0x5f, 0x62, 0x31, 0x10, 0x29, 0x12, 0x0f,
	0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x69, 0x6f, 0x6a, 0x69, 0x5f, 0x62, 0x32, 0x10, 0x2a, 0x32,
	0x4c, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x79, 0x6e,
	0x63, 0x12, 0x3a, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_block_proto_rawDescOnce sync.Once
	file_proto_block_proto_rawDescData = file_proto_block_proto_rawDesc
)

func file_proto_block_proto_rawDescGZIP() []byte {
	file_proto_block_proto_rawDescOnce.Do(func() {
		file_proto_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_block_proto_rawDescData)
	})
	return file_proto_block_proto_rawDescData
}

var file_proto_block_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_block_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_block_proto_goTypes = []interface{}{
	(NotifyStateRequest_State)(0),     // 0: NotifyStateRequest.State
	(NotifyStateResponse_Response)(0), // 1: NotifyStateResponse.Response
	(Blocks_BlockId)(0),               // 2: Blocks.BlockId
	(*NotifyStateRequest)(nil),        // 3: NotifyStateRequest
	(*NotifyStateResponse)(nil),       // 4: NotifyStateResponse
	(*Blocks)(nil),                    // 5: Blocks
}
var file_proto_block_proto_depIdxs = []int32{
	0, // 0: NotifyStateRequest.state:type_name -> NotifyStateRequest.State
	5, // 1: NotifyStateRequest.block:type_name -> Blocks
	1, // 2: NotifyStateResponse.response:type_name -> NotifyStateResponse.Response
	2, // 3: Blocks.blockId:type_name -> Blocks.BlockId
	3, // 4: BlockStateSync.NotifyState:input_type -> NotifyStateRequest
	4, // 5: BlockStateSync.NotifyState:output_type -> NotifyStateResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_block_proto_init() }
func file_proto_block_proto_init() {
	if File_proto_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyStateRequest); i {
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
		file_proto_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyStateResponse); i {
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
		file_proto_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blocks); i {
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
			RawDescriptor: file_proto_block_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_block_proto_goTypes,
		DependencyIndexes: file_proto_block_proto_depIdxs,
		EnumInfos:         file_proto_block_proto_enumTypes,
		MessageInfos:      file_proto_block_proto_msgTypes,
	}.Build()
	File_proto_block_proto = out.File
	file_proto_block_proto_rawDesc = nil
	file_proto_block_proto_goTypes = nil
	file_proto_block_proto_depIdxs = nil
}
