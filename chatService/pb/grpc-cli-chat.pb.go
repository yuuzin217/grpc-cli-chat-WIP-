// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: chatService/grpc-cli-chat.proto

package pb

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

// クライアントコード
type ClientCode int32

const (
	// 0 番は使用しない
	ClientCode_Unknown ClientCode = 0
	// 汎用エラー
	ClientCode_SystemError ClientCode = 1
	// ルーム番号が不明
	ClientCode_UnknownRoomNumber ClientCode = 2
	// 名前が未入力
	ClientCode_NoNameEntered ClientCode = 3
)

// Enum value maps for ClientCode.
var (
	ClientCode_name = map[int32]string{
		0: "Unknown",
		1: "SystemError",
		2: "UnknownRoomNumber",
		3: "NoNameEntered",
	}
	ClientCode_value = map[string]int32{
		"Unknown":           0,
		"SystemError":       1,
		"UnknownRoomNumber": 2,
		"NoNameEntered":     3,
	}
)

func (x ClientCode) Enum() *ClientCode {
	p := new(ClientCode)
	*p = x
	return p
}

func (x ClientCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientCode) Descriptor() protoreflect.EnumDescriptor {
	return file_chatService_grpc_cli_chat_proto_enumTypes[0].Descriptor()
}

func (ClientCode) Type() protoreflect.EnumType {
	return &file_chatService_grpc_cli_chat_proto_enumTypes[0]
}

func (x ClientCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientCode.Descriptor instead.
func (ClientCode) EnumDescriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{0}
}

// ルームリスト取得
type RoomListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoomListRequest) Reset() {
	*x = RoomListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListRequest) ProtoMessage() {}

func (x *RoomListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListRequest.ProtoReflect.Descriptor instead.
func (*RoomListRequest) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{0}
}

type RoomListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// map[RoomNumber]RoomName
	RoomList map[int32]string `protobuf:"bytes,1,rep,name=RoomList,proto3" json:"RoomList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RoomListResponse) Reset() {
	*x = RoomListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomListResponse) ProtoMessage() {}

func (x *RoomListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomListResponse.ProtoReflect.Descriptor instead.
func (*RoomListResponse) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{1}
}

func (x *RoomListResponse) GetRoomList() map[int32]string {
	if x != nil {
		return x.RoomList
	}
	return nil
}

// ルーム参加
type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	RoomNumber int32  `protobuf:"varint,2,opt,name=RoomNumber,proto3" json:"RoomNumber,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JoinRequest) GetRoomNumber() int32 {
	if x != nil {
		return x.RoomNumber
	}
	return 0
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Error  *Error `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{3}
}

func (x *JoinResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *JoinResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// チャット通信
type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Message        string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	RegisterStream bool   `protobuf:"varint,3,opt,name=RegisterStream,proto3" json:"RegisterStream,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ConnectRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConnectRequest) GetRegisterStream() bool {
	if x != nil {
		return x.RegisterStream
	}
	return false
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ClientCode `protobuf:"varint,1,opt,name=Code,proto3,enum=grpc.cli.chat.ClientCode" json:"Code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatService_grpc_cli_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_chatService_grpc_cli_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_chatService_grpc_cli_chat_proto_rawDescGZIP(), []int{6}
}

func (x *Error) GetCode() ClientCode {
	if x != nil {
		return x.Code
	}
	return ClientCode_Unknown
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_chatService_grpc_cli_chat_proto protoreflect.FileDescriptor

var file_chatService_grpc_cli_chat_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x22, 0x11, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x41, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x52, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x22, 0x3f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x54, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x6f, 0x6f,
	0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x64, 0x10, 0x03, 0x32, 0xf0, 0x01, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatService_grpc_cli_chat_proto_rawDescOnce sync.Once
	file_chatService_grpc_cli_chat_proto_rawDescData = file_chatService_grpc_cli_chat_proto_rawDesc
)

func file_chatService_grpc_cli_chat_proto_rawDescGZIP() []byte {
	file_chatService_grpc_cli_chat_proto_rawDescOnce.Do(func() {
		file_chatService_grpc_cli_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatService_grpc_cli_chat_proto_rawDescData)
	})
	return file_chatService_grpc_cli_chat_proto_rawDescData
}

var file_chatService_grpc_cli_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chatService_grpc_cli_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chatService_grpc_cli_chat_proto_goTypes = []interface{}{
	(ClientCode)(0),          // 0: grpc.cli.chat.ClientCode
	(*RoomListRequest)(nil),  // 1: grpc.cli.chat.RoomListRequest
	(*RoomListResponse)(nil), // 2: grpc.cli.chat.RoomListResponse
	(*JoinRequest)(nil),      // 3: grpc.cli.chat.JoinRequest
	(*JoinResponse)(nil),     // 4: grpc.cli.chat.JoinResponse
	(*ConnectRequest)(nil),   // 5: grpc.cli.chat.ConnectRequest
	(*ConnectResponse)(nil),  // 6: grpc.cli.chat.ConnectResponse
	(*Error)(nil),            // 7: grpc.cli.chat.Error
	nil,                      // 8: grpc.cli.chat.RoomListResponse.RoomListEntry
}
var file_chatService_grpc_cli_chat_proto_depIdxs = []int32{
	8, // 0: grpc.cli.chat.RoomListResponse.RoomList:type_name -> grpc.cli.chat.RoomListResponse.RoomListEntry
	7, // 1: grpc.cli.chat.JoinResponse.Error:type_name -> grpc.cli.chat.Error
	0, // 2: grpc.cli.chat.Error.Code:type_name -> grpc.cli.chat.ClientCode
	1, // 3: grpc.cli.chat.ChatService.GetRoomList:input_type -> grpc.cli.chat.RoomListRequest
	3, // 4: grpc.cli.chat.ChatService.JoinRoom:input_type -> grpc.cli.chat.JoinRequest
	5, // 5: grpc.cli.chat.ChatService.Connect:input_type -> grpc.cli.chat.ConnectRequest
	2, // 6: grpc.cli.chat.ChatService.GetRoomList:output_type -> grpc.cli.chat.RoomListResponse
	4, // 7: grpc.cli.chat.ChatService.JoinRoom:output_type -> grpc.cli.chat.JoinResponse
	6, // 8: grpc.cli.chat.ChatService.Connect:output_type -> grpc.cli.chat.ConnectResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chatService_grpc_cli_chat_proto_init() }
func file_chatService_grpc_cli_chat_proto_init() {
	if File_chatService_grpc_cli_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatService_grpc_cli_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListRequest); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomListResponse); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_chatService_grpc_cli_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_chatService_grpc_cli_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatService_grpc_cli_chat_proto_goTypes,
		DependencyIndexes: file_chatService_grpc_cli_chat_proto_depIdxs,
		EnumInfos:         file_chatService_grpc_cli_chat_proto_enumTypes,
		MessageInfos:      file_chatService_grpc_cli_chat_proto_msgTypes,
	}.Build()
	File_chatService_grpc_cli_chat_proto = out.File
	file_chatService_grpc_cli_chat_proto_rawDesc = nil
	file_chatService_grpc_cli_chat_proto_goTypes = nil
	file_chatService_grpc_cli_chat_proto_depIdxs = nil
}
