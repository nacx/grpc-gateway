// Code generated by protoc-gen-go.
// source: examples/examplepb/stream.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import gengo_grpc_gateway_examples_sub "github.com/gengo/grpc-gateway/examples/sub"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for StreamService service

type StreamServiceClient interface {
	BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error)
	List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error)
	BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BulkCreate(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkCreateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[0], c.cc, "/gengo.grpc.gateway.examples.examplepb.StreamService/BulkCreate", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkCreateClient{stream}
	return x, nil
}

type StreamService_BulkCreateClient interface {
	Send(*ABitOfEverything) error
	CloseAndRecv() (*google_protobuf1.Empty, error)
	grpc.ClientStream
}

type streamServiceBulkCreateClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkCreateClient) Send(m *ABitOfEverything) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkCreateClient) CloseAndRecv() (*google_protobuf1.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(google_protobuf1.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[1], c.cc, "/gengo.grpc.gateway.examples.examplepb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*ABitOfEverything, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BulkEcho(ctx context.Context, opts ...grpc.CallOption) (StreamService_BulkEchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[2], c.cc, "/gengo.grpc.gateway.examples.examplepb.StreamService/BulkEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBulkEchoClient{stream}
	return x, nil
}

type StreamService_BulkEchoClient interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ClientStream
}

type streamServiceBulkEchoClient struct {
	grpc.ClientStream
}

func (x *streamServiceBulkEchoClient) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBulkEchoClient) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for StreamService service

type StreamServiceServer interface {
	BulkCreate(StreamService_BulkCreateServer) error
	List(*google_protobuf1.Empty, StreamService_ListServer) error
	BulkEcho(StreamService_BulkEchoServer) error
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_BulkCreate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkCreate(&streamServiceBulkCreateServer{stream})
}

type StreamService_BulkCreateServer interface {
	SendAndClose(*google_protobuf1.Empty) error
	Recv() (*ABitOfEverything, error)
	grpc.ServerStream
}

type streamServiceBulkCreateServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkCreateServer) SendAndClose(m *google_protobuf1.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkCreateServer) Recv() (*ABitOfEverything, error) {
	m := new(ABitOfEverything)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(google_protobuf1.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*ABitOfEverything) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *ABitOfEverything) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BulkEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BulkEcho(&streamServiceBulkEchoServer{stream})
}

type StreamService_BulkEchoServer interface {
	Send(*gengo_grpc_gateway_examples_sub.StringMessage) error
	Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error)
	grpc.ServerStream
}

type streamServiceBulkEchoServer struct {
	grpc.ServerStream
}

func (x *streamServiceBulkEchoServer) Send(m *gengo_grpc_gateway_examples_sub.StringMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBulkEchoServer) Recv() (*gengo_grpc_gateway_examples_sub.StringMessage, error) {
	m := new(gengo_grpc_gateway_examples_sub.StringMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gengo.grpc.gateway.examples.examplepb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BulkCreate",
			Handler:       _StreamService_BulkCreate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BulkEcho",
			Handler:       _StreamService_BulkEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}

func init() { proto.RegisterFile("examples/examplepb/stream.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x90, 0x3b, 0x4e, 0x2b, 0x31,
	0x14, 0x86, 0xe5, 0x7b, 0x11, 0x02, 0x23, 0x1a, 0x17, 0x14, 0x06, 0x29, 0x22, 0x02, 0xf1, 0x10,
	0x3a, 0x0e, 0x50, 0x20, 0xd1, 0x11, 0x94, 0x0e, 0x44, 0x91, 0x8e, 0x26, 0xb2, 0xa3, 0x13, 0xc7,
	0x22, 0x33, 0xb6, 0xc6, 0x4e, 0x20, 0x6d, 0x56, 0x80, 0xc4, 0x02, 0x58, 0x0f, 0x35, 0x5b, 0x60,
	0x21, 0x38, 0xf3, 0xaa, 0x02, 0x04, 0xba, 0x19, 0x9f, 0xd7, 0xf7, 0x7f, 0xb4, 0x81, 0x4f, 0x32,
	0x71, 0x23, 0xf4, 0xa2, 0xfc, 0x70, 0x4a, 0xf8, 0x90, 0xa1, 0x4c, 0xc0, 0x65, 0x36, 0x58, 0xb6,
	0xaf, 0x31, 0xd5, 0x16, 0x74, 0xe6, 0xfa, 0xa0, 0x65, 0xc0, 0x47, 0x39, 0x85, 0x6a, 0x06, 0xea,
	0x19, 0xbe, 0xa3, 0xad, 0xd5, 0x23, 0x14, 0xd2, 0x19, 0x21, 0xd3, 0xd4, 0x06, 0x19, 0x8c, 0x4d,
	0x7d, 0xb1, 0x84, 0x6f, 0x97, 0xd5, 0xfc, 0x4f, 0x8d, 0x07, 0x02, 0x13, 0x17, 0xa6, 0x65, 0xf1,
	0x64, 0x01, 0x82, 0xec, 0x29, 0x13, 0x7a, 0x76, 0xd0, 0xc3, 0x09, 0x66, 0xd3, 0x30, 0x34, 0xa9,
	0x2e, 0xbb, 0x79, 0xdd, 0xed, 0xc7, 0x4a, 0x24, 0xe8, 0xbd, 0xd4, 0x58, 0xd4, 0xce, 0xde, 0xfe,
	0xd3, 0xcd, 0x6e, 0x0e, 0xdf, 0xc5, 0x6c, 0x62, 0xfa, 0xc8, 0x9e, 0x09, 0xa5, 0xed, 0xf1, 0xe8,
	0xe1, 0x3a, 0x3e, 0x06, 0x64, 0x17, 0xb0, 0x54, 0x1a, 0xb8, 0x6a, 0x9b, 0x70, 0x37, 0xe8, 0xd4,
	0xb7, 0xf9, 0x16, 0x14, 0x09, 0xa0, 0x4a, 0x00, 0x9d, 0x79, 0x82, 0xa6, 0x98, 0xbd, 0x7f, 0xbc,
	0xfc, 0x3b, 0x6a, 0xee, 0x89, 0xc9, 0x69, 0x85, 0xbf, 0x08, 0x5e, 0xa8, 0x08, 0x71, 0x49, 0x8e,
	0x0f, 0x09, 0x9b, 0x11, 0xba, 0x72, 0x63, 0x7c, 0x60, 0x5f, 0xec, 0xe4, 0x7f, 0x85, 0x6c, 0x1e,
	0xe4, 0x30, 0xbb, 0xac, 0xf1, 0x03, 0x4c, 0x8b, 0xb0, 0x57, 0x42, 0xd7, 0xe6, 0x5e, 0x3a, 0xfd,
	0xa1, 0x65, 0xf0, 0xed, 0xc1, 0xa8, 0x19, 0xa2, 0xd4, 0x38, 0x77, 0x5b, 0xc8, 0xe6, 0xbf, 0xec,
	0x5f, 0x5e, 0x12, 0x46, 0x9a, 0x5c, 0x52, 0x8b, 0xb4, 0x37, 0xee, 0xd7, 0xeb, 0xa8, 0x6a, 0x35,
	0x57, 0x74, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xde, 0x54, 0xd1, 0xb6, 0xae, 0x02, 0x00, 0x00,
}