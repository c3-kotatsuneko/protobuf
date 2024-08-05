// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: game/rpc/game.proto

package rpc

import (
	resources "github.com/c3-kotatsuneko/protobuf/gen/game/resources"
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

type GameStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId  string            `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Event   resources.Event   `protobuf:"varint,2,opt,name=event,proto3,enum=game.resources.Event" json:"event,omitempty"`
	Mode    resources.Mode    `protobuf:"varint,3,opt,name=mode,proto3,enum=game.resources.Mode" json:"mode,omitempty"`
	Players *resources.Player `protobuf:"bytes,4,opt,name=players,proto3" json:"players,omitempty"`
}

func (x *GameStatusRequest) Reset() {
	*x = GameStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatusRequest) ProtoMessage() {}

func (x *GameStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatusRequest.ProtoReflect.Descriptor instead.
func (*GameStatusRequest) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameStatusRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GameStatusRequest) GetEvent() resources.Event {
	if x != nil {
		return x.Event
	}
	return resources.Event(0)
}

func (x *GameStatusRequest) GetMode() resources.Mode {
	if x != nil {
		return x.Mode
	}
	return resources.Mode(0)
}

func (x *GameStatusRequest) GetPlayers() *resources.Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type GameStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId  string              `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Event   resources.Event     `protobuf:"varint,2,opt,name=event,proto3,enum=game.resources.Event" json:"event,omitempty"`
	Players []*resources.Player `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty"`
	Time    int32               `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GameStatusResponse) Reset() {
	*x = GameStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStatusResponse) ProtoMessage() {}

func (x *GameStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStatusResponse.ProtoReflect.Descriptor instead.
func (*GameStatusResponse) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameStatusResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GameStatusResponse) GetEvent() resources.Event {
	if x != nil {
		return x.Event
	}
	return resources.Event(0)
}

func (x *GameStatusResponse) GetPlayers() []*resources.Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameStatusResponse) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type PhysicsInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string              `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RoomId   string              `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Objects  []*resources.Object `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *PhysicsInitRequest) Reset() {
	*x = PhysicsInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicsInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicsInitRequest) ProtoMessage() {}

func (x *PhysicsInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicsInitRequest.ProtoReflect.Descriptor instead.
func (*PhysicsInitRequest) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{2}
}

func (x *PhysicsInitRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PhysicsInitRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PhysicsInitRequest) GetObjects() []*resources.Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

type PhysicsInitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string              `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RoomId   string              `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	State    resources.GameState `protobuf:"varint,3,opt,name=state,proto3,enum=game.resources.GameState" json:"state,omitempty"`
	Objects  []*resources.Object `protobuf:"bytes,4,rep,name=objects,proto3" json:"objects,omitempty"`
	Hands    []*resources.Hand   `protobuf:"bytes,5,rep,name=hands,proto3" json:"hands,omitempty"`
}

func (x *PhysicsInitResponse) Reset() {
	*x = PhysicsInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicsInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicsInitResponse) ProtoMessage() {}

func (x *PhysicsInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicsInitResponse.ProtoReflect.Descriptor instead.
func (*PhysicsInitResponse) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{3}
}

func (x *PhysicsInitResponse) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PhysicsInitResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PhysicsInitResponse) GetState() resources.GameState {
	if x != nil {
		return x.State
	}
	return resources.GameState(0)
}

func (x *PhysicsInitResponse) GetObjects() []*resources.Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *PhysicsInitResponse) GetHands() []*resources.Hand {
	if x != nil {
		return x.Hands
	}
	return nil
}

type PhysicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string          `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RoomId   string          `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Hands    *resources.Hand `protobuf:"bytes,3,opt,name=hands,proto3" json:"hands,omitempty"`
}

func (x *PhysicsRequest) Reset() {
	*x = PhysicsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicsRequest) ProtoMessage() {}

func (x *PhysicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicsRequest.ProtoReflect.Descriptor instead.
func (*PhysicsRequest) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{4}
}

func (x *PhysicsRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PhysicsRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PhysicsRequest) GetHands() *resources.Hand {
	if x != nil {
		return x.Hands
	}
	return nil
}

type PhysicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId string              `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RoomId   string              `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Objects  []*resources.Object `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *PhysicsResponse) Reset() {
	*x = PhysicsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_rpc_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhysicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhysicsResponse) ProtoMessage() {}

func (x *PhysicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_rpc_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhysicsResponse.ProtoReflect.Descriptor instead.
func (*PhysicsResponse) Descriptor() ([]byte, []int) {
	return file_game_rpc_game_proto_rawDescGZIP(), []int{5}
}

func (x *PhysicsResponse) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *PhysicsResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *PhysicsResponse) GetObjects() []*resources.Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_game_rpc_game_proto protoreflect.FileDescriptor

var file_game_rpc_game_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x19, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x11, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x30, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x12, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73,
	0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x13, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x49,
	0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x68, 0x61, 0x6e, 0x64, 0x73,
	0x22, 0x72, 0x0a, 0x0e, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x05, 0x68, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x68,
	0x61, 0x6e, 0x64, 0x73, 0x22, 0x79, 0x0a, 0x0f, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42,
	0x8b, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x42, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x33, 0x2d, 0x6b, 0x6f, 0x74,
	0x61, 0x74, 0x73, 0x75, 0x6e, 0x65, 0x6b, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02,
	0x03, 0x47, 0x52, 0x58, 0xaa, 0x02, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x70, 0x63, 0xca,
	0x02, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x14, 0x47, 0x61, 0x6d,
	0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_rpc_game_proto_rawDescOnce sync.Once
	file_game_rpc_game_proto_rawDescData = file_game_rpc_game_proto_rawDesc
)

func file_game_rpc_game_proto_rawDescGZIP() []byte {
	file_game_rpc_game_proto_rawDescOnce.Do(func() {
		file_game_rpc_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_rpc_game_proto_rawDescData)
	})
	return file_game_rpc_game_proto_rawDescData
}

var file_game_rpc_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_game_rpc_game_proto_goTypes = []any{
	(*GameStatusRequest)(nil),   // 0: game.rpc.GameStatusRequest
	(*GameStatusResponse)(nil),  // 1: game.rpc.GameStatusResponse
	(*PhysicsInitRequest)(nil),  // 2: game.rpc.PhysicsInitRequest
	(*PhysicsInitResponse)(nil), // 3: game.rpc.PhysicsInitResponse
	(*PhysicsRequest)(nil),      // 4: game.rpc.PhysicsRequest
	(*PhysicsResponse)(nil),     // 5: game.rpc.PhysicsResponse
	(resources.Event)(0),        // 6: game.resources.Event
	(resources.Mode)(0),         // 7: game.resources.Mode
	(*resources.Player)(nil),    // 8: game.resources.Player
	(*resources.Object)(nil),    // 9: game.resources.Object
	(resources.GameState)(0),    // 10: game.resources.GameState
	(*resources.Hand)(nil),      // 11: game.resources.Hand
}
var file_game_rpc_game_proto_depIdxs = []int32{
	6,  // 0: game.rpc.GameStatusRequest.event:type_name -> game.resources.Event
	7,  // 1: game.rpc.GameStatusRequest.mode:type_name -> game.resources.Mode
	8,  // 2: game.rpc.GameStatusRequest.players:type_name -> game.resources.Player
	6,  // 3: game.rpc.GameStatusResponse.event:type_name -> game.resources.Event
	8,  // 4: game.rpc.GameStatusResponse.players:type_name -> game.resources.Player
	9,  // 5: game.rpc.PhysicsInitRequest.objects:type_name -> game.resources.Object
	10, // 6: game.rpc.PhysicsInitResponse.state:type_name -> game.resources.GameState
	9,  // 7: game.rpc.PhysicsInitResponse.objects:type_name -> game.resources.Object
	11, // 8: game.rpc.PhysicsInitResponse.hands:type_name -> game.resources.Hand
	11, // 9: game.rpc.PhysicsRequest.hands:type_name -> game.resources.Hand
	9,  // 10: game.rpc.PhysicsResponse.objects:type_name -> game.resources.Object
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_game_rpc_game_proto_init() }
func file_game_rpc_game_proto_init() {
	if File_game_rpc_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_rpc_game_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GameStatusRequest); i {
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
		file_game_rpc_game_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GameStatusResponse); i {
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
		file_game_rpc_game_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PhysicsInitRequest); i {
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
		file_game_rpc_game_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PhysicsInitResponse); i {
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
		file_game_rpc_game_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PhysicsRequest); i {
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
		file_game_rpc_game_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PhysicsResponse); i {
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
			RawDescriptor: file_game_rpc_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_rpc_game_proto_goTypes,
		DependencyIndexes: file_game_rpc_game_proto_depIdxs,
		MessageInfos:      file_game_rpc_game_proto_msgTypes,
	}.Build()
	File_game_rpc_game_proto = out.File
	file_game_rpc_game_proto_rawDesc = nil
	file_game_rpc_game_proto_goTypes = nil
	file_game_rpc_game_proto_depIdxs = nil
}
