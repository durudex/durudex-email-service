// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	types "github.com/Durudex/durudex-notif-service/internal/delivery/grpc/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotifServiceClient is the client API for NotifService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifServiceClient interface {
	UserVerifyCode(ctx context.Context, in *UserVerifyCodeRequest, opts ...grpc.CallOption) (*types.Status, error)
}

type notifServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifServiceClient(cc grpc.ClientConnInterface) NotifServiceClient {
	return &notifServiceClient{cc}
}

func (c *notifServiceClient) UserVerifyCode(ctx context.Context, in *UserVerifyCodeRequest, opts ...grpc.CallOption) (*types.Status, error) {
	out := new(types.Status)
	err := c.cc.Invoke(ctx, "/durudex.notif.NotifService/UserVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifServiceServer is the server API for NotifService service.
// All implementations must embed UnimplementedNotifServiceServer
// for forward compatibility
type NotifServiceServer interface {
	UserVerifyCode(context.Context, *UserVerifyCodeRequest) (*types.Status, error)
	mustEmbedUnimplementedNotifServiceServer()
}

// UnimplementedNotifServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotifServiceServer struct {
}

func (UnimplementedNotifServiceServer) UserVerifyCode(context.Context, *UserVerifyCodeRequest) (*types.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserVerifyCode not implemented")
}
func (UnimplementedNotifServiceServer) mustEmbedUnimplementedNotifServiceServer() {}

// UnsafeNotifServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifServiceServer will
// result in compilation errors.
type UnsafeNotifServiceServer interface {
	mustEmbedUnimplementedNotifServiceServer()
}

func RegisterNotifServiceServer(s grpc.ServiceRegistrar, srv NotifServiceServer) {
	s.RegisterService(&NotifService_ServiceDesc, srv)
}

func _NotifService_UserVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifServiceServer).UserVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/durudex.notif.NotifService/UserVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifServiceServer).UserVerifyCode(ctx, req.(*UserVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotifService_ServiceDesc is the grpc.ServiceDesc for NotifService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotifService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "durudex.notif.NotifService",
	HandlerType: (*NotifServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserVerifyCode",
			Handler:    _NotifService_UserVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/delivery/grpc/protobuf/notif.proto",
}