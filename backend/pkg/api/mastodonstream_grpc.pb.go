// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: mastodonstream.proto

package hubapiv1

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

const (
	PubSubService_PublishPost_FullMethodName = "/hubapiv1.PubSubService/PublishPost"
)

// PubSubServiceClient is the client API for PubSubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PubSubServiceClient interface {
	PublishPost(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishPostClient, error)
}

type pubSubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubSubServiceClient(cc grpc.ClientConnInterface) PubSubServiceClient {
	return &pubSubServiceClient{cc}
}

func (c *pubSubServiceClient) PublishPost(ctx context.Context, opts ...grpc.CallOption) (PubSubService_PublishPostClient, error) {
	stream, err := c.cc.NewStream(ctx, &PubSubService_ServiceDesc.Streams[0], PubSubService_PublishPost_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubServicePublishPostClient{stream}
	return x, nil
}

type PubSubService_PublishPostClient interface {
	Send(*PublicPost) error
	CloseAndRecv() (*PublishResponse, error)
	grpc.ClientStream
}

type pubSubServicePublishPostClient struct {
	grpc.ClientStream
}

func (x *pubSubServicePublishPostClient) Send(m *PublicPost) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubServicePublishPostClient) CloseAndRecv() (*PublishResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServiceServer is the server API for PubSubService service.
// All implementations must embed UnimplementedPubSubServiceServer
// for forward compatibility
type PubSubServiceServer interface {
	PublishPost(PubSubService_PublishPostServer) error
	mustEmbedUnimplementedPubSubServiceServer()
}

// UnimplementedPubSubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPubSubServiceServer struct {
}

func (UnimplementedPubSubServiceServer) PublishPost(PubSubService_PublishPostServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishPost not implemented")
}
func (UnimplementedPubSubServiceServer) mustEmbedUnimplementedPubSubServiceServer() {}

// UnsafePubSubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PubSubServiceServer will
// result in compilation errors.
type UnsafePubSubServiceServer interface {
	mustEmbedUnimplementedPubSubServiceServer()
}

func RegisterPubSubServiceServer(s grpc.ServiceRegistrar, srv PubSubServiceServer) {
	s.RegisterService(&PubSubService_ServiceDesc, srv)
}

func _PubSubService_PublishPost_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServiceServer).PublishPost(&pubSubServicePublishPostServer{stream})
}

type PubSubService_PublishPostServer interface {
	SendAndClose(*PublishResponse) error
	Recv() (*PublicPost, error)
	grpc.ServerStream
}

type pubSubServicePublishPostServer struct {
	grpc.ServerStream
}

func (x *pubSubServicePublishPostServer) SendAndClose(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubServicePublishPostServer) Recv() (*PublicPost, error) {
	m := new(PublicPost)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubService_ServiceDesc is the grpc.ServiceDesc for PubSubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PubSubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hubapiv1.PubSubService",
	HandlerType: (*PubSubServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishPost",
			Handler:       _PubSubService_PublishPost_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "mastodonstream.proto",
}
