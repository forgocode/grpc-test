// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/logService/proto/log.proto

package proto

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

// RecordLogClient is the client API for RecordLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordLogClient interface {
	RecordLogMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Reponse, error)
}

type recordLogClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordLogClient(cc grpc.ClientConnInterface) RecordLogClient {
	return &recordLogClient{cc}
}

func (c *recordLogClient) RecordLogMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Reponse, error) {
	out := new(Reponse)
	err := c.cc.Invoke(ctx, "/proto.RecordLog/RecordLogMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordLogServer is the server API for RecordLog service.
// All implementations must embed UnimplementedRecordLogServer
// for forward compatibility
type RecordLogServer interface {
	RecordLogMsg(context.Context, *Msg) (*Reponse, error)
	mustEmbedUnimplementedRecordLogServer()
}

// UnimplementedRecordLogServer must be embedded to have forward compatible implementations.
type UnimplementedRecordLogServer struct {
}

func (UnimplementedRecordLogServer) RecordLogMsg(context.Context, *Msg) (*Reponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordLogMsg not implemented")
}
func (UnimplementedRecordLogServer) mustEmbedUnimplementedRecordLogServer() {}

// UnsafeRecordLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordLogServer will
// result in compilation errors.
type UnsafeRecordLogServer interface {
	mustEmbedUnimplementedRecordLogServer()
}

func RegisterRecordLogServer(s grpc.ServiceRegistrar, srv RecordLogServer) {
	s.RegisterService(&RecordLog_ServiceDesc, srv)
}

func _RecordLog_RecordLogMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordLogServer).RecordLogMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RecordLog/RecordLogMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordLogServer).RecordLogMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordLog_ServiceDesc is the grpc.ServiceDesc for RecordLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RecordLog",
	HandlerType: (*RecordLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordLogMsg",
			Handler:    _RecordLog_RecordLogMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/logService/proto/log.proto",
}