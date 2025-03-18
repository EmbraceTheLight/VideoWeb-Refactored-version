// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/captcha/captcha.proto

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
	Captcha_GetImageCaptcha_FullMethodName = "/gateway.api.v1.captcha.Captcha/GetImageCaptcha"
	Captcha_GetCodeCaptcha_FullMethodName  = "/gateway.api.v1.captcha.Captcha/GetCodeCaptcha"
)

// CaptchaClient is the client API for Captcha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptchaClient interface {
	GetImageCaptcha(ctx context.Context, in *GetImageCaptchaRequest, opts ...grpc.CallOption) (*GetImageCaptchaResp, error)
	GetCodeCaptcha(ctx context.Context, in *GetCodeCaptchaReq, opts ...grpc.CallOption) (*GetCodeCaptchaResp, error)
}

type captchaClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaClient(cc grpc.ClientConnInterface) CaptchaClient {
	return &captchaClient{cc}
}

func (c *captchaClient) GetImageCaptcha(ctx context.Context, in *GetImageCaptchaRequest, opts ...grpc.CallOption) (*GetImageCaptchaResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetImageCaptchaResp)
	err := c.cc.Invoke(ctx, Captcha_GetImageCaptcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaClient) GetCodeCaptcha(ctx context.Context, in *GetCodeCaptchaReq, opts ...grpc.CallOption) (*GetCodeCaptchaResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCodeCaptchaResp)
	err := c.cc.Invoke(ctx, Captcha_GetCodeCaptcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServer is the server API for Captcha service.
// All implementations must embed UnimplementedCaptchaServer
// for forward compatibility.
type CaptchaServer interface {
	GetImageCaptcha(context.Context, *GetImageCaptchaRequest) (*GetImageCaptchaResp, error)
	GetCodeCaptcha(context.Context, *GetCodeCaptchaReq) (*GetCodeCaptchaResp, error)
	mustEmbedUnimplementedCaptchaServer()
}

// UnimplementedCaptchaServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaptchaServer struct{}

func (UnimplementedCaptchaServer) GetImageCaptcha(context.Context, *GetImageCaptchaRequest) (*GetImageCaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageCaptcha not implemented")
}
func (UnimplementedCaptchaServer) GetCodeCaptcha(context.Context, *GetCodeCaptchaReq) (*GetCodeCaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeCaptcha not implemented")
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

func _Captcha_GetImageCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).GetImageCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_GetImageCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).GetImageCaptcha(ctx, req.(*GetImageCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Captcha_GetCodeCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeCaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).GetCodeCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_GetCodeCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).GetCodeCaptcha(ctx, req.(*GetCodeCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Captcha_ServiceDesc is the grpc.ServiceDesc for Captcha service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Captcha_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.api.v1.captcha.Captcha",
	HandlerType: (*CaptchaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImageCaptcha",
			Handler:    _Captcha_GetImageCaptcha_Handler,
		},
		{
			MethodName: "GetCodeCaptcha",
			Handler:    _Captcha_GetCodeCaptcha_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/captcha/captcha.proto",
}
