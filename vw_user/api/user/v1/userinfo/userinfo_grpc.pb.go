// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: user/v1/userinfo/userinfo.proto

package v1info

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
	Userinfo_ModifyUsername_FullMethodName  = "/user.v1.info.Userinfo/ModifyUsername"
	Userinfo_GetUserDetail_FullMethodName   = "/user.v1.info.Userinfo/GetUserDetail"
	Userinfo_ModifyEmail_FullMethodName     = "/user.v1.info.Userinfo/ModifyEmail"
	Userinfo_ModifySignature_FullMethodName = "/user.v1.info.Userinfo/ModifySignature"
	Userinfo_UploadAvatar_FullMethodName    = "/user.v1.info.Userinfo/UploadAvatar"
	Userinfo_ModifyPassword_FullMethodName  = "/user.v1.info.Userinfo/ModifyPassword"
	Userinfo_ForgetPassword_FullMethodName  = "/user.v1.info.Userinfo/ForgetPassword"
)

// UserinfoClient is the client API for Userinfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserinfoClient interface {
	ModifyUsername(ctx context.Context, in *ModifyUsernameRequest, opts ...grpc.CallOption) (*ModifyUsernameResp, error)
	GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResp, error)
	ModifyEmail(ctx context.Context, in *ModifyEmailRequest, opts ...grpc.CallOption) (*ModifyEmailResp, error)
	ModifySignature(ctx context.Context, in *ModifySignatureRequest, opts ...grpc.CallOption) (*ModifySignatureResp, error)
	UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResp, error)
	ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResp, error)
	ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordResp, error)
}

type userinfoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserinfoClient(cc grpc.ClientConnInterface) UserinfoClient {
	return &userinfoClient{cc}
}

func (c *userinfoClient) ModifyUsername(ctx context.Context, in *ModifyUsernameRequest, opts ...grpc.CallOption) (*ModifyUsernameResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyUsernameResp)
	err := c.cc.Invoke(ctx, Userinfo_ModifyUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDetailResp)
	err := c.cc.Invoke(ctx, Userinfo_GetUserDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) ModifyEmail(ctx context.Context, in *ModifyEmailRequest, opts ...grpc.CallOption) (*ModifyEmailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyEmailResp)
	err := c.cc.Invoke(ctx, Userinfo_ModifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) ModifySignature(ctx context.Context, in *ModifySignatureRequest, opts ...grpc.CallOption) (*ModifySignatureResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifySignatureResp)
	err := c.cc.Invoke(ctx, Userinfo_ModifySignature_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadAvatarResp)
	err := c.cc.Invoke(ctx, Userinfo_UploadAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModifyPasswordResp)
	err := c.cc.Invoke(ctx, Userinfo_ModifyPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userinfoClient) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...grpc.CallOption) (*ForgetPasswordResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForgetPasswordResp)
	err := c.cc.Invoke(ctx, Userinfo_ForgetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserinfoServer is the server API for Userinfo service.
// All implementations must embed UnimplementedUserinfoServer
// for forward compatibility.
type UserinfoServer interface {
	ModifyUsername(context.Context, *ModifyUsernameRequest) (*ModifyUsernameResp, error)
	GetUserDetail(context.Context, *GetUserDetailRequest) (*GetUserDetailResp, error)
	ModifyEmail(context.Context, *ModifyEmailRequest) (*ModifyEmailResp, error)
	ModifySignature(context.Context, *ModifySignatureRequest) (*ModifySignatureResp, error)
	UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResp, error)
	ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResp, error)
	ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordResp, error)
	mustEmbedUnimplementedUserinfoServer()
}

// UnimplementedUserinfoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserinfoServer struct{}

func (UnimplementedUserinfoServer) ModifyUsername(context.Context, *ModifyUsernameRequest) (*ModifyUsernameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUsername not implemented")
}
func (UnimplementedUserinfoServer) GetUserDetail(context.Context, *GetUserDetailRequest) (*GetUserDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetail not implemented")
}
func (UnimplementedUserinfoServer) ModifyEmail(context.Context, *ModifyEmailRequest) (*ModifyEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyEmail not implemented")
}
func (UnimplementedUserinfoServer) ModifySignature(context.Context, *ModifySignatureRequest) (*ModifySignatureResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySignature not implemented")
}
func (UnimplementedUserinfoServer) UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedUserinfoServer) ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPassword not implemented")
}
func (UnimplementedUserinfoServer) ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgetPassword not implemented")
}
func (UnimplementedUserinfoServer) mustEmbedUnimplementedUserinfoServer() {}
func (UnimplementedUserinfoServer) testEmbeddedByValue()                  {}

// UnsafeUserinfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserinfoServer will
// result in compilation errors.
type UnsafeUserinfoServer interface {
	mustEmbedUnimplementedUserinfoServer()
}

func RegisterUserinfoServer(s grpc.ServiceRegistrar, srv UserinfoServer) {
	// If the following call pancis, it indicates UnimplementedUserinfoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Userinfo_ServiceDesc, srv)
}

func _Userinfo_ModifyUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).ModifyUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_ModifyUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).ModifyUsername(ctx, req.(*ModifyUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_GetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).GetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_GetUserDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).GetUserDetail(ctx, req.(*GetUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_ModifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).ModifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_ModifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).ModifyEmail(ctx, req.(*ModifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_ModifySignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).ModifySignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_ModifySignature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).ModifySignature(ctx, req.(*ModifySignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_UploadAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).UploadAvatar(ctx, req.(*UploadAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_ModifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).ModifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_ModifyPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).ModifyPassword(ctx, req.(*ModifyPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Userinfo_ForgetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserinfoServer).ForgetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Userinfo_ForgetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserinfoServer).ForgetPassword(ctx, req.(*ForgetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Userinfo_ServiceDesc is the grpc.ServiceDesc for Userinfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Userinfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.info.Userinfo",
	HandlerType: (*UserinfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyUsername",
			Handler:    _Userinfo_ModifyUsername_Handler,
		},
		{
			MethodName: "GetUserDetail",
			Handler:    _Userinfo_GetUserDetail_Handler,
		},
		{
			MethodName: "ModifyEmail",
			Handler:    _Userinfo_ModifyEmail_Handler,
		},
		{
			MethodName: "ModifySignature",
			Handler:    _Userinfo_ModifySignature_Handler,
		},
		{
			MethodName: "UploadAvatar",
			Handler:    _Userinfo_UploadAvatar_Handler,
		},
		{
			MethodName: "ModifyPassword",
			Handler:    _Userinfo_ModifyPassword_Handler,
		},
		{
			MethodName: "ForgetPassword",
			Handler:    _Userinfo_ForgetPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/userinfo/userinfo.proto",
}
