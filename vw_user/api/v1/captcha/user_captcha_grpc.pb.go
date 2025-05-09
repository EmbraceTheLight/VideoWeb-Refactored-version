// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/captcha/user_captcha.proto

package captv1

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
	Captcha_DeleteCodeFromCache_FullMethodName = "/user.v1.captcha.Captcha/DeleteCodeFromCache"
	Captcha_SetCodeToCache_FullMethodName      = "/user.v1.captcha.Captcha/SetCodeToCache"
)

// CaptchaClient is the client API for Captcha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptchaClient interface {
	DeleteCodeFromCache(ctx context.Context, in *DeleteCodeFromCacheReq, opts ...grpc.CallOption) (*DeleteCodeFromCacheResp, error)
	SetCodeToCache(ctx context.Context, in *SetCodeToCacheReq, opts ...grpc.CallOption) (*SetCodeToCacheResp, error)
}

type captchaClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaClient(cc grpc.ClientConnInterface) CaptchaClient {
	return &captchaClient{cc}
}

func (c *captchaClient) DeleteCodeFromCache(ctx context.Context, in *DeleteCodeFromCacheReq, opts ...grpc.CallOption) (*DeleteCodeFromCacheResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCodeFromCacheResp)
	err := c.cc.Invoke(ctx, Captcha_DeleteCodeFromCache_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaClient) SetCodeToCache(ctx context.Context, in *SetCodeToCacheReq, opts ...grpc.CallOption) (*SetCodeToCacheResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCodeToCacheResp)
	err := c.cc.Invoke(ctx, Captcha_SetCodeToCache_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServer is the server API for Captcha service.
// All implementations must embed UnimplementedCaptchaServer
// for forward compatibility.
type CaptchaServer interface {
	DeleteCodeFromCache(context.Context, *DeleteCodeFromCacheReq) (*DeleteCodeFromCacheResp, error)
	SetCodeToCache(context.Context, *SetCodeToCacheReq) (*SetCodeToCacheResp, error)
	mustEmbedUnimplementedCaptchaServer()
}

// UnimplementedCaptchaServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaptchaServer struct{}

func (UnimplementedCaptchaServer) DeleteCodeFromCache(context.Context, *DeleteCodeFromCacheReq) (*DeleteCodeFromCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCodeFromCache not implemented")
}
func (UnimplementedCaptchaServer) SetCodeToCache(context.Context, *SetCodeToCacheReq) (*SetCodeToCacheResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCodeToCache not implemented")
}
func (UnimplementedCaptchaServer) mustEmbedUnimplementedCaptchaServer() {}
func (UnimplementedCaptchaServer) testEmbeddedByValue()                 {}

// UnsafeCaptchaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptchaServer will
// result in compilation errors.
type UnsafeCaptchaServer interface {
	mustEmbedUnimplementedCaptchaServer()
}

func RegisterCaptchaServer(s grpc.ServiceRegistrar, srv CaptchaServer) {
	// If the following call pancis, it indicates UnimplementedCaptchaServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Captcha_ServiceDesc, srv)
}

func _Captcha_DeleteCodeFromCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCodeFromCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).DeleteCodeFromCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_DeleteCodeFromCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).DeleteCodeFromCache(ctx, req.(*DeleteCodeFromCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Captcha_SetCodeToCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCodeToCacheReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).SetCodeToCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_SetCodeToCache_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).SetCodeToCache(ctx, req.(*SetCodeToCacheReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Captcha_ServiceDesc is the grpc.ServiceDesc for Captcha service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Captcha_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.captcha.Captcha",
	HandlerType: (*CaptchaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteCodeFromCache",
			Handler:    _Captcha_DeleteCodeFromCache_Handler,
		},
		{
			MethodName: "SetCodeToCache",
			Handler:    _Captcha_SetCodeToCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/captcha/user_captcha.proto",
}
