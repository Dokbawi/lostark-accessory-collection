// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/accessory/accessory.proto

package accessory

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccessoryService_GetAverageTradePrice_FullMethodName = "/accessory.AccessoryService/GetAverageTradePrice"
)

// AccessoryServiceClient is the client API for AccessoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessoryServiceClient interface {
	GetAverageTradePrice(ctx context.Context, in *GetAverageTradePriceRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetAverageTradePriceResponseList], error)
}

type accessoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessoryServiceClient(cc grpc.ClientConnInterface) AccessoryServiceClient {
	return &accessoryServiceClient{cc}
}

func (c *accessoryServiceClient) GetAverageTradePrice(ctx context.Context, in *GetAverageTradePriceRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetAverageTradePriceResponseList], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &AccessoryService_ServiceDesc.Streams[0], AccessoryService_GetAverageTradePrice_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetAverageTradePriceRequest, GetAverageTradePriceResponseList]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessoryService_GetAverageTradePriceClient = grpc.ServerStreamingClient[GetAverageTradePriceResponseList]

// AccessoryServiceServer is the server API for AccessoryService service.
// All implementations must embed UnimplementedAccessoryServiceServer
// for forward compatibility.
type AccessoryServiceServer interface {
	GetAverageTradePrice(*GetAverageTradePriceRequest, grpc.ServerStreamingServer[GetAverageTradePriceResponseList]) error
	mustEmbedUnimplementedAccessoryServiceServer()
}

// UnimplementedAccessoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessoryServiceServer struct{}

func (UnimplementedAccessoryServiceServer) GetAverageTradePrice(*GetAverageTradePriceRequest, grpc.ServerStreamingServer[GetAverageTradePriceResponseList]) error {
	return status.Errorf(codes.Unimplemented, "method GetAverageTradePrice not implemented")
}
func (UnimplementedAccessoryServiceServer) mustEmbedUnimplementedAccessoryServiceServer() {}
func (UnimplementedAccessoryServiceServer) testEmbeddedByValue()                          {}

// UnsafeAccessoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessoryServiceServer will
// result in compilation errors.
type UnsafeAccessoryServiceServer interface {
	mustEmbedUnimplementedAccessoryServiceServer()
}

func RegisterAccessoryServiceServer(s grpc.ServiceRegistrar, srv AccessoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccessoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccessoryService_ServiceDesc, srv)
}

func _AccessoryService_GetAverageTradePrice_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAverageTradePriceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccessoryServiceServer).GetAverageTradePrice(m, &grpc.GenericServerStream[GetAverageTradePriceRequest, GetAverageTradePriceResponseList]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type AccessoryService_GetAverageTradePriceServer = grpc.ServerStreamingServer[GetAverageTradePriceResponseList]

// AccessoryService_ServiceDesc is the grpc.ServiceDesc for AccessoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accessory.AccessoryService",
	HandlerType: (*AccessoryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAverageTradePrice",
			Handler:       _AccessoryService_GetAverageTradePrice_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/accessory/accessory.proto",
}
