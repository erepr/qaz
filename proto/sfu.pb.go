// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.2
// source: cmd/server/grpc/proto/sfu.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SignalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*SignalRequest_Join
	//	*SignalRequest_Negotiate
	//	*SignalRequest_Trickle
	Payload isSignalRequest_Payload `protobuf_oneof:"payload"`
}

func (x *SignalRequest) Reset() {
	*x = SignalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalRequest) ProtoMessage() {}

func (x *SignalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalRequest.ProtoReflect.Descriptor instead.
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{0}
}

func (m *SignalRequest) GetPayload() isSignalRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *SignalRequest) GetJoin() *JoinRequest {
	if x, ok := x.GetPayload().(*SignalRequest_Join); ok {
		return x.Join
	}
	return nil
}

func (x *SignalRequest) GetNegotiate() *SessionDescription {
	if x, ok := x.GetPayload().(*SignalRequest_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (x *SignalRequest) GetTrickle() *Trickle {
	if x, ok := x.GetPayload().(*SignalRequest_Trickle); ok {
		return x.Trickle
	}
	return nil
}

type isSignalRequest_Payload interface {
	isSignalRequest_Payload()
}

type SignalRequest_Join struct {
	Join *JoinRequest `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalRequest_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalRequest_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalRequest_Join) isSignalRequest_Payload() {}

func (*SignalRequest_Negotiate) isSignalRequest_Payload() {}

func (*SignalRequest_Trickle) isSignalRequest_Payload() {}

type SignalReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*SignalReply_Join
	//	*SignalReply_Negotiate
	//	*SignalReply_Trickle
	Payload isSignalReply_Payload `protobuf_oneof:"payload"`
}

func (x *SignalReply) Reset() {
	*x = SignalReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalReply) ProtoMessage() {}

func (x *SignalReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalReply.ProtoReflect.Descriptor instead.
func (*SignalReply) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{1}
}

func (m *SignalReply) GetPayload() isSignalReply_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *SignalReply) GetJoin() *JoinReply {
	if x, ok := x.GetPayload().(*SignalReply_Join); ok {
		return x.Join
	}
	return nil
}

func (x *SignalReply) GetNegotiate() *SessionDescription {
	if x, ok := x.GetPayload().(*SignalReply_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (x *SignalReply) GetTrickle() *Trickle {
	if x, ok := x.GetPayload().(*SignalReply_Trickle); ok {
		return x.Trickle
	}
	return nil
}

type isSignalReply_Payload interface {
	isSignalReply_Payload()
}

type SignalReply_Join struct {
	Join *JoinReply `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalReply_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalReply_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalReply_Join) isSignalReply_Payload() {}

func (*SignalReply_Negotiate) isSignalReply_Payload() {}

func (*SignalReply_Trickle) isSignalReply_Payload() {}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid   string              `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Offer *SessionDescription `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[2]
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
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *JoinRequest) GetOffer() *SessionDescription {
	if x != nil {
		return x.Offer
	}
	return nil
}

type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid    string              `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Answer *SessionDescription `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{3}
}

func (x *JoinReply) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *JoinReply) GetAnswer() *SessionDescription {
	if x != nil {
		return x.Answer
	}
	return nil
}

type Trickle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Init string `protobuf:"bytes,1,opt,name=init,proto3" json:"init,omitempty"`
}

func (x *Trickle) Reset() {
	*x = Trickle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trickle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trickle) ProtoMessage() {}

func (x *Trickle) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trickle.ProtoReflect.Descriptor instead.
func (*Trickle) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{4}
}

func (x *Trickle) GetInit() string {
	if x != nil {
		return x.Init
	}
	return ""
}

type SessionDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // "answer" | "offer" | "pranswer" | "rollback"
	Sdp  []byte `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
}

func (x *SessionDescription) Reset() {
	*x = SessionDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionDescription) ProtoMessage() {}

func (x *SessionDescription) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_server_grpc_proto_sfu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionDescription.ProtoReflect.Descriptor instead.
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP(), []int{5}
}

func (x *SessionDescription) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SessionDescription) GetSdp() []byte {
	if x != nil {
		return x.Sdp
	}
	return nil
}

var File_cmd_server_grpc_proto_sfu_proto protoreflect.FileDescriptor

var file_cmd_server_grpc_proto_sfu_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x66, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x73, 0x66, 0x75, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e,
	0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09,
	0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x72, 0x69,
	0x63, 0x6b, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x66, 0x75,
	0x2e, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa1,
	0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24,
	0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x66, 0x75, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x04,
	0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x09, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x07, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x73, 0x66, 0x75, 0x2e, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x07,
	0x74, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x4e, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x22, 0x4e, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x22, 0x1d, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x63, 0x6b, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x69,
	0x74, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x64, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x64, 0x70, 0x32, 0x3b, 0x0a,
	0x03, 0x53, 0x46, 0x55, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x12,
	0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x66, 0x75, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6f,
	0x6e, 0x2d, 0x73, 0x66, 0x75, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cmd_server_grpc_proto_sfu_proto_rawDescOnce sync.Once
	file_cmd_server_grpc_proto_sfu_proto_rawDescData = file_cmd_server_grpc_proto_sfu_proto_rawDesc
)

func file_cmd_server_grpc_proto_sfu_proto_rawDescGZIP() []byte {
	file_cmd_server_grpc_proto_sfu_proto_rawDescOnce.Do(func() {
		file_cmd_server_grpc_proto_sfu_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_server_grpc_proto_sfu_proto_rawDescData)
	})
	return file_cmd_server_grpc_proto_sfu_proto_rawDescData
}

var file_cmd_server_grpc_proto_sfu_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_server_grpc_proto_sfu_proto_goTypes = []interface{}{
	(*SignalRequest)(nil),      // 0: sfu.SignalRequest
	(*SignalReply)(nil),        // 1: sfu.SignalReply
	(*JoinRequest)(nil),        // 2: sfu.JoinRequest
	(*JoinReply)(nil),          // 3: sfu.JoinReply
	(*Trickle)(nil),            // 4: sfu.Trickle
	(*SessionDescription)(nil), // 5: sfu.SessionDescription
}
var file_cmd_server_grpc_proto_sfu_proto_depIdxs = []int32{
	2, // 0: sfu.SignalRequest.join:type_name -> sfu.JoinRequest
	5, // 1: sfu.SignalRequest.negotiate:type_name -> sfu.SessionDescription
	4, // 2: sfu.SignalRequest.trickle:type_name -> sfu.Trickle
	3, // 3: sfu.SignalReply.join:type_name -> sfu.JoinReply
	5, // 4: sfu.SignalReply.negotiate:type_name -> sfu.SessionDescription
	4, // 5: sfu.SignalReply.trickle:type_name -> sfu.Trickle
	5, // 6: sfu.JoinRequest.offer:type_name -> sfu.SessionDescription
	5, // 7: sfu.JoinReply.answer:type_name -> sfu.SessionDescription
	0, // 8: sfu.SFU.Signal:input_type -> sfu.SignalRequest
	1, // 9: sfu.SFU.Signal:output_type -> sfu.SignalReply
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cmd_server_grpc_proto_sfu_proto_init() }
func file_cmd_server_grpc_proto_sfu_proto_init() {
	if File_cmd_server_grpc_proto_sfu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalRequest); i {
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
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalReply); i {
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
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trickle); i {
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
		file_cmd_server_grpc_proto_sfu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionDescription); i {
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
	file_cmd_server_grpc_proto_sfu_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SignalRequest_Join)(nil),
		(*SignalRequest_Negotiate)(nil),
		(*SignalRequest_Trickle)(nil),
	}
	file_cmd_server_grpc_proto_sfu_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SignalReply_Join)(nil),
		(*SignalReply_Negotiate)(nil),
		(*SignalReply_Trickle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_server_grpc_proto_sfu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_server_grpc_proto_sfu_proto_goTypes,
		DependencyIndexes: file_cmd_server_grpc_proto_sfu_proto_depIdxs,
		MessageInfos:      file_cmd_server_grpc_proto_sfu_proto_msgTypes,
	}.Build()
	File_cmd_server_grpc_proto_sfu_proto = out.File
	file_cmd_server_grpc_proto_sfu_proto_rawDesc = nil
	file_cmd_server_grpc_proto_sfu_proto_goTypes = nil
	file_cmd_server_grpc_proto_sfu_proto_depIdxs = nil
}