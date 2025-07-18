// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: mcsports/queue/v1/interaction.proto

package queuev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnqueueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	QueueName     string                 `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	PlayerIds     []string               `protobuf:"bytes,2,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnqueueRequest) Reset() {
	*x = EnqueueRequest{}
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnqueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueRequest) ProtoMessage() {}

func (x *EnqueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueRequest.ProtoReflect.Descriptor instead.
func (*EnqueueRequest) Descriptor() ([]byte, []int) {
	return file_mcsports_queue_v1_interaction_proto_rawDescGZIP(), []int{0}
}

func (x *EnqueueRequest) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *EnqueueRequest) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type EnqueueResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	QueueId       string                 `protobuf:"bytes,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnqueueResponse) Reset() {
	*x = EnqueueResponse{}
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnqueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnqueueResponse) ProtoMessage() {}

func (x *EnqueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnqueueResponse.ProtoReflect.Descriptor instead.
func (*EnqueueResponse) Descriptor() ([]byte, []int) {
	return file_mcsports_queue_v1_interaction_proto_rawDescGZIP(), []int{1}
}

func (x *EnqueueResponse) GetQueueId() string {
	if x != nil {
		return x.QueueId
	}
	return ""
}

type DequeueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerIds     []string               `protobuf:"bytes,1,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
	Forced        *bool                  `protobuf:"varint,2,opt,name=forced,proto3,oneof" json:"forced,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DequeueRequest) Reset() {
	*x = DequeueRequest{}
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DequeueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueRequest) ProtoMessage() {}

func (x *DequeueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueRequest.ProtoReflect.Descriptor instead.
func (*DequeueRequest) Descriptor() ([]byte, []int) {
	return file_mcsports_queue_v1_interaction_proto_rawDescGZIP(), []int{2}
}

func (x *DequeueRequest) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

func (x *DequeueRequest) GetForced() bool {
	if x != nil && x.Forced != nil {
		return *x.Forced
	}
	return false
}

type DequeueResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DequeueResponse) Reset() {
	*x = DequeueResponse{}
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DequeueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueResponse) ProtoMessage() {}

func (x *DequeueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mcsports_queue_v1_interaction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueResponse.ProtoReflect.Descriptor instead.
func (*DequeueResponse) Descriptor() ([]byte, []int) {
	return file_mcsports_queue_v1_interaction_proto_rawDescGZIP(), []int{3}
}

var File_mcsports_queue_v1_interaction_proto protoreflect.FileDescriptor

const file_mcsports_queue_v1_interaction_proto_rawDesc = "" +
	"\n" +
	"#mcsports/queue/v1/interaction.proto\x12\x11mcsports.queue.v1\"N\n" +
	"\x0eEnqueueRequest\x12\x1d\n" +
	"\n" +
	"queue_name\x18\x01 \x01(\tR\tqueueName\x12\x1d\n" +
	"\n" +
	"player_ids\x18\x02 \x03(\tR\tplayerIds\",\n" +
	"\x0fEnqueueResponse\x12\x19\n" +
	"\bqueue_id\x18\x01 \x01(\tR\aqueueId\"W\n" +
	"\x0eDequeueRequest\x12\x1d\n" +
	"\n" +
	"player_ids\x18\x01 \x03(\tR\tplayerIds\x12\x1b\n" +
	"\x06forced\x18\x02 \x01(\bH\x00R\x06forced\x88\x01\x01B\t\n" +
	"\a_forced\"\x11\n" +
	"\x0fDequeueResponse2\xb6\x01\n" +
	"\x10QueueInteraction\x12P\n" +
	"\aEnqueue\x12!.mcsports.queue.v1.EnqueueRequest\x1a\".mcsports.queue.v1.EnqueueResponse\x12P\n" +
	"\aDequeue\x12!.mcsports.queue.v1.DequeueRequest\x1a\".mcsports.queue.v1.DequeueResponseB\xcb\x01\n" +
	"\x15com.mcsports.queue.v1B\x10InteractionProtoP\x01Z:github.com/bufbuild/buf-tour/gen/mcsports/queue/v1;queuev1\xa2\x02\x03MQX\xaa\x02\x11Mcsports.Queue.V1\xca\x02\x11Mcsports\\Queue\\V1\xe2\x02\x1dMcsports\\Queue\\V1\\GPBMetadata\xea\x02\x13Mcsports::Queue::V1b\x06proto3"

var (
	file_mcsports_queue_v1_interaction_proto_rawDescOnce sync.Once
	file_mcsports_queue_v1_interaction_proto_rawDescData []byte
)

func file_mcsports_queue_v1_interaction_proto_rawDescGZIP() []byte {
	file_mcsports_queue_v1_interaction_proto_rawDescOnce.Do(func() {
		file_mcsports_queue_v1_interaction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mcsports_queue_v1_interaction_proto_rawDesc), len(file_mcsports_queue_v1_interaction_proto_rawDesc)))
	})
	return file_mcsports_queue_v1_interaction_proto_rawDescData
}

var file_mcsports_queue_v1_interaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mcsports_queue_v1_interaction_proto_goTypes = []any{
	(*EnqueueRequest)(nil),  // 0: mcsports.queue.v1.EnqueueRequest
	(*EnqueueResponse)(nil), // 1: mcsports.queue.v1.EnqueueResponse
	(*DequeueRequest)(nil),  // 2: mcsports.queue.v1.DequeueRequest
	(*DequeueResponse)(nil), // 3: mcsports.queue.v1.DequeueResponse
}
var file_mcsports_queue_v1_interaction_proto_depIdxs = []int32{
	0, // 0: mcsports.queue.v1.QueueInteraction.Enqueue:input_type -> mcsports.queue.v1.EnqueueRequest
	2, // 1: mcsports.queue.v1.QueueInteraction.Dequeue:input_type -> mcsports.queue.v1.DequeueRequest
	1, // 2: mcsports.queue.v1.QueueInteraction.Enqueue:output_type -> mcsports.queue.v1.EnqueueResponse
	3, // 3: mcsports.queue.v1.QueueInteraction.Dequeue:output_type -> mcsports.queue.v1.DequeueResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcsports_queue_v1_interaction_proto_init() }
func file_mcsports_queue_v1_interaction_proto_init() {
	if File_mcsports_queue_v1_interaction_proto != nil {
		return
	}
	file_mcsports_queue_v1_interaction_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mcsports_queue_v1_interaction_proto_rawDesc), len(file_mcsports_queue_v1_interaction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mcsports_queue_v1_interaction_proto_goTypes,
		DependencyIndexes: file_mcsports_queue_v1_interaction_proto_depIdxs,
		MessageInfos:      file_mcsports_queue_v1_interaction_proto_msgTypes,
	}.Build()
	File_mcsports_queue_v1_interaction_proto = out.File
	file_mcsports_queue_v1_interaction_proto_goTypes = nil
	file_mcsports_queue_v1_interaction_proto_depIdxs = nil
}
