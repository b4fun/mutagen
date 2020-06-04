// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/forwarding/forwarding.proto

package forwarding

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/mutagen-io/mutagen/pkg/forwarding"
	selection "github.com/mutagen-io/mutagen/pkg/selection"
	url "github.com/mutagen-io/mutagen/pkg/url"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// CreationSpecification contains the metadata required for a new session.
type CreationSpecification struct {
	// Source is the source endpoint URL for the session.
	Source *url.URL `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Destination is the destination endpoint URL for the session.
	Destination *url.URL `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	// Configuration is the base session configuration. It is the result of
	// merging the global configuration (unless disabled), any manually
	// specified configuration file, and any command line configuration
	// parameters.
	Configuration *forwarding.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// ConfigurationSource is the source-specific session configuration. It is
	// determined based on command line configuration parameters.
	ConfigurationSource *forwarding.Configuration `protobuf:"bytes,4,opt,name=configurationSource,proto3" json:"configurationSource,omitempty"`
	// ConfigurationDestination is the destination-specific session
	// configuration. It is determined based on command line configuration
	// parameters.
	ConfigurationDestination *forwarding.Configuration `protobuf:"bytes,5,opt,name=configurationDestination,proto3" json:"configurationDestination,omitempty"`
	// Name is the name for the session object.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Labels are the labels for the session object.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Paused indicates whether or not to create the session pre-paused.
	Paused               bool     `protobuf:"varint,8,opt,name=paused,proto3" json:"paused,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreationSpecification) Reset()         { *m = CreationSpecification{} }
func (m *CreationSpecification) String() string { return proto.CompactTextString(m) }
func (*CreationSpecification) ProtoMessage()    {}
func (*CreationSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{0}
}

func (m *CreationSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreationSpecification.Unmarshal(m, b)
}
func (m *CreationSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreationSpecification.Marshal(b, m, deterministic)
}
func (m *CreationSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreationSpecification.Merge(m, src)
}
func (m *CreationSpecification) XXX_Size() int {
	return xxx_messageInfo_CreationSpecification.Size(m)
}
func (m *CreationSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_CreationSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_CreationSpecification proto.InternalMessageInfo

func (m *CreationSpecification) GetSource() *url.URL {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CreationSpecification) GetDestination() *url.URL {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *CreationSpecification) GetConfiguration() *forwarding.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *CreationSpecification) GetConfigurationSource() *forwarding.Configuration {
	if m != nil {
		return m.ConfigurationSource
	}
	return nil
}

func (m *CreationSpecification) GetConfigurationDestination() *forwarding.Configuration {
	if m != nil {
		return m.ConfigurationDestination
	}
	return nil
}

func (m *CreationSpecification) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreationSpecification) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CreationSpecification) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

// CreateRequest encodes a request for session creation.
type CreateRequest struct {
	// Prompter is the prompter identifier to use for creating sessions.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Specification is the creation specification.
	Specification        *CreationSpecification `protobuf:"bytes,2,opt,name=specification,proto3" json:"specification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *CreateRequest) GetSpecification() *CreationSpecification {
	if m != nil {
		return m.Specification
	}
	return nil
}

// CreateResponse encodes a session creation response.
type CreateResponse struct {
	// Session is the resulting session identifier.
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

// ListRequest encodes a request for session metadata.
type ListRequest struct {
	// Selection is the session selection criteria.
	Selection *selection.Selection `protobuf:"bytes,1,opt,name=selection,proto3" json:"selection,omitempty"`
	// PreviousStateIndex is the previously seen state index. 0 may be provided
	// to force an immediate state listing.
	PreviousStateIndex   uint64   `protobuf:"varint,2,opt,name=previousStateIndex,proto3" json:"previousStateIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

func (m *ListRequest) GetPreviousStateIndex() uint64 {
	if m != nil {
		return m.PreviousStateIndex
	}
	return 0
}

// ListResponse encodes session metadata.
type ListResponse struct {
	// StateIndex is the state index associated with the session metadata.
	StateIndex uint64 `protobuf:"varint,1,opt,name=stateIndex,proto3" json:"stateIndex,omitempty"`
	// SessionStates are the session metadata states.
	SessionStates        []*forwarding.State `protobuf:"bytes,2,rep,name=sessionStates,proto3" json:"sessionStates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{4}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetStateIndex() uint64 {
	if m != nil {
		return m.StateIndex
	}
	return 0
}

func (m *ListResponse) GetSessionStates() []*forwarding.State {
	if m != nil {
		return m.SessionStates
	}
	return nil
}

// PauseRequest encodes a request to pause sessions.
type PauseRequest struct {
	// Prompter is the prompter to use for status message updates.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection            *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{5}
}

func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (m *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(m, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *PauseRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

// PauseResponse indicates completion of pause operation(s).
type PauseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{6}
}

func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (m *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(m, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

// ResumeRequest encodes a request to resume sessions.
type ResumeRequest struct {
	// Prompter is the prompter identifier to use for resuming sessions.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection            *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResumeRequest) Reset()         { *m = ResumeRequest{} }
func (m *ResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ResumeRequest) ProtoMessage()    {}
func (*ResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{7}
}

func (m *ResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeRequest.Unmarshal(m, b)
}
func (m *ResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeRequest.Marshal(b, m, deterministic)
}
func (m *ResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeRequest.Merge(m, src)
}
func (m *ResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ResumeRequest.Size(m)
}
func (m *ResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeRequest proto.InternalMessageInfo

func (m *ResumeRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *ResumeRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

// ResumeResponse indicates completion of resume operation(s).
type ResumeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResumeResponse) Reset()         { *m = ResumeResponse{} }
func (m *ResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ResumeResponse) ProtoMessage()    {}
func (*ResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{8}
}

func (m *ResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResumeResponse.Unmarshal(m, b)
}
func (m *ResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResumeResponse.Marshal(b, m, deterministic)
}
func (m *ResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResumeResponse.Merge(m, src)
}
func (m *ResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ResumeResponse.Size(m)
}
func (m *ResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResumeResponse proto.InternalMessageInfo

// TerminateRequest encodes a request to terminate sessions.
type TerminateRequest struct {
	// Prompter is the prompter to use for status message updates.
	Prompter string `protobuf:"bytes,1,opt,name=prompter,proto3" json:"prompter,omitempty"`
	// Selection is the session selection criteria.
	Selection            *selection.Selection `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TerminateRequest) Reset()         { *m = TerminateRequest{} }
func (m *TerminateRequest) String() string { return proto.CompactTextString(m) }
func (*TerminateRequest) ProtoMessage()    {}
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{9}
}

func (m *TerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateRequest.Unmarshal(m, b)
}
func (m *TerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateRequest.Marshal(b, m, deterministic)
}
func (m *TerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateRequest.Merge(m, src)
}
func (m *TerminateRequest) XXX_Size() int {
	return xxx_messageInfo_TerminateRequest.Size(m)
}
func (m *TerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateRequest proto.InternalMessageInfo

func (m *TerminateRequest) GetPrompter() string {
	if m != nil {
		return m.Prompter
	}
	return ""
}

func (m *TerminateRequest) GetSelection() *selection.Selection {
	if m != nil {
		return m.Selection
	}
	return nil
}

// TerminateResponse indicates completion of termination operation(s).
type TerminateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminateResponse) Reset()         { *m = TerminateResponse{} }
func (m *TerminateResponse) String() string { return proto.CompactTextString(m) }
func (*TerminateResponse) ProtoMessage()    {}
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3507425a8852e9f1, []int{10}
}

func (m *TerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminateResponse.Unmarshal(m, b)
}
func (m *TerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminateResponse.Marshal(b, m, deterministic)
}
func (m *TerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminateResponse.Merge(m, src)
}
func (m *TerminateResponse) XXX_Size() int {
	return xxx_messageInfo_TerminateResponse.Size(m)
}
func (m *TerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TerminateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreationSpecification)(nil), "forwarding.CreationSpecification")
	proto.RegisterMapType((map[string]string)(nil), "forwarding.CreationSpecification.LabelsEntry")
	proto.RegisterType((*CreateRequest)(nil), "forwarding.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "forwarding.CreateResponse")
	proto.RegisterType((*ListRequest)(nil), "forwarding.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "forwarding.ListResponse")
	proto.RegisterType((*PauseRequest)(nil), "forwarding.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "forwarding.PauseResponse")
	proto.RegisterType((*ResumeRequest)(nil), "forwarding.ResumeRequest")
	proto.RegisterType((*ResumeResponse)(nil), "forwarding.ResumeResponse")
	proto.RegisterType((*TerminateRequest)(nil), "forwarding.TerminateRequest")
	proto.RegisterType((*TerminateResponse)(nil), "forwarding.TerminateResponse")
}

func init() {
	proto.RegisterFile("service/forwarding/forwarding.proto", fileDescriptor_3507425a8852e9f1)
}

var fileDescriptor_3507425a8852e9f1 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x4f, 0xdb, 0x4c,
	0x14, 0x25, 0x0f, 0x4c, 0x72, 0x83, 0xf9, 0xe0, 0xc2, 0x47, 0x1d, 0xab, 0x45, 0xa9, 0xbb, 0x89,
	0x90, 0x70, 0xa4, 0xb4, 0xea, 0x53, 0x6a, 0x45, 0x29, 0xad, 0xda, 0xb2, 0xa8, 0x26, 0x65, 0xd3,
	0x45, 0x2b, 0xc7, 0xb9, 0xa4, 0x16, 0xf1, 0x83, 0x19, 0x9b, 0x96, 0xdf, 0xda, 0x3f, 0xd2, 0x65,
	0xe5, 0xf1, 0x23, 0x63, 0x08, 0x82, 0x0d, 0xbb, 0xfb, 0x38, 0xe7, 0x3e, 0xe6, 0xcc, 0x68, 0xe0,
	0x91, 0x20, 0x7e, 0xee, 0xb9, 0x34, 0x38, 0x09, 0xf9, 0x2f, 0x87, 0x4f, 0xbc, 0x60, 0xaa, 0x98,
	0x76, 0xc4, 0xc3, 0x38, 0x44, 0x98, 0x47, 0xcc, 0xae, 0xa0, 0x19, 0xb9, 0xb1, 0x17, 0x06, 0x83,
	0xd2, 0xca, 0x60, 0xe6, 0x8e, 0x52, 0xc3, 0x0d, 0x83, 0x13, 0x6f, 0x9a, 0x70, 0x47, 0xc9, 0x6f,
	0x2b, 0x79, 0x11, 0x3b, 0x31, 0xe5, 0x71, 0x3d, 0xe1, 0xb3, 0x41, 0xc2, 0x67, 0x99, 0x6b, 0xfd,
	0x6d, 0xc0, 0xff, 0x07, 0x9c, 0x24, 0x73, 0x14, 0x91, 0xeb, 0x9d, 0x78, 0xae, 0x74, 0xb0, 0x07,
	0x9a, 0x08, 0x13, 0xee, 0x92, 0x51, 0xeb, 0xd5, 0xfa, 0x9d, 0x61, 0xcb, 0x4e, 0x59, 0xc7, 0xec,
	0x88, 0xe5, 0x71, 0xdc, 0x85, 0xce, 0x84, 0x44, 0xec, 0x05, 0x92, 0x60, 0xd4, 0x2f, 0xc1, 0xd4,
	0x24, 0xbe, 0x01, 0xbd, 0x32, 0xa5, 0xd1, 0x90, 0xe8, 0xae, 0xad, 0xec, 0x7f, 0xa0, 0x02, 0x58,
	0x15, 0x8f, 0x9f, 0x61, 0xb3, 0x12, 0x18, 0x65, 0xb3, 0x35, 0x6f, 0x2a, 0xb3, 0x88, 0x85, 0xc7,
	0x60, 0x54, 0xc2, 0xef, 0x94, 0x35, 0x96, 0x6f, 0xaa, 0x78, 0x2d, 0x15, 0x11, 0x9a, 0x81, 0xe3,
	0x93, 0xa1, 0xf5, 0x6a, 0xfd, 0x36, 0x93, 0x36, 0x1e, 0x82, 0x36, 0x73, 0xc6, 0x34, 0x13, 0xc6,
	0x4a, 0xaf, 0xd1, 0xef, 0x0c, 0xf7, 0x2a, 0x85, 0x17, 0x9d, 0xbc, 0x7d, 0x24, 0xf1, 0x87, 0x41,
	0xcc, 0x2f, 0x58, 0x4e, 0xc6, 0x6d, 0xd0, 0x22, 0x27, 0x11, 0x34, 0x31, 0x5a, 0xbd, 0x5a, 0xbf,
	0xc5, 0x72, 0xcf, 0x7c, 0x01, 0x1d, 0x05, 0x8e, 0xeb, 0xd0, 0x38, 0xa5, 0x0b, 0xa9, 0x58, 0x9b,
	0xa5, 0x26, 0x6e, 0xc1, 0xf2, 0xb9, 0x33, 0x4b, 0x48, 0xca, 0xd3, 0x66, 0x99, 0xf3, 0xb2, 0xfe,
	0xbc, 0x66, 0xc5, 0xa0, 0xcb, 0xfe, 0xc4, 0xe8, 0x2c, 0x21, 0x11, 0xa3, 0x09, 0xad, 0x88, 0x87,
	0x7e, 0x14, 0x13, 0xcf, 0x2b, 0x94, 0x3e, 0x7e, 0x00, 0x5d, 0xa8, 0x43, 0xe6, 0x6a, 0x3f, 0xbc,
	0x71, 0x1b, 0x56, 0xe5, 0x59, 0xbb, 0xb0, 0x56, 0x74, 0x15, 0x51, 0x18, 0x08, 0x42, 0x03, 0x56,
	0x04, 0x09, 0x91, 0x16, 0xcd, 0xba, 0x16, 0xae, 0x75, 0x06, 0x9d, 0x23, 0x4f, 0xc4, 0xc5, 0x7c,
	0x43, 0x68, 0x97, 0xaf, 0x20, 0xbf, 0x94, 0x5b, 0xf6, 0xfc, 0x5d, 0x8c, 0x0a, 0x8b, 0xcd, 0x61,
	0x68, 0x03, 0x46, 0x9c, 0xce, 0xbd, 0x30, 0x11, 0xa3, 0xf4, 0x15, 0x7c, 0x0c, 0x26, 0xf4, 0x5b,
	0x0e, 0xdf, 0x64, 0x0b, 0x32, 0xd6, 0x14, 0x56, 0xb3, 0x96, 0xf9, 0x70, 0x3b, 0x00, 0x62, 0xce,
	0xab, 0x49, 0x9e, 0x12, 0xc1, 0x67, 0xa0, 0xe7, 0xd3, 0xca, 0x22, 0xc2, 0xa8, 0x4b, 0x95, 0x37,
	0xd4, 0x73, 0x91, 0x19, 0x56, 0xc5, 0x59, 0xdf, 0x61, 0xf5, 0x4b, 0x2a, 0xe1, 0x6d, 0x0e, 0xbf,
	0xb2, 0x78, 0xfd, 0x56, 0x8b, 0x5b, 0xff, 0x81, 0x9e, 0xd7, 0xcf, 0x36, 0xb1, 0x7e, 0x80, 0xce,
	0x48, 0x24, 0xfe, 0x9d, 0x75, 0x5c, 0x87, 0xb5, 0xa2, 0x41, 0xde, 0x72, 0x0c, 0xeb, 0x5f, 0x89,
	0xfb, 0xe9, 0xf3, 0xb8, 0xb3, 0xae, 0x9b, 0xb0, 0xa1, 0xf4, 0xc8, 0x1a, 0x0f, 0xff, 0xd4, 0x01,
	0xde, 0x97, 0x02, 0xe0, 0x3e, 0x68, 0xd9, 0x9d, 0xc3, 0xee, 0x95, 0xfb, 0x5a, 0x0c, 0x66, 0x9a,
	0x8b, 0x52, 0xf9, 0x22, 0x4b, 0xf8, 0x0a, 0x9a, 0xe9, 0xbd, 0xc0, 0x7b, 0x2a, 0x4a, 0xb9, 0x9c,
	0xa6, 0x71, 0x35, 0x51, 0x92, 0x5f, 0xc3, 0xb2, 0xd4, 0x02, 0x2b, 0x20, 0x55, 0x7e, 0xb3, 0xbb,
	0x20, 0x53, 0xf2, 0xf7, 0x41, 0xcb, 0x4e, 0xb6, 0x3a, 0x7f, 0x45, 0xce, 0xea, 0xfc, 0x97, 0x84,
	0x58, 0xc2, 0x4f, 0xd0, 0x2e, 0x8f, 0x09, 0xef, 0xab, 0xd0, 0xcb, 0x0a, 0x99, 0x0f, 0xae, 0xc9,
	0x16, 0xb5, 0xde, 0x3e, 0xfd, 0xf6, 0x64, 0xea, 0xc5, 0x3f, 0x93, 0xb1, 0xed, 0x86, 0xfe, 0xc0,
	0x4f, 0x62, 0x67, 0x4a, 0xc1, 0x9e, 0x17, 0x16, 0xe6, 0x20, 0x3a, 0x9d, 0x0e, 0xae, 0x7e, 0x75,
	0x63, 0x4d, 0x7e, 0x39, 0x8f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x4d, 0xfb, 0x22, 0x07,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ForwardingClient is the client API for Forwarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ForwardingClient interface {
	// Create creates a new session.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Pause pauses sessions.
	Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error)
	// Terminate terminates sessions.
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
}

type forwardingClient struct {
	cc grpc.ClientConnInterface
}

func NewForwardingClient(cc grpc.ClientConnInterface) ForwardingClient {
	return &forwardingClient{cc}
}

func (c *forwardingClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error) {
	out := new(PauseResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Resume(ctx context.Context, in *ResumeRequest, opts ...grpc.CallOption) (*ResumeResponse, error) {
	out := new(ResumeResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *forwardingClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/forwarding.Forwarding/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ForwardingServer is the server API for Forwarding service.
type ForwardingServer interface {
	// Create creates a new session.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// List returns metadata for existing sessions.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Pause pauses sessions.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)
	// Resume resumes paused or disconnected sessions.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)
	// Terminate terminates sessions.
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
}

// UnimplementedForwardingServer can be embedded to have forward compatible implementations.
type UnimplementedForwardingServer struct {
}

func (*UnimplementedForwardingServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedForwardingServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedForwardingServer) Pause(ctx context.Context, req *PauseRequest) (*PauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (*UnimplementedForwardingServer) Resume(ctx context.Context, req *ResumeRequest) (*ResumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (*UnimplementedForwardingServer) Terminate(ctx context.Context, req *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}

func RegisterForwardingServer(s *grpc.Server, srv ForwardingServer) {
	s.RegisterService(&_Forwarding_serviceDesc, srv)
}

func _Forwarding_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Pause(ctx, req.(*PauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Resume(ctx, req.(*ResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Forwarding_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ForwardingServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/forwarding.Forwarding/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ForwardingServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Forwarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "forwarding.Forwarding",
	HandlerType: (*ForwardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Forwarding_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Forwarding_List_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Forwarding_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _Forwarding_Resume_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _Forwarding_Terminate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/forwarding/forwarding.proto",
}
