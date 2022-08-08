// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: service/prompting/prompting.proto

package prompting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PromptingClient is the client API for Prompting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromptingClient interface {
	// Host allows clients to perform prompt hosting.
	Host(ctx context.Context, opts ...grpc.CallOption) (Prompting_HostClient, error)
	// Prompt performs prompting using a specific prompter.
	Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error)
}

type promptingClient struct {
	cc grpc.ClientConnInterface
}

func NewPromptingClient(cc grpc.ClientConnInterface) PromptingClient {
	return &promptingClient{cc}
}

func (c *promptingClient) Host(ctx context.Context, opts ...grpc.CallOption) (Prompting_HostClient, error) {
	stream, err := c.cc.NewStream(ctx, &Prompting_ServiceDesc.Streams[0], "/prompting.Prompting/Host", opts...)
	if err != nil {
		return nil, err
	}
	x := &promptingHostClient{stream}
	return x, nil
}

type Prompting_HostClient interface {
	Send(*HostRequest) error
	Recv() (*HostResponse, error)
	grpc.ClientStream
}

type promptingHostClient struct {
	grpc.ClientStream
}

func (x *promptingHostClient) Send(m *HostRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *promptingHostClient) Recv() (*HostResponse, error) {
	m := new(HostResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *promptingClient) Prompt(ctx context.Context, in *PromptRequest, opts ...grpc.CallOption) (*PromptResponse, error) {
	out := new(PromptResponse)
	err := c.cc.Invoke(ctx, "/prompting.Prompting/Prompt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromptingServer is the server API for Prompting service.
// All implementations must embed UnimplementedPromptingServer
// for forward compatibility
type PromptingServer interface {
	// Host allows clients to perform prompt hosting.
	Host(Prompting_HostServer) error
	// Prompt performs prompting using a specific prompter.
	Prompt(context.Context, *PromptRequest) (*PromptResponse, error)
	mustEmbedUnimplementedPromptingServer()
}

// UnimplementedPromptingServer must be embedded to have forward compatible implementations.
type UnimplementedPromptingServer struct {
}

func (UnimplementedPromptingServer) Host(Prompting_HostServer) error {
	return status.Errorf(codes.Unimplemented, "method Host not implemented")
}
func (UnimplementedPromptingServer) Prompt(context.Context, *PromptRequest) (*PromptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prompt not implemented")
}
func (UnimplementedPromptingServer) mustEmbedUnimplementedPromptingServer() {}

// UnsafePromptingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromptingServer will
// result in compilation errors.
type UnsafePromptingServer interface {
	mustEmbedUnimplementedPromptingServer()
}

func RegisterPromptingServer(s grpc.ServiceRegistrar, srv PromptingServer) {
	s.RegisterService(&Prompting_ServiceDesc, srv)
}

func _Prompting_Host_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PromptingServer).Host(&promptingHostServer{stream})
}

type Prompting_HostServer interface {
	Send(*HostResponse) error
	Recv() (*HostRequest, error)
	grpc.ServerStream
}

type promptingHostServer struct {
	grpc.ServerStream
}

func (x *promptingHostServer) Send(m *HostResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *promptingHostServer) Recv() (*HostRequest, error) {
	m := new(HostRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Prompting_Prompt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromptingServer).Prompt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prompting.Prompting/Prompt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromptingServer).Prompt(ctx, req.(*PromptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Prompting_ServiceDesc is the grpc.ServiceDesc for Prompting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Prompting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "prompting.Prompting",
	HandlerType: (*PromptingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prompt",
			Handler:    _Prompting_Prompt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Host",
			Handler:       _Prompting_Host_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/prompting/prompting.proto",
}
