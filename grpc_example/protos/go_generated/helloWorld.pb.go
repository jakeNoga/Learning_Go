// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: helloWorld.proto

package go_generated

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloWorld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloWorld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_helloWorld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings.
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloWorld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloWorld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_helloWorld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_helloWorld_proto protoreflect.FileDescriptor

var file_helloWorld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb5, 0x02,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x41, 0x67, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x13, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x4f, 0x0a, 0x12, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x42, 0x69, 0x64, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloWorld_proto_rawDescOnce sync.Once
	file_helloWorld_proto_rawDescData = file_helloWorld_proto_rawDesc
)

func file_helloWorld_proto_rawDescGZIP() []byte {
	file_helloWorld_proto_rawDescOnce.Do(func() {
		file_helloWorld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloWorld_proto_rawDescData)
	})
	return file_helloWorld_proto_rawDescData
}

var file_helloWorld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_helloWorld_proto_goTypes = []any{
	(*HelloRequest)(nil),  // 0: helloWorld.HelloRequest
	(*HelloResponse)(nil), // 1: helloWorld.HelloResponse
}
var file_helloWorld_proto_depIdxs = []int32{
	0, // 0: helloWorld.Greeter.SayHello:input_type -> helloWorld.HelloRequest
	0, // 1: helloWorld.Greeter.SayHelloAgain:input_type -> helloWorld.HelloRequest
	0, // 2: helloWorld.Greeter.SayHelloStreamReply:input_type -> helloWorld.HelloRequest
	0, // 3: helloWorld.Greeter.SayHelloBidiStream:input_type -> helloWorld.HelloRequest
	1, // 4: helloWorld.Greeter.SayHello:output_type -> helloWorld.HelloResponse
	1, // 5: helloWorld.Greeter.SayHelloAgain:output_type -> helloWorld.HelloResponse
	1, // 6: helloWorld.Greeter.SayHelloStreamReply:output_type -> helloWorld.HelloResponse
	1, // 7: helloWorld.Greeter.SayHelloBidiStream:output_type -> helloWorld.HelloResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloWorld_proto_init() }
func file_helloWorld_proto_init() {
	if File_helloWorld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloWorld_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HelloRequest); i {
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
		file_helloWorld_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_helloWorld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloWorld_proto_goTypes,
		DependencyIndexes: file_helloWorld_proto_depIdxs,
		MessageInfos:      file_helloWorld_proto_msgTypes,
	}.Build()
	File_helloWorld_proto = out.File
	file_helloWorld_proto_rawDesc = nil
	file_helloWorld_proto_goTypes = nil
	file_helloWorld_proto_depIdxs = nil
}
