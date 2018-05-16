// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/daemon/service/daemon.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_a255a8a3561fb2df, []int{0}
}
func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (dst *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(dst, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	// TODO: Should we encapsulate these inside a Version message type, perhaps
	// in the mutagen package?
	Major                uint64   `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor                uint64   `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
	Patch                uint64   `protobuf:"varint,3,opt,name=patch" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_a255a8a3561fb2df, []int{1}
}
func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (dst *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(dst, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetMajor() uint64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionResponse) GetMinor() uint64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionResponse) GetPatch() uint64 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type ShutdownRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownRequest) Reset()         { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()    {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_a255a8a3561fb2df, []int{2}
}
func (m *ShutdownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownRequest.Unmarshal(m, b)
}
func (m *ShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownRequest.Marshal(b, m, deterministic)
}
func (dst *ShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownRequest.Merge(dst, src)
}
func (m *ShutdownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownRequest.Size(m)
}
func (m *ShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownRequest proto.InternalMessageInfo

type ShutdownResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownResponse) Reset()         { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()    {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_daemon_a255a8a3561fb2df, []int{3}
}
func (m *ShutdownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownResponse.Unmarshal(m, b)
}
func (m *ShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownResponse.Marshal(b, m, deterministic)
}
func (dst *ShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownResponse.Merge(dst, src)
}
func (m *ShutdownResponse) XXX_Size() int {
	return xxx_messageInfo_ShutdownResponse.Size(m)
}
func (m *ShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionRequest)(nil), "service.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "service.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "service.ShutdownResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Daemon service

type DaemonClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/service.Daemon/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := grpc.Invoke(ctx, "/service.Daemon/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Daemon service

type DaemonServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Daemon/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Daemon/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Daemon_Version_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _Daemon_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/havoc-io/mutagen/pkg/daemon/service/daemon.proto",
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/daemon/service/daemon.proto", fileDescriptor_daemon_a255a8a3561fb2df)
}

var fileDescriptor_daemon_a255a8a3561fb2df = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x1d, 0x7f, 0x5a, 0xb9, 0x0b, 0x5b, 0x83, 0xe0, 0xd8, 0x95, 0xcc, 0xca, 0x8d, 0x13,
	0xd0, 0xa5, 0x20, 0x08, 0x3e, 0x41, 0x0b, 0xee, 0xd3, 0xf4, 0x32, 0x89, 0x92, 0xdc, 0x98, 0x9f,
	0xfa, 0x10, 0xbe, 0xb4, 0x34, 0x49, 0x5b, 0xd4, 0x2e, 0xcf, 0x77, 0xc3, 0x39, 0x1f, 0x81, 0xa7,
	0x41, 0x47, 0x95, 0x96, 0xbd, 0x24, 0xc3, 0x95, 0x58, 0x93, 0xbc, 0xd7, 0xc4, 0x4d, 0x8a, 0x62,
	0x40, 0xcb, 0xdd, 0xc7, 0xc0, 0x57, 0x02, 0x0d, 0x59, 0x1e, 0xd0, 0xaf, 0xb5, 0xc4, 0x1a, 0x7b,
	0xe7, 0x29, 0x12, 0x1b, 0x57, 0xda, 0x4d, 0xe1, 0xe2, 0x0d, 0x7d, 0xd0, 0x64, 0xe7, 0xf8, 0x99,
	0x30, 0xc4, 0x6e, 0x01, 0x93, 0x1d, 0x09, 0x8e, 0x6c, 0x40, 0x76, 0x05, 0x67, 0x46, 0xbc, 0x93,
	0x6f, 0x9b, 0xdb, 0xe6, 0xee, 0x74, 0x5e, 0x42, 0xa6, 0xda, 0x92, 0x6f, 0x8f, 0x2b, 0xdd, 0x84,
	0x0d, 0x75, 0x22, 0x4a, 0xd5, 0x9e, 0x14, 0x9a, 0x43, 0x77, 0x09, 0x93, 0x85, 0x4a, 0x71, 0x45,
	0x5f, 0xbb, 0x1d, 0x06, 0xd3, 0x3d, 0x2a, 0x43, 0x0f, 0xdf, 0x0d, 0x8c, 0x5e, 0xb3, 0x27, 0x7b,
	0x86, 0x71, 0xd5, 0x60, 0xd7, 0x7d, 0xb5, 0xed, 0x7f, 0xab, 0xce, 0xda, 0xff, 0x87, 0x52, 0xd4,
	0x1d, 0xb1, 0x17, 0x38, 0xdf, 0xd6, 0xb3, 0xfd, 0xbb, 0x3f, 0x12, 0xb3, 0x9b, 0x03, 0x97, 0x6d,
	0xc5, 0x72, 0x94, 0xff, 0xea, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xc0, 0xd6, 0xe3, 0x6a,
	0x01, 0x00, 0x00,
}
