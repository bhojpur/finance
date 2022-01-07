// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// FinanceUIClient is the client API for FinanceUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FinanceUIClient interface {
	// ListAlgorithmSpecs returns a list of Algorithm that can be started through the UI.
	ListAlgorithmSpecs(ctx context.Context, in *ListAlgorithmSpecsRequest, opts ...grpc.CallOption) (FinanceUI_ListAlgorithmSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type financeUIClient struct {
	cc grpc.ClientConnInterface
}

func NewFinanceUIClient(cc grpc.ClientConnInterface) FinanceUIClient {
	return &financeUIClient{cc}
}

func (c *financeUIClient) ListAlgorithmSpecs(ctx context.Context, in *ListAlgorithmSpecsRequest, opts ...grpc.CallOption) (FinanceUI_ListAlgorithmSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &FinanceUI_ServiceDesc.Streams[0], "/v1.FinanceUI/ListAlgorithmSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &financeUIListAlgorithmSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FinanceUI_ListAlgorithmSpecsClient interface {
	Recv() (*ListAlgorithmSpecsResponse, error)
	grpc.ClientStream
}

type financeUIListAlgorithmSpecsClient struct {
	grpc.ClientStream
}

func (x *financeUIListAlgorithmSpecsClient) Recv() (*ListAlgorithmSpecsResponse, error) {
	m := new(ListAlgorithmSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *financeUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.FinanceUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinanceUIServer is the server API for FinanceUI service.
// All implementations must embed UnimplementedFinanceUIServer
// for forward compatibility
type FinanceUIServer interface {
	// ListAlgorithmSpecs returns a list of Algorithm that can be started through the UI.
	ListAlgorithmSpecs(*ListAlgorithmSpecsRequest, FinanceUI_ListAlgorithmSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedFinanceUIServer()
}

// UnimplementedFinanceUIServer must be embedded to have forward compatible implementations.
type UnimplementedFinanceUIServer struct {
}

func (UnimplementedFinanceUIServer) ListAlgorithmSpecs(*ListAlgorithmSpecsRequest, FinanceUI_ListAlgorithmSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAlgorithmSpecs not implemented")
}
func (UnimplementedFinanceUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedFinanceUIServer) mustEmbedUnimplementedFinanceUIServer() {}

// UnsafeFinanceUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FinanceUIServer will
// result in compilation errors.
type UnsafeFinanceUIServer interface {
	mustEmbedUnimplementedFinanceUIServer()
}

func RegisterFinanceUIServer(s grpc.ServiceRegistrar, srv FinanceUIServer) {
	s.RegisterService(&FinanceUI_ServiceDesc, srv)
}

func _FinanceUI_ListAlgorithmSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListAlgorithmSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FinanceUIServer).ListAlgorithmSpecs(m, &financeUIListAlgorithmSpecsServer{stream})
}

type FinanceUI_ListAlgorithmSpecsServer interface {
	Send(*ListAlgorithmSpecsResponse) error
	grpc.ServerStream
}

type financeUIListAlgorithmSpecsServer struct {
	grpc.ServerStream
}

func (x *financeUIListAlgorithmSpecsServer) Send(m *ListAlgorithmSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FinanceUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinanceUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.FinanceUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinanceUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FinanceUI_ServiceDesc is the grpc.ServiceDesc for FinanceUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FinanceUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FinanceUI",
	HandlerType: (*FinanceUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _FinanceUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAlgorithmSpecs",
			Handler:       _FinanceUI_ListAlgorithmSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "finance-ui.proto",
}
