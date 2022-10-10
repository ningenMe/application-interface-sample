// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: interface-grpc/chat.proto

package interface_grpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_grpc_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interface_grpc_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_interface_grpc_chat_proto_rawDescGZIP(), []int{0}
}

func (x *PostRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatList []*Chat `protobuf:"bytes,1,rep,name=chatList,proto3" json:"chatList,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_grpc_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interface_grpc_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_interface_grpc_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetChatList() []*Chat {
	if x != nil {
		return x.ChatList
	}
	return nil
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body      string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interface_grpc_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_interface_grpc_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_interface_grpc_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Chat) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Chat) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_interface_grpc_chat_proto protoreflect.FileDescriptor

var file_interface_grpc_chat_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x21, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x04,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x7c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13, 0x2e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6e, 0x67, 0x65, 0x6e, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interface_grpc_chat_proto_rawDescOnce sync.Once
	file_interface_grpc_chat_proto_rawDescData = file_interface_grpc_chat_proto_rawDesc
)

func file_interface_grpc_chat_proto_rawDescGZIP() []byte {
	file_interface_grpc_chat_proto_rawDescOnce.Do(func() {
		file_interface_grpc_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_interface_grpc_chat_proto_rawDescData)
	})
	return file_interface_grpc_chat_proto_rawDescData
}

var file_interface_grpc_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interface_grpc_chat_proto_goTypes = []interface{}{
	(*PostRequest)(nil), // 0: sample.PostRequest
	(*GetResponse)(nil), // 1: sample.GetResponse
	(*Chat)(nil),        // 2: sample.Chat
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_interface_grpc_chat_proto_depIdxs = []int32{
	2, // 0: sample.GetResponse.chatList:type_name -> sample.Chat
	0, // 1: sample.ChatService.Post:input_type -> sample.PostRequest
	3, // 2: sample.ChatService.Get:input_type -> google.protobuf.Empty
	3, // 3: sample.ChatService.Post:output_type -> google.protobuf.Empty
	1, // 4: sample.ChatService.Get:output_type -> sample.GetResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_interface_grpc_chat_proto_init() }
func file_interface_grpc_chat_proto_init() {
	if File_interface_grpc_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interface_grpc_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequest); i {
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
		file_interface_grpc_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_interface_grpc_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
			RawDescriptor: file_interface_grpc_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interface_grpc_chat_proto_goTypes,
		DependencyIndexes: file_interface_grpc_chat_proto_depIdxs,
		MessageInfos:      file_interface_grpc_chat_proto_msgTypes,
	}.Build()
	File_interface_grpc_chat_proto = out.File
	file_interface_grpc_chat_proto_rawDesc = nil
	file_interface_grpc_chat_proto_goTypes = nil
	file_interface_grpc_chat_proto_depIdxs = nil
}
