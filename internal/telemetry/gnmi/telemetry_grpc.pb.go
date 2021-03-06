// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gnmi

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

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelemetryClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (Telemetry_PublishClient, error)
}

type telemetryClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryClient(cc grpc.ClientConnInterface) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) Publish(ctx context.Context, opts ...grpc.CallOption) (Telemetry_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &Telemetry_ServiceDesc.Streams[0], "/telemetry2.Telemetry/publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &telemetryPublishClient{stream}
	return x, nil
}

type Telemetry_PublishClient interface {
	Send(*SubscribeResponse) error
	CloseAndRecv() (*TelemetryStreamResponse, error)
	grpc.ClientStream
}

type telemetryPublishClient struct {
	grpc.ClientStream
}

func (x *telemetryPublishClient) Send(m *SubscribeResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *telemetryPublishClient) CloseAndRecv() (*TelemetryStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TelemetryStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TelemetryServer is the server API for Telemetry service.
// All implementations must embed UnimplementedTelemetryServer
// for forward compatibility
type TelemetryServer interface {
	Publish(Telemetry_PublishServer) error
	mustEmbedUnimplementedTelemetryServer()
}

// UnimplementedTelemetryServer must be embedded to have forward compatible implementations.
type UnimplementedTelemetryServer struct {
}

func (UnimplementedTelemetryServer) Publish(Telemetry_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedTelemetryServer) mustEmbedUnimplementedTelemetryServer() {}

// UnsafeTelemetryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelemetryServer will
// result in compilation errors.
type UnsafeTelemetryServer interface {
	mustEmbedUnimplementedTelemetryServer()
}

func RegisterTelemetryServer(s grpc.ServiceRegistrar, srv TelemetryServer) {
	s.RegisterService(&Telemetry_ServiceDesc, srv)
}

func _Telemetry_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TelemetryServer).Publish(&telemetryPublishServer{stream})
}

type Telemetry_PublishServer interface {
	SendAndClose(*TelemetryStreamResponse) error
	Recv() (*SubscribeResponse, error)
	grpc.ServerStream
}

type telemetryPublishServer struct {
	grpc.ServerStream
}

func (x *telemetryPublishServer) SendAndClose(m *TelemetryStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *telemetryPublishServer) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Telemetry_ServiceDesc is the grpc.ServiceDesc for Telemetry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Telemetry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry2.Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "publish",
			Handler:       _Telemetry_Publish_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "telemetry.proto",
}
