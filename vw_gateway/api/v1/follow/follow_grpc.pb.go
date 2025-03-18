// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/follow/follow.proto

package followv1

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
	Follow_FollowUser_FullMethodName      = "/gateway.api.v1.follow.Follow/FollowUser"
	Follow_UnfollowUser_FullMethodName    = "/gateway.api.v1.follow.Follow/UnfollowUser"
	Follow_GetFolloweeInfo_FullMethodName = "/gateway.api.v1.follow.Follow/GetFolloweeInfo"
)

// FollowClient is the client API for Follow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowClient interface {
	FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*FollowUserResp, error)
	UnfollowUser(ctx context.Context, in *UnfollowUserReq, opts ...grpc.CallOption) (*UnfollowUserResp, error)
	GetFolloweeInfo(ctx context.Context, in *GetFolloweeInfoReq, opts ...grpc.CallOption) (*GetFolloweeInfoResp, error)
}

type followClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowClient(cc grpc.ClientConnInterface) FollowClient {
	return &followClient{cc}
}

func (c *followClient) FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*FollowUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FollowUserResp)
	err := c.cc.Invoke(ctx, Follow_FollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) UnfollowUser(ctx context.Context, in *UnfollowUserReq, opts ...grpc.CallOption) (*UnfollowUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnfollowUserResp)
	err := c.cc.Invoke(ctx, Follow_UnfollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followClient) GetFolloweeInfo(ctx context.Context, in *GetFolloweeInfoReq, opts ...grpc.CallOption) (*GetFolloweeInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFolloweeInfoResp)
	err := c.cc.Invoke(ctx, Follow_GetFolloweeInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServer is the server API for Follow service.
// All implementations must embed UnimplementedFollowServer
// for forward compatibility.
type FollowServer interface {
	FollowUser(context.Context, *FollowUserReq) (*FollowUserResp, error)
	UnfollowUser(context.Context, *UnfollowUserReq) (*UnfollowUserResp, error)
	GetFolloweeInfo(context.Context, *GetFolloweeInfoReq) (*GetFolloweeInfoResp, error)
	mustEmbedUnimplementedFollowServer()
}

// UnimplementedFollowServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFollowServer struct{}

func (UnimplementedFollowServer) FollowUser(context.Context, *FollowUserReq) (*FollowUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedFollowServer) UnfollowUser(context.Context, *UnfollowUserReq) (*UnfollowUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedFollowServer) GetFolloweeInfo(context.Context, *GetFolloweeInfoReq) (*GetFolloweeInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolloweeInfo not implemented")
}
func (UnimplementedFollowServer) mustEmbedUnimplementedFollowServer() {}
func (UnimplementedFollowServer) testEmbeddedByValue()                {}

// UnsafeFollowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServer will
// result in compilation errors.
type UnsafeFollowServer interface {
	mustEmbedUnimplementedFollowServer()
}

func RegisterFollowServer(s grpc.ServiceRegistrar, srv FollowServer) {
	// If the following call pancis, it indicates UnimplementedFollowServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Follow_ServiceDesc, srv)
}

func _Follow_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).FollowUser(ctx, req.(*FollowUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_UnfollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).UnfollowUser(ctx, req.(*UnfollowUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Follow_GetFolloweeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFolloweeInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServer).GetFolloweeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Follow_GetFolloweeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServer).GetFolloweeInfo(ctx, req.(*GetFolloweeInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Follow_ServiceDesc is the grpc.ServiceDesc for Follow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Follow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.api.v1.follow.Follow",
	HandlerType: (*FollowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowUser",
			Handler:    _Follow_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _Follow_UnfollowUser_Handler,
		},
		{
			MethodName: "GetFolloweeInfo",
			Handler:    _Follow_GetFolloweeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/follow/follow.proto",
}
