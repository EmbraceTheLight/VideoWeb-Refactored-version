// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: v1/userfile/file.proto

package filev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileService_UploadAvatar_FullMethodName = "/gateway.api.v1.file.FileService/UploadAvatar"
	FileService_UpdateAvatar_FullMethodName = "/gateway.api.v1.file.FileService/UpdateAvatar"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadAvatar(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UploadAvatarResp, error)
	UpdateAvatar(ctx context.Context, in *UpdateAvatarReq, opts ...grpc.CallOption) (*UpdateAvatarResp, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadAvatar(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UploadAvatarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadAvatarResp)
	err := c.cc.Invoke(ctx, FileService_UploadAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateAvatar(ctx context.Context, in *UpdateAvatarReq, opts ...grpc.CallOption) (*UpdateAvatarResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAvatarResp)
	err := c.cc.Invoke(ctx, FileService_UpdateAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility.
type FileServiceServer interface {
	UploadAvatar(context.Context, *emptypb.Empty) (*UploadAvatarResp, error)
	UpdateAvatar(context.Context, *UpdateAvatarReq) (*UpdateAvatarResp, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileServiceServer struct{}

func (UnimplementedFileServiceServer) UploadAvatar(context.Context, *emptypb.Empty) (*UploadAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedFileServiceServer) UpdateAvatar(context.Context, *UpdateAvatarReq) (*UpdateAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatar not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}
func (UnimplementedFileServiceServer) testEmbeddedByValue()                     {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	// If the following call pancis, it indicates UnimplementedFileServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_UploadAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadAvatar(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileService_UpdateAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateAvatar(ctx, req.(*UpdateAvatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.api.v1.file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadAvatar",
			Handler:    _FileService_UploadAvatar_Handler,
		},
		{
			MethodName: "UpdateAvatar",
			Handler:    _FileService_UpdateAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/userfile/file.proto",
}
