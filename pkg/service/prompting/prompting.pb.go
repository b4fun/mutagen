// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: service/prompting/prompting.proto

package prompting

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

// HostRequest encodes either an initial request to perform prompt hosting or a
// follow-up response to a message or prompt.
type HostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AllowPrompts indicates whether or not the hoster will allow prompts. If
	// not, it will only receive message requests. This field may only be set on
	// the initial request.
	AllowPrompts bool `protobuf:"varint,1,opt,name=allowPrompts,proto3" json:"allowPrompts,omitempty"`
	// Response is the prompt response, if any. On the initial request, this
	// must be an empty string. When responding to a prompt, it may be any
	// value. When responding to a message, it must be an empty string.
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *HostRequest) Reset() {
	*x = HostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_prompting_prompting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostRequest) ProtoMessage() {}

func (x *HostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_prompting_prompting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostRequest.ProtoReflect.Descriptor instead.
func (*HostRequest) Descriptor() ([]byte, []int) {
	return file_service_prompting_prompting_proto_rawDescGZIP(), []int{0}
}

func (x *HostRequest) GetAllowPrompts() bool {
	if x != nil {
		return x.AllowPrompts
	}
	return false
}

func (x *HostRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// HostResponse encodes either an initial response to perform prompt hosting or
// a follow-up request for messaging or prompting.
type HostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier is the prompter identifier. It is only set in the initial
	// response sent after the initial request.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// IsPrompt indicates if the response is requesting a prompt (as opposed to
	// simple message display).
	IsPrompt bool `protobuf:"varint,2,opt,name=isPrompt,proto3" json:"isPrompt,omitempty"`
	// Message is the message associated with the prompt or message.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HostResponse) Reset() {
	*x = HostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_prompting_prompting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostResponse) ProtoMessage() {}

func (x *HostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_prompting_prompting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostResponse.ProtoReflect.Descriptor instead.
func (*HostResponse) Descriptor() ([]byte, []int) {
	return file_service_prompting_prompting_proto_rawDescGZIP(), []int{1}
}

func (x *HostResponse) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *HostResponse) GetIsPrompt() bool {
	if x != nil {
		return x.IsPrompt
	}
	return false
}

func (x *HostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// PromptRequest encodes a request for prompting by a specific prompter.
type PromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompter is the prompter identifier.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Prompt is the prompt to present.
	Prompt string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
}

func (x *PromptRequest) Reset() {
	*x = PromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_prompting_prompting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRequest) ProtoMessage() {}

func (x *PromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_prompting_prompting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRequest.ProtoReflect.Descriptor instead.
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return file_service_prompting_prompting_proto_rawDescGZIP(), []int{2}
}

func (x *PromptRequest) GetPrompter() string {
	if x != nil {
		return x.Prompter
	}
	return ""
}

func (x *PromptRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

// PromptResponse encodes the response from a prompter.
type PromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response is the response returned by the prompter.
	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PromptResponse) Reset() {
	*x = PromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_prompting_prompting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptResponse) ProtoMessage() {}

func (x *PromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_prompting_prompting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptResponse.ProtoReflect.Descriptor instead.
func (*PromptResponse) Descriptor() ([]byte, []int) {
	return file_service_prompting_prompting_proto_rawDescGZIP(), []int{3}
}

func (x *PromptResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_service_prompting_prompting_proto protoreflect.FileDescriptor

var file_service_prompting_prompting_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x4d,
	0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x0a,
	0x0c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x43, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x2c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75,
	0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_prompting_prompting_proto_rawDescOnce sync.Once
	file_service_prompting_prompting_proto_rawDescData = file_service_prompting_prompting_proto_rawDesc
)

func file_service_prompting_prompting_proto_rawDescGZIP() []byte {
	file_service_prompting_prompting_proto_rawDescOnce.Do(func() {
		file_service_prompting_prompting_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_prompting_prompting_proto_rawDescData)
	})
	return file_service_prompting_prompting_proto_rawDescData
}

var file_service_prompting_prompting_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_prompting_prompting_proto_goTypes = []interface{}{
	(*HostRequest)(nil),    // 0: prompting.HostRequest
	(*HostResponse)(nil),   // 1: prompting.HostResponse
	(*PromptRequest)(nil),  // 2: prompting.PromptRequest
	(*PromptResponse)(nil), // 3: prompting.PromptResponse
}
var file_service_prompting_prompting_proto_depIdxs = []int32{
	0, // 0: prompting.Prompting.Host:input_type -> prompting.HostRequest
	2, // 1: prompting.Prompting.Prompt:input_type -> prompting.PromptRequest
	1, // 2: prompting.Prompting.Host:output_type -> prompting.HostResponse
	3, // 3: prompting.Prompting.Prompt:output_type -> prompting.PromptResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_prompting_prompting_proto_init() }
func file_service_prompting_prompting_proto_init() {
	if File_service_prompting_prompting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_prompting_prompting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostRequest); i {
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
		file_service_prompting_prompting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostResponse); i {
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
		file_service_prompting_prompting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRequest); i {
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
		file_service_prompting_prompting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptResponse); i {
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
			RawDescriptor: file_service_prompting_prompting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_prompting_prompting_proto_goTypes,
		DependencyIndexes: file_service_prompting_prompting_proto_depIdxs,
		MessageInfos:      file_service_prompting_prompting_proto_msgTypes,
	}.Build()
	File_service_prompting_prompting_proto = out.File
	file_service_prompting_prompting_proto_rawDesc = nil
	file_service_prompting_prompting_proto_goTypes = nil
	file_service_prompting_prompting_proto_depIdxs = nil
}
