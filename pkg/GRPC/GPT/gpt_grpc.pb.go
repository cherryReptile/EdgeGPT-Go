// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: proto/gpt.proto

package GPT

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

// GptServiceClient is the client API for GptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GptServiceClient interface {
	Ask(ctx context.Context, in *AskRequest, opts ...grpc.CallOption) (GptService_AskClient, error)
}

type gptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGptServiceClient(cc grpc.ClientConnInterface) GptServiceClient {
	return &gptServiceClient{cc}
}

func (c *gptServiceClient) Ask(ctx context.Context, in *AskRequest, opts ...grpc.CallOption) (GptService_AskClient, error) {
	stream, err := c.cc.NewStream(ctx, &GptService_ServiceDesc.Streams[0], "/gpt.GptService/Ask", opts...)
	if err != nil {
		return nil, err
	}
	x := &gptServiceAskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GptService_AskClient interface {
	Recv() (*AskResponse, error)
	grpc.ClientStream
}

type gptServiceAskClient struct {
	grpc.ClientStream
}

func (x *gptServiceAskClient) Recv() (*AskResponse, error) {
	m := new(AskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GptServiceServer is the server API for GptService service.
// All implementations must embed UnimplementedGptServiceServer
// for forward compatibility
type GptServiceServer interface {
	Ask(*AskRequest, GptService_AskServer) error
	mustEmbedUnimplementedGptServiceServer()
}

// UnimplementedGptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGptServiceServer struct {
}

func (UnimplementedGptServiceServer) Ask(*AskRequest, GptService_AskServer) error {
	return status.Errorf(codes.Unimplemented, "method Ask not implemented")
}
func (UnimplementedGptServiceServer) mustEmbedUnimplementedGptServiceServer() {}

// UnsafeGptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GptServiceServer will
// result in compilation errors.
type UnsafeGptServiceServer interface {
	mustEmbedUnimplementedGptServiceServer()
}

func RegisterGptServiceServer(s grpc.ServiceRegistrar, srv GptServiceServer) {
	s.RegisterService(&GptService_ServiceDesc, srv)
}

func _GptService_Ask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GptServiceServer).Ask(m, &gptServiceAskServer{stream})
}

type GptService_AskServer interface {
	Send(*AskResponse) error
	grpc.ServerStream
}

type gptServiceAskServer struct {
	grpc.ServerStream
}

func (x *gptServiceAskServer) Send(m *AskResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GptService_ServiceDesc is the grpc.ServiceDesc for GptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gpt.GptService",
	HandlerType: (*GptServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ask",
			Handler:       _GptService_Ask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gpt.proto",
}
